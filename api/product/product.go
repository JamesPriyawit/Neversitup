package product

import (
	"net/http"
	"neversitup/authen"
	"neversitup/constant"
	"neversitup/common"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// CreateProduct godoc
// @Summary CreateProduct
// @Description CreateProduct
// @Tags Product
// @Accept json
// @Produce json
// @Param Product body Product true "name,desc,type is require"
// @Success 201 string string
// @Failure 400 string string
// @Failure 500 string string
// @Router /api/createProduct [post]
func CreateProduct(c echo.Context) error {
	req := new(Product)
		if err := c.Bind(req); err != nil {
			return c.String(http.StatusBadRequest, "Error while bind HTTP Request to struct")
		}
	message, err := common.CheckRequest(*req)
	if err != nil {
		return c.String(http.StatusBadRequest, message)
	}

	checkProduct, err := getProductByName(req.ProductName)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if len(checkProduct) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "ProductName already exist")
	}
	statement, bindValue := prepareSQLInsertProduct(*req)
	err = common.ExecuteStatement(statement, bindValue)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusCreated, constant.Created)
}

// GetProduct godoc
// @Summary GetProduct
// @Description GetProduct
// @Tags Product
// @Accept json
// @Produce json
// @Param productId query string false "optional"
// @Param productType query string false "optional"
// @Param page query string false "optional"
// @Param size query string false "optional"
// @Success 200 {object} []Product
// @Failure 400 string string
// @Failure 500 string string
// @Router /api/getProduct [get]
// @Security ApiKeyAuth
// @securityDefinitions.basic BasicAuth
// @Param Authorization header string true "Bearer"
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
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, products)
}
