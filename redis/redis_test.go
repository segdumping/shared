package redis

import (
	"encoding/xml"
	"io/ioutil"
	"math"
	"math/rand"
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

func TestLoadConfig(t *testing.T) {
	t.Log(loadConfig())
}

func TestNewPool(t *testing.T) {
	config, err := loadConfig()
	if err != nil {
		t.Logf("load config error: %s", err.Error())
		return
	}

	pool := newPool(&config)
	conn := pool.Get()
	r, err := conn.Do("SET", "test_set", 3)
	t.Log(r, err)
}

func TestScanHash(t *testing.T) {
	for i := 0; i < math.MaxUint8; i++ {
		_, err := Do("HSET", "test_hset", i, rand.Intn(1000))
		if err != nil {
			t.Logf("hset error: %s", err.Error())
		}
	}

	m := make(map[int]int)
	r, err := ScanHash("test_hset")
	if err != nil {
		t.Logf("scan hash error: %s", err.Error())
		return
	}

	err = ArrayToMap(r, m)
	if err != nil {
		t.Logf("array to map error: %s", err.Error())
		return
	}

	for k, v := range m {
		t.Logf("key: value [%d: %d]", k, v)
	}
}
