package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //初始化
)

type dao struct {
	db *sql.DB
}

func New() (d *dao) {
	//todo

	return
}

type Dao interface {
	Close()
}

// Close close the resource.
func (d *dao) Close() {

}
