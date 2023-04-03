package mongodbRepository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	model "task3_4/user-management/internal/entity"
	"time"
)

type UserRepository struct {
	collection *mongo.Collection
	con        context.Context
}

func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
	return &UserRepository{collection: db.Collection(collection), con: context.Background()}
}

func (ur *UserRepository) FindAll(page, pageSize int64, u []*model.User) ([]*model.User, error) {
	opts := options.Find().SetSkip((page - 1) * pageSize).SetLimit(pageSize)
	cursor, err := ur.collection.Find(ur.con, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ur.con, &u); err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *UserRepository) FindByID(stringID string) (u *model.User, err error) {
	ID, err := primitive.ObjectIDFromHex(stringID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %s", stringID)
	}

	f := bson.D{{"_id", ID}}

	err = ur.collection.FindOne(ur.con, f).Decode(&u)

	if err == mongo.ErrNoDocuments {
		return u, fmt.Errorf("could not found user with ID %s", stringID)
	}

	if err != nil {
		return nil, err
	}

	return
}

func (ur *UserRepository) FindByNickname(nickname string) (u *model.User, err error) {
	f := bson.D{{"nickname", nickname}}

	err = ur.collection.FindOne(ur.con, f).Decode(&u)

	if err == mongo.ErrNoDocuments {
		return u, fmt.Errorf("could not found user with nickname %s", nickname)
	}

	if err != nil {
		return nil, err
	}

	return
}

func (ur *UserRepository) FindLasHourtUserVoteByVoteID(voterID string) (u *model.User, err error) {
	vID, err := primitive.ObjectIDFromHex(voterID)
	if err != nil {
		return nil, fmt.Errorf("invalid vote ID: %s", voterID)
	}

	filter := bson.M{
		"votes.voted_at": bson.M{
			"$gte": time.Now().Add(-time.Hour).Unix(),
		},
		"votes.voter_id": vID,
	}

	sort := bson.M{
		"votes.voted_at": -1,
	}

	err = ur.collection.FindOne(ur.con, filter, options.FindOne().SetSort(sort)).Decode(&u)

	if err != nil {
		return nil, err
	}

	return
}

func (ur *UserRepository) Create(u model.User) (*model.User, error) {
	user, err := ur.collection.InsertOne(ur.con, &u)

	if err != nil {
		return nil, err
	}

	u.ID = user.InsertedID.(primitive.ObjectID)

	return &u, nil
}

func (ur *UserRepository) Update(u *model.User) (*model.User, error) {
	if err := ur.collection.FindOneAndUpdate(
		context.Background(),
		bson.D{
			{"_id", u.ID},
		},
		bson.D{{"$set", u}},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&u); err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *UserRepository) ConvertObjectIDFromHex(ID string) (primitive.ObjectID, error) {
	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("invalid vote ID: %s", ID)
	}

	return oID, nil
}
