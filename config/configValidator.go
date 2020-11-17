package config

import (
	"github.com/fabricorgi/cmd/orgchecker"
	"github.com/go-playground/validator"
)

//Validate ...
var Validate *validator.Validate = validator.New()

//ValidateOrdererConfig ...
func ValidateOrdererConfig(data *orgchecker.OrdererConfig) error {

	err := Validate.Struct(data)
	if err != nil {
		return err
	}
	return nil
}

//ValidateOrgConfig ...
func ValidateOrgConfig(data *orgchecker.OrganizationConfig) error {

	err := Validate.Struct(data)
	if err != nil {
		return err
	}
	return nil
}

//ValidateOrgRemoveConfig ...
func ValidateOrgRemoveConfig(data *orgchecker.OrganizationRemove) error {

	err := Validate.Struct(data)
	if err != nil {
		return err
	}
	return nil
}
