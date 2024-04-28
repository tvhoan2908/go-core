package services

import (
	"errors"
	Config "go-core/config"
	"go-core/databases/models"
	"go-core/types"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ELearningClaims struct {
	UserID uint `json:"userId"`
	jwt.StandardClaims
}

type UserInterface struct {
	ID          uint
	AccountType int
	Permissions []*string
}

type AuthService interface {
	Register(request types.RegisterRequest) (*models.User, error)
	Login(request types.LoginRequest) (*models.User, error)
}

type authService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{db: db}
}

func InitUser(user *types.UserInfoDTO) *UserInterface {
	return &UserInterface{
		ID:          user.ID,
		AccountType: user.AccountType,
		Permissions: user.Permissisons,
	}
}

func (s *authService) Register(request types.RegisterRequest) (*models.User, error) {
	user := models.User{
		Username: request.Username,
		Password: request.Password,
		FullName: request.FullName,
		Email:    request.Email,
	}

	result := s.db.Create(&user)

	return &user, result.Error
}

func (s *authService) Login(request types.LoginRequest) (*models.User, error) {
	var user models.User
	result := s.db.Where(&models.User{Username: request.Username, Status: Config.USER_VISIBLE}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	checkPass := CheckPasswordHash(request.Password, user.Password)
	if !checkPass {
		return nil, errors.New("password invalid")
	}

	return &user, result.Error
}

func CreateToken(user models.User) (string, error) {
	secretKey := []byte(Config.SECRET_KEY)
	// Create the Claims
	claims := ELearningClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * Config.TOKEN_EXPIRED_TIME).Unix(),
			Issuer:    user.Username,
			IssuedAt:  time.Now().Unix(),
		},
	}

	tokenAt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAt.SignedString(secretKey)

	return token, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (user *UserInterface) HasPermission(permissionName string) bool {
	if user.AccountType == Config.SUPER_ADMIN {
		return true
	}

	permissions := user.Permissions
	hasPermission := false
	for _, permission := range permissions {
		if *permission == permissionName {
			hasPermission = true
			break
		}
	}

	return hasPermission
}
