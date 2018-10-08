package main

import (
	"github.com/spf13/pflag"
	"fmt"
)
var inputName = pflag.StringP("name", "n","CHENJIAN", "Input Your Name.")
var inputAge = pflag.IntP("age", "a",27, "Input Your Age")
var inputGender = pflag.StringP("gender","g","female", "Input Your Gender")
var inputFlagvar int

func Init() {
	pflag.IntVarP(&inputFlagvar, "flagname","f", 1234, "Help")
	// pflag.MarkShorthandDeprecated("name", "please use --name only")
}
func main() {
	Init()
	pflag.Parse()
	// func Args() []string
	// Args returns the non-flag command-line arguments.
	// func NArg() int
	// NArg is the number of arguments remaining after flags have been processed.
	fmt.Printf("args=%s, num=%d\n", pflag.Args(), pflag.NArg())
	for i := 0; i != pflag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, pflag.Arg(i))
	}
	//pointer
	fmt.Println("name=", *inputName)
	fmt.Println("age=", *inputAge)
	fmt.Println("gender=", *inputGender)
	//value
	fmt.Println("flagname=", inputFlagvar)
}