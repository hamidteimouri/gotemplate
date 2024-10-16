package postgres

import (
	"github.com/hamidteimouri/gommon/htenvier"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres(db *gorm.DB) *Postgres {
	p := &Postgres{db: db}
	if htenvier.IsDebugMode() {
		p.db = p.db.Debug()
	}

	return p
}
