package main

import (
	"fmt"
	"testing"
	"time"
)

func TestFib(t *testing.T) {
	comeco := time.Now()
	have := fib(40)
	duracao := time.Since(comeco)

	fmt.Println("VALOR:", have)
	fmt.Println("Duração:", duracao)
	//t.Fail()
}
