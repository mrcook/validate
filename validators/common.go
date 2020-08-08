package validators

import (
	"fmt"

	"github.com/gobuffalo/flect"
)

var CustomKeys = map[string]string{}

func GenerateKey(s string) string {
	key := CustomKeys[s]
	if key != "" {
		return key
	}
	return flect.Underscore(s)
}

func GenerateObjectKey(o, k string) string {
	key := CustomKeys[k]
	if key != "" {
		return key
	}

	k = flect.Underscore(k)

	if o != "" {
		k = fmt.Sprintf("%s.%s", o, k)
	}

	return k
}
