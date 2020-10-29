package config

import (
	"log"
	"os"

	"github.com/fabricorgi/orgchecker"
	"github.com/go-playground/validator"
)

//Validate ...
var Validate *validator.Validate = validator.New()

//ValidateOrgConfig ...
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

//GetEnvironmentVariables looking for valid config data...
func GetEnvironmentVariables() error {
	var exists bool

	ordererIP, exists := os.LookupEnv("FABRICORGI_ORDERER_IP")
	if !exists {
		log.Printf("Cannot find variable FABRICORGI_ORDERER_IP. Exists: %v", exists)
	} else {
		log.Printf("Orderer ip is: %v", ordererIP)
	}

	return nil
}
