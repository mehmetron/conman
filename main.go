package main

import (
	"fmt"
	"github.com/mehmetron/conman/handlers"
)

func main() {
	fmt.Println("Conman running...")

	handlers.Routes()

}
