package parsing

import (
	"github.com/darkraiden/switch/pkg/handlers"
	"github.com/go-yaml/yaml"
)

// BuildMap builds a map of strings out of a Routes struct
func BuildMap(r Routes) map[string]string {
	paths := make(map[string]string)
	for _, path := range r.Paths {
		paths[path.Source] = path.Destination
	}
	return paths
}

// ParseYAML parses a slice of bytes of a YAML file
// and returns back a Routes type
func ParseYAML(data []byte) Routes {
	var routes Routes
	err := yaml.Unmarshal(data, &routes)
	handlers.CheckErrors(err)
	return routes
}
