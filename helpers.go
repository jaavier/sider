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

func notEnoughArguments(fields []string) bool {
	if len(fields) > 0 {
		fmt.Println("Missing required arguments:")
		for index, field := range fields {
			fmt.Printf("\t[%d] %s\n", index+1, field)
		}
	}
	return len(fields) > 0 
}
