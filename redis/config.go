package redis

import (
	"encoding/xml"
	"io/ioutil"
)

type Config struct {
	MaxIdle      int    `xml:"maxIdle"`
	MaxActive    int    `xml:"maxActive"`
	IdleTimeout  int    `xml:"idleTimeout"`
	ExhaustWait  bool   `xml:"exhaustWait"`
	Address      string `xml:"address"`
	Password     string `xml:"password"`
	ConnTimeout  int    `xml:"connTimeout"`
	ReadTimeout  int    `xml:"readTimeout"`
	WriteTimeout int    `xml:"writeTimeout"`
}

func loadConfig() (Config, error){
	b, err := ioutil.ReadFile("config.xml")
	if err != nil {
		return Config{}, err
	}

	var conf Config
	err = xml.Unmarshal(b, &conf)

	return conf, err
}