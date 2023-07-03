package user

import (
	"net/http"
	"neversitup/constant"
	"neversitup/common"
	"strings"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"neversitup/authen"
)

// CreateUser godoc
// @Summary CreateUser
// @Description CreateUser
// @Tags User
// @Accept json
// @Produce json
// @Param User body User true "except created every field are require and validate length"
// @Success 201 string string
// @Failure 400 string string
// @Failure 500 string string
// @Router /api/register [post]
func CreateUser(c echo.Context) error {
	req := new(User)
		if err := c.Bind(req); err != nil {
			return c.String(http.StatusBadRequest, "Error while bind HTTP Request to struct.")
		}
	message, err := common.CheckRequest(*req)
	if err != nil {
		return c.String(http.StatusBadRequest, message)
	}

	checkUser, err := getUserByUsername(req.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if len(checkUser) > 0 {
		return c.String(http.StatusBadRequest, "Username already exist.")
	}

	if req.Password != req.RePassword {
		return c.String(http.StatusBadRequest, "Password not match.")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	// pass := string(hashedPassword)
	statement, bindValue := prepareSQLInsertUser(*req, hashedPassword)
	err = common.ExecuteStatement(statement, bindValue)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusCreated, constant.Created)
}

// GetUser godoc
// @Summary GetUser
// @Description GetUser
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Failure 400 string string
// @Failure 500 string string
// @Router /api/getUser [get]
// @Security ApiKeyAuth
// @securityDefinitions.basic BasicAuth
// @Param Authorization header string true "Bearer"
func GetUser(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	userId, errMsg := authen.ValidateToken(token)
	if len(errMsg) > 0 {
		return c.String(http.StatusUnauthorized, errMsg)
	}
	userInfo, err := getUserById(userId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userInfo)
}