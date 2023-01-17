package model

import (
	"gorm.io/gorm"
)

type Conta struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey; autoIncrement:true" json: "id"`
	UsuarioID   string `gorm:"type:varchar(255); NOT NULL; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"idUsuario"`
	Usuario     Usuario
	NumeroConta string  `gorm:"type:varchar(10); NOT NULL;" json:"numeroConta"`
	DigitoConta string  `gorm:"type:varchar(2); NOT NULL;" json:"digitoConta"`
	Saldo       float64 `gorm:"type:float; NOT NULL;" json:"saldo"`
	TipoConta   int     `gorm:"type:int; NOT NULL;" json:"tipoConta"`
}
