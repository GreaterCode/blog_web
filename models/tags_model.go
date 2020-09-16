package models

import "strings"

func HandleTagsListData(tags []string) map[string]int {
	var tagMap = make(map[string]int)

	for _, tag := range tags {
		tagsList := strings.Split(tag, "&")
		for _, value := range tagsList {
			tagMap[value]++
		}
	}
	return tagMap
}
