package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file crea ion\n")
		return
	}
	//在函数执行结束前，一定要关闭，谨记
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "a中国b\U0010FFFF"
	fmt.Println(outputString)
	for _, b := range []byte(outputString) {
		fmt.Printf("xxxxxxx %x\n", b)
	}
	for _, b := range outputString {
		fmt.Printf("yyyyyyy %x\n", b)
	}
	for _, b := range []rune(outputString) {
		fmt.Printf("zzzzzzz %x\n", b)
	}
	outputWriter.WriteString(outputString)
	outputWriter.Flush()
}
