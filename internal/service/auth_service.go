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
		return "", errors.New("invalid credentials")
	}

	// Compare Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid credentials")
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

func (s *AuthService) Register(req request.RegisterRequest) (*domain.User, error) { // Fungsi services untuk Register
	// Cek email / phone
	if existing, _ := s.UserRepo.FindByEmailOrPhone(req.Email, req.Phone); existing != nil {
		return nil, errors.New("email atau nomor telepon telah digunakan")
	}

	// Hash password
	hashedPass, _ := utils.HashPassword(req.Password)

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPass,
		Role:     "user", // default role
	}

	if err := s.UserRepo.Create(&user); err != nil {
		return nil, err
	}

	// Auto-create store
	store := domain.Store{
		Name:   user.Name + "'s Store",
		UserID: user.ID,
	}

	if err := s.StoreRepo.Create(&store); err != nil {
		return nil, err
	}

	return &user, nil
}
