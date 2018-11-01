package state

import (
	"strings"
)

// AsLowerCase is a StateFunc from helper/schema that converts the
// supplied value to lower before saving to state for consistency.
func AsLowerCase(val interface{}) string {
	return strings.ToLower(val.(string))
}
