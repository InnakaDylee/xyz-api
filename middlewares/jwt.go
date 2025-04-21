package middlewares

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id, email, err := ExtractToken(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
			}
			c.Set("id", id)
			c.Set("email", email)
			return next(c)
		}
	}
}

func GenerateAuthToken(id int, email string) (string, error) {
	log.Printf("Generate Token: ID: %d, Email: %s", id, email)

	tokenClaims := jwt.MapClaims{
		"authorized": true,
		"id":         id,
		"email":       email,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateActivateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 10).Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token, err
}

func ExtractToken(c echo.Context) (int, string, error) {
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return 0, "", errors.New("missing authorization token")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return 0, "", errors.New("invalid authorization token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New("invalid authorization token")
	}

	idFloat, idOk := claims["id"].(float64)
	email, emailOk := claims["email"].(string)

	if !idOk || !emailOk {
		return 0, "", errors.New("invalid token claims")
	}

	id := int(idFloat)

	fmt.Println("ID:", id)
	fmt.Println("email:", email)


	return id, email, nil
}

func ExtractTokenEmail(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		return email, nil
	}

	return "", jwt.NewValidationError("invalid token", jwt.ValidationErrorSignatureInvalid)
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorSignatureInvalid)
}

func IsTokenExpired(tokenString string) (bool, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return false, err
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return false, jwt.NewValidationError("invalid exp claim", jwt.ValidationErrorClaimsInvalid)
	}

	return time.Unix(int64(exp), 0).Before(time.Now()), nil
}

