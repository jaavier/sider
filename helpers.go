package sider

func isList(key string) bool {
	for listName, _ := range lists {
		if listName == key {
			return true
		}
	}
	return false
}
