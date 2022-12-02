package data

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rongpengju/delivery_problem/internal/pkg/etc"
)

type Database struct {
	db *gorm.DB
}

var ProviderSet = wire.NewSet(
	NewDatabase,
)

func NewDatabase(config *etc.Config) (*Database, error) {
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return &Database{db: db}, nil
}

func (d *Database) Create() {

}

func (d *Database) Get() {

}
