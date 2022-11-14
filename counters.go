package sider

var counter = make(map[string]int64) 

func Incr(key string) int64 {
	counter[key]++
	return counter[key]
}

func Decr(key string) int64 {
	counter[key]--
	return counter[key]
}

