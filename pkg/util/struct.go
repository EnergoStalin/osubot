package util

import (
	"fmt"
	"strings"

	"go.devnw.com/structs"
)

func SToMap(s any) map[string]string {
	m := map[string]string{}
	for k, v := range structs.Map(s) {
		m[strings.ToLower(k[0:1])+k[1:]] = fmt.Sprint(v)
	}

	return m
}
