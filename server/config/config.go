package config

type Server struct {
	Sys System   `mapstructure:"system"`
	DB  MySQL    `mapstructure:"db"`
	Log BeegoLog `mapstructure:"log"`
}
