package util

import "encoding/json"

func UnmarshalJson[T any](data []byte) (res T, err error) {
	err = json.Unmarshal(data, &res)

	return
}
