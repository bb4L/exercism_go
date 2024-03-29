package flatten

import (
	"reflect"
)

// Flatten flattens arbitrary nested arrays
func Flatten(data interface{}) []interface{} {
	result := []interface{}{}

	for _, d := range data.([]interface{}) {
		if d == nil {
			continue
		}
		rt := reflect.TypeOf(d)
		switch rt.Kind() {
		case reflect.Slice:
			res := Flatten(d)
			if len(res) == 0 {
				continue
			}
			result = append(result, res...)
		default:
			result = append(result, d)
		}
	}
	return result
}
