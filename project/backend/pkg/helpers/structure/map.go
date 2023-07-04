package structure

import "encoding/json"

// Map convert a struct into map.
func Map(object any) (mapped map[string]any) {
	marshaled, _ := json.Marshal(&object)
	json.Unmarshal(marshaled, &mapped)
	return
}
