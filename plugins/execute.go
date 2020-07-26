package plugins

//Vcs is the main
type Vcs struct {
	Git []Git `yaml:"git"`
}

//Data primary object construct for the plugin
type Data struct {
	Vcs   `yaml:",inline"`
	Steps struct {
		Execute []Execute `yaml:"execute"`
	} `yaml:"steps"`
}
