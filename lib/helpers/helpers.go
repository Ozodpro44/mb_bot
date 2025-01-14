package helpers

import (
	"bot/models"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
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

	// Check for international format (starts with +, followed by country code)
	if len(digits) >= 10 {
		if strings.HasPrefix(digits, "998") && len(digits) == 12 {
			// Uzbekistan format: +998 xx xxx xx xx
			return fmt.Sprintf("+%s %s %s %s %s",
				digits[:3],         // +998
				digits[3:5],        // xx
				digits[5:8],        // xxx
				digits[8:10],       // xx
				digits[10:12]), nil // xx
		} else if strings.HasPrefix(digits, "90") && len(digits) == 12 {
			// Turkey format: +90 xxx xxx xxxx
			return fmt.Sprintf("+%s %s %s %s",
				digits[:2],        // +90
				digits[2:5],        // xxx
				digits[5:8],        // xxx
				digits[8:12]), nil  // xxxx
		} else if len(digits) > 10 {
			// General international format: +CC xxxx...xxxx
			return fmt.Sprintf("+%s %s",
				digits[:len(digits)-10], // Country code (variable length)
				digits[len(digits)-10:]), nil
		}
	}

	return "", fmt.Errorf("invalid or unsupported phone number format")
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

func fGetAddressFromCoordinates(lat, lon float32) (string, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?lat=%f&lon=%f&format=json", lat, lon)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to call geocoding API: %v", err)
	}
	defer resp.Body.Close()

	var geocodeResp models.GeocodingResponse
	if err := json.NewDecoder(resp.Body).Decode(&geocodeResp); err != nil {
		return "", fmt.Errorf("failed to decode geocoding response: %v", err)
	}

	address := fmt.Sprintf("📍 %s, %s, %s, улица %s",
		geocodeResp.Address.Country,
		geocodeResp.Address.State,
		geocodeResp.Address.Suburb,
		geocodeResp.Address.Road,
	)
	return address, nil
}

func GetAddressFromCoordinates2(lat, lon float32, lang string) (string, error) {
	switch lang {
	case "ru":
		lang = "ru_RU"
	case "uz":
		lang = "uz_UZ"
	case "en":
		lang = "en_US"
	case "tr":
		lang = "tr_TR"
	default:
		lang = "ru_RU"
	}

	apiKey := "5509ec40-9d43-49cd-b1d7-775a86e210e0"
	url := fmt.Sprintf("https://geocode-maps.yandex.ru/1.x/?apikey=%s&geocode=%f,%f&format=json&lang=%s", apiKey, lon, lat, lang)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("не удалось выполнить запрос к API Яндекс.Карт: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("не удалось прочитать ответ API: %v", err)
	}

	var geocodeResp map[string]interface{}
	if err := json.Unmarshal(body, &geocodeResp); err != nil {
		return "", fmt.Errorf("не удалось декодировать ответ API: %v", err)
	}

	// Извлечение адреса из ответа
	featureMembers, ok := geocodeResp["response"].(map[string]interface{})["GeoObjectCollection"].(map[string]interface{})["featureMember"].([]interface{})
	if !ok || len(featureMembers) == 0 {
		return "", fmt.Errorf("адрес по указанным координатам не найден")
	}

	geoObject := featureMembers[0].(map[string]interface{})["GeoObject"].(map[string]interface{})
	addressDetails, ok := geoObject["metaDataProperty"].(map[string]interface{})["GeocoderMetaData"].(map[string]interface{})["AddressDetails"].(map[string]interface{})["Country"].(map[string]interface{})["AdministrativeArea"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("не удалось извлечь адрес из ответа API")
	}

	city := addressDetails["AdministrativeAreaName"].(string)
	locality := addressDetails["Locality"].(map[string]interface{})["LocalityName"].(string)
	street := addressDetails["Locality"].(map[string]interface{})["Thoroughfare"].(map[string]interface{})["ThoroughfareName"].(string)
	house := addressDetails["Locality"].(map[string]interface{})["Thoroughfare"].(map[string]interface{})["Premise"].(map[string]interface{})["PremiseNumber"].(string)

	address := fmt.Sprintf("📍 %s, %s, %s, дом %s", city, locality, street, house)
	return address, nil
}

func GetAddressFromCoordinates(lat, lon float32, lang string) (string, error) {
	switch lang {
	case "ru":
		lang = "ru_RU"
	case "uz":
		lang = "uz_UZ"
	case "en":
		lang = "en_US"
	case "tr":
		lang = "tr_TR"
	default:
		lang = "ru_RU"
	}
	apiKey := "5509ec40-9d43-49cd-b1d7-775a86e210e0"
	url := fmt.Sprintf("https://geocode-maps.yandex.ru/1.x/?apikey=%s&geocode=%f,%f&format=json&lang=%s", apiKey, lon, lat, lang)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("не удалось выполнить запрос к API Яндекс.Карт: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("не удалось прочитать ответ API: %v", err)
	}

	var geocodeResp map[string]interface{}
	if err := json.Unmarshal(body, &geocodeResp); err != nil {
		return "", fmt.Errorf("не удалось декодировать ответ API: %v", err)
	}

	// Проверка наличия featureMember
	response, ok := geocodeResp["response"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("ответ API не содержит ключ 'response'")
	}

	geoObjectCollection, ok := response["GeoObjectCollection"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("ответ API не содержит ключ 'GeoObjectCollection'")
	}

	featureMembers, ok := geoObjectCollection["featureMember"].([]interface{})
	if !ok || len(featureMembers) == 0 {
		return "", fmt.Errorf("адрес по указанным координатам не найден")
	}

	geoObject, ok := featureMembers[0].(map[string]interface{})["GeoObject"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("не удалось извлечь 'GeoObject'")
	}

	metaDataProperty, ok := geoObject["metaDataProperty"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("не удалось извлечь 'metaDataProperty'")
	}

	geocoderMetaData, ok := metaDataProperty["GeocoderMetaData"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("не удалось извлечь 'GeocoderMetaData'")
	}

	address, ok := geocoderMetaData["Address"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("не удалось извлечь 'AddressDetails'")
	}

	addressDetails, ok := address["formatted"].(string)
	if !ok {
		return "", fmt.Errorf("не удалось извлечь 'Country'")
	}

	add := fmt.Sprintf("📍 %s", addressDetails)
	return add, nil
}

func Haversine(lat1, lon1, lat2, lon2 float64) bool {
	const R = 6371 // Радиус Земли в
	// Переводим градусы в радианы
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180
	lat1 = lat1 * math.Pi / 180
	lat2 = lat2 * math.Pi / 180

	// Формула гаверсинуса
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c <= 0.6
}

