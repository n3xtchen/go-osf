package main

import (
    "fmt"
    "github.com/n3xtchen/go-osf/register"
)

func main() {
	r := register.NewRegistryConf("test", "127.0.0.1")
	fmt.Printf("%s", r)
}
