package routers

import (
	"github.com/ziyeziye/framework/pkg/utli"
	"strings"
)

func splitBy(str, tag string) []string {
	return strings.Split(str, tag)
}

func splitJson(json, index, tag string) []string {
	maps, err := utli.Json2Map(json)
	if err != nil {
		return []string{}
	}

	if item, ok := maps[index]; ok {
		switch value := item.(type) {
		case string:
			return splitBy(value, tag)
		case []string:
			return value
		case []interface{}:
			strs := []string{}
			for _, val := range value {
				switch value := val.(type) {
				case string:
					strs = append(strs,value)
				}
			}
			return strs
		}
	}

	return []string{}
}
