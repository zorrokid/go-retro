module github.com/zorrokid/go-retro/ui

go 1.17

require (
	fyne.io/fyne/v2 v2.1.2
	github.com/zorrokid/go-retro/database v0.0.0-20220107225113-b9b770b3cc85
	github.com/zorrokid/go-retro/repository v0.0.0-20220107225113-b9b770b3cc85
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fredbi/uri v0.0.0-20181227131451-3dcfdacbaaf3 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20211213063430-748e38ca8aec // indirect
	github.com/godbus/dbus/v5 v5.0.6 // indirect
	github.com/goki/freetype v0.0.0-20181231101311-fa8a33aabaff // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/mattn/go-sqlite3 v1.14.10 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/srwiley/oksvg v0.0.0-20211120171407-1837d6608d8c // indirect
	github.com/srwiley/rasterx v0.0.0-20210519020934-456a8d69b780 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/yuin/goldmark v1.4.4 // indirect
	github.com/zorrokid/go-retro/archive v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410 // indirect
	golang.org/x/net v0.0.0-20220107192237-5cfca573fb4d // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/driver/sqlite v1.2.6 // indirect
	gorm.io/gorm v1.22.4 // indirect
)

replace github.com/zorrokid/go-retro/repository => ../repository

replace github.com/zorrokid/go-retro/database => ../database

replace github.com/zorrokid/go-retro/archive => ../archive
