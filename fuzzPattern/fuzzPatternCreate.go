package main

import (
   "fmt"
   "gopkg.in/alecthomas/kingpin.v2"
   "bytes"
)

func main() {
   charCountPointer:= kingpin.Arg("char-count", "Number of characters to generate").Required().Int()
   kingpin.Parse()

   var charCount int=*charCountPointer
   var ans bytes.Buffer
   var tempCombi string
   var combiArr [4]int
   charsetCaps := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
   charsetSmall := "abcdefgjijklmnopqrustuvwxyz"
   charsetNum := "0123456789"

   for ans.Len()<charCount {
      tempCombi=string(charsetCaps[combiArr[0]]) + string(charsetSmall[combiArr[1]]) + string(charsetSmall[combiArr[2]]) + string(charsetNum[combiArr[3]])
      if ans.Len()+len(tempCombi) < charCount {
         ans.WriteString(tempCombi)
      } else  {
         ans.WriteString(tempCombi[:charCount-ans.Len()])
      }

      addCombi(combiArr[:])
   }

   fmt.Println(ans.String())
}

func addCombi(combiArr []int) {
   combiMax := [4]int{26, 26, 26, 10}

   combiArr[len(combiArr)-1]+=1

   for i:= len(combiArr)-1; i>0; i-- {
      if combiArr[i] >= combiMax[i] {
         combiArr[i] = 0
         combiArr[i-1] += 1
      }
   }

   if combiArr[0] >= combiMax[0] {
      combiArr[0] = 0
   }
}
