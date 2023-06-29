package product

import (
	"net/http"
	"neversitup/authen"
	"neversitup/util"
	"strconv"
	"strings"
	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	req := new(Product)
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Error while bind HTTP Request to struct")
		}
	message, err := util.CheckRequest(*req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message)
	}

	checkProduct, err := getProductByName(req.ProductName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if len(checkProduct) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "ProductName already exist")
	}
	statement, bindValue := prepareSQLInsertProduct(*req)
	err = util.ExecuteStatement(statement, bindValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent,nil)
}

func GetProduct(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	err := authen.CheckToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}
	id := c.QueryParam("id")
	proType := c.QueryParam("type")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	size, _ := strconv.Atoi(c.QueryParam("size"))
	offset := size * (page - 1)
	products, err := getAllProduct(proType, id, offset, size)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, products)
}
