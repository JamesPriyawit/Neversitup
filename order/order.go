package order

import (
	"net/http"
	"neversitup/util"
	"strconv"
	"strings"
	"github.com/labstack/echo/v4"
	"neversitup/authen"
)

func CreateOrder(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	req := new(Order)
		if err := c.Bind(req); err != nil {
			return c.String(http.StatusBadRequest, "Error while bind HTTP Request to struct")
		}
	message, err := util.CheckRequest(*req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message)
	}
	// get userId from token
	userId, errMsg := authen.ValidateToken(token)
	if len(errMsg) > 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, errMsg)
	}
	statement, bindValue := prepareSQLInsertOrder(req.ProductId, userId)
	err = util.ExecuteStatement(statement, bindValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}

func GetOrder(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	orderId := c.QueryParam("orderId")
	status := c.QueryParam("status")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	size, _ := strconv.Atoi(c.QueryParam("size"))
	offset := size * (page - 1)
	// get userId from token
	userId, errMsg := authen.ValidateToken(token)
	if len(errMsg) > 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, errMsg)
	}
	orders, err := getOrderById(orderId, userId, status, offset, size)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orders)
}

func DeleteOrder(c echo.Context) error {
	id := c.Param("id")
	statement, bindValue := prepareSQLDeleteOrder(id)
	err := util.ExecuteStatement(statement, bindValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent,nil)
}