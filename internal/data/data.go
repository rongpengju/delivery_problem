package data

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	db *gorm.DB
}

var ProviderSet = wire.NewSet(
	NewDatabase,
)

func NewDatabase() (*Database, error) {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	db.AutoMigrate(&Order{})
	return &Database{db: db}, nil
}

func (d *Database) Insert(order *Order) error {
	if err := d.db.Create(&Order{
		OrderId:   order.OrderId,
		Uid:       order.Uid,
		Weight:    order.Weight,
		CreatedAt: order.CreatedAt,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (d *Database) Get(uid int) ([]*Order, error) {
	var res []*Order
	if err := d.db.Where("uid = ?", uid).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
