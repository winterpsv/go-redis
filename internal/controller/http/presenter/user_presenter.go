package presenter

import (
	"task3_4/user-management/internal/controller/http/dto"
	model "task3_4/user-management/internal/entity"
)

type UserPresenter struct {
}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

func (up *UserPresenter) ResponseUsers(us []*model.User) []*dto.UserDTO {
	var pUsers []*dto.UserDTO

	for i := range us {
		newUser := dto.UserDTO{
			ID:           us[i].ID,
			Nickname:     us[i].Nickname,
			FirstName:    us[i].FirstName,
			LastName:     us[i].LastName,
			PasswordHash: "********",
			CreatedAt:    us[i].CreatedAt,
			UpdatedAt:    us[i].UpdatedAt,
			DeletedAt:    us[i].DeletedAt,
			Role:         us[i].Role,
			Active:       us[i].Active,
			Votes:        up.GetVotesSum(us[i]),
		}

		pUsers = append(pUsers, &newUser)
	}

	return pUsers
}

func (up *UserPresenter) ResponseUser(us *model.User) *dto.UserDTO {
	newUser := dto.UserDTO{
		ID:           us.ID,
		Nickname:     us.Nickname,
		FirstName:    us.FirstName,
		LastName:     us.LastName,
		PasswordHash: "********",
		CreatedAt:    us.CreatedAt,
		UpdatedAt:    us.UpdatedAt,
		DeletedAt:    us.DeletedAt,
		Role:         us.Role,
		Active:       us.Active,
		Votes:        up.GetVotesSum(us),
	}

	return &newUser
}

func (up *UserPresenter) ResponseToken(token string) string {
	return token
}

func (up *UserPresenter) ResponseError(err error) error {
	return err
}

func (up *UserPresenter) GetVotesSum(us *model.User) int {
	var sum = 0
	for _, vote := range us.Votes {
		sum += vote.VoteValue
	}
	return sum
}
