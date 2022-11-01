package sider

import (
	"fmt"
	"strconv"
)

var lists = make(map[string][]string)

func GetList(options ...string) ([]string, error) {
	var listName string

	if len(options) == 0 {
		return nil, fmt.Errorf("not enough arguments")
	}

	listName = options[0]

	if !isList(listName) {
		return []string{}, fmt.Errorf("'%v' is not a list", listName)
	}

	switch {
	case len(options) == 2:
		if start, err := strconv.Atoi(options[1]); err == nil {
			return lists[listName][start:], nil
		}
	case len(options) == 3:
		if start, errStart := strconv.Atoi(options[1]); errStart == nil {
			if stop, errStop := strconv.Atoi(options[2]); errStop == nil {
				if stop > len(lists[listName]) {
					stop = len(lists[listName])
				} else {
					stop += 1
				}
				return lists[listName][start:stop], nil
			}
		}
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

	list, err := GetList(key)

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
