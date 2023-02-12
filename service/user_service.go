package service

import (
	"log"

	"github.com/LebrancWorkshop/Go-Hexagonal-Architecture-Codebangkok/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) userService {
	return userService{userRepository: userRepository};
}

func (s userService) GetUsers() ([]UserResponse, error) {
	users, err := s.userRepository.GetUsers();
	if err != nil {
		log.Println(err);
		return nil, err;
	}
	userResponses := []UserResponse{};
	for _, user := range users {
		userResponse := UserResponse{
			ID: user.ID,
			Username: user.Username,
			Displayname: user.Displayname,
		}
		userResponses = append(userResponses, userResponse);
	}
	return userResponses, nil;
}

func (s userService) GetUserByID(id int) (*UserResponse, error) {
	user, err := s.userRepository.GetUserByID(id);
	if err != nil {
		log.Println(err);
		return nil, err;
	}

	userResponse := UserResponse{
		ID: user.ID,
		Username: user.Username,
		Displayname: user.Displayname,
	}

	return &userResponse, nil; 
}