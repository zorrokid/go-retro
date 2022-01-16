package archive

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const FileChunk = 4096

func ReadFileInfo(filePath string) FileInfo {

	// open file for reading
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// get fileinfo
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fileSize := info.Size()

	// prepare to read in the file in blocks of bytes
	blocks := uint64(math.Ceil(float64(fileSize) / float64(FileChunk)))

	fmt.Printf("fileSize: %d\n", fileSize)
	fmt.Printf("FileChunkSize: %d\n", FileChunk)
	fmt.Printf("blocks: %d\n", blocks)

	// prepare the hash calculation
	hash := sha1.New()

	for i := uint64(0); i < blocks; i++ {
		// calculate current block size
		bytesRead := i * FileChunk
		fmt.Printf("bytes read: %d\n", bytesRead)
		remainingBytes := fileSize - int64(bytesRead)
		fmt.Printf("remaining bytes: %d\n", remainingBytes)
		blocksize := int(math.Min(FileChunk, float64(remainingBytes)))
		fmt.Printf("current blockSize: %d\n", blocksize)
		buf := make([]byte, blocksize)
		f.Read(buf)
		// write buffer to hash
		io.WriteString(hash, string(buf))
	}

	sum := hash.Sum(nil)
	fmt.Printf("hash sum is %x\n", sum)
	fmt.Printf("file size is %d\n", fileSize)
	fmt.Printf("file name is %s\n", info.Name())

	fileInfo := FileInfo{
		FileSizeInBytes: fileSize,
		CheckSum:        hash.Sum(nil),
		FileName:        info.Name(),
	}
	return fileInfo
}
