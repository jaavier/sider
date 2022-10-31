package sider

import (
	"encoding/json"
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

func ImportData() bool {
	listsContent, errLists := ioutil.ReadFile(pwd + "/db/lists.json")
	keysContent, errKeys := ioutil.ReadFile(pwd + "/db/keys.json")

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

func SaveData() {
	for {
		time.Sleep(2 * time.Second)
		SaveFile(pwd+"/db/lists.json", lists)
		SaveFile(pwd+"/db/keys.json", keys)
	}
}
