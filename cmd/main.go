package main

import "log"

func main() {
	if err := migrate(); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

func migrate() error {
	return nil
}
