package sql

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func loadConfig() (Config, error) {
	b, err := ioutil.ReadFile("config.xml")
	if err != nil {
		return Config{}, err
	}

	var conf Config
	err = xml.Unmarshal(b, &conf)

	return conf, err
}

func TestMysql(t *testing.T) {
	conf, err := loadConfig()
	if err != nil {
		t.Log(err)
		return
	}

	dbPool, err := NewDBPool(&conf)
	if err != nil {
		t.Log(err)
		return
	}

	r := dbPool.QueryRow("select count(*) from srinfo")
	var count int
	if err := r.Scan(&count); err != nil {
		t.Log(err)
	}

	t.Log(count)
}
