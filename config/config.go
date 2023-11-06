package config

type Config struct {
	Path string
	UserFileName string
	ProductFileName string
	ShopcartFileName string
}

func Load() Config {
	
	cfg := Config{}

	cfg.Path = "./data/"
	cfg.UserFileName = "users.json"
	cfg.ProductFileName = "products.json"
	cfg.ShopcartFileName = "shopcart.json"

	return cfg
}