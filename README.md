# spidee
Uncomplicated CLI tool for executing Discord webhooks from the terminal or in shell scripts. Can send, edit, delete messages with plain text, embeds, and/or files with a custom username and avatar.

## Usage
Download the [latest release](https://github.com/rowlul/spidee/releases), open the terminal in the same directory as the executable and run:
```
$ ./spidee
```

## Installation
### Go
- Install [Go](https://go.dev/dl/)
- Ensure `$GOBIN` is in `PATH`
- Run `go install github.com/rowlul/spidee@latest`

## Examples
Supply ID and token every time you execute a webhook or store in the environment:
```
$ spidee --id ID --token TOKEN [...]

# ..or just once

$ export SPIDEE_WEBHOOK_ID="ID"
$ export SPIDEE_WEBHOOK_TOKEN="TOKEN"
```
Send a plain text message:
```
$ spidee send --content "Hello world"
```
Send a plain text message with attachments:
```
$ spidee send --content "Here are the files" --file file1.txt --file file2.png
```
Send an embed with fields: (properties go by the order `name`,`value`,`inline`)
```
$ spidee send --embed --embed-title "Embeds are cool" --embed-field "Some field","Some value",false --embed-field "Another field","Yet another value",false
```
Send an embed with a timestamp:
```
$ spidee send --embed --embed-title "Some embed" --embed-timestamp now

# ..or custom RFC3339 timestamp

$ spidee send --embed --embed-title "Some embed" --embed-timestamp 2015-12-31T12:00:00.000Z
```
