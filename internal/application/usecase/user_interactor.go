package interactor

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	repository "task3_4/user-management/internal/adapter/db/mongodb"
	"task3_4/user-management/internal/application/service"
	"task3_4/user-management/internal/controller/http/dto"
	"task3_4/user-management/internal/controller/http/presenter"
	model "task3_4/user-management/internal/entity"
	"time"
)

type UserInteractor struct {
	UserMongoRepository repository.UserRepositoryInterface
	UserPresenter       presenter.UserPresenterInterface
	Auth                service.AuthInterface
	Cache               service.CacheInterface
}

func NewUserInteractor(r repository.UserRepositoryInterface, p presenter.UserPresenterInterface, a service.AuthInterface, c service.CacheInterface) *UserInteractor {
	return &UserInteractor{r, p, a, c}
}

func (us *UserInteractor) GetAll(page, pageSize int64, key string) ([]*dto.UserDTO, error) {
	var u []*model.User
	expiration := 1 * time.Minute

	u, err := us.UserMongoRepository.FindAll(page, pageSize, u)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	err = us.Cache.AddToCache(key, data, expiration)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}

func (us *UserInteractor) Get(id, key string) (*dto.UserDTO, error) {
	expiration := 1 * time.Minute

	u, err := us.UserMongoRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	err = us.Cache.AddToCache(key, data, expiration)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUser(u), nil
}

func (us *UserInteractor) UpdateVote(userForm *dto.VoteUserDTO, ID string, token *jwt.Token) (*dto.UserDTO, error) {
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	u, err := us.UserMongoRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	if userId == u.ID.Hex() {
		return nil, fmt.Errorf("can't vote for yourself")
	}

	if !u.Active {
		return nil, fmt.Errorf("user %s is deleted", u.Nickname)
	}

	for _, vote := range u.Votes {
		if vote.VoterID.Hex() == userId && vote.VoteValue == userForm.Value {
			return nil, fmt.Errorf("you have already voted for this user")
		}
	}

	lastUser, _ := us.UserMongoRepository.FindLasHourtUserVoteByVoteID(userId)
	if lastUser != nil {
		return nil, fmt.Errorf("can only vote once per hour")
	}

	VoterID, err := us.UserMongoRepository.ConvertObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	newVote := model.UserVote{
		VoterID:   VoterID,
		VoteValue: userForm.Value,
		VotedAt:   time.Now().Unix(),
	}

	u.Votes = append(u.Votes, newVote)

	u, err = us.UserMongoRepository.Update(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUser(u), nil
}

func (us *UserInteractor) Update(userForm *dto.UpdateUserDTO, ID string) (*dto.UserDTO, error) {
	u, err := us.UserMongoRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	if !u.Active {
		return nil, fmt.Errorf("user %s is deleted", u.Nickname)
	}

	u.FirstName = userForm.FirstName
	u.LastName = userForm.LastName
	u.Role = userForm.Role
	u.UpdatedAt = time.Now().Unix()

	u, err = us.UserMongoRepository.Update(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUser(u), nil
}

func (us *UserInteractor) Deactivate(ID string) (*dto.UserDTO, error) {
	u, err := us.UserMongoRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	if !u.Active {
		return nil, fmt.Errorf("user %s is deleted", u.Nickname)
	}

	u.DeletedAt = time.Now().Unix()
	u.Active = false

	u, err = us.UserMongoRepository.Update(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUser(u), nil
}

func (us *UserInteractor) GetUserByToken(token *jwt.Token, key string) (*dto.UserDTO, error) {
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	expiration := 1 * time.Minute

	u, err := us.UserMongoRepository.FindByID(userId)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	err = us.Cache.AddToCache(key, data, expiration)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUser(u), nil
}

func (us *UserInteractor) CacheGet(key string) ([]byte, error) {
	data, err := us.Cache.GetFromCache(key)
	if err != nil {
		return nil, err
	}
	return data, nil
}
