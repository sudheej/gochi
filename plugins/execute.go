package plugins

/*
//Vcs is the main
type Vcs struct {
	Operations struct {
		Git []Git `yaml:"operations"`
	} `yaml:"vcs"`
}

*/
//Vcs version control settings
type Vcs struct {
	Git []Git `yaml:"git"`
}

//Data defines structure of yaml
type Data struct {
	Vcs   `yaml:"vcs"`
	Steps struct {
		Execute []Execute `yaml:"execute"`
	} `yaml:"steps"`
}
