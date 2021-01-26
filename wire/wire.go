package wire

import (
	"blog-api/handler"
	"blog-api/infrastructure/persistence"
	"blog-api/usecase"
	"gorm.io/gorm"
)

func InitUserAPI(db *gorm.DB) handler.UserHandler {
	userRepository := persistence.NewUserPersistence(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	return handler.NewUserHandler(userUseCase)
}
