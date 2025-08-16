package main

import (
	"fmt"

	"github.com/biancarosa/go-learning-resources/structs/models"
)

func main() {
	p := models.NewPerson("Zé", "das Couves")

	fmt.Println(p.FullName)
}
