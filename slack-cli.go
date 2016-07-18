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

// Conf represents Slack configuration data
type Conf struct {
	Token    string `json:"token"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
}

var conf Conf

func main() {

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

	err = json.Unmarshal(file, &conf)
	if err != nil {
		os.Stderr.WriteString("Can't parse ~/.slack-cli.json. Is the format correct?\n")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		os.Stderr.WriteString("Usage: slack-cli MESSAGE-TO-SEND\n")
		os.Exit(1)
	}

	query := fmt.Sprintf("token=%s&channel=%s&username=%s&text=%s",
		conf.Token,
		conf.Channel,
		conf.Username,
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
