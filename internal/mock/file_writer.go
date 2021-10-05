package mock

import "github.com/stretchr/testify/mock"

type FileWriter struct {
	mock.Mock
}

func (fw *FileWriter) WriteDataToFile(fName string, path string, data string) error {
	args := fw.Called(fName, path, data)
	err := args.Error(0)
	return err
}
