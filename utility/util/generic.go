package util

import (
	"github.com/flosch/pongo2/v6"
	"github.com/gogf/gf/v2/encoding/gjson"
)

func GetPointer[T any](v T) *T {
	return &v
}

func SqlLikeStr(str string) string {
	return "%" + str + "%"
}

func InSlice[T comparable](slice []T, v T) bool {
	for _, item := range slice {
		if item == v {
			return true
		}
	}
	return false
}

func ToPointer[T any](val T) *T {
	return &val
}

func Pongo2Parse(content string, data *gjson.Json) (result string, err error) {
	var contentTpl *pongo2.Template
	contentTpl, err = pongo2.FromString(content)
	if err != nil {
		return
	}

	result, err = contentTpl.Execute(data.Map())
	return
}
