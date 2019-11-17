package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DBPool struct {
	*sql.DB
}

func NewDBPool(conf *Config) (*DBPool, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		conf.Username,
		conf.Password,
		conf.Address,
		conf.Database,
	)

	db, err := sql.Open(conf.Driver, dns)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(conf.MaxLifeTime) * time.Second)

	//create conns
	for i := 0; i < conf.MaxIdleConns; i++ {
		err = db.Ping()
		if err != nil {
			return nil, err
		}
	}

	return &DBPool{db}, nil
}
