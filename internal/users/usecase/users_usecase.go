package usecase

import (
	"context"
	"time"
	"transaction_service/internal/commons"
	models "transaction_service/internal/domain/model"
	"transaction_service/internal/domain/repository"
	"transaction_service/internal/domain/request"

	"github.com/golang-jwt/jwt"
)

type UserUseCase struct {
	userRepository repository.UserRepository
	ctx            context.Context
}

func NewUserUseCase(
	ctx context.Context,
	userRepostory repository.UserRepository,
) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepostory,
		ctx:            ctx,
	}
}

func (u *UserUseCase) Login(ctx context.Context, request request.LoginUserRequest) (token *models.TokenDetails, err error) {
	user, err := u.userRepository.FindUserByEmail(ctx, request.UserName)
	if err != nil {
		return token, err
	}
	isTrue := commons.CheckPasswordHash(request.Password, user.Password)
	if !isTrue {
		return nil, err
	}
	JWTConfig := u.ctx.Value("JWT_config").(map[string]string)
	token, err = u.CreateToken(JWTConfig["Signature_Key"], &user)
	if err != nil {
		return nil, err
	}
	return token, err
}

func (u *UserUseCase) CreateToken(signKey string, user *models.User) (tokenDetails *models.TokenDetails, err error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return tokenDetails, err
	}
	timeNow := time.Now().In(location)
	expiresAt := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()+1, 0, 0, 0, -1, timeNow.Location()).Unix()

	claims := models.JWTClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			Audience:  "user.access_token",
			ExpiresAt: expiresAt,
			Issuer:    "ProjectPos",
			IssuedAt:  timeNow.Unix(),
			NotBefore: timeNow.Unix(),
			Subject:   user.Name,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = "v1"

	tokenString, err := token.SignedString([]byte(signKey))
	if err != nil {
		return nil, err
	}
	tokenDetails = &models.TokenDetails{
		AccessToken: tokenString,
		ExpiresAt:   expiresAt,
		UserID:      user.ID,
	}
	return tokenDetails, nil
}
