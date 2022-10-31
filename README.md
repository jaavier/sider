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


# IMPORTANT UPDATE 👾

I updated all functions to return a value and a error (if there's one). If you cloned this project before, please update your dependencies executing `go get -u -d ./...` in your project's folder

# Add Key

```golang
sider.AddKey(key, value string) (bool, error)
```

# Read Key
```golang
sider.ReadKey(key string) (string, error)
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
sider.ReadList(listName string) ([]string, error)
```

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
- Add function to find elements in lists
- Add os.Args to give the option to don't persist data

