package main

// methods to print stuff on the screen
import "fmt"

func main() {
  x := 5
  y := 10

  if x < y {
    fmt.Printf("%d is less than %d\n", x, y)
  } else {
    fmt.Printf("%d is less than %d", y, x)
  }

  // else if
  color := "red"

  if color == "red" {
    fmt.Println("color is red")
  } else if color == "blue" {
    fmt.Println("color is blue")
  } else {
    fmt.Println("color is Not blue or red")
  }
  
  // Simple hello world
  fmt.Println("Hello World")
}