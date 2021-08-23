package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/nghiatrandev/template/back/dataservice"
)



func (server *Server)Register(ctx echo.Context) error {

	dto := new(dataservice.UserDTO)
	ctx.Bind(dto)

	user := new(dataservice.User)
	user.FromDTO(dto)

	server.db.Begin()
	if err := user.Create(server.db); err !=  nil {
		server.db.Rollback()
	}
	server.db.Commit()


	return ctx.JSON(200,user)

}

func
