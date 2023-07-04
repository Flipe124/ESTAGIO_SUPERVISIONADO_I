package structure

import (
	"reflect"
	"regexp"
)

// Assign assign a struct object into struct target.
// OBS: the passed structs are needed struct pointers but NOT pass like params your pointer reference (&)!
// OBS: the skips param is options.
func Assign(object, target any, skips ...string) {
	objectStruct := reflect.ValueOf(object).Elem()
	targetStruct := reflect.ValueOf(target).Elem()
	for objectIndex := 0; objectIndex < objectStruct.NumField(); objectIndex++ {
		objectField := objectStruct.Type().Field(objectIndex).Name
		for targetIndex := 0; targetIndex < targetStruct.NumField(); targetIndex++ {
			targetField := targetStruct.Type().Field(targetIndex).Name
			if targetField == objectField {
				if len(skips) > 0 && regexp.MustCompile(skips[len(skips)-1]).MatchString(targetField) {
					continue
				}
				targetStruct.Field(targetIndex).Set(objectStruct.Field(objectIndex))
			}
		}
	}
}
