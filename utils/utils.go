package utils

import (
	"log"
	"os"
	"path"
	"runtime"
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