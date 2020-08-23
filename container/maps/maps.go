package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":     "a",
		"language": "golang",
		"site":     "ok",
		"quality":  "good",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m, m2, m3)

	/// 不会保证顺序
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")

	if cName, ok := m["name"]; ok {
		fmt.Println(cName)
	} else {
		fmt.Println("not error key")
	}

	if notName, ok := m["fu"]; ok {
		fmt.Println(notName)
	} else {
		fmt.Println("key dose not exist")
	}

	fmt.Println("Deleting values")
	delete(m, "name")
	name, ok := m["name"]
	fmt.Println(name, ok)

}
