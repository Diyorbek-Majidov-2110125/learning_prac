package config

type Config struct {
	Path string
	UserFileName string
	ProductFileName string
	ShopcartFileName string
	CommissionFileName string
}

func Load() Config {
	
	cfg := Config{}

	cfg.Path = "./data/"
	cfg.UserFileName = "users.json"
	cfg.ProductFileName = "products.json"
	cfg.ShopcartFileName = "shopcart.json"
	cfg.CommissionFileName = "commission.json"

	return cfg
}