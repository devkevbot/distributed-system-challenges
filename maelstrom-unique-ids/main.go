package main

import (
	"encoding/json"
	"fmt"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		// Unmarshal the message body as a loosely-typed map
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// The node id + the msg_id will always be unique
		body["id"] = fmt.Sprintf("%s-%v", n.ID(), body["msg_id"])
		body["type"] = "generate_ok"

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
