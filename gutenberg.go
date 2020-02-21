package gutenberg

import (
	"fmt"
	"strings"

	"github.com/cloudfoundry/packit"
	"github.com/cloudfoundry/packit/scribe"
)

type Press struct {
}

func NewPress() Press {
	return Press{}
}

func (p Press) MapEnv(env packit.Environment) scribe.FormattedMap {
	envMap := scribe.FormattedMap{}
	for key, value := range env {
		envName, operation := splitEnv(key)
		switch {
		case operation == "override" || operation == "default":
			envMap[envName] = fmt.Sprintf("%q", value)
		case operation == "prepend":
			delim := env[envName+".delim"]
			envMap[envName] = fmt.Sprintf("%q", strings.Join([]string{value, "$" + envName}, delim))
		case operation == "append":
			delim := env[envName+".delim"]
			envMap[envName] = fmt.Sprintf("%q", strings.Join([]string{"$" + envName, value}, delim))
		}
	}

	return envMap
}

func splitEnv(env string) (string, string) {
	splitEnv := strings.SplitN(env, ".", 2)
	return splitEnv[0], splitEnv[1]
}
