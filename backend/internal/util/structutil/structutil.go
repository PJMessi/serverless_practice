package structutil

import "encoding/json"

func ConvertStrToStruct(strVal string, placeholderStruct interface{}) error {
    return json.Unmarshal([]byte(strVal), placeholderStruct)
}
