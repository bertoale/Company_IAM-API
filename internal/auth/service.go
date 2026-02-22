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
	GenerateRefreshToken(user *user.User, roles []string, permissions []string, applications []string) (string, error)
	Login(req *LoginRequest) (string, string, *user.UserResponse, error)
	RefreshToken(refreshToken string) (string, string, error)
	ValidateRefreshToken(refreshToken string) (*Claims, error)
}

type service struct {
	repo Repository
	cfg  config.Config
}

type Claims struct {
	ID           uint     `json:"id"`
	Roles        []string `json:"roles"`
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
		Applications: applications,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWTSecret))
}

// GenerateRefreshToken implements [Service] - Stateless refresh token using JWT.
func (s *service) GenerateRefreshToken(user *user.User, roles []string, permissions []string, applications []string) (string, error) {
	duration, err := time.ParseDuration(s.cfg.RefreshTokenExpires)
	if err != nil {
		duration = 7 * 24 * time.Hour // default 7 days
	}
	claims := Claims{
		ID:           user.ID,
		Roles:        roles,
		Applications: applications,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.RefreshTokenSecret))
}

// ValidateRefreshToken implements [Service] - Validate and parse refresh token.
func (s *service) ValidateRefreshToken(refreshToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validasi signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.cfg.RefreshTokenSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

// RefreshToken implements [Service] - Generate new tokens from refresh token.
func (s *service) RefreshToken(refreshToken string) (string, string, error) {
	// Validate refresh token
	claims, err := s.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	// Get fresh user data from database by ID
	userObj, err := s.repo.GetUserByID(claims.ID)
	if err != nil {
		return "", "", fmt.Errorf("failed to get user: %w", err)
	}

	// Get fresh roles, permissions, and applications
	roles, err := s.repo.GetUserRoles(userObj.ID)
	if err != nil {
		return "", "", fmt.Errorf("failed to get user roles: %w", err)
	}

	permissions, err := s.repo.GetUserPermissions(userObj.ID)
	if err != nil {
		return "", "", fmt.Errorf("failed to get user permissions: %w", err)
	}

	applications, err := s.repo.GetUserApplications(userObj.ID)
	if err != nil {
		return "", "", fmt.Errorf("failed to get user applications: %w", err)
	}

	// Generate new access token
	newToken, err := s.GenerateToken(userObj, roles, permissions, applications)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate token: %w", err)
	}

	// Generate new refresh token
	newRefreshToken, err := s.GenerateRefreshToken(userObj, roles, permissions, applications)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return newToken, newRefreshToken, nil
}

// Login implements [Service].
func (s *service) Login(req *LoginRequest) (string, string, *user.UserResponse, error) {
	userObj, err := s.repo.GetUserByIdentifier(req.Identifier)
	if err != nil {
		return "", "", nil, fmt.Errorf("failed to get user by identifier: %w", err)
	}
	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(userObj.Password), []byte(req.Password)); err != nil {
		return "", "", nil, fmt.Errorf("invalid password: %w", err)
	}

	// Ambil roles, permissions, dan applications
	roles, err := s.repo.GetUserRoles(userObj.ID)
	if err != nil {
		return "", "", nil, fmt.Errorf("failed to get user roles: %w", err)
	}

	permissions, err := s.repo.GetUserPermissions(userObj.ID)
	if err != nil {
		return "", "", nil, fmt.Errorf("failed to get user permissions: %w", err)
	}

	applications, err := s.repo.GetUserApplications(userObj.ID)
	if err != nil {
		return "", "", nil, fmt.Errorf("failed to get user applications: %w", err)
	}

	token, err := s.GenerateToken(userObj, roles, permissions, applications)
	if err != nil {
		return "", "", nil, fmt.Errorf("failed to generate token: %w", err)
	}

	refreshToken, err := s.GenerateRefreshToken(userObj, roles, permissions, applications)
	if err != nil {
		return "", "", nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return token, refreshToken, user.ToUserResponse(userObj), nil
}

func NewService(repo Repository, cfg *config.Config) Service {
	return &service{
		repo: repo,
		cfg:  *cfg,
	}
}
