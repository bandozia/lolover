package fileservice

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func RenderDir() {
	// if files, e := ioutil.ReadDir("/home/bruno"); e == nil {

	// 	for _, f := range files {
	// 		println(f.Name())
	// 	}
	// }
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
