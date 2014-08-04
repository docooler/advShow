package adutils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// exists returns whether the given file or directory exists or not
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//delete file if the error delay after 10s and try it again
func Unlink(file string) error {
	count := 5
	var ret error
	for i := 0; i < count; i++ {
		ret = os.Remove(file)
		if ret == nil {
			return nil
		}
		time.Sleep(10)
	}
	return ret
}

//rename file if error delay after 10s and try it again
func Rename(oldfile string, newfile string) error {
	count := 5
	for i := 0; i < count; i++ {
		err := os.Rename(oldfile, newfile)
		if err == nil {
			return nil
		}
		log.Println(err)
		log.Println("Rename error!")
		time.Sleep(10)
	}
	return nil
}
func GetAdvFilename(index string) (string, int, error) {
	i, err := strconv.Atoi(index)
	if err != nil {
		// handle error
		fmt.Println(err)
		return " ", 0, err
	}
	fmt.Println(index)
	page, err := ContentParse()
	if err != nil {
		// handle error
		fmt.Println(err)
		return " ", 0, err
	}

	disLen := len(page.Display)

	findex := i % disLen
	if page.Display[findex].Type == "link" {
		return page.Display[findex].Link, 1, nil
	}

	filename := "../static/" + page.Display[findex].Link

	ret, _ := Exists(filename)
	if ret != true {
		fmt.Println(filename + "not exist in the file")
	}
	return filename, 0, nil
}
