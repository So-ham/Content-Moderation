package tisane

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type tisane struct {
	baseURL string
}

type TiSane interface {
	CheckIfContentFlagged(content string) (bool, string, error)
}

func New() TiSane {

	baseURL := os.Getenv("TISANE_BASE_URL")

	return &tisane{
		baseURL: baseURL,
	}
}

const (
	apiUrl = "https://api.tisane.ai/parse" // API URL with environment variable
)

type Request struct {
	Language string                 `json:"language"`
	Content  string                 `json:"content"`
	Settings map[string]interface{} `json:"settings"`
}

type Abuse struct {
	SentenceIndex int    `json:"sentence_index"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	Type          string `json:"type"`
	Severity      string `json:"severity"`
}

type Response struct {
	Text  string  `json:"text"`
	Abuse []Abuse `json:"abuse"`
}

// CheckIfContentFlagged checks if abuse exists and returns boolean and severity
func (m *tisane) CheckIfContentFlagged(content string) (bool, string, error) {
	// Prepare the request body
	requestBody := Request{
		Language: "en",
		Content:  content,
		Settings: make(map[string]interface{}),
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return false, "", fmt.Errorf("error marshaling request body: %v", err)
	}

	// Make HTTP POST request
	req, err := http.NewRequest("POST", m.baseURL+"/parse", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, "", fmt.Errorf("error creating request: %v", err)
	}

	// Set the correct header with the Subscription Key
	apiKey := os.Getenv("TISANE_KEY")
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey) // Correct header key name
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, "", fmt.Errorf("error reading response body: %v", err)
	}

	// Parse the response JSON into the Response struct
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, "", fmt.Errorf("error unmarshaling response body: %v", err)
	}

	// Check if there is any abusive content in the "abuse" field
	if len(response.Abuse) > 0 {
		// Log the abuse details if necessary
		severity := response.Abuse[0].Severity // Get the severity of the first abuse found
		return true, severity, nil             // Content is flagged, return severity
	}

	return false, "", nil // No abuse found
}
