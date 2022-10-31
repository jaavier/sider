package sider

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
