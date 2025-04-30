package utils

import (
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

// Returns the information about the weather in today
//
// In case the data (weather info) is not cached at local.
// The data will be fetched from API and save at local.
func Weather() (string, error) {
	filePath, err := weatherTempFilePath()
	if err != nil {
		return "", err
	}
	// Check weather file exists
	_, err = os.Stat(filePath)
	if err == nil {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return "", errors.New("read error:" + err.Error())
		}
		return string(data), nil
	}

	res, err := http.Get("https://wttr.in/?format=3")
	if err != nil {
		return "", errors.New(err.Error())
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.New(err.Error())
	}

	err = os.WriteFile(filePath, body, 0644)
	if err != nil {
		return "", errors.New("write error:" + err.Error())
	}
	return string(body), nil
}

// Return the local path of weather data file in a specific day (today)
func weatherTempFilePath() (string, error) {
	tempFolder := os.TempDir() + "weather/"
	err := os.MkdirAll(tempFolder, 0755)
	if err != nil {
		return "", errors.New("create temp folder: " + err.Error())
	}
	today := time.Now().Format("2006-01-02")
	return tempFolder + today + ".txt", nil
}
