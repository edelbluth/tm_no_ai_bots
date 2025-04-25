package tm_no_ai_bots

type Config struct {
	BotPatterns []string `json:"botPatterns,omitempty" yaml:"botPatterns,omitempty"`
}

func CreateConfig() *Config {
	return &Config{
		BotPatterns: []string{},
	}
}
