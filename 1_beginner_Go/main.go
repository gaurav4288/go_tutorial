package main

import (
	"fmt"
	"unicode/utf8"
)

func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

type Person struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	// variables
	var name string = "Gaurav"
	age := 23 //walrus operator
	var balance float64 = 300.458

	fmt.Printf("This is my name: %s and I'm %d year old. My balance is: %.2f\n", name, age, balance)

	var (
		isEmployed bool   = true
		salary     int64  = 50000
		position   string = "developer"
	)

	fmt.Printf("isEmployed: %t this is my salary: %d and this is my position: %s\n", isEmployed, salary, position)

	//rune
	const mypet = "🐻" //compiler doesn't yell if const is not accessed
	// for index, runeValue := range mypet {
	//     fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	// }
	fmt.Printf("constant 'name' byte length: %d\n", len(mypet))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(mypet))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", mypet)

	//Enum
	const (
		Jan = iota + 1 //1
		Feb            //2
		Mar            //3
		Apr            //4
	)

	fmt.Printf("jan - %d feb - %d mar - %d apr - %d\n", Jan, Feb, Mar, Apr)

	result := add(3, 4)
	fmt.Println("This is the result", result)

	sum, product := calculateSumAndProduct(5, 6)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)

	for i := 0; i < 5; i++ {
		fmt.Println("This is iteration number\n", i)
		// fmt.Printf("This is iteration number %d\n", i)
	}

	counter := 0
	for counter < 5 {
		// if counter == 4 {
		// 	continue
		// }
		fmt.Println("This is iteration number\n", counter)
		counter++
	}

	// numbers := [5]int{10, 20, 30, 40, 50}
	numbers := [...]int{10, 20, 30, 40, 50}
	fmt.Println("This is the array", numbers)

	n := len(numbers)
	fmt.Println("this is the last value", numbers[n-1])

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("This is the matrix", matrix)

	allNumbers := numbers[:]
	firstTwo := numbers[:2] //go slices are like views into the array, they don't copy the data, they just reference it
	//size allocated to slice doubles when we append to it, so it can accommodate more elements without needing to copy the data every time we append
	fmt.Printf("This is the length of the slice: %d and this is the capacity of the slice: %d\n", len(firstTwo), cap(firstTwo))
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Printf("This is the length of the slice: %d and this is the capacity of the slice: %d\n", len(fruits), cap(fruits))
	fruits = append(fruits, "kiwi")
	fmt.Printf("This is the length of the slice: %d and this is the capacity of the slice: %d\n", len(fruits), cap(fruits))

	moreFruits := []string{"orange", "grape"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("This is the length of the slice: %d and this is the capacity of the slice: %d\n", len(fruits), cap(fruits))
	fmt.Println("This is the first two numbers", firstTwo)
	fmt.Println("This is the slice", allNumbers)

	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	//maps
	capitalCities := map[string]string{
		"USA":    "Washington, D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
		"UK":     "London",
	}

	fmt.Println("This is the map", capitalCities)
	fmt.Printf("The capital of USA is %s\n", capitalCities["USA"])
	capital, exists := capitalCities["India"]
	if exists {
		fmt.Printf("The capital of India is %s\n", capital)
	} else {
		fmt.Println("Capital of India not found in the map")
	}

	delete(capitalCities, "UK")
	fmt.Println("This is the map after deletion", capitalCities)

	for country, capital := range capitalCities {
		fmt.Printf("Country: %s, Capital: %s\n", country, capital)
	}

	//structs
	person := Person{Name: "gaurav", Age: 24, Salary: 2500.98}
	fmt.Printf("This is out Person:%v\n", person)

	employee := struct {
		Name   string
		id	   int
	}{
		Name: "alice",
		id: 12345,
	}
	fmt.Printf("This is our employee: %v\n", employee)

	//nested struct
	type address struct {
		street string
		city   string
	}
	type Contact struct{
		Name string
		Address address
		Phone string
	}

	contact := Contact{
		Name: "Marc",
		Address: address{
			street: "123 Main St",
			city:   "Anytown",
		},
		Phone: "555-1234",
	}

	fmt.Printf("This is our contact: %v\n", contact)

	//passing by reference
	fmt.Println("Name before", person.Name)
	// modifyPersonName(&person)
	person.modifyPersonName("chris")

	x := 20
	ptr := &x
	*ptr = 30
	fmt.Printf("value of x: %d and address of x: %p\n", x, ptr)
}

// func modifyPersonName(person *Person) {
// 	person.Name = "Chris"
// 	fmt.Println("inside scope: new name:", person.Name)
// }

func (p *Person) modifyPersonName(name string) {
	p.Name = name
	fmt.Println("inside scope: new name:", p.Name)
}
