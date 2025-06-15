package config

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}
