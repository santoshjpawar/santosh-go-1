package main

import (
	"strings"
	"fmt"
)

func main() {
	if parts := strings.SplitN("consul:8500", "unix://", 2); len(parts) == 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}