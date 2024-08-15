package main

import (
	"fmt"
)

func main(){
	var name string 
	fmt.Println("Enter your name")

	fmt.Scanln(&name, name);
	/* reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n') */
	fmt.Println("Hello", name)
}