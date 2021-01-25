package persistence

import (
	"blog-api/domain/entity"
	"blog-api/domain/repository"
	"gorm.io/gorm"
)

type userPersistence struct {
	Db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{Db: db}
}

func (up *userPersistence) FindAll() ([]entity.User, error) {
	var users []entity.User
	result := up.Db.Find(&users)
	if err := result.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (up *userPersistence) Create(user *entity.User) (*entity.User, error) {
	result := up.Db.Create(user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (up *userPersistence) Find(id int) (*entity.User, error) {
	user := &entity.User{}
	result := up.Db.First(user, id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (up *userPersistence) Update(user *entity.User) (*entity.User, error) {
	result := up.Db.Save(user)
	if err := result.Error; err != nil {
		return nil, result.Error
	}
	return user, nil
}

func (up *userPersistence) Delete(id int) error {
	result := up.Db.Delete(&entity.User{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
