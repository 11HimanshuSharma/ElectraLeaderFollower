package config

type Config struct {
	ID string `json:"id"`
	Address string `json:"address"`
	Peers []string `json:"peers"`
}


func Load(path string) *Config {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	var cfg Config

	if err := json.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}
	return &cfg
}