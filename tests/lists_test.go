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
		var list []string = sider.ReadList(key)
		if list[0] != item4 {
			t.Errorf("Error pushing left item %s", item4)
		}
	})

	t.Run("Test RPush", func(t *testing.T) {
		var key string = "my-list"
		var newKey string = "0"
		sider.RPush(key, newKey)
		var list []string = sider.ReadList(key)
		if list[len(list)-1] != newKey {
			t.Errorf("Error pushing right item %s", newKey)
		}
	})

	t.Run("Read List", func(t *testing.T) {
		var key string = "my-list"
		var expected []string = []string{"4", "3", "2", "1"}
		var list []string = sider.ReadList(key)
		for index := range expected {
			if expected[index] != list[index] {
				t.Errorf("Error reading list %s", key)
			}
		}
	})

	t.Run("Test Pop Left", func(t *testing.T) {
		var key string = "my-list"
		if sider.Pop(key, "left") != "4" {
			t.Errorf("Error pop left item %s", key)
		}
		if sider.Pop(key, "left") != "3" {
			t.Errorf("Error pop left item %s", key)
		}
	})

	t.Run("Test Pop Right", func(t *testing.T) {
		var key string = "my-list"
		if sider.Pop(key) != "0" {
			t.Errorf("Error pop right item %s", key)
		}
		if sider.Pop(key) != "1" {
			t.Errorf("Error pop right item %s", key)
		}
	})
}