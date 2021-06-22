package main

import (
	"fmt"
)

func main() {
	//Array
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	var nos [5]int = [...]int{3, 1, 4, 2, 5}
	fmt.Println(nos)
	fmt.Println("Size of the nos array = ", len(nos))

	//iterate an array (using the indexer)
	/*
		for idx := 0; idx < len(nos); idx++ {
			fmt.Println(nos[idx])
		}
	*/

	//iterate an array (using range construct)
	/*
		for idx, no := range nos {
			fmt.Printf("Value at index %d is %d\n", idx, no)
		}
	*/
	for _, no := range nos {
		fmt.Println(no)
	}

	//creating a copy of an array
	newNos := nos
	newNos[0] = 200
	fmt.Println(newNos, nos)

	//slice
	var names []string = []string{"Magesh", "Ramesh"}
	fmt.Println(names, len(names))

	//adding a new value
	names = append(names, "Suresh", "Girish")
	fmt.Println(names)

	newNames := []string{"Ganesh", "Rajesh"}
	names = append(names, newNames...)
	fmt.Println(names)

	//slicing
	/*
		[lo : hi] => from low to hi - 1
		[lo : ] => all the values from lo
		[ : hi] => all the values from 0 to hi - 1
		[lo : lo] => empty slice
		[lo : lo+1] => item at lo
		[:] => copy of the slice
	*/

	fmt.Println("names[1:3] => ", names[1:3])
	fmt.Println("names[3:] => ", names[3:])
	fmt.Println("names[:3] => ", names[:3])
	fmt.Println("names[3:3] => ", names[3:3])
	fmt.Println("names[3:4] => ", names[3:4])

	dupNames := names[:]
	dupNames[0] = "Ram"
	fmt.Println(names, dupNames)

	dupNames = append(dupNames, "Laksh")
	fmt.Println(names, dupNames)

	//map
	/* var cityRanks map[string]int = map[string]int{
		"Bengaluru": 5,
		"Udupi":     1,
		"Mangaluru": 3,
		"Mysuru":    2,
	} */

	/* var cityRanks = map[string]int{
		"Bengaluru": 5,
		"Udupi":     1,
		"Mangaluru": 3,
		"Mysuru":    2,
	} */

	cityRanks := map[string]int{
		"Bengaluru": 5,
		"Udupi":     1,
		"Mangaluru": 3,
		"Mysuru":    2,
	}
	fmt.Println(cityRanks, len(cityRanks))
	cityRanks["Chennai"] = 4
	fmt.Println(cityRanks, len(cityRanks))
	fmt.Println("Rank of Mysuru => ", cityRanks["Mysuru"])

	//check if key exists
	if rank, exists := cityRanks["Washington"]; exists {
		fmt.Println("Rank of Washington => ", rank)
	} else {
		fmt.Println("Washington is not ranked yet!")
	}

	//removing a key
	delete(cityRanks, "Chennai")
	fmt.Println("After removing Chennai, cityRanks => ", cityRanks)

	//iterating a map
	for city, rank := range cityRanks {
		fmt.Printf("Rank of %s is %d\n", city, rank)
	}

}
