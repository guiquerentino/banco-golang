package model

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	ID                    uint   `gorm:"primaryKey; autoIncrement:true;" json:"id"`
	NomeProprietario      string `gorm:"type:varchar(255); NOT NULL;" json:"nomeProprietario"`
	DocumentoProprietario string `gorm:"type:varchar(20); NOT NULL;" json:"documentoProprietario"`
	Endereco              string `gorm:"type:varchar(255); NOT NULL;" json:"endereco"`
}
