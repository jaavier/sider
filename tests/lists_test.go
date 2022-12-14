package tests

import (
	"testing"

	"github.com/jaavier/sider"
)

func TestLists(t *testing.T) {
	t.Run("Test LPush", func(t *testing.T) {
		var key string = "my-list"
		var item1 string = "1"
		var item2 string = "2"
		var item3 string = "3"
		var item4 string = "4"
		sider.LPush(key, item1)
		sider.LPush(key, item2)
		sider.LPush(key, item3)
		sider.LPush(key, item4)
		if list, err := sider.GetList(key); err != nil && list[0] != item4 {
			t.Errorf("Error pushing left item %s", item4)
		}
	})

	t.Run("Test RPush", func(t *testing.T) {
		var key string = "my-list"
		var newItem string = "0"
		sider.RPush(key, newItem)
		if list, err := sider.GetList(key); err != nil && list[len(list)-1] != newItem {
			t.Errorf("Error pushing right item %s", newItem)
		}
	})

	t.Run("Read List", func(t *testing.T) {
		var key string = "my-list"
		var expected []string = []string{"4", "3", "2", "1"}
		if list, err := sider.GetList(key); err != nil {
			t.Errorf("Error reading list %s", key)
		} else {
			for index := range expected {
				if expected[index] != list[index] {
					t.Errorf("Error reading list %s", key)
				}
			}
		}
	})

	t.Run("Read List with start", func(t *testing.T) {
		var key string = "my-list-start"
		var elements []string = []string{"4", "3", "2", "1"}
		var expected []string = []string{"3", "2", "1"}
		var start string = "1"
		for _, e := range elements {
			sider.RPush(key, e)
		}
		if list, err := sider.GetList(key, start); err != nil {
			t.Errorf("Error reading list %s", key)
		} else {
			if len(list) != len(expected) {
				t.Errorf("Error reading list '%s' with start '%s'", key, start)
			}
			for index := range expected {
				if expected[index] != list[index] {
					t.Errorf("Error reading list %s", key)
				}
			}
		}
	})

	t.Run("Read List with start and stop", func(t *testing.T) {
		var key string = "my-list-start-stop"
		var elements []string = []string{"4", "3", "2", "1"}
		var expected []string = []string{"3", "2"}
		var start string = "1"
		var stop string = "2"
		for _, e := range elements {
			sider.RPush(key, e)
		}
		if list, err := sider.GetList(key, start, stop); err != nil {
			t.Errorf("Error reading list %s", key)
		} else {
			if len(list) != len(expected) {
				t.Errorf("Error reading list '%s' with start '%s' and stop '%s", key, start, stop)
			}
			for index := range expected {
				if expected[index] != list[index] {
					t.Errorf("Error reading list %s", key)
				}
			}
		}
	})

	t.Run("Test Pop Left", func(t *testing.T) {
		var key string = "my-list"
		if result, err := sider.Pop(key, "left"); err != nil && result != "4" {
			t.Errorf("Error pop left item %s", key)
		}
		if result, err := sider.Pop(key, "left"); err != nil && result != "3" {
			t.Errorf("Error pop left item %s", key)
		}
	})

	t.Run("Test Pop Right", func(t *testing.T) {
		var key string = "my-list"
		if result, err := sider.Pop(key, "right"); err != nil && result != "0" {
			t.Errorf("Error pop right item %s", key)
		}
		if result, err := sider.Pop(key, "right"); err != nil && result != "1" {
			t.Errorf("Error pop right item %s", key)
		}
	})

	t.Run("Test Get List Length", func(t *testing.T) {
		var key string = "elements"
		var expected int = 2
		sider.LPush(key, "item1")
		sider.LPush(key, "item2")
		if result, err := sider.LLen(key); err != nil && result != expected {
			t.Errorf("Error get length")
		}
	})

	t.Run("Test IndexOf", func(t *testing.T) {
		var key string = "new-list-index-of"
		var element1 string = "1"
		var element2 string = "2"
		var element3 string = "3"
		var expected int = 2
		sider.LPush(key, element1)
		sider.LPush(key, element2)
		sider.LPush(key, element3)
		if result, err := sider.IndexOf(key, element1); err != nil && result != expected {
			t.Errorf("Error getting index %v", err)
		}
	})

	t.Run("Test Replace List", func(t *testing.T) {
		var listName string = "replace-list"
		var expected string = "jaavier"
		sider.LPush(listName, "1")
		sider.LPush(listName, "2")
		sider.LPush(listName, "3")
		sider.LPush(listName, "4")
		sider.LPush(listName, "5")
		if _, err := sider.ReplaceList(listName, 4, expected); err != nil {
			t.Errorf("Error replacing %v", err)
		}
		if res, err := sider.Pop(listName, "right"); err != nil {
			t.Errorf("Error getting list after replace %v", err)
		} else {
			if res != expected {
				t.Errorf("Error replacing element. %s is different from %s", res, expected)
			}
		}
	})

	t.Run("Test Count List", func(t *testing.T) {
		var listName string = "countable"
		var expected int = 3
		sider.RPush(listName, "1 item")
		sider.RPush(listName, "2 item")
		sider.RPush(listName, "3 item")
		if count, err := sider.CountList(listName); err != nil {
			t.Errorf("Error CountList function %v", err)
		} else {
			if count != expected {
				t.Errorf("Value count %v != %v", count, expected)
			}
		}
	})

	t.Run("Delete List", func(t *testing.T) {
		var newList string = "my-list-to-delete"
		sider.LPush(newList, "element")
		if err := sider.DeleteList(newList); err != nil {
			t.Errorf("error: %v", err)
		}
	})

	t.Run("Delete item from list by content", func(t *testing.T) {
		var listName string = "tmp-list-with-item"
		var newItem string = "delete-me"
		sider.RPush(listName, newItem)
		if sider.DeleteItemByContent(listName, newItem) == false {
			t.Errorf("Error deleting item: %v", newItem)
		}
		if result, _ := sider.GetList(listName); len(result) != 0 {
			t.Errorf("[2] Error deleting item: %s", newItem)
		}
	})

	t.Run("Delete item from list by index", func(t *testing.T) {
		var listName string = "delete-list-by-index"
		var newItem string = "delete-item"
		sider.RPush(listName, newItem)
		if index, err := sider.IndexOf(listName, newItem); err != nil {
			t.Errorf("error finding index for item: %s", newItem)
		} else {
			if sider.DeleteItemByIndex(listName, index) == false {
				t.Errorf("[2] Error deleting item by index: %s", newItem)
			}
			if result, _ := sider.GetList(listName); len(result) != 0 {
				t.Errorf("[3] Error deleting item: %s", newItem)
			}	
		}
	})
}
