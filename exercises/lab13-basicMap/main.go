package main

import "fmt"

// map
// ref-type
// auto increase capacity (len)
// value cases: struct (recommend)

func main() {
	// 3 students name & gender
	// map[sid]map[props:{name, gender}]string
	// init
	students := make(map[string]map[string]string, 3)
	initMap(students)

	fmt.Println(students)

	// update
	updateOrInsert(students, "no1", "name", "Jack-update")
	updateOrInsert(students, "no4", "name", "Savy")

	fmt.Println(students)

	// del
	del(students, "no1")
	fmt.Println(students)

	// print
	printMap(students)
}

func initMap(students map[string]map[string]string) {
	// init layer2
	students["no1"] = make(map[string]string, 2)
	students["no1"]["name"] = "Jack"
	students["no1"]["gender"] = "male"

	students["no2"] = make(map[string]string, 2)
	students["no2"]["name"] = "Alice"
	students["no2"]["gender"] = "female"

	students["no3"] = make(map[string]string, 2)
	students["no3"]["name"] = "Sandy"
	students["no3"]["gender"] = "female"
}

func updateOrInsert(students map[string]map[string]string, noId, attr, value string) {
	// if students[noId] != nil ...
	if m, ok := students[noId]; ok {
		if _, ok := m[attr]; ok {
			m[attr] = value
		}
	} else {
		students[noId] = make(map[string]string)
		students[noId][attr] = value
	}
}

func del(students map[string]map[string]string, noId string) {
	delete(students, noId)

	// remove all
	// delete by key for-loop
	// assign to new make(map[]...)
}

func printMap(students map[string]map[string]string) {
	for no, vs := range students {
		fmt.Printf("%s = ", no)
		for attr, value := range vs {
			fmt.Printf("{%s: %s} ", attr, value)
		}
		fmt.Println()
	}
}
