package main

type Person struct {
	Name string
	Age int
}

func main() {
	// Basics
	// var name string = "Toto"
	// fmt.Printf("Hello, how are you %s\n", name)

	// age := 27
	// fmt.Printf("this is your age %d\n", age)

	// var city string
	// city = "Paris"
	// fmt.Printf("This is your city %s\n", city)

	// var country, continent string = "France", "Eurasia"
	// fmt.Printf("this is the country %s and this is the  continent %s\n", country, continent)

	// var (
	// 	isEmployed bool = true
	// 	salary int			= 50000
	// 	position string = "developer"
	// )
	// fmt.Printf("is employed : %t this is my salary %d and this is my position %s\n", isEmployed, salary, position)

	// //Zero Values
	// var defaultInt int
	// var defaultFloat float64
	// var defaultString string
	// var defaultBool bool

	// fmt.Printf("int : %d float %f string %s and bool %t\n", defaultInt, defaultFloat, defaultString, defaultBool)

	// const pi = 3.14
	// const (
	// 	Monday = 1
	// 	Tuesday = 2
	// 	Wednesday = 3
	// )

	// fmt.Printf("Monday: %d - Tuesday %d - Wednesday %d\n", Monday, Tuesday, Wednesday)

	// const typedAge int = 25
	// const untypedAge = 25

	// const (
	// 	Jan = iota + 1 // 1
	// 	Feb // 2
	// 	Mar // 3
	// 	Apr // 4
	// )

	// fmt.Printf("jan - %d feb - %d - mar %d - apr %d\n", Jan, Feb, Mar, Apr)

	// result := add(3, 4)
	// fmt.Printf("this is the result %d\n", result)

	// sum, product := calculateSumAndProduct(10, 10)
	// fmt.Printf("sum is %d, product is %d\n", sum, product)

	// Conditionnal statements
	// age := 30

	// if age >= 18 {
	// 	fmt.Println("You're an adult")
	// } else if (age >= 13) {
	// 	fmt.Println("You're a teenager")
	// } else {
	// 	fmt.Println("You're a child")
	// }

	// day := "Tuesday"
	// switch day {
	// case "Monday":
	// 	fmt.Println("start of the week")
	// case "Tuesday", "Wednesday", "Thursday":
	// 	fmt.Println("midweek")
	// case "Friday":
	// 	fmt.Println("tgif")
	// default:
	// 	fmt.Println("it's the weekend")
	// }

	// Loops
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("this is i", i)
	// }

	// counter := 0
	// for counter < 3 {
	// 	fmt.Println("this is the counter", counter)
	// 	counter++
	// }

	// // also
	// iterations := 0
	// for {
	// 	if iterations > 3 {
	// 		//some condition is met
	// 		break
	// 	}
	// 	iterations++
	// }

	// Array and Slices
		// numbers := [5]int{10, 20, 30, 40, 50} // array can't change, we can't do numbers[5] = 60 for example
		// fmt.Printf("this is the array %v\n", numbers)
		// fmt.Println("this the last value", numbers[len(numbers)-1])

		// matrix := [2][3]int{
		// 	{1, 2, 3},
		// 	{4, 5, 6},
		// }
		// fmt.Printf("this is the matrix: %v\n", matrix)

		// // Slices
		// // allNumbers := numbers[:]
		// // firstThree := numbers[0:3]

		// fruits := []string{"apple", "banana", "strawberry"}
		// fmt.Printf("these are my fruits %v\n", fruits)
		
		// fruits = append(fruits, "kiwi")
		// fmt.Printf("fruits with kiwi %v\n", fruits)

		// for index, value := range numbers {
		// 	fmt.Printf("index %d and value %d\n", index, value)
		// }

		// Maps
		// capitalCities := map[string]string {
		// 	"USA": "Washington D.C.",
		// 	"India": "New Delhi",
		// 	"UK": "London", 
		// }
		// fmt.Println(capitalCities["USA"])

		// capital, exists := capitalCities["Germany"]
		// if (exists) {
		// 	fmt.Println("this is the capital", capital)
		// } else {
		// 	fmt.Println("doesn't exist")
		// }

		// delete(capitalCities,"UK")
		// fmt.Printf("this is the new deleted map %v\n", capitalCities)


		//Struct
		// person := Person{Name: "John", Age: 25}
		// fmt.Printf("this is the person %+v\n", person)

		// // anonymous struct
		// employee := struct {
		// 	name string
		// 	id int
		// }{
		// 	name: "Alice",
		// 	id: 123,
		// }
		
		// fmt.Println("this is employee", employee)

		// type Address struct {
		// 	Street string
		// 	City string
		// }

		// type Contact struct {
		// 	Name string
		// 	Address Address
		// 	Phone string
		// }

		// contact := Contact{
		// 	Name: "Marc",
		// 	Address: Address{
		// 		Street: "123 main street",
		// 		City: "Town",
		// 	},
		// }

		// fmt.Println("this is contact", contact)

		// fmt.Println("name before", person.Name)
		// modifyPersonName(&person)
		// fmt.Println("name after", person.Name)

		// x := 20
		// ptr := &x // pointer (address) of x
		// fmt.Printf("value of x: %d and address of x %p\n", x, ptr)
		// *ptr = 30
		// fmt.Printf("value of new x: %d and address of x %p\n", x, ptr)
}
// person is a pointer when using *Person instead of Person
// func modifyPersonName(person *Person) { 
// 	person.Name = "Mickey"
// 	fmt.Println("inside scope: name after", person.Name)

// }

// func add(a int, b int) int {
// 	return a + b
// }

// func calculateSumAndProduct(a, b int) (int, int) {
// 	return a + b, a * b
// }