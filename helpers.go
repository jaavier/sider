package sider

import "fmt"

func isList(key string) bool {
	for listName, _ := range lists {
		if listName == key {
			return true
		}
	}
	return false
}

func isKey(key string) bool {
	for k, _ := range lists {
		if k == key {
			return true
		}
	}
	return false
}

func notEnoughArguments(fields []string) {
	if len(fields) > 0 {
		fmt.Println("Next arguments are not present or they're invalid:")
		for index, field := range fields {
			fmt.Printf("\t[%d] %s\n", index+1, field)
		}
	}
}
