module github.com/zorrokid/go-retro/tester

go 1.17

replace github.com/zorrokid/go-retro/database => ../database

require (
	github.com/zorrokid/go-retro/database v0.0.0-20220102220145-c1a27ea4f28f
	github.com/zorrokid/go-retro/repository v0.0.0-00010101000000-000000000000
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/mattn/go-sqlite3 v1.14.10 // indirect
	gorm.io/driver/sqlite v1.2.6 // indirect
	gorm.io/gorm v1.22.4 // indirect
)

replace github.com/zorrokid/go-retro/repository => ../repository
