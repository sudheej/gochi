package plugins

//Data primary object construct for the plugin
type Data struct {
	Steps struct {
		Execute []Execute `yaml:"execute"`
	} `yaml:"steps"`
}
