package utils

import (
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
)



func SetBasePathToCurrentDir(){
	_, filename, _, _ := runtime.Caller(1) // skip one stack to ensure we get dir of the calling file and not the utils
	err := os.Chdir(path.Dir(filename))
	HandleError(err, "updating current working dir")
}

func HandleError(err error, action string) {
	if err!= nil {
		log.Fatalf("Error while %s: %s", action, err)
	}
}

func Binary2Int(binary string) int64 {
	n, err := strconv.ParseInt(binary, 2, 64)
	HandleError(err, "trying to convert binary to int")
	return n
}

func String2Int(s string) int {
	n, err := strconv.Atoi(s)
	HandleError(err, "trying to convert string to int")
	return n
}

func getFile(s string) *os.File {
	file, err := os.Open("input_numbers_drawn.txt")
	HandleError(err, "opening input file")
	defer func(file *os.File) {
		err := file.Close()
		HandleError(err, "closing input file")
	}(file)

	return file
}