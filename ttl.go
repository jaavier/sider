package main

import (
	"fmt"
	"time"
)

var ttlKeys = make(map[string]int64)
var ttlLists = make(map[string]int64)

func CheckTTL() {
	for range time.Tick(1 * time.Second){
		for keyName, timestamp := range ttlKeys {
			if time.Now().Unix() > timestamp {
				fmt.Printf("Key '%v' expired", keyName)
				delete(keys, keyName)
				delete(ttlKeys, keyName)
			}
		}	

		for listName, timestamp := range ttlLists {
			if time.Now().Unix() > timestamp {
				fmt.Printf("List '%v' expired", listName)
				delete(lists, listName)
				delete(ttlLists, listName)
			}
		}	
	}
}

// TODO: Refactor "ExpireKey" and "ExpireList" to use only 1 function for expire

func ExpireKey(keyName string, expirationTimestamp int64) bool {
	ttlKeys[keyName] = expirationTimestamp
	return true
}

func ExpireList(listName string, expirationTimestamp int64) bool {
	ttlLists[listName] = expirationTimestamp
	return true
}