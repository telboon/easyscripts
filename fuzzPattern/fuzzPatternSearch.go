package main

import (
   "fmt"
   "gopkg.in/alecthomas/kingpin.v2"
   "encoding/hex"
   "strings"
   "bytes"
   "os"
)

func main() {
   patternPointer:= kingpin.Arg("search", "Hex(eg '0x42424242') or String to be search").Required().String()
   littleEndianPointer:= kingpin.Flag("little-endian", "Assumes the string to be little endian (default)").Default("true").Bool()
   bigEndianPointer:= kingpin.Flag("big-endian", "Assumes the string to be big endian").Default("false").Bool()
   pattLengthPointer:= kingpin.Flag("length", "Pattern length").Short(rune('l')).Int()
   kingpin.Parse()

   var fullPattern string
   var pattern string
   var combiArr []int = []int{0,0,0,0}
   combiMax := []int{26, 26, 26, 10}
   var fullCombi int
   charsetCaps := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
   charsetSmall := "abcdefgjijklmnopqrustuvwxyz"
   charsetNum := "0123456789"
   pattLength := *pattLengthPointer
   littleEndian := *littleEndianPointer
   bigEndian := *bigEndianPointer
   finalPos := 0
   var startPos int

   if len(*patternPointer)<4 {
      fmt.Println("Input must have at least 4 bytes")
      os.Exit(1)
   }

   if (*patternPointer)[:2] == "0x" {
      fmt.Println("Hex detected")
      fmt.Println("Assumed litle endian")
      fullBytes, _ := hex.DecodeString((*patternPointer)[2:])
      fullPattern = string(fullBytes)
   } else  {
      fullPattern = *patternPointer
   }

   if len(fullPattern)<4 {
      fmt.Println("Input must have at least 4 bytes")
      os.Exit(1)
   }

   //change if little endian
   if !bigEndian && littleEndian {
      fullPattern = strReverse(fullPattern)
   }

   pattern = fullPattern[:4]

   //check if combi is valid
   bigCheck:=0
   for i:=0; i<len(pattern); i++ {
      if strings.Contains(charsetCaps, string(pattern[i])) {
         startPos = i
         bigCheck+=1
      }
   }
   if bigCheck!=1 {
      fmt.Println("Your pattern is invalid. Please try again...")
      os.Exit(1)
   }

   //find combi
   combiArr[0]=strings.IndexByte(charsetCaps, pattern[(startPos+0)%4])
   combiArr[1]=strings.IndexByte(charsetSmall, pattern[(startPos+1)%4])
   combiArr[2]=strings.IndexByte(charsetSmall, pattern[(startPos+2)%4])
   combiArr[3]=strings.IndexByte(charsetNum, pattern[(startPos+3)%4])

   //check if combi is valid
   for _,i := range combiArr {
      if i<0 {
         fmt.Println("Your pattern is invalid. Please try again...")
         os.Exit(1)
      }
   }

   finalPos = findPos(combiArr, startPos)

   if pattLength == 0 {
      fmt.Println("Pattern found:",finalPos)
   } else {
      fullCombi = combiMax[0] * combiMax[1] * combiMax[2] * combiMax[3]
      fmt.Println("Full cycle:", fullCombi,"*", len(combiArr), "=", fullCombi*len(combiArr))
      fmt.Println("")

      for finalPos < pattLength {
         fmt.Println("Pattern found:",finalPos)
         finalPos+=fullCombi*len(combiArr)
      }
   }
}

func findPos(combiArr []int, startPos int) int {
   combiMax := []int{26, 26, 26, 10}
   permuteCombi := []int{10*26*26, 26*10, 10, 1}

   var answerLength int

   if startPos > 0 {
      combiArr[len(combiArr)-1]+=1

      for i:= len(combiArr)-1; i>0; i-- {
         if (combiArr[i] >= combiMax[i]) && (i >= (len(combiArr)-startPos)) {
            combiArr[i] = 0
            if (i-1 >= (len(combiArr)-startPos)) {
               combiArr[i-1] += 1
            }
         }
      }
   }

   for i:=0; i<len(combiArr); i++ {
      answerLength += combiArr[i] * permuteCombi[i]
   }

   answerLength = answerLength * len(combiArr)
   answerLength -= startPos

   return answerLength
}

func strReverse(before string) string {
   var newBuff bytes.Buffer
   for i,_ := range before {
      newBuff.WriteByte(before[len(before)-i-1])
   }
   return newBuff.String()
}
