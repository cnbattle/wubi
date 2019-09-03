package main

import (
	"github.com/cnbattle/wubi"
	"log"
)

func main() {
	got, err := wubi.Get("朱文倩")
	log.Println(got, err)
}
