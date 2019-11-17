package redis

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
