package main

import (
	"fmt"
	"os"

	"github.com/vcokltfre/smelt/libsmelt/repository"
)

func main() {
	switch os.Args[1] {
	case "init":
		repo, err := repository.Init("test", ".")
		if err != nil {
			panic(err)
		}

		fmt.Printf("Initialized repository %s at %s\n", repo.Name, repo.Root)
	case "load":
		repo, err := repository.Load(".")
		if err != nil {
			panic(err)
		}

		fmt.Printf("Loaded repository %s at %s\n", repo.Name, repo.Root)
	}
}
