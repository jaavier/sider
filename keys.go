package main

var keys = make(map[string]string)

func AddKey(key, value string) bool {
	if len(key) == 0 || len(value) == 0 {
		return false
	}

	keys[key] = value

	return true
}

func ReadKey(key string) interface{} {
	for k := range keys {
		if k == key {
			return keys[key]
		}
	}
	return false
}