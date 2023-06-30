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

// Login godoc
// @Summary Login
// @Description Login
// @Tags Authen
// @Accept json
// @Produce json
// @Param username formData string true "require"
// @Param password formData string true "require"
// @Success 200 {object} LoginRes
// @Failure 400 string string
// @Failure 500 string string
// @Router /login [post]
func Login(c echo.Context) error {
	user := c.FormValue("username")
	pass := c.FormValue("password")
	id, username, password, err := getUserByUsername(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if len(username) == 0 {
		return c.String(http.StatusBadRequest, "Invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(pass))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid username or password")
	}
	token := generateJWT(id)
	var resp LoginRes
	resp.Token = token
	resp.UserId = id
	return c.JSON(http.StatusOK, resp)
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

func ValidateToken(tokenString string) (id, message string) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // Replace "secret" with your actual secret key
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			message = "invalid token signature"
			return "",message
		}
		message = "invalid token"
		return "",message
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok || !token.Valid {
		message = "invalid token claims"
		return "",message
	}

	// if claims.UserID != userId {
	// 	message = "user ID mismatch"
	// 	return "",message
	// }

	return claims.UserID, message
}

func CheckToken(token string) (err error) {
	claims := &jwtClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil 
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return err
		}
		return err
	}
	if !tkn.Valid {
		return err
	}
	return err
}