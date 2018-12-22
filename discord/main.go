package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	webhookID    = os.Getenv("WEBHOOK_ID")
	webhookToken = os.Getenv("WEBHOOK_TOKEN")
	message      = os.Getenv("MESSAGE")
	username     = os.Getenv("USERNAME")
	avatarURL    = os.Getenv("AVATAR_URL")
)

const (
	// Endpoint defines the Discord Webhook Endpoint used to send the message
	Endpoint = "https://discordapp.com/api/webhooks/%s/%s"
)

// Params represents the parameters sent as the payload to the Discord Webhook Endpoint
type Params struct {
	Content   string `json:"content,omitempty"`
	Username  string `json:"username,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

func main() {
	uri := fmt.Sprintf(Endpoint, webhookID, webhookToken)

	payload, err := json.Marshal(&Params{
		Content:   message,
		Username:  username,
		AvatarURL: avatarURL,
	})
	check(err)

	fmt.Println("Sending Discord Message to Webhook ID " + webhookID)
	_, err = http.Post(uri, "application/json", bytes.NewBuffer(payload))
	check(err)
}

// check panics if err is not nil
func check(err error) {
	if err != nil {
		panic(err)
	}
}
