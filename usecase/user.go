package usecase

import (
	"blog-api/domain/entity"
	"blog-api/domain/repository"
)

type UserUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return UserUseCase{
		userRepository: ur,
	}
}

func (uu *UserUseCase) GetAllUser() ([]entity.User, error) {
	return uu.userRepository.FindAll()
}

func (uu *UserUseCase) GetUser(id int) (*entity.User, error) {
	return uu.userRepository.Find(id)
}
