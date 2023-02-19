package user

import (
	"testing"

	"github.com/Takao-Yamasaki/golang-unit-test/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryStub struct {
	// mockツールを埋め込む
	mock.Mock
}

func (r *UserRepositoryStub) Add(user entity.User) error {
	args := r.Called(user)

	return args.Error(0)
}

type BadWordsRepositoryStub struct {
	mock.Mock
}

func (r *BadWordsRepositoryStub) FindAll() ([]string, error) {
	args := r.Called()

	return args.Get(0).([]string), args.Error(1)
}

func TestShouldSuccessfullyRegistrateAnUser(t *testing.T) {

	user := entity.User{
		Name:        "Vinicius",
		Email:       "vinicius@example.com",
		Description: "Software Developer",
	}

	userRepository := &UserRepositoryStub{}
	userRepository.On("Add", user).Return(nil)
	badWordsRepository := &BadWordsRepositoryStub{}
	badWordsRepository.On("FindAll").Return([]string{"tomato", "potato"}, nil)
	// ここで実際の依存関係をインスタンス化できない
	userService := NewUserService(userRepository, badWordsRepository)
	err := userService.Register(user)

	userRepository.AssertCalled(t, "Add", user)
	assert.Nil(t, err)
}

func TestShouldNotRegistrateTheUserWhenBadWordIsFound(t *testing.T) {

	user := entity.User{
		Name:        "Vinicius",
		Email:       "vinicius@example.com",
		Description: "Software potato Developer",
	}

	userRepository := &UserRepositoryStub{}
	userRepository.On("Add", user).Return(nil)
	badWordsRepository := &BadWordsRepositoryStub{}
	badWordsRepository.On("FindAll").Return([]string{"tomato", "potato"}, nil)
	// ここで実際の依存関係をインスタンス化できない
	userService := NewUserService(userRepository, badWordsRepository)
	err := userService.Register(user)

	userRepository.AssertNotCalled(t, "Add", user)
	assert.Error(t, err)
}
