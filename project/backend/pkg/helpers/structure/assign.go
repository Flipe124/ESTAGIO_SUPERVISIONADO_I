package structure

import (
	"reflect"
	"regexp"
	"strings"
	"time"
)

// Assign assign a struct object into struct target.
// OBS: the passed structs are needed struct pointers but NOT pass like params your pointer reference (&)!
// OBS: the skips param is options.
// OBS: if in equivalent fields (with equal names), we have *any in the object, the target field must be a pointer too, if it is not a pointer in the object, in the target it must not be either.
func Assign(object, target any, skips ...string) {
	objectStruct := reflect.ValueOf(object).Elem()
	targetStruct := reflect.ValueOf(target).Elem()
	for objectIndex := 0; objectIndex < objectStruct.NumField(); objectIndex++ {
		objectStructField := objectStruct.Field(objectIndex)
		objectStructFieldTypeKind := objectStructField.Kind()
		objectStructFieldTypeString := objectStructField.Type().String()
		for targetIndex := 0; targetIndex < targetStruct.NumField(); targetIndex++ {
			targetStructField := targetStruct.Field(targetIndex)
			targetStructFieldName := targetStruct.Type().Field(targetIndex).Name
			if objectStruct.Type().Field(objectIndex).Name == targetStructFieldName && !regexp.MustCompile(targetStructFieldName).MatchString(strings.Join(skips, " ")) {
				objectStructFieldElem := objectStructField.Elem()
				objectStructFieldInterface := objectStructField.Interface()
				if objectStructFieldTypeKind == reflect.Interface {
					switch targetStructField.Kind() {
					case reflect.String:
						targetStructField.SetString(objectStructFieldElem.String())
					case reflect.Int64:
						targetStructField.SetInt(objectStructFieldElem.Int())
					case reflect.Float64:
						targetStructField.SetFloat(objectStructFieldElem.Float())
					case reflect.Bool:
						targetStructField.SetBool(objectStructFieldElem.Bool())
					case reflect.Struct:
						if targetStructField.Type() == reflect.TypeOf(time.Time{}) {
							datetime, _ := time.Parse("2006-01-02 15:04:05", objectStructFieldInterface.(any).(string))
							targetStructField.Set(reflect.ValueOf(any(datetime)))
						}
					default:
						targetStructField.Set(objectStructField)
					}
				} else if objectStructFieldTypeString == reflect.TypeOf((*interface{})(nil)).String() {
					switch targetStructField.Type() {
					case reflect.TypeOf((*string)(nil)):
						targetStructField.SetString(objectStructFieldInterface.(string))
					case reflect.TypeOf((*int)(nil)):
						targetStructField.SetInt(objectStructFieldInterface.(int64))
					case reflect.TypeOf((*float64)(nil)):
						targetStructField.SetFloat(objectStructFieldInterface.(float64))
					case reflect.TypeOf((*bool)(nil)):
						targetStructField.SetBool(objectStructFieldInterface.(bool))
					case reflect.TypeOf(&time.Time{}):
						if nil == objectStructFieldInterface.(*any) {
							continue
						}
						datetime, _ := time.Parse("2006-01-02 15:04:05", (*(objectStructFieldInterface.(*any))).(string))
						targetStructField.Set(reflect.ValueOf(any(&datetime)))
					default:
						targetStructField.Set(objectStructField)
					}
				} else {
					targetStructField.Set(objectStructField)
				}
			}
		}
	}
}
