package logs

type Config struct {
	Level        string                 `json:"level" default:"debug"`
	Encode       string                 `json:"encode" default:"console"`
	LevelPort    int                    `json:"levelPort" default:"0"`
	LevelPattern string                 `json:"levelPattern" default:""`
	Output       string                 `json:"output" default:"console"`
	InitFields   map[string]interface{} `json:"initFields"`
	File         FileSinkConfig         `json:"file"`
}

type FileSinkConfig struct {
	Path       string `json:"path" default:"./logs/app.log"`
	MaxSize    int    `json:"maxSize" default:"100"`
	MaxBackups int    `json:"maxBackups" default:"10"`
	MaxAge     int    `json:"maxAge" default:"30"`
	Encode     string `json:"encode"  default:"json"`
}

func NewConfig() *Config {
	return &Config{
		Level:        "debug",
		Encode:       "console",
		Output:       "console",
	}
}