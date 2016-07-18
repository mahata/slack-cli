package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"
)

// Config represents Slack configuration data
type Config struct {
	Token    string `json:"token"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
}

func main() {
	var config Config

	usr, err := user.Current()
	if err != nil {
		os.Stderr.WriteString("Can't get current user info... Are you an alien?\n")
		os.Exit(1)
	}
	confFile := usr.HomeDir + "/.slack-cli.json"

	file, err := ioutil.ReadFile(confFile)
	if err != nil {
		os.Stderr.WriteString("Can't find ~/.slack-cli.json.\n")
		os.Exit(1)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		os.Stderr.WriteString("Can't parse config.json file. Is the format correct?\n")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		os.Stderr.WriteString("Usage: slack-cli MESSAGE-TO-SEND\n")
		os.Exit(1)
	}

	query := fmt.Sprintf("token=%s&channel=%s&username=%s&text=%s",
		config.Token,
		config.Channel,
		config.Username,
		strings.Join(os.Args[1:], " "))

	body := strings.NewReader(query)
	req, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", body)
	if err != nil {
		os.Stderr.WriteString("Can't post to Slack API.\n")
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Stderr.WriteString("Can't post to Slack API.\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
}
