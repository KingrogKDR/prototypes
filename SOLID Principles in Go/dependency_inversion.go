package main

import "fmt"

// SOLID -> D
// Dependency Inversion
type StoreHandler interface {
	Store(string)
}

type Database struct{}

func (db Database) Store(data string) {
	fmt.Println("Data:", data)
}

type BusinessLogic struct {
	storeHandler StoreHandler // not using db Database directly since we only need the store function of db and not the whole db itself
}

func (b BusinessLogic) SaveData(data string) {
	b.storeHandler.Store(data)
}

func DependencyInversion() {
	db := Database{}
	b := BusinessLogic{storeHandler: db}
	b.SaveData("Saved data!")
}
