package tests

import (
	"testing"

	"github.com/jaavier/sider"
)


func TestKeys(t *testing.T) {
	t.Run("Add Keys", func(t *testing.T) {
		var key string = "my-key"
		var value string = "my-value"
		var result bool = sider.AddKey(key, value)
		if !result {
			t.Errorf("Error adding key '%v'", key)
		}
	})

	t.Run("Read Key", func(t *testing.T) {
		var key string = "my-key"
		var expected string = "my-value"
		var result interface{} = sider.ReadKey(key)
		if result != expected {
			t.Errorf("Error reading key '%v'", key)
		}
	})

}