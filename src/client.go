package main

import (
	"fmt"
	"mac/src/imap"
	"net"
)

type Client struct {
    Conn net.Conn
    Mbox *imap.Mailbox
}

func NewClient(server string, port uint16) (*Client, error) {
    conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", server, port))

    if err != nil {
        return nil, err
    }

    return &Client{
        Conn: conn,
    }, nil
}

func DeleteMailbox(mbox *imap.Mailbox) {}

func RenameMailbox(mbox *imap.Mailbox) {}
