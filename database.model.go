package godab

import "strconv"

type DatabaseConfig struct {
	// Hostname or IP address
	Hostname string `json:"hostname" docker:"DATABASE_HOSTNAME"`
	Port     int    `json:"port" docker:"DATABASE_PORT"`
	// Database driver
	Driver string `json:"driver" docker:"DATABASE_DRIVER"`
	// Database name
	Name      string `json:"name" docker:"DATABASE_NAME"`
	Username  string `json:"username" docker:"DATABASE_USERNAME"`
	Password  string `json:"password" docker:"DATABASE_PASSWORD"`
	ParseTime bool   `json:"parse_time" docker:"DATABASE_PARSE_TIME"`
}

func (this *DatabaseConfig) GetDriver() string {
	return this.Driver
}

func (this *DatabaseConfig) GetConnectionString() string {
	return this.Username +
		":" +
		this.Password +
		"@tcp(" +
		this.Hostname +
		":" + strconv.Itoa(this.Port) +
		")/" + this.Name +
		"?parseTime=" + this.GetParseTime()
}

func (this *DatabaseConfig) GetConnectionStringWithoutDatabase() string {
	return this.Username +
		":" +
		this.Password +
		"@tcp(" +
		this.Hostname +
		":" + strconv.Itoa(this.Port) +
		")/?parseTime=" + this.GetParseTime()
}

func (this *DatabaseConfig) GetDatabasePort() string {
	return ":" + strconv.Itoa(this.Port)
}

func (this *DatabaseConfig) GetParseTime() string {
	return strconv.FormatBool(this.ParseTime)
}
