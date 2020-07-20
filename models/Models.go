package models

import "time"

type model struct {
	CriatoEm     time.Time  `gorm:"Column:criado_em;type:TIMESTAMP;"`
	AtualizadoEm time.Time  `gorm:"Column:atualizado_em;type:TIMESTAMP;"`
	DeletadoEm   *time.Time `gorm:"Column:deletado_em;type:TIMESTAMP;"`
}
