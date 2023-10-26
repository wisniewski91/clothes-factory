package utils

import "fmt"

type ValidateError struct {
	Message string
	Code    int
	Object  string
}

func (ve ValidateError) Error() string {
	result := fmt.Sprintf("Error: %d\n Object: %s\n Message: %s", ve.Code, ve.Object, ve.Message)
	return result
}

type Validator interface {
	Validate() error
}

func Validate(v Validator) error {
	err := v.Validate()
	if err != nil {
		return err
	}
	return nil
}
