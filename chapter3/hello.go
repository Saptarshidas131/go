// package main declaration
package main

// import the fmt package for printing
import "fmt"

// this is a comment

// function main declaration, this is function that gets called on program execution
func main(){
    // print length of Hello World using length string property
    fmt.Println(len("Hello World"))
    // print the second character in byte from the string
    fmt.Println("Hello World"[1])
    // print by concatenating the strings together
    fmt.Println("Hello " + "World")
}
