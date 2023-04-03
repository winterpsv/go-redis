package mongodbRepository

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	model "task3_4/user-management/internal/entity"
	"testing"
)

var (
	cases = []*model.User{
		{
			ID:           getObjectIDFromString("6403d2a6caeea3db93cbfe86"),
			Nickname:     "test",
			FirstName:    "test first_name",
			LastName:     "test last_name",
			PasswordHash: "Qwertyuiop123qaz",
			CreatedAt:    1677972134,
			UpdatedAt:    0,
			DeletedAt:    0,
			Active:       true,
		},
		{
			ID:           getObjectIDFromString("6403c7ef1b636731aac1b080"),
			Nickname:     "sfgsdfg",
			FirstName:    "sfgsdfg first_name",
			LastName:     "sfgsdfg last_name",
			PasswordHash: "Asdfghjkl123azq",
			CreatedAt:    1677973456,
			UpdatedAt:    0,
			DeletedAt:    0,
			Active:       true,
		},
		{
			ID:           getObjectIDFromString("6403dd6760c678a8c0a17fb6"),
			Nickname:     "tghdg",
			FirstName:    "tghdg first_name",
			LastName:     "tghdg last_name",
			PasswordHash: "Zxcvbnm123",
			CreatedAt:    1677973456,
			UpdatedAt:    0,
			DeletedAt:    0,
			Active:       true,
		},
		{
			ID:           getObjectIDFromString("6403c7ef1b636731aac1b080"),
			Nickname:     "Kdjvd",
			FirstName:    "Kdjvd first_name",
			LastName:     "Kdjvd last_name",
			PasswordHash: "Qazxswedcvfr123",
			CreatedAt:    1677973456,
			UpdatedAt:    0,
			DeletedAt:    0,
			Active:       true,
		},
		{
			ID:           getObjectIDFromString("6403c7ef1b636731aac1b080"),
			Nickname:     "Jhdgvsd",
			FirstName:    "Jhdgvsd first_name",
			LastName:     "Jhdgvsd last_name",
			PasswordHash: "Zaqwsxcder123",
			CreatedAt:    1677973456,
			UpdatedAt:    0,
			DeletedAt:    0,
			Active:       true,
		},
	}
)

func getObjectIDFromString(id string) primitive.ObjectID {
	objID, _ := primitive.ObjectIDFromHex(id)
	return objID
}

func TestUserRepository_FindAll(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	for _, el := range cases {
		var userModels []*model.User
		mt.Run("success", func(mt *mtest.T) {
			db := &UserRepository{collection: mt.Coll}
			mocUser := model.User{
				ID:           el.ID,
				Nickname:     el.Nickname,
				FirstName:    el.FirstName,
				LastName:     el.LastName,
				PasswordHash: el.PasswordHash,
				CreatedAt:    el.CreatedAt,
				UpdatedAt:    el.UpdatedAt,
				DeletedAt:    el.DeletedAt,
				Active:       el.Active,
			}
			u := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
				{"_id", mocUser.ID},
				{"nickname", mocUser.Nickname},
				{"first_name", mocUser.FirstName},
				{"last_name", mocUser.LastName},
				{"password", mocUser.PasswordHash},
				{"created_at", mocUser.CreatedAt},
				{"updated_at", mocUser.UpdatedAt},
				{"deleted_at", mocUser.DeletedAt},
				{"active", mocUser.Active},
			})
			killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
			mt.AddMockResponses(u, killCursors)
			users, err := db.FindAll(1, 10, userModels)

			assert.Nil(t, err)
			assert.Equal(t, []*model.User{
				&mocUser,
			}, users)
		})

		mt.Run("error", func(mt *mtest.T) {
			db := &UserRepository{collection: mt.Coll}
			mt.AddMockResponses(bson.D{{"ok", 0}})
			_, err := db.FindAll(1, 10, userModels)

			assert.NotNil(t, err)
		})
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	for _, el := range cases {
		mt.Run("success", func(mt *mtest.T) {
			db := &UserRepository{collection: mt.Coll}
			mocUser := &model.User{
				ID:           el.ID,
				Nickname:     el.Nickname,
				FirstName:    el.FirstName,
				LastName:     el.LastName,
				PasswordHash: el.PasswordHash,
				CreatedAt:    el.CreatedAt,
				UpdatedAt:    el.UpdatedAt,
				DeletedAt:    el.DeletedAt,
				Active:       el.Active,
			}

			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
				{"_id", mocUser.ID},
				{"nickname", mocUser.Nickname},
				{"first_name", mocUser.FirstName},
				{"last_name", mocUser.LastName},
				{"password", mocUser.PasswordHash},
				{"created_at", mocUser.CreatedAt},
				{"updated_at", mocUser.UpdatedAt},
				{"deleted_at", mocUser.DeletedAt},
				{"active", mocUser.Active},
			}))

			responseUser, err := db.FindByID(el.ID.Hex())

			assert.Nil(t, err)
			assert.Equal(t, mocUser, responseUser)
		})

		mt.Run("error", func(mt *mtest.T) {
			db := &UserRepository{collection: mt.Coll}
			mt.AddMockResponses(bson.D{{"ok", 0}})

			_, err := db.FindByID(el.ID.String())

			assert.NotNil(t, err)
		})
	}
}

func TestUserRepository_FindByNickname(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	for _, el := range cases {
		mt.Run("success", func(mt *mtest.T) {
			db := &UserRepository{collection: mt.Coll}
			mocUser := &model.User{
				ID:           el.ID,
				Nickname:     el.Nickname,
				FirstName:    el.FirstName,
				LastName:     el.LastName,
				PasswordHash: el.PasswordHash,
				CreatedAt:    el.CreatedAt,
				UpdatedAt:    el.UpdatedAt,
				DeletedAt:    el.DeletedAt,
				Active:       el.Active,
			}

			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
				{"_id", mocUser.ID},
				{"nickname", mocUser.Nickname},
				{"first_name", mocUser.FirstName},
				{"last_name", mocUser.LastName},
				{"password", mocUser.PasswordHash},
				{"created_at", mocUser.CreatedAt},
				{"updated_at", mocUser.UpdatedAt},
				{"deleted_at", mocUser.DeletedAt},
				{"active", mocUser.Active},
			}))

			responseUser, err := db.FindByNickname(el.Nickname)

			assert.Nil(t, err)
			assert.Equal(t, mocUser, responseUser)
		})

		mt.Run("error", func(mt *mtest.T) {
			db := &UserRepository{collection: mt.Coll}
			mt.AddMockResponses(bson.D{{"ok", 0}})

			_, err := db.FindByNickname(el.Nickname)

			assert.NotNil(t, err)
		})
	}
}

func TestUserRepository_Create(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	for _, el := range cases {
		mocUser := model.User{
			ID:           el.ID,
			Nickname:     el.Nickname,
			FirstName:    el.FirstName,
			LastName:     el.LastName,
			PasswordHash: el.PasswordHash,
			CreatedAt:    el.CreatedAt,
			UpdatedAt:    el.UpdatedAt,
			DeletedAt:    el.DeletedAt,
			Active:       el.Active,
		}

		mt.Run("success", func(mt *mtest.T) {
			db := &UserRepository{collection: mt.Coll}
			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
				{"_id", mocUser.ID},
				{"nickname", mocUser.Nickname},
				{"first_name", mocUser.FirstName},
				{"last_name", mocUser.LastName},
				{"password", mocUser.PasswordHash},
				{"created_at", mocUser.CreatedAt},
				{"updated_at", mocUser.UpdatedAt},
				{"deleted_at", mocUser.DeletedAt},
				{"active", mocUser.Active},
			}))

			responseUser, err := db.Create(mocUser)

			assert.Nil(t, err)
			assert.Equal(t, &mocUser, responseUser)
		})

		mt.Run("error", func(mt *mtest.T) {
			db := &UserRepository{collection: mt.Coll}
			mt.AddMockResponses(bson.D{{"ok", 0}})

			_, err := db.Create(mocUser)

			assert.NotNil(t, err)
		})
	}
}

func TestUserRepository_Update(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	for _, el := range cases {
		mocUser := &model.User{
			ID:           el.ID,
			Nickname:     el.Nickname,
			FirstName:    el.FirstName,
			LastName:     el.LastName,
			PasswordHash: el.PasswordHash,
			CreatedAt:    el.CreatedAt,
			UpdatedAt:    el.UpdatedAt,
			DeletedAt:    el.DeletedAt,
			Active:       el.Active,
		}

		mt.Run("success", func(mt *mtest.T) {
			db := &UserRepository{collection: mt.Coll}
			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"value", bson.D{
					{"_id", mocUser.ID},
					{"nickname", mocUser.Nickname},
					{"first_name", mocUser.FirstName},
					{"last_name", mocUser.LastName},
					{"password", mocUser.PasswordHash},
					{"created_at", mocUser.CreatedAt},
					{"updated_at", mocUser.UpdatedAt},
					{"deleted_at", mocUser.DeletedAt},
					{"active", mocUser.Active},
				}},
			})

			updatedUser, err := db.Update(mocUser)

			assert.Nil(t, err)
			assert.Equal(t, mocUser, updatedUser)
		})

		mt.Run("error", func(mt *mtest.T) {
			db := &UserRepository{collection: mt.Coll}
			mt.AddMockResponses(bson.D{{"ok", 0}})
			_, err := db.Update(mocUser)

			assert.NotNil(t, err)
		})
	}
}
