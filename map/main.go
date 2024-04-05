package main

import (
	"fmt"
)

func main() {
	var myMap map[string]int = make(map[string]int);

	myMap["one"] = 1;
	myMap["two"] = 2;

	fmt.Println(myMap); 

	delete(myMap, "one");

	for key, val := range myMap {
		fmt.Println(key, val);
	}

	val, ok := myMap["one"];

	if ok {
		fmt.Println(val);
	} else {
		fmt.Println("Key not found");
	}
}