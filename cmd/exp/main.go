package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/sviut/photo-locker/models"
)

type User struct {
	Name string
	Age  int
	Meta Meta
}

type Meta struct {
	Visits int
}

func generateCSRFKey() string {
	key := make([]byte, 32)
	rand.Read(key)
	return base64.StdEncoding.EncodeToString(key)
}

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database")

	us := models.UserService{DB: db}
	user, err := us.Create("bob@test.com", "32456")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
