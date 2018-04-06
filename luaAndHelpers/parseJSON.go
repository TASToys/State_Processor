package luaAndHelpers

import (
	"encoding/json"
	"strconv"
	"strings"
	//"fmt"
)

//GetJSONPiece gets the location as split by ":" from the JSON.
func GetJSONPiece(JSON, location string) string {

	stuffs := strings.Split(location, ":")

	//fmt.Printf("Stuffs: %v\n", stuffs)
	var output map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &output)
	if err != nil {
		return ""
	}
	return getElementFromJSON(output, stuffs, 0)

}

//gets elements from a json object.
func getElementFromJSON(input map[string]interface{}, args []string, location int) (output string) {
	switch input[args[location]].(type) {
	case string:
		return input[args[location]].(string)
	case float64:
		return strconv.FormatFloat((input[args[location]].(float64)), 'f', -1, 64)
	case bool:
		if input[args[location]].(bool) {
			return "true"
		}
		return "false"
	case map[string]interface{}:
		return getElementFromJSON((input[args[location]].(map[string]interface{})), args, location+1)
	case []interface{}:
		return getElementFromJSONArray((input[args[location]].([]interface{})), args, location+1)
	}
	return ""
}

//gets elements from a JSON array.
func getElementFromJSONArray(input []interface{}, args []string, location int) (output string) {
	val, err := strconv.Atoi(args[location])
	if err != nil {
		return ""
	}
	switch input[val].(type) {
	case string:
		return input[val].(string)
	case float64:
		return strconv.FormatFloat((input[val].(float64)), 'f', -1, 64)
	case bool:
		if input[val].(bool) {
			return "true"
		}
		return "false"
	case map[string]interface{}:
		return getElementFromJSON((input[val].(map[string]interface{})), args, location+1)
	case []interface{}:
		return getElementFromJSONArray((input[val].([]interface{})), args, location+1)
	}
	return ""
}
