package main

import (
  "fmt"
)

func main() {
  mySlice := []int{10, 20, 30, 40, 50} // A slice of integers

  fmt.Println(mySlice)          // Output: [10 20 30 40 50]
  fmt.Println(len(mySlice))     // Length of the slice: 5
  fmt.Println(cap(mySlice))     // Capacity of the slice: 5

  // Slicing a slice
  subSlice := mySlice[1:3]      // Slice from index 1 to 2
  fmt.Println(subSlice)         // Output: [20 30]
  fmt.Println(len(subSlice))
  fmt.Println(cap(subSlice))

  // Appending to a slice
  mySlice = append(mySlice, 60) // Append 60 to the slice

  // Covert an array to a slice
  myArray := [5]int{15, 25, 35, 45, 55}

  mySlice = myArray[:]          // Convert the array to a slice
  fmt.Println(mySlice)          // Output: [10 20 30 40 50]
  mySlice = append(mySlice, 60) // Append 60 to the slice
}
