package main

import (
	"fmt"
	"context"

	raftpb "go.etcd.io/etcd/raft/v3/raftpb"
)

func main() {

	fmt.Println("### Golang Snippets ###")

	// create new Raft message
	msg := &raftpb.Message{
		Type:    raftpb.MsgHup,
		From:    1,
		To:      2,
		Term:    3,
		LogTerm: 4,
		Index:   5,
		Entries: []raftpb.Entry{{Index: 6, Term: 7}},
		Commit:  8,
	}

	// Use the raftpb.Message to send a message from one Raft node to another.
	sendMessage(context.Background(), msg)
}

func sendMessage(ctx context.Context, msg *raftpb.Message) {
	// Send the message...
}