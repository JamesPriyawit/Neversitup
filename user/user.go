package user

import (
	"net/http"
	"neversitup/util"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c echo.Context) error {
	req := new(User)
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Error while bind HTTP Request to struct")
		}
	message, err := util.CheckRequest(*req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message)
	}

	checkUser, err := getUserByUsername(req.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if len(checkUser) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Username already exist")
	}

	if req.Password != req.RePassword {
		return echo.NewHTTPError(http.StatusBadRequest, "Password not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// pass := string(hashedPassword)
	statement, bindValue := prepareSQLInsertUser(*req, hashedPassword)
	err = util.ExecuteStatement(statement, bindValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}

func GetUser(c echo.Context) error {
	userId := c.Param("id")
	userInfo, err := getUserById(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if len(userInfo.Username) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "user not found.")
	}
	return c.JSON(http.StatusOK, userInfo)
}