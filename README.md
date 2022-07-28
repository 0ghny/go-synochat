# go-synochat

Synology Chat client library in Golang

## Why ?

Needs of connect using go code with [Synology Chat Server](https://www.synology.com/es-es/dsm/feature/chat).

It's follow [Synology Chat integration](https://kb.synology.com/en-in/DSM/help/Chat/chat_integration?version=7).

## Installs

```shell
go get github.com/0ghny/go-synochat
```

## Usage

```go
// Send a message using an incoming hook to a channel
c, _ := synochat.NewClient("https://your_synology_url") // or your_synology_url:port
err := c.SendMessage(&synochat.ChatMessage{Text: "Hello from automated test"}, "your incoming webhook token")
```

To learn how to configure an incoming webhook [read this](https://kb.synology.com/en-in/DSM/tutorial/How_to_configure_webhooks_and_slash_commands_in_Chat_Integration#x_anchor_id5).

Read [Incoming webhooks](https://kb.synology.com/en-in/DSM/help/Chat/chat_integration?version=7) to see how format "text" variable to print links, etc.
