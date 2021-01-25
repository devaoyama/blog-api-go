package persistence

import (
	"blog-api/domain"
	"blog-api/domain/repository"
	"gorm.io/gorm"
)

type userPersistence struct {
	Db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{Db: db}
}

func (up *userPersistence) FindAll() ([]domain.User, error) {
	var users []domain.User
	result := up.Db.Find(&users)
	if err := result.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (up *userPersistence) Create(user *domain.User) (*domain.User, error) {
	result := up.Db.Create(user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (up *userPersistence) Find(id int) (*domain.User, error) {
	user := &domain.User{}
	result := up.Db.First(user, id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (up *userPersistence) Update(user *domain.User) (*domain.User, error) {
	result := up.Db.Save(user)
	if err := result.Error; err != nil {
		return nil, result.Error
	}
	return user, nil
}

func (up *userPersistence) Delete(id int) error {
	result := up.Db.Delete(&domain.User{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
