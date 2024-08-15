package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){
	fmt.Println("what is your name ?")
	reader :=bufio.NewReader(os.Stdin);
	name ,err:=reader.ReadString('\n');
	if(err==nil){
		fmt.Println("hello",name);
	}else{
		log.Fatal(err)
	}

}