package main

import (
	"fmt"

	"github.com/zorrokid/go-retro/database"
	"github.com/zorrokid/go-retro/repository"
)

func main() {
	db := database.NewDatabase()
	db.InitDB()
	repo := repository.NewRepository(db)
	repo.AddTitle("Zorro")
	title1 := repo.GetTitle(1)
	title2 := repo.GetTitleByName("Zorro")
	fmt.Printf("Title: %s, id: %d\n", title1.Name, title1.ID)
	fmt.Printf("Title: %s, id: %d\n", title2.Name, title2.ID)
}
