package main

import (
	"fmt"
	"github.com/nghiatrandev/template/back/config"
	"github.com/nghiatrandev/template/back/handler"
	log "github.com/sirupsen/logrus"
)

var cfg *config.Config

func main() {
	fmt.Println("hello")


	server := handler.NerServer()
	log.Fatal(server.Start(":8080"))

}



