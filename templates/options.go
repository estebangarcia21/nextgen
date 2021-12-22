package templates

import "flag"

// Options is the configurable aspects of a template.
type Options struct {
	TemplateName string // The name of the template to generate
	ArtifactName string // The name of the artifact being generated
	Typescript   bool   // Is the artifact being generated typescript?
}

// ReadOptions reads the options from the command line.
func ReadOptions() Options {
	flag.Parse()

	return Options{
		TemplateName: flag.Arg(0),
		ArtifactName: flag.Arg(1),
		Typescript:   flag.Arg(2) == "--ts",
	}
}
