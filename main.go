package main

import (
	"github.com/shisiying/go-tools/cmd"
	"log"
)

func main()  {
	err := cmd.Execute()
	if err != nil {
		log.Printf("cmd.Execute err: %v",err)
	}
}