package Sender

import (
	"os"

	"github.com/sijiramakun/seapick/utils"
)

type FileMetaData struct {
	name     string
	fileSize uint32
	reps     uint32
}

//INFO:This code checks if the file at the given path exists and is a regular file.
func isFile(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		// File does not exist
		return false
	}
	return !info.IsDir() // Returns true if it's not a directory
}

func prepareFIleMetaData(file *os.File) FileMetaData {
	fileInfo, err := file.Stat()

	utils.CheckError(err)

	size := fileInfo.Size()

	header := FileMetaData{
		name:     fileInfo.Name(),
		fileSize: uint32(size),
		reps:     uint32(size/1014) + 1,
	}

	return header
}
