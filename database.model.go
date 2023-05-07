package godab

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (this *DatabaseConfig) GetDriver() string {
	return this.Driver
}

func (this *DatabaseConfig) GetConnectionString() string {
	return this.Username + ":" + this.Password + "@/" + this.Name + "?parseTime=true"
}

func (this *DatabaseConfig) GetConnectionStringWithoutDatabase() string {
	return this.Username + ":" + this.Password + "@/?parseTime=true"
}