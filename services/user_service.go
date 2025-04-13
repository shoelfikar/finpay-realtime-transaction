package services

import (
	"github.com/shoelfikar/finpay-realtime-transaction/model"
	"github.com/shoelfikar/finpay-realtime-transaction/repository"
	"github.com/shoelfikar/finpay-realtime-transaction/utils"
)

type userService struct {
   UserRepository repository.UserRepository
}

type UserService interface {
   CreateUser(user *model.User) *model.User
   GetUserByEmail(email string) *model.User
   GetUserDetail(userId string) *model.User
}

func NewUserService(user repository.UserRepository) UserService {
   return &userService{
      UserRepository: user,
   }
}

func (u *userService) CreateUser(user *model.User) *model.User {
   passwordHash, err := utils.HashPassword(*user.Password)
   if err != nil {
      panic("Error encrypt password"+ err.Error())
   }

   user.Password = &passwordHash

   result := u.UserRepository.CreateUser(user)
   return result
}

func (u *userService) GetUserByEmail(email string) *model.User {
   user := u.UserRepository.GetUserByEmail(email)
   return user
}

func (u *userService) GetUserDetail(userId string) *model.User {
   user := u.UserRepository.GetUserByID(userId)
   return user
}