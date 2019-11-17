package sql

type Config struct {
	Driver       string `xml:"driver"`
	Address      string `xml:"address"`
	Database     string `xml:"database"`
	Username     string `xml:"username"`
	Password     string `xml:"password"`
	MaxLifeTime  int    `xml:"max_life_time"`
	MaxOpenConns int    `xml:"max_open_conns"`
	MaxIdleConns int    `xml:"max_idle_conns"`
}
