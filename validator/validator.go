package validator

import validate "gopkg.in/go-playground/validator.v9"

func CheckRequest(bean interface{}) (err error) {
	valid := validate.New()
	err = valid.Struct(bean)

	if err != nil {
		return
	}

	return
}
