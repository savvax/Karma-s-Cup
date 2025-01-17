package config

type Config struct {
	MorningHour int
	MorningMin  int
	NoonHour    int
	NoonMin     int
	EveningHour int
	EveningMin  int
}

func New() *Config {
	return &Config{
		MorningHour: 8,
		MorningMin:  35,
		NoonHour:    13,
		NoonMin:     45,
		EveningHour: 20,
		EveningMin:  21,
	}
}
