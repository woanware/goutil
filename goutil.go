package goutil

import (
	"github.com/mgutz/ansi"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
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

func ParseNameValue(data string) (string, string, error) {
	name := ""
	value := ""
	remainder := ""
	var err error

	// Name starts with a quote so find next quote
	if data[0:1] == "\"" {
		name, remainder, err = getQuotedString(data[1:])
		if err != nil {
			return "", "", err
		}

		// Ensure that there is a space in between the name and value
		index := strings.Index(remainder, " ")
		if index > -1 {
			remainder = remainder[1:]

			if remainder[0:1] == "\"" {
				value, remainder, err = getQuotedString(remainder[1:])
				if err != nil {
					return "", "", err
				}
			} else {
				value = remainder
			}
		}
	} else {
		// Name doesn't start with a quote so find the next space
		index := strings.Index(data, " ")
		// No space identified so the "name" part is the entire string, with a blank "value"
		if (index == -1) {
			name = data
		} else {
			name = data[0:index]
			if data[index + 1:index + 2] == "\"" {
				value, remainder, err = getQuotedString(data[index + 2:])
				if err != nil {
					return "", "", err
				}
			} else {
				value = data[index + 1:]
			}
		}
	}

	return name, value, nil
}

func getQuotedString(data string) (string, string, error) {
	index := strings.Index(data[1:], "\"")
	if (index == -1) {
		return "", "", errors.New("Invalid name value pair, no second quote")
	}

	return data[0:index + 1], data[index + 2:], nil
}

func GetStringSlicePosition(data []string, term string) (int) {
	for i, v := range data {
		if v == term {
			return i
		}
	}

	return -1
}

type NopCloser struct {
	io.Reader
}
func (NopCloser) Close() (err error) { return nil }

