package main

import (
	"fmt"
	"mac/src/imap"
	"net"
)

type Client struct {
    Conn net.Conn
    Mailbox *imap.Mailbox
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

func DeleteMailbox(mailbox *imap.Mailbox) {}

func RenameMailbox(mailbox *imap.Mailbox) {}
