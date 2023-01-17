package model

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey; autoIncrement:true;" json:"id"`
	Nome      string `gorm:"type:varchar(255); NOT NULL;" json:"nome"`
	Documento string `gorm:"type:varchar(20); NOT NULL;" json:"documento"`
	Endereco  string `gorm:"type:varchar(255); NOT NULL;" json:"endereco"`
}
