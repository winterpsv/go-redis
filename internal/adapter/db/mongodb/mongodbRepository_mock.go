package mongodbRepository

import (
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	model "task3_4/user-management/internal/entity"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindAll(page, pageSize int64, u []*model.User) ([]*model.User, error) {
	args := m.Called(page, pageSize, u)
	return args.Get(0).([]*model.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(ID string) (*model.User, error) {
	args := m.Called(ID)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) FindByNickname(nickname string) (*model.User, error) {
	args := m.Called(nickname)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) FindLasHourtUserVoteByVoteID(voterID string) (u *model.User, err error) {
	args := m.Called(voterID)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) Create(u model.User) (*model.User, error) {
	args := m.Called(u)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) Update(u *model.User) (*model.User, error) {
	args := m.Called(u)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) ConvertObjectIDFromHex(ID string) (primitive.ObjectID, error) {
	args := m.Called(ID)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}
