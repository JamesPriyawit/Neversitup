package authen

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type jwtClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}


func Login(c echo.Context) error {
	user := c.FormValue("username")
	pass := c.FormValue("password")
	id, username, password, err := getUserByUsername(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if len(username) == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(pass))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
	}
	token := generateJWT(id)
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
		"userId": id,
	})
}

func generateJWT(userID string) string {
	claims := &jwtClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("secret")) 

	return signedToken
}

func ValidateToken(tokenString string, userID string) (message string) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // Replace "secret" with your actual secret key
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			message = "invalid token signature"
			return message
		}
		message = "invalid token"
		return message
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok || !token.Valid {
		message = "invalid token claims"
		return message
	}

	if claims.UserID != userID {
		message = "user ID mismatch"
		return message
	}

	return message
}

func CheckToken(token string) error {
	claims := &jwtClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil 
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return echo.ErrUnauthorized
		}
		return echo.ErrInternalServerError
	}
	if !tkn.Valid {
		return echo.ErrUnauthorized
	}
	return nil
}