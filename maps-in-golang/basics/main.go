package main

import "fmt"

func main() {
	fmt.Println("Let's learn about maps in golang")

	// will show error if we do this and run
	//var mp map[string]int
	//mp["oka"] = 100

	// it is slightly different then above one, and it will work fine
	var mp map[string]int
	mp = map[string]int{
		"oka": 100,
	}
	fmt.Println(mp)

	//this one is done by using make, see readme for the definition of new and make
	mp2 := make(map[string]int)
	mp2["oka"] = 100
	fmt.Println(mp2)

	val, ok := mp2["oka"]
	fmt.Println(val)
	fmt.Println(ok)

	for key, value := range mp {
		fmt.Println(key)
		fmt.Println(value)
	}

	// check every time you can different order of key-value pair, there is no guarantee of getting the same order in golang map
	mp2["new one"] = 200
	mp2["another one"] = 300
	for key, value := range mp2 {
		fmt.Println(key, value)
	}
	for key, value := range mp2 {
		fmt.Println(key, value)
	}
	for key, value := range mp2 {
		fmt.Println(key, value)
	}
}
