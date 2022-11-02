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

func notEnoughArguments(errors []string) bool {
	if len(errors) > 0 {
		fmt.Println("Missing required arguments:")
		for index, e := range errors {
			fmt.Printf("\t[%d] %s\n", index+1, e)
		}
	}
	return len(errors) > 0
}
