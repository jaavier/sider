package sider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var path = "/tmp"

func SaveFile(file string, content interface{}) bool {
	parse, err := json.Marshal(content)
	if err != nil {
		return false
	} else {
		ioutil.WriteFile(file, parse, 0644)
		return true
	}
}

func ImportData(customPath ...string) bool {
	if len(customPath) > 0 {
		if len(customPath[0]) > 0 {
			path = customPath[0]
		}
	}
	listsContent, errLists := ioutil.ReadFile(fmt.Sprintf("%s/lists.json", path))
	keysContent, errKeys := ioutil.ReadFile(fmt.Sprintf("%s/keys.json", path))
	countersContent, errCounters := ioutil.ReadFile(fmt.Sprintf("%s/counters.json", path))

	if errLists != nil || errKeys != nil || errCounters != nil{
		return false
	}

	json_lists := make(map[string][]string)
	json_keys := make(map[string]string)
	json_counters := make(map[string]int64) 

	json.Unmarshal(listsContent, &json_lists)
	json.Unmarshal(keysContent, &json_keys)
	json.Unmarshal(countersContent, &json_counters)

	lists = json_lists
	keys = json_keys
	counters = json_counters

	return true
}

func SaveData(customPath ...string) { // execute as goroutine
	if len(customPath) > 0 {
		if len(customPath[0]) > 0 {
			path = customPath[0]
		}
	}
	fmt.Printf("Store data at folder '%s'\n", path)
	SaveFile(fmt.Sprintf("%s/lists.json", path), lists)
	SaveFile(fmt.Sprintf("%s/keys.json", path), keys)
	SaveFile(fmt.Sprintf("%s/counters.json", path), counters)
}
