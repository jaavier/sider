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
sider.AddKey(key string, value string) bool
```

# Read Key
```golang
sider.ReadKey(key string) interface{}
```

# Push an item at left
```golang
sider.LPush(listName string, value string) bool
```

# Push an item at right
```golang
sider.RPush(listName string, value string) bool
```

# Get length of a list
```golang
sider.LLen(listName string) int
```

# Get index of element in list
```golang
sider.IndexOf(listName string, element string) (int, error)
```

# Read List
```golang
sider.ReadList(listName string) []string
```

# Expire Key
```golang
sider.ExpireKey(key string, timestamp int64) bool
```

# Expire List
```golang
sider.ExpireList(listName string, timestamp int64) bool
```

# Pop list

Pop at right
```golang
sider.Pop(listName string) string
```

Pop at left
```golang
sider.Pop(listName string, "left") string
```


# TODO

- Add function to reverse lists
- Add function to find elements in lists
- Add os.Args to give the option to don't persist data

