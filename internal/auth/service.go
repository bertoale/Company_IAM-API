package auth

import (
	"company_iam/internal/user"
	"company_iam/pkg/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GenerateToken(user *user.User, roles []string, permissions []string, applications []string) (string, error)
	Login(req *LoginRequest) (string, *user.UserResponse, error)
}

type service struct {
	repo Repository
	cfg  config.Config
}

type Claims struct {
	ID           uint     `json:"id"`
	Roles        []string `json:"roles"`
	Permissions  []string `json:"permissions"`
	Applications []string `json:"applications"`
	jwt.RegisteredClaims
}

// GenerateToken implements [Service].
func (s *service) GenerateToken(user *user.User, roles []string, permissions []string, applications []string) (string, error) {
	duration, err := time.ParseDuration(s.cfg.JWTExpires)
	if err != nil {
		duration = 168 * time.Hour // default 7 days
	}
	claims := Claims{
		ID:           user.ID,
		Roles:        roles,
		Permissions:  permissions,
		Applications: applications,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWTSecret))
}

// Login implements [Service].
func (s *service) Login(req *LoginRequest) (string, *user.UserResponse, error) {
	userObj, err := s.repo.GetUserByIdentifier(req.Identifier)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get user by identifier: %w", err)
	}
	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(userObj.Password), []byte(req.Password)); err != nil {
		return "", nil, fmt.Errorf("invalid password: %w", err)
	}

	// Ambil roles, permissions, dan applications
	roles, err := s.repo.GetUserRoles(userObj.ID)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get user roles: %w", err)
	}

	permissions, err := s.repo.GetUserPermissions(userObj.ID)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get user permissions: %w", err)
	}

	applications, err := s.repo.GetUserApplications(userObj.ID)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get user applications: %w", err)
	}

	token, err := s.GenerateToken(userObj, roles, permissions, applications)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate token: %w", err)
	}
	return token, user.ToUserResponse(userObj), nil
}

func NewService(repo Repository, cfg *config.Config) Service {
	return &service{
		repo: repo,
		cfg:  *cfg,
	}
}
