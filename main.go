package main

import (
	"fmt"
)

func main() {
  var name string

  fmt.Scanln(&name)

  fmt.Println("name :", name)

  age := 123
  fmt.Println(age)

  for i := 1; i < 5; i++ {
    fmt.Println(i)
  }

  for age < 124 {
    age++
    fmt.Println("hi")
  }

  for {
    fmt.Println("infinity")
    if age == 124 {
      age++
      continue
    }
    break
  }
  
  오종진()
}

func 오종진() {
  fmt.Println("나는 완도사람입니다")
}
