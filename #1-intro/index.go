/*
* This Go program prints the message "Hello, World!" to the terminal.
* This program is a basic example for beginners.
*
* Basic steps:
* 1. Define the main package with `package main`.
* 2. Call the formatting library with `import "fmt"`.
* 3. Create the `func main()` function — the starting point of Go programs.
* 4. "var" is a special keyword for creating a new variable.
* 5. world is the name of this variable to always call it, and string is the type name indicating that the value of this variable is of type text
* 6. Print text to the terminal with `fmt.Println()`.
*
* We used fmt in this program because fmt is a special package for input/output operations.
* Actually this program is very simple, you can try to make this program yourself.
*
* Find answers to specific questions about the program:
* 1. Why are we calling and using the main package in this program?
* 2. What is the difference println() and fmt.Println()
* 3. Can we change the main function name to myMain?
* 4. Can we change the type of world variable to bool?
*
 */

package main

import "fmt"
 
 func main() {
	 var world string = "World"
	 fmt.Println("Hello", world)
 }
 