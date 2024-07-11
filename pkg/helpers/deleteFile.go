package helpers

import (
	"fmt"
	"os"
	"strings"
)

func DeleteFile(imgUrl string) bool {
	fileName := strings.Split(imgUrl, "static/photo/")[1]

	err := os.Remove(fmt.Sprintf("uploads/photo/%s", fileName))
	if err != nil {
		return false
	}

	fmt.Println("File deleted")
	return true
}
