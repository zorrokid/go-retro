package archive

type ArchiveFile struct {
	FileName        string
	CheckSum        []byte
	FileSizeInBytes int64
}
