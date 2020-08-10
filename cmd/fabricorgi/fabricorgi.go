package main

import (
	"log"

	"github.com/fabricorgi/api"
	"github.com/fabricorgi/config"
)

func main() {
	initController()
}

func initController() {
	log.Println(
		`
________ ________  ________  ________  ___  ________  ________  ________  ________  ___     
|\  _____\\   __  \|\   __  \|\   __  \|\  \|\   ____\|\   __  \|\   __  \|\   ____\|\  \    
\ \  \__/\ \  \|\  \ \  \|\ /\ \  \|\  \ \  \ \  \___|\ \  \|\  \ \  \|\  \ \  \___|\ \  \   
 \ \   __\\ \   __  \ \   __  \ \   _  _\ \  \ \  \    \ \  \\\  \ \   _  _\ \  \  __\ \  \  
  \ \  \_| \ \  \ \  \ \  \|\  \ \  \\  \\ \  \ \  \____\ \  \\\  \ \  \\  \\ \  \|\  \ \  \ 
   \ \__\   \ \__\ \__\ \_______\ \__\\ _\\ \__\ \_______\ \_______\ \__\\ _\\ \_______\ \__\
    \|__|    \|__|\|__|\|_______|\|__|\|__|\|__|\|_______|\|_______|\|__|\|__|\|_______|\|__|`)

	log.Println(`--------- Getting ENV-variables ---------`)
	initialiseControllerConfig()
	log.Println("--------- Complete ENV-variables ---------")
	log.Println("--------- Starting serve requests ---------")
	api.InitialiseAPI()
}

func initialiseControllerConfig() {
	err := config.GetEnvironmentVariables()
	if err != nil {
		log.Fatalf("Error while check Environment Vars, %v", err)
	}
}
