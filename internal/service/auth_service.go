package service

import (
	"errors"
	"final-project/dto/request"
	"final-project/internal/domain"
	"final-project/internal/repository"
	"final-project/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(req request.LoginRequest) (string, error) {
	user, err := s.UserRepo.FindByEmailOrPhoneLogin(req.EmailOrPhone)
	if err != nil {
		return "", errors.New("Invalid Credentials")
	}

	// Compare Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("Invalid Credentials")
	}

	// Generate Token
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

type AuthService struct {
	UserRepo  *repository.UserRepository
	StoreRepo *repository.StoreRepository
}

func NewAuthService(userRepo *repository.UserRepository, storeRepo *repository.StoreRepository) *AuthService {
	return &AuthService{UserRepo: userRepo, StoreRepo: storeRepo}
}

func (s *AuthService) Register(req request.RegisterRequest) error {
	// Cek email / phone
	if existing, _ := s.UserRepo.FindByEmailOrPhone(req.Email, req.Phone); existing != nil {
		return errors.New("email or phone already used")
	}

	// Hash password
	hashedPass, _ := utils.HashPassword(req.Password)

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPass,
	}

	if err := s.UserRepo.Create(&user); err != nil {
		return err
	}

	// Auto-create store
	store := domain.Store{
		Name:   user.Name + "'s Store",
		UserID: user.ID,
	}
	return s.StoreRepo.Create(&store)
}
