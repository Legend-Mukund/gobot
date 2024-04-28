package helper

import (
	"encoding/json"

	"github.com/MukundSinghRajput/gorunner"
)

type Output struct {
	CompileErrors string `json:"compile_errors"`
	Output        string `json:"output"`
}

func Runner(code string) (Output, error) {
	res, err := gorunner.RunCode(code, "true")
	if err != nil {
		return Output{"", ""}, err
	}
	var jsonData Output
	err = json.Unmarshal([]byte(res), &jsonData)
	if err != nil {
		return Output{"", ""}, err
	}
	return jsonData, nil
}
