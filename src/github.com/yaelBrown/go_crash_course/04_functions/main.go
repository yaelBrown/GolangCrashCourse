package main

// methods to print stuff on the screen
import "fmt"

func greeting(name string) string {
  return "Hello " + name 
}

func getSum(num1 int, num2 int) int {
  return num1 + num2
}

func main() {
  fmt.Println(getSum(3,2))
}