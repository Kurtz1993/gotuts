package main

import "fmt"

func main() {
	// The following two lines are equivalent
	// var grades map[string]float32
	grades := make(map[string]float32)

	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 67

	fmt.Println(grades)

	// Access the map by key
	TimsGrade := grades["Timmy"]

	fmt.Println(TimsGrade)

	// Remove an entry from a map
	delete(grades, "Timmy")

	fmt.Println(grades)

	// Iterate over maps
	for key, value := range(grades) {
		fmt.Printf("%s has a grade of %f\n", key, value)
	}
}
