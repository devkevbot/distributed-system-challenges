package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastRequest struct {
	Type    string `json:"type"`
	Message int    `json:"message"`
}

type BroadcastResponse struct {
	Type string `json:"type"`
}

type ReadRequest struct {
	Type string `json:"type"`
}

type ReadResponse struct {
	Type     string `json:"type"`
	Messages []int  `json:"messages"`
}

type ToplogyRequest struct {
	Type     string              `json:"type"`
	Topology map[string][]string `json:"topology"`
}

type TopologyResponse struct {
	Type string `json:"type"`
}

func main() {
	n := maelstrom.NewNode()

	messages := []int{}

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body BroadcastRequest
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		message := body.Message
		messages = append(messages, message)

		return n.Reply(msg, BroadcastResponse{Type: "broadcast_ok"})
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		var body ReadRequest
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		return n.Reply(msg, ReadResponse{Type: "read_ok", Messages: messages})
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		var body ToplogyRequest
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		return n.Reply(msg, TopologyResponse{Type: "topology_ok"})
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
