package sider

import (
	"fmt"
)

var lists = make(map[string][]string)

func ReadList(listName string) []string {
	return lists[listName]
}

func LPush(key string, value string) bool {
	if len(key) == 0 {
		return false
	}
	var slice = []string{value}
	lists[key] = append(slice, lists[key]...)
	return true
}

func RPush(key string, value string) bool {
	if len(key) == 0 {
		return false
	}
	var slice = []string{value}
	lists[key] = append(lists[key], slice...)
	return true
}

func Pop(options... string) interface{} {
	var key string
	var list []string
	var popped string

	if len(options) == 0 {
		return false
	}
	
	key = options[0]
	list = ReadList(key)

	if len(options) == 1 {		
		popped = list[len(list) - 1]
		lists[key] = append([]string{}, list[:len(list)-1]...)
		return popped
	}

	var direction string = options[1]

	if direction == "left" {
		popped = list[0]
		lists[key] = append([]string{}, list[1:]...)
		return popped
	}
	return false
}

func LLen(key string) int {
	if !isList(key) {
		return 0
	}
	return len(lists[key])
}

func IndexOf(listName, element string) (int, error) {
	if isList(listName) {
		for index, value := range lists[listName] {
			if value == element {
				return index, nil
			}
		}
	}	
	return 0, fmt.Errorf("error getting index of '%v'", element)
}