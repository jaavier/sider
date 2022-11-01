# What is Sider?
Sider is a in-memory database with persistence option. This is a personal project I'm coding to practice **Golang** by doing.

# What can Sider do?

- Read/Add keys 
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
sider.Pop(options... string) (string, error)
```

Pop at left
```golang
sider.Pop(listName string, "left") (string, error)
```


# TODO

- Add function to reverse lists
- Add os.Args to give the option to don't persist data
- Add function to replace element in list (LSET)