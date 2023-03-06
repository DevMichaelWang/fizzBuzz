package handler

import (
	"backend/src/model"
	"backend/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetFizzBuzz(c echo.Context) (err error) {

	data := new(model.Request)
	res := new(model.Response)

	if err = c.Bind(data); err != nil {
		res.Status = 400
		res.Message = "'count' param is required and should be integer"
		return c.JSON(http.StatusBadRequest, res)
	}
	if err = c.Validate(data); err != nil {
		res.Status = 400
		res.Message = "'count' param is required and should be integer"
		return c.JSON(http.StatusBadRequest, res)
	}

	fizz := utils.GoDotEnvVariable("THREE_STRING")
	buzz := utils.GoDotEnvVariable("FIVE_STRING")
	fizz_buzz := utils.GoDotEnvVariable("THREE_AND_FIVE_STRING")
	other := utils.GoDotEnvVariable("OTHER")

	if (fizz == "env_err") || (buzz == "env_err") || (fizz_buzz == "env_err") || (other == "env_err") {
		res.Status = 500
		res.Message = "Server Internal Error"
		return c.JSON(http.StatusInternalServerError, res)
	} else {
		res.Status = 200
		if (data.Count%3 == 0) && (data.Count%5 == 0) {
			res.Message = fizz_buzz
		} else if data.Count%3 == 0 {
			res.Message = fizz
		} else if data.Count%5 == 0 {
			res.Message = buzz
		} else {
			res.Message = other
		}
	}

	return c.JSON(http.StatusOK, res)
}
