package providers

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

const TextbeltAPIEndpoint = "https://textbelt.com/text"
const TextbeltAPIKey = "textbelt"

func GenCaptchaCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var codes [6]byte
	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + r.Intn(10))
	}

	return string(codes[:])
}

func SendSMSWithTextbelt(phone string, message string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"phone":   phone,
		"message": message,
		"key":     TextbeltAPIKey,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(TextbeltAPIEndpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
