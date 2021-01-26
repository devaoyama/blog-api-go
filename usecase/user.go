package usecase

import (
	"blog-api/domain/entity"
	"blog-api/domain/repository"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return UserUseCase{
		userRepository: ur,
	}
}

func (uu *UserUseCase) SignUp(user *entity.User) (string, error) {
	user.Password = toHashPassword(user.PlainPassword)
	_, err := uu.userRepository.Create(user)
	if err != nil {
		return "", err
	}

	return generateJwtToken(user)
}

func (uu *UserUseCase) Login(email string, password string) (string, error) {
	user, err := uu.userRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if compareHashedPassword(user.Password, password) {
		return generateJwtToken(user)
	}

	return "", errors.New("credential is not correct")
}

func (uu *UserUseCase) GetAllUser() ([]entity.User, error) {
	return uu.userRepository.FindAll()
}

func (uu *UserUseCase) GetUser(id int) (*entity.User, error) {
	return uu.userRepository.Find(id)
}

func toHashPassword(pass string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hashed)
}

func compareHashedPassword(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err == nil {
		return true
	}
	return false
}

func generateJwtToken(user *entity.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["email"] = user.Email
	claims["role"] = "user"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte("secret"))
}
