package main

import (
	"fmt"
	"log"
	"github.com/arjunsingh14/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	cfg.SetUser("arjun")
	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("db_url: %s\ndb_name: %s", cfg.Db_url, cfg.User)
}