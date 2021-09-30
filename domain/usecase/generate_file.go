package usecase

import (
	"fmt"
	"github.com/ISKalsi/leet-scrape/v2/domain/model"
	"log"
	"os"
	"path/filepath"
)

type generateFile struct {
	question *model.Question
	path     string
}

func (uc *generateFile) writeDataToFile(fName string, data string) error {
	f, err := os.OpenFile(filepath.Join(uc.path, fName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	_, err = f.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created file " + fName + " successfully")
	return nil
}
