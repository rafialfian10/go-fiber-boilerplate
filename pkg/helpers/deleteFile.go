package helpers

import (
	"fmt"
	"os"
	"strings"
)

func DeleteFile(imgUrl string) bool {
	var fileName string

	if strings.Contains(imgUrl, "static/photo/") {
		fileName = strings.Split(imgUrl, "static/photo/")[1]
		err := os.Remove(fmt.Sprintf("uploads/photo/%s", fileName))
		if err != nil {
			return false
		}
	} else if strings.Contains(imgUrl, "static/image/") {
		fileName = strings.Split(imgUrl, "static/image/")[1]
		err := os.Remove(fmt.Sprintf("uploads/image/%s", fileName))
		if err != nil {
			return false
		}
	} else {
		return false
	}

	fmt.Println("File deleted")
	return true
}

// func DeleteFile(imgUrl string) bool {
// 	fileName := strings.Split(imgUrl, "static/photo/")[1]

// 	err := os.Remove(fmt.Sprintf("uploads/photo/%s", fileName))
// 	if err != nil {
// 		return false
// 	}

// 	fmt.Println("File deleted")
// 	return true
// }
