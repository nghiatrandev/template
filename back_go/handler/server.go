package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/nghiatrandev/template/back/dataservice"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func NerServer() *echo.Echo  {
	db, err := dataservice.NewDB()
	if err != nil {
		log.Fatalln(err)
	}
	server := Server{db: db}

	e := server.buildRoutes()
	return e
}
