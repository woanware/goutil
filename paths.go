package goutil

import (
	"os"
	"path/filepath"
	"github.com/mgutz/ansi"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

//
func RemoveIllegalPathCharacters(path string) string {
	re, _ := regexp.Compile("[\\:*?\"<>|]")
	var buffer bytes.Buffer
	for _, c := range path {
		if re.MatchString(string(c))  == false {
			buffer.WriteRune(c)
		}
	}

	temp := buffer.String()
	temp = strings.Replace(temp, "/", "_", -1)

	// Remove the first underscore as there will be one after the port
	if len(temp) > 0 {
		if string(temp[0:1]) == "_" {
			temp = temp[1:]
		}
	}

	// Remove the last underscore as not necessary for the output file
	if len(temp) > 0 {
		if string(temp[len(temp) - 1:]) == "_" {
			temp = temp[0:len(temp) - 1]
		}
	}

	return temp
}

// Ensure that the user supplied path exists as a directory
func DoesDirectoryExist(path string) (bool) {
	file_info, err := os.Stat(path)
	if err == nil {
		if file_info.IsDir() == false {
			fmt.Println(ansi.Color("The item is not a directory", "red"))
			return false
		}

		return true
	} else {
		fmt.Println(ansi.Color(err.Error(), "red"))
	}

	if os.IsNotExist(err) { return false}
	return false
}

// Ensure that the user supplied path exists as a file
func DoesFileExist(path string) (bool) {
	file_info, err := os.Stat(path)
	if err == nil {
		if file_info.IsDir() == true {
			fmt.Println(ansi.Color("The item is not a file", "red"))
			return false
		}

		return true
	} else {
		fmt.Println(ansi.Color(err.Error(), "red"))
	}

	if os.IsNotExist(err) { return false}
	return false
}

func GetApplicationDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}

	return dir
}
