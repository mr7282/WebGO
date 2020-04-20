package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// * Напишите функцию, которая получает на вход публичную ссылку на файл с «Яндекс.Диска» и сохраняет полученный файл на диск пользователя.
// Подсказка: документация к нужному запросу.

// FileInfoFields - структура вида ответа с yandex.ru
type FileInfoFields struct {
	Link string `json:"file"`
	Name string `json:"name"`
	Size int    `json:"size"`
}

func storedFile(link string) {
	resp, err := http.Get(link)
	if err != nil {
		log.Error(err)
		return
	}
	defer resp.Body.Close()

	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	readInfo := &FileInfoFields{}
	err = json.Unmarshal(read, readInfo)
	if err != nil {
		log.Error(err)
	}

	fileResp, err := http.Get(readInfo.Link)
	if err != nil {
		log.Error(err)
	}
	defer fileResp.Body.Close()

	openFile, err := ioutil.ReadAll(fileResp.Body)
	if err != nil {
		log.Error(err)
	}
	
	err = ioutil.WriteFile("./test.mp4", openFile, 0644)
	if err != nil {
		log.Error(err)
	}
	log.Info("the file is stored")
}

func main() {
	storedFile("https://cloud-api.yandex.net/v1/disk/public/resources?public_key=https://yadi.sk/i/OcXmlJRIvpH7-A")
}
