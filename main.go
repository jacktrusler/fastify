package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	// "sync"

	"github.com/dixonwille/wmenu/v5"
	_ "github.com/mattn/go-sqlite3"
)

const file string = "fastfood.db"

const create string = `
  CREATE TABLE IF NOT EXISTS fast_food (
  id INTEGER NOT NULL PRIMARY KEY,
  name TEXT,
  description TEXT
);`

func makeFastifyTable() (*fastFood, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}
	return &fastFood{
		name: "lmao",
	}, nil
}

type fastFood struct {
	id          string
	name        string
	description string
}

func handleFunc(db *sql.DB, opts []wmenu.Opt) {

	switch opts[0].Value {

	case 0:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a fast food joint: ")
		name, _ := reader.ReadString('\n')
		fmt.Print("Enter a description: ")
		description, _ := reader.ReadString('\n')

		newFF := fastFood{
			name:        name,
			description: description,
		}

		addFoodJoint(db, newFF)
		break

	case 1:
		fmt.Println("Finding a Person")
	case 2:
		fmt.Println("Update a Person's information")
	case 3:
		fmt.Println("Deleting a person by ID")
	case 4:
		fmt.Println("Quitting application")
	}
}

func addFoodJoint(db *sql.DB, newFF fastFood) {
	stmt, _ := db.Prepare("INSERT INTO fast_food (id, name, description) VALUES (?, ?, ?)")
	stmt.Exec(nil, newFF.name, newFF.description)
	defer stmt.Close()

	fmt.Printf("Added %v %v \n", newFF.name, newFF.description)
}

func main() {

	db, err := sql.Open("sqlite3", "./fastfood.db")
	if err != nil {
		log.Fatal(err)
	}

	makeFastifyTable()

	menu := wmenu.NewMenu("What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error {
		handleFunc(db, opts)
		return nil
	})

	menu.Option("Add a new Person", 0, true, nil)
	menu.Option("Find a Person", 1, false, nil)
	menu.Option("Update a Person's information", 2, false, nil)
	menu.Option("Delete a person by ID", 3, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		fmt.Println("lmao")
		log.Fatal(menuerr)
	}

	defer db.Close()
}
