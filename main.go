package main // Define the main package (required for standalone Go programs)

// Import the fmt package for input and output

//////////////******Integers*******//////////////////////////////////
// func main() {
// 	var a, b int // Declare two integer variables to store user input

// 	// Prompt the user to enter the first number
// 	fmt.Print("Enter first number: ")
// 	fmt.Scanln(&a) // Read the input and store it in variable 'a'

// 	// Prompt the user to enter the second number
// 	fmt.Print("Enter second number: ")
// 	fmt.Scanln(&b) // Read the input and store it in variable 'b'

// 	sum := a + b // Add the two numbers and store the result in 'sum'

//		// Display the result
//		fmt.Println("Sum:", sum)
//	}
// func main() {

// 	var a int = -5
// 	var b uint = 10

//		fmt.Println("Signed int:", a)
//		fmt.Println("USigned int:", b)
//	}
//////////////******Complex Numbers*******//////////////////////////////////
// func main() {

// 	c1 := complex(2, 3)
// 	c2 := complex(1, 4)

// 	sum := c1 + c2
// 	diff := c1 - c2
// 	multi := c1 * c2

// 	fmt.Println("Sum =", sum)
// 	fmt.Println("Diffarntiate =", diff)
// 	fmt.Println("multiplyed", multi)

// }
// Write a Go program that:

// Declares two complex numbers:

// c1 = 3 + 4i

// c2 = 1 - 2i

// Calculates and prints the following:

// Sum of c1 and c2

// Difference of c1 and c2

// Product of c1 and c2

// Real and imaginary parts of the product

// func main() {

// 	c1 := complex(3, 4)
// 	c2 := complex(1, -2)

// 	sum := c1 + c2
// 	diff := c1 - c2
// 	prod := c1 * c2

// 	real_part := real(prod)
// 	image_part := imag(prod)

//		fmt.Println("Sum:", sum)
//		fmt.Println("Diffarance:", diff)
//		fmt.Println("Product:", prod)
//		fmt.Println("Real part Of product:", real_part)
//		fmt.Println("Imaganariy part Of product:", image_part)
//	}

// Write a Go program that:

// Defines two boolean variables:

// isRaining (e.g., true)

// hasUmbrella (e.g., false)

// Uses logic to determine if the person can go outside safely.

// ////////////******Boolean*******//////////////////////////////////
// func main() {

// 	var israining bool = false
// 	var hasumbrella bool = true

// 	fmt.Println("Is Raining:", israining)
// 	fmt.Println("has Umbrella:", hasumbrella)

// 	fmt.Println("Can Go out side saftely:", !israining || hasumbrella)

// }
// ////////////******Strings (Raw , Interpeted)*******//////////////////////////////////

// func main() {

// 	raw := `Go is a Fun language it is the best language i ever learned so much `

//		fmt.Print(raw)
//	}
// func main() {

// 	var1 := "Hello Sadumina, Welcome to Go!\nLearn with fun"
// 	var2 := `Hello Sadumina,Welcome to Go!\nLearn with fun`

// 	fmt.Println("Interped :", var1)
// 	fmt.Print("Raw :", var2)
// }

// ////////////******Type conversions*******//////////////////////////////////

// func main() {

// 	var i int = 42
// 	var f float64 = float64(i)
// 	var u uint = uint(i)

// 	fmt.Println("Integer:", i)
// 	fmt.Println("Float64:", f)
// 	fmt.Println("Unit:", u)
// }

// Write a Go program that:

// Declares two variables:

// a as an int with value 10

// b as a float64 with value 5.5

// Converts a to float64 and adds it to b

// Converts b to int and subtracts it from a

// Prints both results with explanatory messages

// func main() {

// 	var a int = 10
// 	var b float64 = 5.5

// 	fmt.Println("Sum of float64(a) + b:", float64(a)+b)
// 	fmt.Println("Difference of a - int(b):", a-int(b))

// }

// ////////////******Scope and Shadowing*******//////////////////////////////////
//Scope
// package main

// import "fmt"

// var packageVar = "I am package scoped" // package scope

// func main() {
//     funcVar := "I am function scoped" // function scope

//     if true {
//         blockVar := "I am block scoped" // block scope
//         fmt.Println(blockVar)           // Accessible here
//     }

//     // fmt.Println(blockVar) // Error: blockVar undefined here

//     fmt.Println(funcVar)        // Accessible here
//     fmt.Println(packageVar)     // Accessible here
// }

//Shadowing
// package main

// import "fmt"

// var x = 10 // package scope

// func main() {
//     x := 20 // shadows package variable x

//     fmt.Println(x) // prints 20, not 10

//     {
//         x := 30 // shadows x in main
//         fmt.Println(x) // prints 30
//     }

//     fmt.Println(x) // prints 20 again
//     fmt.Println(main.x) // invalid — cannot access outer variable through package in this way
// }
