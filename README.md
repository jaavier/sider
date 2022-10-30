# What is Sider?
Sider is a in-memory database with persistence option. This is a personal project I'm coding to practice **Golang** by doing.

# What can Sider do?

- Read/Add keys 
- Push/Pop lists (LEFT and RIGHT)
- Expire lists/keys 
- Backup **lists** and **keys** in JSON files.
- Import last backup when starting Sider again

# How to use Sider?

1. Go to your project directory: `cd ~/my-project`
2. Install dependency with `go get -u github.com/jaavier/sider`
3. Import in your code `import "github.com/jaavier/sider"`

# Add Key
`sider.AddKey(key string, value string)`

# Read Key
`sider.ReadKey(key string)`

# Push an item at left
`sider.LPush(listName string, value string)`

# Push an item at right
`sider.RPush(listName string, value string)`

# Pop list
Pop at right
`sider.Pop(listName string)`

Pop at left
`sider.Pop(listName string, "left")`


# TODO

- Add function to reverse lists
- Add function to find elements in lists
- Add function to check TTL of lists or keys
- Add os.Args to give the option to don't persist data

