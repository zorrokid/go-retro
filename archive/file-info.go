package archive

type FileInfo struct {
	FileName        string
	CheckSum        []byte
	FileSizeInBytes int64
}
