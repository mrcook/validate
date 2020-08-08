package validators

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/validate/v3"
)

type StringsMatch struct {
	Object  string
	Name    string
	Field   string
	Field2  string
	Message string
}

// IsValid performs the validation equality of two strings.
func (v *StringsMatch) IsValid(errors *validate.Errors) {
	if strings.TrimSpace(v.Field) != strings.TrimSpace(v.Field2) {
		if v.Message == "" {
			v.Message = fmt.Sprintf("%s does not equal %s.", v.Field, v.Field2)
		}
		errors.Add(GenerateObjectKey(v.Object, v.Name), v.Message)
	}
}
