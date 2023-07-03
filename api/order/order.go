package order

import (
	"net/http"
	"neversitup/authen"
	"neversitup/constant"
	"neversitup/common"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// CreateOrder godoc
// @Summary CreateOrder
// @Description CreateOrder
// @Tags Order
// @Accept json
// @Produce json
// @Param productId formData string true "productId"
// @Success 201 string string
// @Failure 400 string string
// @Failure 401 string string
// @Failure 500 string string
// @Router /api/createOrder [post]
// @Security ApiKeyAuth
// @securityDefinitions.basic BasicAuth
// @Param Authorization header string true "Bearer"
func CreateOrder(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	productId := c.FormValue("productId")
	// req := new(Order)
	// 	if err := c.Bind(req); err != nil {
	// 		return c.String(http.StatusBadRequest, "Error while bind HTTP Request to struct")
	// 	}
	// message, err := util.CheckRequest(*req)
	// if err != nil {
	// 	return c.String(http.StatusBadRequest, message)
	// }
	// get userId from token
	userId, errMsg := authen.ValidateToken(token)
	if len(errMsg) > 0 {
		return c.String(http.StatusUnauthorized, errMsg)
	}
	statement, bindValue := prepareSQLInsertOrder(productId, userId)
	err := common.ExecuteStatement(statement, bindValue)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusCreated, constant.Created)
}

// GetOrder godoc
// @Summary GetOrder
// @Description GetOrder
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} []Order
// @Param orderId query string false "optional"
// @Param status query string false "optional"
// @Param page query string false "optional"
// @Param size query string false "optional"
// @Failure 400 string string
// @Failure 401 string string
// @Failure 500 string string
// @Router /api/getOrder [get]
// @Security ApiKeyAuth
// @securityDefinitions.basic BasicAuth
// @Param Authorization header string true "Bearer"
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
		return c.String(http.StatusUnauthorized, errMsg)
	}
	orders, err := getOrderById(orderId, userId, status, offset, size)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orders)
}

// DeleteOrder godoc
// @Summary DeleteOrder
// @Description DeleteOrder
// @Tags Order
// @Accept json
// @Produce json
// @Success 204 
// @Failure 400 string string
// @Failure 500 string string
// @Router /api/deleteOrder [delete]
// @Param id path string true "productId"
func DeleteOrder(c echo.Context) error {
	id := c.Param("id")
	statement, bindValue := prepareSQLDeleteOrder(id)
	err := common.ExecuteStatement(statement, bindValue)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent,nil)
}