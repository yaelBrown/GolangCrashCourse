package main

// methods to print stuff on the screen
// Example of how to import multiple packages
import (
  "fmt"
  "math"
  "github.com/yaelBrown/go_crash_course/03_packages/strutil" // use full domain name to import pages you created
)

func main() {
 
  fmt.Println(math.Floor(2.5))
  fmt.Println(math.Ceil(2.6))
  
}