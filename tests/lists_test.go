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
		if list, err := sider.ReadList(key); err != nil && list[0] != item4 {
			t.Errorf("Error pushing left item %s", item4)
		}
	})

	t.Run("Test RPush", func(t *testing.T) {
		var key string = "my-list"
		var newItem string = "0"
		sider.RPush(key, newItem)
		if list, err := sider.ReadList(key); err != nil && list[len(list)-1] != newItem {
			t.Errorf("Error pushing right item %s", newItem)
		}
	})

	t.Run("Read List", func(t *testing.T) {
		var key string = "my-list"
		var expected []string = []string{"4", "3", "2", "1"}
		if list, err := sider.ReadList(key); err != nil {
			t.Errorf("Error reading list %s", key)
		} else {
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
		if result, err := sider.Pop(key); err != nil && result != "0" {
			t.Errorf("Error pop right item %s", key)
		}
		if result, err := sider.Pop(key); err != nil && result != "1" {
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

	t.Run("Test IndexOf", func(t *testing.T){
		var key string = "new-list-index-of"
		var element1 string = "1"
		var element2 string = "2"
		var element3 string = "3"
		var expected int = 0
		sider.LPush(key, element1)
		sider.LPush(key, element2)
		sider.LPush(key, element3)
		if result, err := sider.IndexOf(key, element3); err != nil && result != expected {
			t.Errorf("Error getting index %v", err)
		}
	})
}
