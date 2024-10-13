package main

type Profile struct {
	Name      string
	Age       int
	IsMarried bool
}

type ProfileQueryKey struct {
	Name string
	Age  int
}

// 同一个哈希值可能指向不同的数据，mapper的value为Profile的数组
var mapper = make(map[int][]Profile)

func main() {
	list := []*Profile{
		{Name: "Alice", Age: 30, IsMarried: true},
		{Name: "Bob", Age: 40},
		{Name: "Charlie", Age: 60, IsMarried: false},
		{Name: "David", Age: 80},
	}

	buildIndex(list)
	//queryData("Alice", 30)
}

func buildIndex(list []*Profile) {
	for _, profile := range list {
		currentKey := ProfileQueryKey{profile.Name, profile.Age}
		existValue := mapper[currentKey.hash()]
		existValue = append(existValue, *profile)
		mapper[currentKey.hash()] = existValue
	}
}

func (q *ProfileQueryKey) hash() int {
	return simpleHash(q.Name) + q.Age*1000000
}

func simpleHash(str string) (total int) {
	for i := 0; i < len(str); i++ {
		c := str[i]
		total += int(c)
	}
	return
}
