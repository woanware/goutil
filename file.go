package goutil

import (
	"path/filepath"
	"os"
	"archive/zip"
	"archive/tar"
	"compress/gzip"
	"io"
	"fmt"
	"errors"
	"crypto/md5"
)

//
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() error {
		if err := r.Close(); err != nil {
			return fmt.Errorf("Error unzipping: (%s) %v", src, err)
		}

		return nil
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() error {
			if err := rc.Close(); err != nil {
				return fmt.Errorf("Error unzipping: (%s) %v", src, err)
			}

			return nil
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() error {
				if err := f.Close(); err != nil {
					return fmt.Errorf("Error unzipping: (%s) %v", src, err)
				}

				return nil
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

//
func UnzipFile(src, dest, fileName string) error {
	zipReader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() error {
		if err := zipReader.Close(); err != nil {
			return fmt.Errorf("Error unzipping: (%s) %v", src, err)
		}

		return nil
	}()

	for _, zf := range zipReader.File {
		if zf.Name != fileName {
			continue
		}

		dst, err := os.Create(dest)
		if err != nil {
			return err
		}

		defer dst.Close()
		src, err := zf.Open()
		if err != nil {
			// err
		}
		defer src.Close()

		io.Copy(dst, src)
	}

	return nil
}

//
func Ungunzip(src, dest string) error {
	gzipfile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer gzipfile.Close()

	reader, err := gzip.NewReader(gzipfile)
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer writer.Close()

	if _, err = io.Copy(writer, reader); err != nil {
		return err
	}

	return nil
}

//
func UnTarGunzipFile(src, dest, fileName string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	gzipReader, err := gzip.NewReader(f)
	defer func() error {
		if err := gzipReader.Close(); err != nil {
			return fmt.Errorf("Error tar gunzipping: (%s) %v", src, err)
		}

		return nil
	}()

	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		name := header.Name
		if name != fileName {
			continue
		}

		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			dst, err := os.Create(dest)
			if err != nil {
				return err
			}
			defer dst.Close()

			io.Copy(dst, tarReader)
			break

		default:
		}
	}

	return nil
}

//
func Md5File(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error opening file for MD5 computation: %v (%s)", err, filePath))
	}

	hasher := md5.New()
	io.Copy(hasher, f)

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

// Ensure that the user supplied path exists as a file
func DoesFileExist(path string) (bool) {
	file_info, err := os.Stat(path)
	if err == nil {
		if file_info.IsDir() == true {
			fmt.Println("The item is not a file")
			return false
		}

		return true
	}

	if os.IsNotExist(err) { return false}
	return false
}

