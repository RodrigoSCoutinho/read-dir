package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	dir, err := os.Open("/tmp")
	if err != nil {
		panic(err)
	}

	//*LOOP INFINITE

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			//*VALIDATE
			//* EOF => END OF FILE
			if err == io.EOF {
				break
			}
			fmt.Printf("ERROR: %v\n", err)
			continue
		}

		//* RETURN THE FILE NAME

		fmt.Println(files[0].Name())
	}
}
