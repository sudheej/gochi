package plugins

//Execute structure for mvn temporarily
type Execute struct {
	MVN        string `yaml:"mvn"`
	Goals      string `yaml:"goals"`
	Concurrent bool   `yaml:"concurrent"`
}
