package util

import(
	"neversitup/constant"
	"neversitup/db"
	validate "gopkg.in/go-playground/validator.v9"
	"log"
)

func mapErrorResponse(tag string, field string, param string) (message string) {
	switch tag {
	case constant.Required:
		message = "Error : " + field + " is required."
	case constant.Max:
		message = "Error : " + field + " is not over " + param
	case constant.Min:
		message = "Error : " + field + " must be greater than " + param
	case constant.Length:
		message = "Error : " + field + " length not more than " + param
	case constant.Numeric:
		message = "Error : " + field + " is not numeric."
	default:
		message = "Other error."
	}
	return message
}

func CheckRequest(bean interface{}) (message string, err error) {
	valid := validate.New()
	err = valid.Struct(bean)

	if err != nil {
		for _, e := range err.(validate.ValidationErrors) {
			message := mapErrorResponse(e.Tag(), e.Field(), e.Param())
			return message, err
		}
	}
	return
}

func ExecuteStatement(statement string, bindValue []interface{}) (err error) {
	dbPool := db.GetConnectPostgresDB()

	prepare, err := dbPool.Prepare(statement)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer prepare.Close()

	response, err := prepare.Exec(bindValue...)
	if err != nil {
		log.Fatal(err.Error())		
		return
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected == 0 {
		log.Println("rowsAffected :", rowsAffected)
	}

	return err
}