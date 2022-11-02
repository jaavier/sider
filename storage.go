package sider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var pwd, _ = os.Getwd()

func SaveFile(file string, content interface{}) bool {
	parse, _ := json.Marshal(content)
	if len(parse) > 0 {
		ioutil.WriteFile(file, parse, 0644)
	}
	return true
}

func ImportData(path string) bool {
	if len(path) > 0 {
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
	} else {
		fmt.Println("Path not specified = don't import data")
		return false
	}
}

func SaveData(path string) { // goroutine
	if len(path) > 0{
		for {
			time.Sleep(2 * time.Second)
			SaveFile(fmt.Sprintf("%s/lists.json", path), lists)
			SaveFile(fmt.Sprintf("%s/keys.json", path), keys)
		}	
	} else {
		fmt.Println("Path not specified = don't save data")
	}
}
