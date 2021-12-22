package templates

import (
	"fmt"
	"log"

	// Embed directive for template loading.
	_ "embed"

	"gopkg.in/yaml.v2"
)

//go:embed src/component.yml
var component string

// Templates is all the templates in the templates directory.
var Templates = map[string]Template{
	"component": parseYml(component),
}

// Parse parses a template from the templates directory.
//
// A template is a yaml file. The yml extension should be omitted.
func parseYml(src string) Template {
	fmt.Printf("src: %s\n", src)
	t := Template{}

	err := yaml.Unmarshal([]byte(src), &t)
	if err != nil {
		log.Fatalf("invalid template yml: %v", err)
		return t
	}

	return t
}
