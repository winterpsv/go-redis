package presenter

import (
	"github.com/stretchr/testify/mock"
	"task3_4/user-management/internal/controller/http/dto"
	model "task3_4/user-management/internal/entity"
)

type MockUserPresenter struct {
	mock.Mock
}

func (m *MockUserPresenter) ResponseUsers(us []*model.User) []*dto.UserDTO {
	args := m.Called(us)
	return args.Get(0).([]*dto.UserDTO)
}

func (m *MockUserPresenter) ResponseUser(us *model.User) *dto.UserDTO {
	args := m.Called(us)
	return args.Get(0).(*dto.UserDTO)
}

func (m *MockUserPresenter) ResponseToken(token string) string {
	args := m.Called(token)
	return args.String(0)
}

func (m *MockUserPresenter) ResponseError(err error) error {
	args := m.Called(err)
	return args.Error(0)
}

func (m *MockUserPresenter) GetVotesSum(us *model.User) int {
	args := m.Called(us)
	return args.Int(0)
}
