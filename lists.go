package sider

import (
	"fmt"
)

var lists = make(map[string][]string)

func ReadList(listName string) ([]string, error) {
	if !isList(listName) {
		return []string{}, fmt.Errorf("'%v' is not a list", listName)
	}
	return lists[listName], nil
}

func LPush(key string, value string) (bool, error) {
	if len(key) == 0 {
		return false, fmt.Errorf("error pushing '%v' at left", key)
	}
	var slice = []string{value}
	lists[key] = append(slice, lists[key]...)
	return true, nil
}

func RPush(key string, value string) (bool, error) {
	if len(key) == 0 {
		return false, fmt.Errorf("error pushing '%v' at right", key)
	}
	var slice = []string{value}
	lists[key] = append(lists[key], slice...)
	return true, nil
}

func Pop(options ...string) (string, error) {
	var key string
	var list []string
	var popped string

	if len(options) == 0 {
		return "", fmt.Errorf("i need more parameters")
	}

	key = options[0]

	if _, err := LLen(key); err != nil {
		return "", err
	}

	list, err := ReadList(key)

	if err != nil {
		return "", err
	}

	if len(options) == 1 {
		popped = list[len(list)-1]
		lists[key] = append([]string{}, list[:len(list)-1]...)
		return popped, nil
	}

	var direction string = options[1]

	if direction == "left" {
		popped = list[0]
		lists[key] = append([]string{}, list[1:]...)
		return popped, nil
	}

	return "", fmt.Errorf("error while pop list '%v' at '%v'", key, direction)
}

func LLen(key string) (int, error) {
	if !isList(key) {
		return 0, fmt.Errorf("'%s' is not a list", key)
	}
	return len(lists[key]), nil
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
