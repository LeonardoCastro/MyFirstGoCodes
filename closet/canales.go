package main

import(
        "fmt"
        "strings"
        //"math"
)

func SearchColor(c chan string, Clothes []string, color1 string, color2 string) {
  for _, cloth := range(Clothes) {
    if strings.Contains(cloth, color1) || strings.Contains(cloth, color2) {
      c <- cloth
    }
  }
}

func main() {

  color1 := "blue"
  color2 := "red"

  Cover := make(chan string)
  //Top := make(chan string)
  //Bottom := make(chan string)

  Array1 := []string{"blue sweater", "red sweater"}
  Array2 := []string{"pink tshirt", "blue shirt"}
  Array3 := []string{"blue jeans", "blue pants", "green olive cargo pants"}

  for _, c := range(Array1) {
    if strings.Contains(c, color1) {
      fmt.Println("yey")
    }
  }

   go SearchColor(Cover, Array1, color1, color2)
   go SearchColor(Cover, Array2, color1, color2)
   go SearchColor(Cover, Array3, color1, color2)
   //, len(Top), len(Bottom)
   fmt.Println(len(Cover))
  //min1 := math.Min( float64(len(Cover)), float64(len(Top)))
  //min2 := int( math.Min(min1, float64(len(Bottom))))

  for i := 0; i < 5; i ++ {
    x := <-Cover
    fmt.Println(x)
  }
}
