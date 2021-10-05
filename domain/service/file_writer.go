package service

type FileWriter interface {
	WriteDataToFile(fName string, path string, data string) error
}
