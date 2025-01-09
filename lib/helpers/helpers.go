package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func FormatPhoneNumber(input string) (string, error) {
	// Remove all non-digit characters
	re := regexp.MustCompile(`\D`)
	digits := re.ReplaceAllString(input, "")

	// Check if the number is valid for Uzbekistan (+998)
	if len(digits) == 12 && strings.HasPrefix(digits, "998") {
		// Format as +998 xx xxx xx xx
		return fmt.Sprintf("+%s %s %s %s %s",
			digits[:3],         // +998
			digits[3:5],        // xx
			digits[5:8],        // xxx
			digits[8:10],       // xx
			digits[10:12]), nil // xx
	} else if len(digits) == 9 {
		// If a local number (without +998), assume Uzbekistan and prepend it
		return fmt.Sprintf("+998 %s %s %s %s",
			digits[:2],       // xx
			digits[2:5],      // xxx
			digits[5:7],      // xx
			digits[7:9]), nil // xx
	}

	return "", fmt.Errorf("invalid phone number")
}

func OwnloadPhoto(fileURL, savePath string) error {
	// Create the save directory if it doesn't exist
	saveDir := filepath.Dir(savePath)
	err := os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Get the photo from the URL
	resp, err := http.Get(fileURL)
	if err != nil {
		return fmt.Errorf("failed to download photo: %v", err)
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	// Write the file contents
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save photo: %v", err)
	}

	return nil
}

func DownloadPhoto(url, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func EscapeMarkdownV2(text string) string {
    replacer := strings.NewReplacer(
        "_", "\\_",
        "*", "\\*",
        "[", "\\[",
        "]", "\\]",
        "(", "\\(",
        ")", "\\)",
        "~", "\\~",
        "`", "\\`",
        ">", "\\>",
        "#", "\\#",
        "+", "\\+",
        "-", "\\-",
        "=", "\\=",
        "|", "\\|",
        "{", "\\{",
        "}", "\\}",
        ".", "\\.",
        "!", "\\!",
    )
    return replacer.Replace(text)
}