package sider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

var path = "/tmp"

func SaveFile(file string, content interface{}) bool {
	parse, _ := json.Marshal(content)
	if len(parse) > 0 {
		ioutil.WriteFile(file, parse, 0644)
	}
	return true
}

func ImportData(customPath ...string) bool {
	if len(customPath) > 0 {
		if len(customPath[0]) > 0 {
			path = customPath[0]
		}
	}
	listsContent, errLists := ioutil.ReadFile(fmt.Sprintf("%s/lists.json", path))
	keysContent, errKeys := ioutil.ReadFile(fmt.Sprintf("%s/keys.json", path))

	if errLists != nil || errKeys != nil {
		return false
	}

	json_lists := make(map[string][]string)
	json_keys := make(map[string]string)

	json.Unmarshal(listsContent, &json_lists)
	json.Unmarshal(keysContent, &json_keys)

	lists = json_lists
	keys = json_keys

	return true
}

func SaveData(customPath ...string) { // execute as goroutine
	if len(customPath) > 0 {
		if len(customPath[0]) > 0 {
			path = customPath[0]
		}
	}
	fmt.Printf("Store data at folder '%s'\n", path)
	for range time.Tick(10 * time.Second) {
		SaveFile(fmt.Sprintf("%s/lists.json", path), lists)
		SaveFile(fmt.Sprintf("%s/keys.json", path), keys)
	}
}
