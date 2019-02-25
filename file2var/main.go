package main

import (
    "gopkg.in/alecthomas/kingpin.v2"
    "io/ioutil"
    "fmt"
    "strings"
    "log"
)

func main() {
    outFilePt := kingpin.Flag("outfile", "File to output the variable").Short('o').String()
    fileNamePt := kingpin.Arg("filename", "File to be converted into variable").Required().String()
    kingpin.Parse()

    outFile := *outFilePt
    fileName := *fileNamePt

    outputBytes, err := ioutil.ReadFile(fileName)
    if err!= nil {
        log.Fatal("Error: Cannot read file")
    }
    convertedStr := convertSanitize(outputBytes)

    if outFile == ""  {
        fmt.Println(convertedStr)
    } else {
        ioutil.WriteFile(outFile, []byte(convertedStr), 0660)
    }
}

func convertSanitize(inByte []byte) string {
    var sb strings.Builder

    sb.WriteString("fileVar := \"")

    for i:=0;i<len(inByte);i++ {
        if inByte[i] == '\n' {
            sb.WriteString("\\n")
        } else if inByte[i] == '\r' {
            sb.WriteString("\\r")
        } else if inByte[i] == '\t' {
            sb.WriteString("\\t")
        } else if inByte[i] == '"' {
            sb.WriteString("\\\"")
        } else if inByte[i] == '\\' {
            sb.WriteString("\\\\")
        } else if inByte[i] > 32 && inByte[i] < 126{
            sb.WriteByte(inByte[i])
        } else {
            sb.WriteString(fmt.Sprintf("\\x%02x",inByte[i]))
        }
        if i%50 ==0 {
            sb.WriteString("\"\nfileVar += \"")
        }
    }
    sb.WriteString("\"\n")
    return sb.String()
}
