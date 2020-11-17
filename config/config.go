package config

import (
	"log"
	"os"
)

//GetEnvironmentVariables looking for valid config data...
func GetEnvironmentVariables() error {
	var exists bool

	_, exists = os.LookupEnv("FABRICORGI_ORDERER_IP")
	if !exists {
		log.Printf("Cannot find variable FABRICORGI_ORDERER_IP. Exists: %v", exists)
	}

	_, exists = os.LookupEnv("CORE_PEER_ADDRESS")
	if !exists {
		log.Printf("Cannot find variable CORE_PEER_ADDRESS. Exists: %v", exists)
	}

	_, exists = os.LookupEnv("CORE_PEER_LOCALMSPID")
	if !exists {
		log.Printf("Cannot find variable CORE_PEER_LOCALMSPID. Exists: %v", exists)
	}

	_, exists = os.LookupEnv("CORE_PEER_ADDRESS")
	if !exists {
		log.Printf("Cannot find variable CORE_PEER_MSPCONFIGPATH. Exists: %v", exists)
	}

	return nil
}
