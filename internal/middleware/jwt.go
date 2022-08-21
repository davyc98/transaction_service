package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	models "transaction_service/internal/domain/model"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	JWTRegexString = "^[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]*$"
)

var (
	JWTRegex = regexp.MustCompile(JWTRegexString)
)

type JWTMiddleware struct {
	JWTConfig string
}

func NewJWTMiddleware(JWTConfig string) *JWTMiddleware {
	return &JWTMiddleware{JWTConfig: JWTConfig}
}

func (m *JWTMiddleware) Middleware() echo.MiddlewareFunc {
	conf := middleware.JWTConfig{
		Skipper:        m.skipper,
		ContextKey:     "user",
		SigningMethod:  jwt.SigningMethodHS256.Name,
		Claims:         &models.JWTClaims{},
		SigningKey:     []byte(m.JWTConfig),
		TokenLookup:    "header:" + echo.HeaderAuthorization,
		AuthScheme:     "Bearer",
		ParseTokenFunc: m.parseTokenFunc,
		ErrorHandler:   m.errorHandler,
		SuccessHandler: m.successHandler,
	}
	return middleware.JWTWithConfig(conf)
}

func (m *JWTMiddleware) parseTokenFunc(authToken string, _ echo.Context) (interface{}, error) {
	var claims models.JWTClaims
	t, err := jwt.ParseWithClaims(authToken, &claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, fmt.Errorf("unexpected access token signing method=%v, expect %v", t.Header["alg"], jwt.SigningMethodHS256)
		}
		if kid, ok := t.Header["kid"].(string); ok {
			if kid == "v1" {
				return []byte(m.JWTConfig), nil
			}
		}
		return nil, fmt.Errorf("unexpected access token kid=%v", t.Header["kid"])
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, errors.New("invalid token")
	}
	if claims.Audience != fmt.Sprintf("user.access_token") {
		return echo.NewHTTPError(http.StatusUnauthorized,
			fmt.Sprintf("Invalid access token, audience mismatch, got %q, expected %q. you may send request to the wrong environment",
				claims.Audience, fmt.Sprintf("user.access_token"),
			)), nil
	}
	return t, nil
}

func (m *JWTMiddleware) errorHandler(err error) error {
	switch {
	case errors.Is(err, jwt.ErrSignatureInvalid):
		return echo.NewHTTPError(http.StatusUnauthorized)
	case errors.Is(err, middleware.ErrJWTMissing):
		return echo.NewHTTPError(http.StatusBadRequest, "missing or malformed token")
	case errors.Is(err, middleware.ErrJWTInvalid):
		return echo.NewHTTPError(http.StatusUnauthorized, "token is expired")
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
}

func (m *JWTMiddleware) skipper(ctx echo.Context) bool {
	if ctx.Request().URL.Path == "/login" {
		return true
	}
	return false
}

func (m *JWTMiddleware) successHandler(ctx echo.Context) {
	switch ctx.Get("user").(type) {
	case *jwt.Token:
		token := ctx.Get("user").(*jwt.Token)
		claims, _ := token.Claims.(*models.JWTClaims)
		ctx.Set("name", claims.Subject)
	}
}
