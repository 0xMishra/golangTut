package main

import (
	"fmt"
	"maps"
)

func mapsTut() {
	// var m map[keyType]valueType

	// var m = make(map[keyType]valueType)

	// mapV := map[keyType]valueType{
	// key1:value1
	// key2:value2
	//}

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["key2"] = 19
	myMap["key3"] = 9
	myMap["key4"] = 7
	myMap["code"] = 10
	fmt.Println(myMap)

	fmt.Println(myMap["code"])
	fmt.Println(myMap["nonExistingKey"])

	_, ok := myMap["key1"]
	fmt.Println(ok)

	delete(myMap, "key1")
	clear(myMap)
	fmt.Println(myMap)

	myMap2 := map[string]int{"a": 1}

	if maps.Equal(myMap, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	}

	for k, v := range myMap2 {
		fmt.Println(k, v)
	}
}
