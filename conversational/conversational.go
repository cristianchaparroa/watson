package watson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Credentials
type Credentials struct {
	Username string
	Password string
	Url      string
}

// ConversationService
type ConversationService struct {
	Credentials *Credentials
	WorkspaceId string
}

// NewConversationService
func NewConversationService(c *Credentials, workspaceId string) *ConversationService {
	return &ConversationService{Credentials: c, WorkspaceId: workspaceId}
}

// NewMessage, send a message for conversational service
func (c *ConversationService) NewMessage(m *MessageRequest) (*MessageResponse, error) {
	if c.Credentials.Username == "" && c.Credentials.Password == "" {
		return nil, errors.New("Credentials not provided")
	}
	url := fmt.Sprintf("%s/v1/workspaces/%s/message?version=2016-09-20", c.Credentials.Url, c.WorkspaceId)
	body, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Credentials.Username, c.Credentials.Password)

	defer req.Body.Close()
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, errors.New(string(bytes))
	}
	if err != nil {
		return nil, err
	}
	var response *MessageResponse
	json.Unmarshal(bytes, &response)
	return response, nil
}
