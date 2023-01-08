package main

import (
	"github.com/arvians-id/go-clean-architecture/cmd/config"
	"github.com/arvians-id/go-clean-architecture/cmd/injection"
	"log"
)

func main() {
	configuration := config.New("../../.env")
	router, err := injection.InitServerAPI(configuration)
	if err != nil {
		log.Fatalln(err)
	}

	err = router.Run(configuration.Get("APP_PORT"))
	if err != nil {
		log.Fatalln(err)
	}
}
