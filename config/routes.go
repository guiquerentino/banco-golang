package config

import (
	"banco-golang/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetaRotas(db *gorm.DB) *gin.Engine {

	rota := gin.Default()

	rota.GET("/getUsuario", func(context *gin.Context) {
		var objetos []model.Usuario

		resultado := db.Find(&objetos)

		fmt.Println(resultado)

		context.JSON(200, objetos)
	})

	return rota

}
