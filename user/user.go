package user

import (
	"errors"
	"strings"

	"github.com/Takao-Yamasaki/golang-unit-test/database"
	"github.com/Takao-Yamasaki/golang-unit-test/entity"
)

type UserSercvice struct {
	userRepository     database.User
	badWordsRepository database.BadWords
}

func NewUserService(userRepository database.User, badWordsRepository database.BadWords) *UserSercvice {
	return &UserSercvice{
		userRepository:     userRepository,
		badWordsRepository: badWordsRepository,
	}
}

func (c *UserSercvice) Register(user entity.User) error {
	badWords, err := c.badWordsRepository.FindAll()
	if err != nil {
		return err
	}

	for _, badWord := range badWords {
		if strings.Contains(badWord, user.Description) {
			return errors.New("bad word found")
		}
	}

	err = c.userRepository.Add(user)
	if err != nil {
		return err
	}
	return nil
}
