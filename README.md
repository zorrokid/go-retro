# go-retro

The goal of go-retro is to be an emulator front end for retro gaming and retro game collectors. It will have a bit different point of view compared to most emulator front ends: go-retro is rather game title centric whereas most front ends are gaming system centric or emulator centric. This means that the main view will not have a list or menuof different gaming systems, it will have a list of different game titles, each game title will have a list of different releases of this title for different gaming or computer platforms. Emulators can be configured for each of these platforms. Also other types of files can be added to each release: manuals, box scans, screenshots etc.

Another goal is to be kind of a "rom manager" tool. go-retro will store the checksums of files added to database and use that information when scanning files. Possibly there will be an online database of checksum files that users will be able to use scanning files for go-retro usage. Possibly there will be also screenshots provided from online API along with file info.

go-retro is build using Go language and fyne.io for UI. go-retro is build from separate modules. The idea using separate modules instead of one module containing everything as packages is that as modules it would be easier to reuse part of this code for example to build Cloud/Web-based database and API for go-retro use.
