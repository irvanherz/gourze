package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/irvanherz/gourze/config"
	"github.com/irvanherz/gourze/modules/auth/dto"
	"github.com/irvanherz/gourze/modules/user"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Signin(input dto.AuthSigninInput) (*dto.AuthResultDto, error)
	Signup(input dto.AuthSignupInput) (*dto.AuthResultDto, error)
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
	GenerateAccessToken(user user.User) (string, error)
	GenerateRefreshToken(user user.User) (string, error)
}

type authService struct {
	Db     *gorm.DB
	Config *config.Config
}

// GenerateAccessToken implements AuthService.
func (s *authService) GenerateAccessToken(user user.User) (string, error) {
	jwtSecret := s.Config.Auth.JWTSecret
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,                          // Subject (user identifier)
		"iss": "gourze",                         // Issuer
		"aud": user.Role,                        // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})
	tokenString, err := claims.SignedString([]byte(jwtSecret))
	return tokenString, err
}

// GenerateRefreshToken implements AuthService.
func (s *authService) GenerateRefreshToken(user user.User) (string, error) {
	panic("unimplemented")
}

// Signin implements AuthService.
func (s *authService) Signin(input dto.AuthSigninInput) (*dto.AuthResultDto, error) {
	// get user
	var user user.User
	if err := s.Db.Where("email = ? OR username = ?", input.UsernameOrEmail, input.UsernameOrEmail).First(&user).Error; err != nil {
		return nil, err
	}
	if err := s.CompareHashAndPassword(user.Password, input.Password); err != nil {
		return nil, err
	}
	accessToken, err := s.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	var authUser dto.AuthUser
	copier.Copy(&authUser, &user)
	return &dto.AuthResultDto{
		AccessToken:          accessToken,
		RefreshToken:         "",
		AccessTokenExpiredAt: time.Now().Add(time.Hour).Unix(),
		User:                 authUser,
	}, nil
}

func (s *authService) Signup(input dto.AuthSignupInput) (*dto.AuthResultDto, error) {
	var user user.User
	copier.Copy(&user, &input)
	// userMeta, _ := json.Marshal(map[string]interface{}{})
	// user.Meta = datatypes.JSON(userMeta)
	hashedPassword, err := s.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	if err := s.Db.Create(&user).Error; err != nil {
		return nil, err
	}

	accessToken, err := s.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	var authUser dto.AuthUser
	copier.Copy(&authUser, &user)
	return &dto.AuthResultDto{
		AccessToken:          accessToken,
		RefreshToken:         "",
		AccessTokenExpiredAt: 0,
		User:                 authUser,
	}, nil
}

func (s *authService) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

func (s *authService) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func NewAuthService(db *gorm.DB, conf *config.Config) AuthService {
	return &authService{Db: db, Config: conf}
}
