package goutil

import (
	"errors"
	"io"
	"strings"
	"io/ioutil"
)

// ##### Types #########################################################################################################

type Tuple struct {
	name  string
	value string
}

type NopCloser struct {
	io.Reader
}

// ##### Public Methods ################################################################################################

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

func GetStringSlicePosition(data []string, term string) (int) {
	for i, v := range data {
		if v == term {
			return i
		}
	}

	return -1
}

func (NopCloser) Close() (err error) { return nil }

// Reads all of the text from a file
func ReadTextFromFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ##### Internal Methods ##############################################################################################

func getQuotedString(data string) (string, string, error) {
	index := strings.Index(data[1:], "\"")
	if (index == -1) {
		return "", "", errors.New("Invalid name value pair, no second quote")
	}

	return data[0:index + 1], data[index + 2:], nil
}
