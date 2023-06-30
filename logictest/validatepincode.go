package logictest

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

// ValidatePinCode godoc
// @Summary ValidatePinCode
// @Description logic test - ValidatePinCode
// @Tags logictest
// @Accept json
// @Produce json
// @Param input query string true "input"
// @Success 200 {object} Result
// @Failure 400 {object} Result
// @Router /validatePinCode [get]
func ValidatePinCode(c echo.Context) error{
	input := c.QueryParam("input")
	var resp Result
	resp.Result = "false"
	_, err := strconv.Atoi(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resp)
	}
	chars := []rune(input)
	if len(chars) < 6 {
		return c.JSON(http.StatusOK, resp)
	}

	for i := 0; i < len(chars) - 2 ; i++ {
		n1, _ := strconv.Atoi(string(chars[i]))
		n2, _ := strconv.Atoi(string(chars[i+1]))
		n3, _ := strconv.Atoi(string(chars[i+2]))
		if n1 == n2 && n2 == n3 {
			return c.JSON(http.StatusOK, resp)
		}
		if n1 + 1 == n2 && n2 + 1 == n3 {
			return c.JSON(http.StatusOK, resp)
		}
	}
	count := 0
	for i := 0; i < len(chars) - 1 ; i++ {
		n1, _ := strconv.Atoi(string(chars[i]))
		n2, _ := strconv.Atoi(string(chars[i+1]))
		if n1 == n2 {
			count++
		}
		if count > 2 {
			return c.JSON(http.StatusOK, resp)
		}
	}
	resp.Result = "true"
	return c.JSON(http.StatusOK, resp)
}