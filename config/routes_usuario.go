package config

import (
	"banco-golang/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetaRotas(db *gorm.DB) *gin.Engine {

	rota := gin.Default()

	rota.POST("/api/v1/usuario", func(context *gin.Context) {
		var objetoJson model.Usuario

		if err := context.ShouldBindJSON(&objetoJson); err != nil {
			context.JSON(400, gin.H{
				"erro": "Objeto não reconhecido!",
			})
		}

		db.Create(&objetoJson)

		context.JSON(200, gin.H{
			"mensagem": "Usuario criado com sucesso!",
		})

	})

	rota.GET("/api/v1/:id", func(context *gin.Context) {
		var objetos []model.Usuario

		id := context.Param("id")

		db.Find(&objetos, id)

		context.JSON(200, objetos)

	})

	rota.GET("/api/v1/getUsuarios", func(context *gin.Context) {
		var objetos []model.Usuario

		resultado := db.First(&objetos)

		fmt.Println(resultado)

		context.JSON(200, objetos)
	})

	rota.PUT("/api/v1/update/:id", func(context *gin.Context) {
		var objeto model.Usuario
		var objetoDTO model.Usuario

		if err := context.ShouldBindJSON(&objetoDTO); err != nil {
			panic("Erro no JSON!")
		}

		id := context.Param("id")

		db.First(&objeto, id)

		objeto.DocumentoProprietario = objetoDTO.DocumentoProprietario
		objeto.NomeProprietario = objetoDTO.NomeProprietario
		objeto.DocumentoProprietario = objetoDTO.NomeProprietario
		objeto.Endereco = objetoDTO.Endereco

		db.Save(&objeto)

		context.JSON(200, gin.H{
			"mensagem": "Usuario atualizado com sucesso!",
		})

	})

	rota.DELETE("/api/v1/delete/:id", func(context *gin.Context) {
		var objeto model.Usuario

		id := context.Param("id")

		db.First(&objeto, id)

		db.Delete(&objeto, id)

		context.JSON(200, gin.H{
			"mensagem": "Usuário deletado com sucesso!",
		})

	})

	return rota

}
