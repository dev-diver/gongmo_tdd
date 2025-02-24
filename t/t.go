package main

import (
	"fmt"

	"github.com/dev-diver/gongmo/domain"
)

func main() {

	m := make(map[domain.AccountId]int)
	m[domain.AccountId("1")] = 1
	m[domain.AccountId("2")] = 2
	m[domain.AccountId("3")] = 3

	fmt.Println(m)
}
