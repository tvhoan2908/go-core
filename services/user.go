package services

import (
	"encoding/json"
	"errors"
	"fmt"
	Config "go-core/config"
	Models "go-core/databases/models"
	"go-core/types"
	"time"

	"gorm.io/gorm"
)

type UserService interface {
	StoreUser(request *types.StoreUserRequest) error
	UpdateUser(request *types.UpdateUserRequest, id uint64) error
	ChangePassword(userId uint64, request *types.ChangePasswordRequest) error
	UserList() ([]*types.UserDTO, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func ValidateUserFromToken(userId uint64, issuedAt int64) (*types.UserInfoDTO, error) {
	var user Models.User
	expiredTime := time.Unix(issuedAt, 0)
	result := Config.DB.Where("id = ? AND (token_expired_at IS NULL OR token_expired_at < ?)", userId, expiredTime).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return UserInfo(userId)
}

func UserInfo(userId uint64) (*types.UserInfoDTO, error) {
	var userInfo types.UserInfoDTO
	redisClient := Config.RedisClient
	userKey := fmt.Sprintf("%s%d", Config.USER_CACHE_KEY, userId)
	redisUser, err := redisClient.Get(userKey).Result()
	if err != nil {
		var user Models.User

		db := Config.DB
		result := db.Preload("Roles.Permissions").First(&user, userId)
		if result.Error != nil {
			return nil, result.Error
		}

		userInfo.ID = user.ID
		userInfo.Username = user.Username
		userInfo.FullName = user.FullName
		userInfo.Email = user.Email
		userInfo.AccountType = user.AccountType
		userInfo.Status = user.Status

		var permissions []*string
		for _, role := range user.Roles {
			for _, permission := range role.Permissions {
				permissions = append(permissions, &permission.Name)
			}
		}

		userInfo.Permissisons = permissions

		userJson, err := json.Marshal(userInfo)
		if err != nil {
			return nil, err
		}

		err = redisClient.Set(userKey, userJson, 0).Err()
		if err != nil {
			return nil, err
		}
	}

	json.Unmarshal([]byte(redisUser), &userInfo)

	return &userInfo, nil
}

func (s *userService) StoreUser(request *types.StoreUserRequest) error {
	db := s.db

	password, err := HashPassword(request.Password)
	if err != nil {
		return err
	}
	user := Models.User{
		Username: request.Username,
		FullName: request.FullName,
		Email:    request.Email,
		Password: password,
	}

	var roles []*Models.Role
	roleRs := db.Find(&roles, request.Roles)
	if roleRs.Error != nil {
		return roleRs.Error
	}

	user.Roles = roles
	rs := db.Create(&user)

	return rs.Error
}

func (s *userService) UpdateUser(request *types.UpdateUserRequest, id uint64) error {
	db := s.db
	var user Models.User

	rs := db.First(&user, id)
	if rs.Error != nil {
		return rs.Error
	}

	if request.Password != "" {
		password, err := HashPassword(request.Password)
		if err != nil {
			return err
		}

		user.Password = password
	}

	user.FullName = request.FullName

	var roles []*Models.Role
	roleRs := db.Find(&roles, request.Roles)
	if roleRs.Error != nil {
		return roleRs.Error
	}

	db.Model(&user).Association("Roles").Clear()
	user.Roles = roles

	result := db.Save(&user)

	return result.Error
}

func (s *userService) ChangePassword(userId uint64, request *types.ChangePasswordRequest) error {
	db := s.db
	var user Models.User
	result := db.First(&user, userId)
	if result.Error != nil {
		return result.Error
	}

	checkPass := CheckPasswordHash(request.OldPassword, user.Password)
	if !checkPass {
		return errors.New("current password invalid")
	}

	password, err := HashPassword(request.NewPassword)
	if err != nil {
		return err
	}

	user.Password = password
	user.TokenExpiredAt = time.Now()
	result = db.Save(&user)

	return result.Error
}

func (s *userService) UserList() ([]*types.UserDTO, error) {
	var resources []*types.UserDTO
	var users []Models.User

	db := s.db
	result := db.Order("id DESC").Preload("Roles").Offset(0).Limit(20).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, user := range users {
		var resource types.UserDTO
		resource.ID = user.ID
		resource.FullName = user.FullName
		resource.Username = user.Username
		resource.Email = user.Email
		resource.CreatedAt = user.CreatedAt
		resource.Status = user.Status
		resource.AccountType = user.AccountType

		var roles []*string
		for _, role := range user.Roles {
			roles = append(roles, &role.Name)
		}

		resource.Roles = roles

		resources = append(resources, &resource)
	}

	return resources, nil
}
