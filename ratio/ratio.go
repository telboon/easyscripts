package main

import "github.com/alecthomas/kingpin"
import "fmt"
import "math"

func main() {
   var (
      x = kingpin.Arg("x", "X value of the ratio").Required().Int64()
      y = kingpin.Arg("y", "Y value of the ratio").Required().Int64()
   )
   kingpin.Parse()

   for i:=int64(2); i<=int64(math.Max(float64(*x), float64(*y))); i++ {
      if (*x % i ==0) && (*y % i ==0) {
         *x = *x / i
         *y = *y / i
         i = 1
      }
   }
   fmt.Println(*x,*y)
}
