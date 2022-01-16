package archive

import (
	"archive/zip"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"log"
	"math"
)

const (
	// IEEE is by far and away the most common CRC-32 polynomial.
	// Used by ethernet (IEEE 802.3), v.42, fddi, gzip, zip, png, ...
	IEEE = 0xedb88320

	// Castagnoli's polynomial, used in iSCSI.
	// Has better error detection characteristics than IEEE.
	// https://dx.doi.org/10.1109/26.231911
	Castagnoli = 0x82f63b78

	// Koopman's polynomial.
	// Also has better error detection characteristics than IEEE.
	// https://dx.doi.org/10.1109/DSN.2002.1028931
	Koopman = 0xeb31d82e
)

func ReadZip(filePath string) []ArchiveFile {

	// open file for reading
	r, err := zip.OpenReader(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	files := make([]ArchiveFile, len(r.File))

	// go through each file in archive and calculate checksumS
	for _, f := range r.File {

		// compressed file has checksum (for decompressed file?)
		fileChecksum := f.CRC32

		// open file for reading
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()

		fileSize := f.UncompressedSize64
		blocks := uint64(math.Ceil(float64(fileSize) / float64(FileChunk)))
		crc32q := crc32.MakeTable(IEEE)
		crc := uint32(0)
		hash := sha1.New()

		for i := uint64(0); i < blocks; i++ {
			// calculate current block size
			bytesRead := i * FileChunk
			fmt.Printf("bytes read: %d\n", bytesRead)
			remainingBytes := int64(fileSize - bytesRead)
			fmt.Printf("remaining bytes: %d\n", remainingBytes)
			blocksize := int(math.Min(FileChunk, float64(remainingBytes)))
			fmt.Printf("current blockSize: %d\n", blocksize)
			buf := make([]byte, blocksize)
			rc.Read(buf)

			crc = crc32.Update(crc, crc32q, buf)
			hash.Write(buf)
		}

		sum := hash.Sum(nil)

		fmt.Printf("original crc32 checksum %d\n", fileChecksum)
		fmt.Printf("calculated crc32 checksum %d\n", crc)
		fmt.Printf("calculated sha1 checksum %s\n", hex.EncodeToString(sum))
		fmt.Printf("calculated sha1 checksum %x\n", sum)

		af := ArchiveFile{
			CheckSum:        sum,
			FileName:        f.Name,
			FileSizeInBytes: int64(fileSize),
		}
		files = append(files, af)
	}
	return files
}
