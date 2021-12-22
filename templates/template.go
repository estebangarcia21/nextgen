package templates

// Template defines the shape of a yaml template.
type Template struct {
	Ts []Artifact `yaml:"ts,omitempty"`
	Js []Artifact `yaml:"js"`
}

// Artifact represents a generated file artifact.
type Artifact struct {
	Name  string `yaml:"name"`            // The name of the file
	Src   string `yaml:"src"`             // The source to be generated
	React bool   `yaml:"react,omitempty"` // Is the artifact being generated react? (JSX / TSX)
}
