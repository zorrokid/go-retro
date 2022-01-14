package archive

import (
	"archive/zip"
	"crypto/sha1"
	"io"
	"log"
)

func ReadZip(filePath string) []ArchiveFile {

	r, err := zip.OpenReader(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	files := make([]ArchiveFile, len(r.File))

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()

		hash := sha1.New()
		bytesCopied, err := io.Copy(hash, rc)

		if err != nil {
			log.Fatal(err)
		}
		sum := hash.Sum(nil)
		af := ArchiveFile{
			CheckSum:        sum,
			FileName:        f.Name,
			FileSizeInBytes: bytesCopied,
		}
		files = append(files, af)
	}
	return files
}
