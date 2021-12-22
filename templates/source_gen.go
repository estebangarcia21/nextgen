package templates

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// GenerateSource generates the final source for a template on the file system.
func GenerateSource(opts Options) error {
	template, ok := Templates[opts.TemplateName]
	if !ok {
		return fmt.Errorf("template %s not found", opts.TemplateName)
	}

	var artifacts []Artifact

	if opts.Typescript {
		artifacts = template.Ts
	} else {
		artifacts = template.Js
	}

	for _, artifact := range artifacts {
		name, src := artifact.Name, artifact.Src
		var fileExt string

		if opts.Typescript {
			fileExt = "ts"
		} else {
			fileExt = "js"
		}

		if artifact.React {
			fileExt = fileExt + "x"
		}

		fileName := fmt.Sprintf("%s.%s", replaceOpt(name, "name", opts.ArtifactName), fileExt)
		interpolatedSrc := replaceOpt(src, "name", opts.ArtifactName)

		err := ioutil.WriteFile(fileName, []byte(interpolatedSrc), 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

// replaceOpt replaces anything in between <placeholderName> with a value.
func replaceOpt(src string, name string, value string) string {
	return strings.ReplaceAll(src, fmt.Sprintf("<%s>", name), value)
}
