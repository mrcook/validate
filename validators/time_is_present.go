package validators

import (
	"fmt"
	"time"

	"github.com/gobuffalo/validate/v3"
)

type TimeIsPresent struct {
	Object  string
	Name    string
	Field   time.Time
	Message string
}

// IsValid adds an error if the field is not a valid time.
func (v *TimeIsPresent) IsValid(errors *validate.Errors) {
	t := time.Time{}
	if v.Field.UnixNano() != t.UnixNano() {
		return
	}

	if len(v.Message) > 0 {
		errors.Add(GenerateKey(v.Name), v.Message)
		return
	}

	errors.Add(GenerateObjectKey(v.Object, v.Name), fmt.Sprintf("%s can not be blank.", v.Name))
}
