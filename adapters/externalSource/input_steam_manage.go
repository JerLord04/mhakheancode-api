package externalsource

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"mime/multipart"
)

type InputSteamManage interface {
	ConvertByteArrayToBase64(source []byte) (string, error)
	ConvertFileToByreArray(file *multipart.FileHeader) ([]byte, error)
}

type InputSteamManageStruct struct {
}

func (b *InputSteamManageStruct) ConvertFileToByreArray(file *multipart.FileHeader) ([]byte, error) {
	fileByteArray, err := file.Open()
	if err != nil {
		return nil, err
	}
	fileByteArray.Close()
	openFile, err := ioutil.ReadAll(fileByteArray)
	if err != nil {
		return nil, err
	}
	return openFile, nil
}

func (b *InputSteamManageStruct) ConvertByteArrayToBase64(source []byte) (string, error) {
	if source == nil || len(source) == 0 {
		return "", errors.New("source is empty or nil")
	}
	base64String := base64.StdEncoding.EncodeToString(source)

	return base64String, nil
}

func NewByteArrayToImageStruct() InputSteamManage {
	return &InputSteamManageStruct{}
}
