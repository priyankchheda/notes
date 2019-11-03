package main

import (
	"fmt"

	"github.com/x899/gopass/gopass"
)

func main() {
	// createFlag := flag.Bool("create", false, "Insert a username/password in to your vault")
	// flag.Parse()

	// if *createFlag {
	// 	create()
	// }

	gopass.Create("main1", "gopher1")

	key := []byte("passphrasewhichneedstobe32bytes!")
	data := gopass.Decrypt("myfile.data", key)
	fmt.Println(string(data))
}
