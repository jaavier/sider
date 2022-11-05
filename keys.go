package sider

import "fmt"

var keys = make(map[string]string)

func Set(key, value string) (bool, error) {
	if len(key) == 0 {
		return false, fmt.Errorf("value '%s' cannot be empty", key)
	}
	if len(value) == 0 {
		return false, fmt.Errorf("value '%s' cannot be empty", value)
	}
	if keys[key] != value {
		keys[key] = value
	}	
	return true, nil
}

func Get(key string) (string, error) {
	for k := range keys {
		if k == key {
			return keys[key], nil
		}
	}
	return "", fmt.Errorf("key '%s' not found", key)
}

func DeleteKey(key string) error {
	value, _ := Get(key)
	if len(value) > 0 {
		delete(keys, key)
		return nil
	}
	return fmt.Errorf("key '%s' doesn't exist", key)
}	