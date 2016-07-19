# Slack Go Client

Slack Go Client. Only posting to a specific channel is supported currently.

## Install

Change `slack-cli.json` and copy that to `.slack-cli.json` under your `$HOME`.

## Run

Provided that you built `slack-cli` by `go build -o slack-cli slack-cli.go`:

```
$ slack-cli MESSAGE YOU WANT TO POST

(or)

$ slack-cli -ch=CHANNEL_TO_POST Hello

(or)

$ slack-cli -ch=CHANNEL_TO_POST -uname=JANE_DOE Hello
```
