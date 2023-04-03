package interactor

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	repository "task3_4/user-management/internal/adapter/db/mongodb"
	"task3_4/user-management/internal/application/service"
	"task3_4/user-management/internal/controller/http/dto"
	"task3_4/user-management/internal/controller/http/presenter"
	model "task3_4/user-management/internal/entity"
	"testing"
	"time"
)

func getObjectIDFromString(id string) primitive.ObjectID {
	objID, _ := primitive.ObjectIDFromHex(id)
	return objID
}

func GenerateHash(password string) string {
	h := sha256.New()

	h.Write([]byte(password))

	return hex.EncodeToString(h.Sum(nil))
}

func TestUserInteractor_GetAll(t *testing.T) {
	var (
		// Создание mock объекта UserRepository
		mockUserRepository = new(repository.MockUserRepository)
		mockUserPresenter  = new(presenter.MockUserPresenter)
		mockAuth           = new(service.MockAuth)

		// Создание объекта interactor с mock mockUserRepository, mockUserPresenter, mockAuth
		interactor = NewUserInteractor(mockUserRepository, mockUserPresenter, mockAuth)
	)
	expectedUsers := []*model.User{
		&model.User{
			ID:           getObjectIDFromString("1"),
			Nickname:     "user1",
			FirstName:    "user1",
			LastName:     "user1",
			PasswordHash: GenerateHash("Password1"),
			CreatedAt:    time.Now().Unix(),
			UpdatedAt:    0,
			DeletedAt:    0,
			Active:       true,
			Votes:        []model.UserVote{},
		},
		&model.User{
			ID:           getObjectIDFromString("2"),
			Nickname:     "user2",
			FirstName:    "user2",
			LastName:     "user2",
			PasswordHash: GenerateHash("Password2"),
			CreatedAt:    time.Now().Unix(),
			UpdatedAt:    0,
			DeletedAt:    0,
			Active:       true,
			Votes:        []model.UserVote{},
		},
	}

	dtodUsers := []*dto.UserDTO{
		&dto.UserDTO{
			ID:           getObjectIDFromString("1"),
			Nickname:     "user1",
			FirstName:    "user1",
			LastName:     "user1",
			PasswordHash: GenerateHash("Password1"),
			CreatedAt:    time.Now().Unix(),
			UpdatedAt:    0,
			DeletedAt:    0,
			Active:       true,
			Votes:        0,
		},
		&dto.UserDTO{
			ID:           getObjectIDFromString("2"),
			Nickname:     "user2",
			FirstName:    "user2",
			LastName:     "user2",
			PasswordHash: GenerateHash("Password2"),
			CreatedAt:    time.Now().Unix(),
			UpdatedAt:    0,
			DeletedAt:    0,
			Active:       true,
			Votes:        0,
		},
	}

	mockUserRepository.On("FindAll", int64(1), int64(10), []*model.User(nil)).Return(expectedUsers, nil)
	mockUserPresenter.On("ResponseUsers", expectedUsers).Return(dtodUsers)

	// Act
	users, err := interactor.GetAll(1, 10)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, dtodUsers, users)

	mockUserRepository.AssertExpectations(t)
	mockUserPresenter.AssertExpectations(t)
}

func TestUserInteractor_Get(t *testing.T) {
	var (
		// Создание mock объекта UserRepository
		mockUserRepository = new(repository.MockUserRepository)
		mockUserPresenter  = new(presenter.MockUserPresenter)
		mockAuth           = new(service.MockAuth)

		// Создание объекта interactor с mock mockUserRepository, mockUserPresenter, mockAuth
		interactor = NewUserInteractor(mockUserRepository, mockUserPresenter, mockAuth)
	)
	mockUser := &model.User{
		ID:           getObjectIDFromString("abc123"),
		Nickname:     "user1",
		FirstName:    "user1",
		LastName:     "user1",
		PasswordHash: GenerateHash("Password1"),
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    0,
		DeletedAt:    0,
		Active:       true,
		Votes:        []model.UserVote{},
	}

	dtoUser := &dto.UserDTO{
		ID:           getObjectIDFromString("abc123"),
		Nickname:     "user1",
		FirstName:    "user1",
		LastName:     "user1",
		PasswordHash: GenerateHash("Password1"),
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    0,
		DeletedAt:    0,
		Active:       true,
		Votes:        0,
	}

	mockUserPresenter.On("ResponseUser", mockUser).Return(dtoUser)

	mockUserRepository.On("FindByID", "abc123").Return(mockUser, nil)

	user, err := interactor.Get("abc123")
	assert.NoError(t, err)
	assert.Equal(t, getObjectIDFromString("abc123"), user.ID)
	assert.Equal(t, "user1", user.FirstName)
	assert.Equal(t, "user1", user.LastName)
	assert.Equal(t, GenerateHash("Password1"), user.PasswordHash)

	mockUserRepository.AssertExpectations(t)
}

func TestUserInteractor_Get_Error(t *testing.T) {
	var (
		// Создание mock объекта UserRepository
		mockUserRepository = new(repository.MockUserRepository)
		mockUserPresenter  = new(presenter.MockUserPresenter)
		mockAuth           = new(service.MockAuth)

		// Создание объекта interactor с mock mockUserRepository, mockUserPresenter, mockAuth
		interactor = NewUserInteractor(mockUserRepository, mockUserPresenter, mockAuth)
	)
	mockUser := &model.User{
		ID:           getObjectIDFromString("abc123"),
		Nickname:     "user1",
		FirstName:    "user1",
		LastName:     "user1",
		PasswordHash: GenerateHash("Password1"),
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    0,
		DeletedAt:    0,
		Active:       true,
	}

	mockUserRepository.On("FindByID", "abc123").Return(mockUser, fmt.Errorf("could not found user with ID abc123"))

	_, err := interactor.Get("abc123")
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("could not found user with ID abc123"), err)
	mockUserRepository.AssertExpectations(t)
}

func TestUserInteractor_Create(t *testing.T) {
	var (
		// Создание mock объекта UserRepository
		mockUserRepository = new(repository.MockUserRepository)
		mockUserPresenter  = new(presenter.MockUserPresenter)
		mockAuth           = new(service.MockAuth)

		// Создание объекта interactor с mock mockUserRepository, mockUserPresenter, mockAuth
		interactor = NewAuthInteractor(mockUserRepository, mockUserPresenter, mockAuth)
	)
	mockUser := dto.CreateUserDTO{
		Nickname:  "testuser",
		FirstName: "testuser",
		LastName:  "testuser",
		Password:  "password",
	}

	mockAuth.On("GenerateHash", mockUser.Password).Return(GenerateHash(mockUser.Password))

	PasswordHash := mockAuth.GenerateHash(mockUser.Password)
	timestamp := time.Now().Unix()

	uModel := model.User{
		Nickname:     mockUser.Nickname,
		FirstName:    mockUser.FirstName,
		LastName:     mockUser.LastName,
		PasswordHash: PasswordHash,
		CreatedAt:    timestamp,
		Active:       true,
		Votes:        []model.UserVote{},
	}

	mUser := &model.User{
		ID:           getObjectIDFromString("1"),
		Nickname:     mockUser.Nickname,
		FirstName:    mockUser.FirstName,
		LastName:     mockUser.LastName,
		PasswordHash: PasswordHash,
		CreatedAt:    timestamp,
		UpdatedAt:    0,
		DeletedAt:    0,
		Active:       true,
		Votes:        []model.UserVote{},
	}

	dtoUser := &dto.UserDTO{
		ID:           getObjectIDFromString("1"),
		Nickname:     mockUser.Nickname,
		FirstName:    mockUser.FirstName,
		LastName:     mockUser.LastName,
		PasswordHash: PasswordHash,
		CreatedAt:    timestamp,
		UpdatedAt:    0,
		DeletedAt:    0,
		Active:       true,
		Votes:        0,
	}

	mockUserRepository.On("FindByNickname", "testuser").Return(&model.User{}, nil)

	mockUserRepository.On("Create", uModel).Return(mUser, nil)

	mockUserPresenter.On("ResponseUser", mUser).Return(dtoUser)

	responseUser, err := interactor.Create(&mockUser)

	assert.Nil(t, err)
	assert.Equal(t, dtoUser, responseUser)
}

func TestUserInteractor_Update(t *testing.T) {
	var (
		// Создание mock объекта UserRepository
		mockUserRepository = new(repository.MockUserRepository)
		mockUserPresenter  = new(presenter.MockUserPresenter)
		mockAuth           = new(service.MockAuth)

		// Создание объекта interactor с mock mockUserRepository, mockUserPresenter, mockAuth
		interactor = NewUserInteractor(mockUserRepository, mockUserPresenter, mockAuth)
	)
	mockUser := &dto.UpdateUserDTO{
		FirstName: "testuser",
		LastName:  "testuser",
	}

	ID := getObjectIDFromString("user123")
	PasswordHash := GenerateHash("Password")
	timestamp := time.Now().Unix()

	uModel := &model.User{
		Nickname:     "User",
		FirstName:    mockUser.FirstName,
		LastName:     mockUser.LastName,
		PasswordHash: PasswordHash,
		CreatedAt:    timestamp,
		UpdatedAt:    0,
		DeletedAt:    0,
		Active:       true,
		Votes:        []model.UserVote{},
	}

	mUser := &model.User{
		ID:           ID,
		Nickname:     "User",
		FirstName:    mockUser.FirstName,
		LastName:     mockUser.LastName,
		PasswordHash: PasswordHash,
		CreatedAt:    timestamp,
		UpdatedAt:    timestamp,
		DeletedAt:    0,
		Active:       true,
		Votes:        []model.UserVote{},
	}

	dtoUser := &dto.UserDTO{
		ID:           ID,
		Nickname:     "User",
		FirstName:    mockUser.FirstName,
		LastName:     mockUser.LastName,
		PasswordHash: PasswordHash,
		CreatedAt:    timestamp,
		UpdatedAt:    timestamp,
		DeletedAt:    0,
		Active:       true,
		Votes:        0,
	}

	mockUserRepository.On("FindByID", "user123").Return(uModel, nil)
	mockUserRepository.On("Update", uModel).Return(mUser, nil)
	mockUserPresenter.On("ResponseUser", uModel).Return(dtoUser)

	u, err := interactor.Update(mockUser, "user123")
	assert.Nil(t, err)
	assert.Equal(t, dtoUser, u)

	mockUserRepository.AssertExpectations(t)
}

func TestUserInteractor_UpdatePassword(t *testing.T) {
	var (
		// Создание mock объекта UserRepository
		mockUserRepository = new(repository.MockUserRepository)
		mockUserPresenter  = new(presenter.MockUserPresenter)
		mockAuth           = new(service.MockAuth)

		// Создание объекта interactor с mock mockUserRepository, mockUserPresenter, mockAuth
		interactor = NewAuthInteractor(mockUserRepository, mockUserPresenter, mockAuth)
	)
	mockUser := &dto.UpdateUserPasswordDTO{
		Password: "newPassword",
	}

	ID := getObjectIDFromString("user123")
	mockAuth.On("GenerateHash", mockUser.Password).Return(GenerateHash(mockUser.Password))
	PasswordHash := mockAuth.GenerateHash(mockUser.Password)
	timestamp := time.Now().Unix()

	uModel := &model.User{
		Nickname:     "User",
		FirstName:    "User",
		LastName:     "User",
		PasswordHash: PasswordHash,
		CreatedAt:    timestamp,
		UpdatedAt:    0,
		DeletedAt:    0,
		Active:       true,
		Votes:        []model.UserVote{},
	}

	mUser := &model.User{
		ID:           ID,
		Nickname:     "User",
		FirstName:    "User",
		LastName:     "User",
		PasswordHash: PasswordHash,
		CreatedAt:    timestamp,
		UpdatedAt:    timestamp,
		DeletedAt:    0,
		Active:       true,
		Votes:        []model.UserVote{},
	}

	dtoUser := &dto.UserDTO{
		ID:           ID,
		Nickname:     "User",
		FirstName:    "User",
		LastName:     "User",
		PasswordHash: PasswordHash,
		CreatedAt:    timestamp,
		UpdatedAt:    timestamp,
		DeletedAt:    0,
		Active:       true,
		Votes:        0,
	}

	mockUserRepository.On("FindByID", "user123").Return(uModel, nil)
	mockUserRepository.On("Update", uModel).Return(mUser, nil)
	mockUserPresenter.On("ResponseUser", mUser).Return(dtoUser)

	u, err := interactor.UpdatePassword(mockUser, "user123")
	assert.Nil(t, err)
	assert.Equal(t, dtoUser, u)

	mockUserRepository.AssertExpectations(t)
}

func TestUserInteractor_Deactivate(t *testing.T) {
	var (
		// Создание mock объекта UserRepository
		mockUserRepository = new(repository.MockUserRepository)
		mockUserPresenter  = new(presenter.MockUserPresenter)
		mockAuth           = new(service.MockAuth)

		// Создание объекта interactor с mock mockUserRepository, mockUserPresenter, mockAuth
		interactor = NewUserInteractor(mockUserRepository, mockUserPresenter, mockAuth)
	)
	ID := getObjectIDFromString("user123")
	PasswordHash := GenerateHash("newPassword")
	timestamp := time.Now().Unix()

	uModel := &model.User{
		Nickname:     "User",
		FirstName:    "testuser",
		LastName:     "testuser",
		PasswordHash: PasswordHash,
		CreatedAt:    timestamp,
		UpdatedAt:    timestamp,
		DeletedAt:    0,
		Active:       true,
		Votes:        []model.UserVote{},
	}

	mUser := &model.User{
		ID:           ID,
		Nickname:     "User",
		FirstName:    "testuser",
		LastName:     "testuser",
		PasswordHash: PasswordHash,
		CreatedAt:    timestamp,
		UpdatedAt:    timestamp,
		DeletedAt:    timestamp,
		Active:       false,
		Votes:        []model.UserVote{},
	}

	dtoUser := &dto.UserDTO{
		ID:           ID,
		Nickname:     "User",
		FirstName:    "testuser",
		LastName:     "testuser",
		PasswordHash: PasswordHash,
		CreatedAt:    timestamp,
		UpdatedAt:    timestamp,
		DeletedAt:    timestamp,
		Active:       false,
		Votes:        0,
	}

	mockUserRepository.On("FindByID", "user123").Return(uModel, nil)
	mockUserRepository.On("Update", uModel).Return(mUser, nil)
	mockUserPresenter.On("ResponseUser", mUser).Return(dtoUser)

	u, err := interactor.Deactivate("user123")
	assert.Nil(t, err)
	assert.Equal(t, dtoUser, u)

	mockUserRepository.AssertExpectations(t)
}

func TestUserInteractor_Authenticate(t *testing.T) {
	var (
		// Создание mock объекта UserRepository
		mockUserRepository = new(repository.MockUserRepository)
		mockUserPresenter  = new(presenter.MockUserPresenter)
		mockAuth           = new(service.MockAuth)

		// Создание объекта interactor с mock mockUserRepository, mockUserPresenter, mockAuth
		interactor = NewAuthInteractor(mockUserRepository, mockUserPresenter, mockAuth)
	)

	// Створення тестового запиту
	req, err := http.NewRequest("GET", "/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	jwtToken := &jwt.Token{}

	mockAuth.On("Authenticate", req).Return(jwtToken, nil)

	// Виклик методу Authenticate()
	token, err := interactor.Authenticate(req)
	if err != nil {
		t.Fatal(err)
	}

	// Перевірка результату
	if token == nil {
		t.Errorf("Expected non-nil token, got nil")
	}
}
