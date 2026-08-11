package main

import (
	"fmt"

	"github.com/matveyarbuzov/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println("it's work!", st)
}
