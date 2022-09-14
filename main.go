package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
  // qualification(2, 2)
  // qualification(2, 4)
  // qualification(2, 5)

  // Defer()

  // fmt.Println("hi")
  // 패닉()
  // fmt.Println("hello")

  test := new(최형우)

  fmt.Println(test)
}

func StandardIo() {
  var f *os.File
  f = os.Stdin
  defer f.Close()

  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    fmt.Println(">", scanner.Text())
  }
}

func repeat() {
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
}

func qualification(a float64, b float64) {
  if c := math.Pow(a, b); c < 10 {
    fmt.Println("hi", c)
  } else if c < 20 {
    백승민(c)
  } else {
    오종진(c)
  }
}

func 백승민(c float64) {
  fmt.Println("나는 백승민이고 강진을 좋아합니다", c)
}

func 오종진(c float64) {
  fmt.Println("나는 오종진이고 완도사람입니다", c)
}

func Defer() {
  defer fmt.Println("World")
  fmt.Println("Hello")
}

func 패닉() {
  defer recover()
  fmt.Println("전")
  panic("으악")
}

type 최형우 struct {
  id string
  안경 bool
  age int
}
