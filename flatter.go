package flatter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Convert function to convert json string to map[string]interface{}
// For nested map, will treated to become flat structured
// example: { "additional_parameters": ["amount":100 ], "opts": [ 1, 2 ] }
// result: map["additional_parameters.amount"]= 100, map["opts[0]"]= 1, map["opts[1]"]= 2
func Flatter(jsonStr string) (map[string]interface{}, error) {
	var mapTemplate interface{}

	d := json.NewDecoder(strings.NewReader(jsonStr))
	d.UseNumber()

	if err := d.Decode(&mapTemplate); err != nil {
		return nil, err
	}

	msgMap := mapTemplate.(map[string]interface{})
	newMap := make(map[string]interface{})
	err := flatMap("", newMap, msgMap)
	if err != nil {
		return nil, err
	}

	return newMap, nil
}

func flatMap(prefix string, cleanMap, dirtyMap map[string]interface{}) error {
	for k, v := range dirtyMap {
		// modify key to be flatten
		idx := k
		if prefix != "" {
			idx = fmt.Sprintf("%s.%s", prefix, k)
		}

		err := cast(idx, v, cleanMap)
		if err != nil {
			return err
		}
	}

	return nil
}

func flatArr(prefix string, cleanMap map[string]interface{}, arr []interface{}) error {
	for k, v := range arr {
		// modify key to be flatten
		idx := fmt.Sprintf("[%d]", k)
		if prefix != "" {
			idx = fmt.Sprintf("%s[%d]", prefix, k)
		}

		err := cast(idx, v, cleanMap)
		if err != nil {
			return err
		}
	}
	return nil
}

func cast(idx string, val interface{}, cleanMap map[string]interface{}) error {
	switch vv := val.(type) {
	case string, bool:
		cleanMap[idx] = val

	case json.Number:
		// convert to respective data type
		intVal, err := vv.Int64()
		if err == nil {
			cleanMap[idx] = intVal
			return nil
		}

		floatVal, err := vv.Float64()
		if err == nil {
			cleanMap[idx] = floatVal
			return nil
		}

		return err

	case map[string]interface{}:
		return flatMap(idx, cleanMap, vv)

	case []interface{}:
		return flatArr(idx, cleanMap, vv)

	default:
		return fmt.Errorf("unknown type of %+v", vv)
	}

	return nil
}
