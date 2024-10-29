//write the "hello world" in golang in this file
package main
import (
	"fmt"
	"os"
)
func main() {
	//ask to the user the language that he want to write
	// 1. english
	// 2. japanese
	// 3. chinese
	//read the user input
	var language int
	fmt.Print("enter 1 for english, 2 for japanese, 3 for chinese: ")
	fmt.Scan(&language)
	if language == 1 {
		//write the println the hello world in english
		fmt.Println("hello world")
	} else if language == 2 {
		//write the println the hello world in jappanese
		fmt.Println("你好世界")
	} else if language == 3 {
		//write the println the hello world in chinese
		fmt.Println("你好，世界")
	}
	//exit
	os.Exit(0)
}