package config

import ()

type Config struct {
	Rules             RulesConfig             `yaml:"rules"`
	SensitivePatterns SensitivePatternsConfig `yaml:"sensitive-patterns"`
	Exclude           ExcludeConfig           `yaml:"exclude"`
}

type RulesConfig struct {
	LowercaseStart bool `yaml:"lowercase-start"`
	EnglishOnly    bool `yaml:"english-only"`
	NoSpecialChars bool `yaml:"no-special-chars"`
	SensitiveData  bool `yaml:"sensitive-data"`
}

type SensitivePatternsConfig struct {
	Builtin bool     `yaml:"builtin"` // Использовать встроенные (password, token, api_key)
	Custom  []string `yaml:"custom"`
}

// Исключаем файлы и какие тоне чувствительные данные
type ExcludeConfig struct {
	Files    []string `yaml:"files"`
	Messages []string `yaml:"messages"`
}

func DefaultConfig() *Config {
	return &Config{
		Rules: RulesConfig{
			LowercaseStart: true,
			EnglishOnly:    true,
			NoSpecialChars: false,
			SensitiveData:  false,
		},
		SensitivePatterns: SensitivePatternsConfig{
			Builtin: true,
			Custom:  []string{},
		},
		Exclude: ExcludeConfig{
			Files:    []string{},
			Messages: []string{},
		},
	}
}
