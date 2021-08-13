package config

type BeegoLog struct {
	Level    string `mapstructure:"level"`
	FilePath string `mapstructure:"file_path"`
}
