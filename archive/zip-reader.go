package archive

import (
	"archive/zip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func ReadZip(filePath string) {

	r, err := zip.OpenReader(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("Checksum of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		hash := sha1.New()
		if _, err := io.Copy(hash, rc); err != nil {
			log.Fatal(err)
		}
		sum := hash.Sum(nil)

		fmt.Printf("%x\n", sum)
		rc.Close()
		fmt.Println()
	}
}
