package inpututils

import (
	"fmt"
	"os"
	"strings"
)

func GetFileInput(filename string) (input string) {
	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		fmt.Print(pwdErr)
	}

	filePath := strings.Join([]string{pwd, "input", filename}, "/")
	fileContent, readFileErr := os.ReadFile(filePath)
	if readFileErr != nil {
		fmt.Print(readFileErr)
	}

	return string(fileContent)
}
