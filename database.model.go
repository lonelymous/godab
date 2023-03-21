package godab

type DatabaseConfig struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (this *DatabaseConfig) GetType() string {
	return this.Type
}

func (this *DatabaseConfig) GetConnectionString() string {
	return this.Username + ":" + this.Password + "@/" + this.Name + "?parseTime=true"
}
