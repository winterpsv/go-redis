package presenter

import (
	"task3_4/user-management/internal/controller/http/dto"
	model "task3_4/user-management/internal/entity"
)

type UserPresenterInterface interface {
	ResponseUsers(u []*model.User) []*dto.UserDTO
	ResponseUser(u *model.User) *dto.UserDTO
	ResponseToken(string) string
	ResponseError(error) error
	GetVotesSum(us *model.User) int
}
