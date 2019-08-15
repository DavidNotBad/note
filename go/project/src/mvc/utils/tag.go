package utils

import (
	"errors"
	"reflect"
	"strings"
)

//GetFieldNameByTagValue 根据结构体tag的值来获取结构体字段名
func GetFieldNameByTagValue(structName interface{}, tagName string, tagVal string) (fieldName []string, err error) {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		err = errors.New("type error: not a struct")
		return
	}
	fieldNum := t.NumField()

	for i := 0; i < fieldNum; i++ {
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			for j := 0; j < len(tags) / 2; j++ {
				structTagName := strings.TrimRight(strings.TrimSpace(tags[j * 2]), ":")
				structTagVal := strings.TrimSpace(tags[j * 2 + 1])
				if (structTagName == tagName) && (tagVal == structTagVal) {
					fieldName = append(fieldName, t.Field(i).Name)
				}
			}
		}
	}

	err = errors.New("no matching field name")
	return
}
