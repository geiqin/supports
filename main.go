package main

import "github.com/geiqin/supports/token"

func main() {
	t :=token.UserToken{}
	t.Decode("sssss")
}