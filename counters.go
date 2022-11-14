package sider

var counters = make(map[string]int64) 

func Incr(key string) int64 {
	counters[key]++
	return counters[key]
}

func Decr(key string) int64 {
	counters[key]--
	return counters[key]
}

func GetCounter(key string) int64 {
	return counters[key]
}