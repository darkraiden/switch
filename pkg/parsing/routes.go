package parsing

// Routes provides the structure of what a route file looks like
type Routes struct {
	Paths []Path `yaml:"paths,omitempty"`
}

// Path represents what a a typical path
// looks like in the provided YAML file
type Path struct {
	Source      string `yaml:"source,omitempty"`
	Destination string `yaml:"destination,omitempty"`
}
