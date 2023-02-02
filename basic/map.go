/**
 * @Author: Noaghzil
 * @Date:   2023-02-02 08:07:11
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-02-02 08:27:23
 */
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// declaration
	var personSalary1 map[string]int
	if personSalary1 == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary1 = make(map[string]int)
	}

	// declaration and initialization
	personSalary2 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary2["mike"] = 9000
	fmt.Println("personSalary2:", personSalary2)
	// create and initialize a new map
	personSalary3 := make(map[string]int)
	personSalary3["steve"] = 12000
	personSalary3["jamie"] = 15000
	fmt.Println("personSalary3:", personSalary3)
	// accessing an item which is not present returns the zero value
	fmt.Println("personSalary3[\"mike\"]:", personSalary3["mike"])
	// delete an item from a map
	delete(personSalary3, "jamie")
	fmt.Println("map after deletion:", personSalary3)
	// checking whether a key is present or not
	value, ok := personSalary3["jamie"]
	if ok == true {
		fmt.Println("jamie is in the map and value is", value)
	} else {
		fmt.Println("jamie is not in the map")
	}
	//iterating over all element using for range
	for key, value := range personSalary2 {
		fmt.Printf("personSalary2[%s] = %d, ", key, value)
	}
	// equality of map
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	map2 := map1

	if reflect.DeepEqual(map1, map2) {
		fmt.Println("map1 and map2 are equal")
	} else {
		fmt.Println("map1 and map2 are not equal")
	}

}
