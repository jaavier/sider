# What is Sider?
Sider is a in-memory database with persistence option. This is a personal project I'm coding to practice **Golang** by doing.

# What can Sider do?

- Read/Add keys 
- Read List (complete or by range)
- Push/Pop lists (LEFT and RIGHT)
- Get length of a list
- Get index of element in list
- Expire lists/keys 
- Backup **lists** and **keys** in JSON files.
- Import last backup when starting Sider again

# How to use Sider?

1. Go to your project directory: `cd ~/my-project`
2. Install dependency with `go get -u github.com/jaavier/sider`
3. Import in your code `import "github.com/jaavier/sider"`


# IMPORTANT UPDATE 👾

I updated all functions to return a value and a error (if there's one). If you cloned this project before, please update your dependencies executing `go get -u -d ./...` in your project's folder

# Persist data on disk
```golang
go sider.SaveData(customPath string) // execute as goroutine
```
***Important***: If you call `sider.SaveData()` without param, it will store data at /tmp

# Import data from last session
```golang
go sider.ImportData(customPath string) // execute as goroutine
```
***Important***: If you call `sider.ImportData()` without param, it will find data at /tmp

# Add Key

```golang
sider.Set(key, value string) (bool, error)
```

# Read Key
```golang
sider.Get(key string) (string, error)
```

# Push an item at left
```golang
sider.LPush(key string, value string) (bool, error)
```

# Push an item at right
```golang
sider.RPush(key string, value string) (bool, error)
```

# Get length of a list
```golang
sider.LLen(listName string) (int, error)
```

# Get index of element in list
```golang
sider.IndexOf(listName string, element string) (int, error)
```

# Replace element in list
```golang
sider.ReplaceList(listName string, index int, element string) (bool, error)
```

# Read List
```golang
sider.GetList(listName string, start string, stop string) ([]string, error)
```
**Parameters _start_ and _stop_ are optionals**

# Expire Key
```golang
sider.ExpireKey(key string, timestamp int64) (bool, error)
```

# Expire List
```golang
sider.ExpireList(listName string, timestamp int64) (bool, error)
```

# Pop list

Pop at right
```golang
sider.Pop(options ...string) (string, error)
```

Pop at left
```golang
sider.Pop(listName string, "left") (string, error)
```


# TODO

- Add function to reverse lists
