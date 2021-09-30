package weichatdat

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func DatToImage(datPath string, imageSavePath string) error {
	fileInfos, err := ioutil.ReadDir(datPath)
	if err != nil {
		return err
	}
	var datFiles []string
	for key, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			log.Println("dir:", key)
		} else {
			if !strings.HasSuffix(fileInfo.Name(), ".dat") {
				break
			}
			infileName := fmt.Sprintf("%s\\%s", datPath, fileInfo.Name())
			datFiles = append(datFiles, infileName)
		}
	}
	if len(datFiles) < 1 {
		return nil
	}
	for key, datFile := range datFiles {
		datBytes, err := ioutil.ReadFile(datFile)
		if err != nil {
			log.Println(err)
		}
		imgBytes, err := wechatDataToImage(datBytes)
		if err != nil {
			log.Println(err)
		}
		datFileName := filepath.Base(datFile) + ".jpg"
		outfilepath := filepath.Join(imageSavePath, datFileName)
		err = checkDirExist(outfilepath)
		err = ioutil.WriteFile(outfilepath, imgBytes, 0777)
		if err != nil {
			log.Println(err)
		}
		log.Printf("completed:%d %s", key, outfilepath)
	}
	return nil
}

func wechatDataToImage(dat []byte) (img []byte, err error) {
	var buffer bytes.Buffer
	for _, val := range dat {
		err = buffer.WriteByte(val ^ 0xF2)
		if err != nil {
			return
		}
	}
	img = buffer.Bytes()
	return
}

func checkDirExist(path string) error {
	dir := filepath.Dir(path)
	exists := isPathExists(path)
	if !exists {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func isPathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
