package util

import (
	"io"
	"os"
)

// CopyFile copyfile from src to dest
func CopyFile(src string, dest string) error {

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	_, err = os.Create(dest)
	if err != nil {
		return err
	}

	destFile, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
