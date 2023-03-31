package mongodbRepository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	model "task3_3_new/user-management/internal/entity"
)

type UserRepositoryInterface interface {
	FindAll(page, pageSize int64, u []*model.User) ([]*model.User, error)
	FindByID(ID string) (*model.User, error)
	FindByNickname(nickname string) (*model.User, error)
	FindLasHourtUserVoteByVoteID(voterID string) (u *model.User, err error)
	Create(u model.User) (*model.User, error)
	Update(u *model.User) (*model.User, error)
	ConvertObjectIDFromHex(ID string) (primitive.ObjectID, error)
}
