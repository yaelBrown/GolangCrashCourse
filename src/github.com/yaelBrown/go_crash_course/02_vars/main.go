package main

// methods to print stuff on the screen
import "fmt"

func main() {
  // MAIN TYPES
  // string
  // bool
  // int
  // int int8 int16 int32 int64
  // uint uint8 uint16 uint32 uint64 uintptr
  // byte - alias for uint8
  // rune - alias for int32
  // float32 float64
  // complex64 comples128

  // Create a variable you have to use it. 

  // Using var (no semi colons)
  var name string = "Yael"
  var age = 32

  fmt.Println(name, age)
  
  // Identify type of variable
  fmt.Printf("%T\n", age)

  //Shorthand
  anotherName := "Brown"

  fmt.Println(anotherName)

  

}