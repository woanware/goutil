package goutil

import (
	"os"
	"net/http"
	"errors"
	"fmt"
	"io"
)

// ##### Methods #############################################################

//
func DownloadToFile(url string, file string) (error) {
	output, err := os.Create(file)
	if err != nil {
		return errors.New(fmt.Sprintf("Error creating download file: %v (%s)", err, file))
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		return errors.New(fmt.Sprintf("Error downloading from URL: %v (%s)", err, url))
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error outputing downloaded file: %v (%s)", err, file))
	}

	return nil
}
