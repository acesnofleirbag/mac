package imap

import "time"

type MsgFlag string

const (
    MsgFlag__Seen MsgFlag = `\Seen`
    MsgFlag__Answered MsgFlag = `\Answered`
    MsgFlag__Flagged MsgFlag = `\Flagged`
    MsgFlag__Deleted MsgFlag = `\Deleted`
    MsgFlag__Draft MsgFlag = `\Draft`
    MsgFlag__Recent MsgFlag = `\Recent`
)

type Keyword string

const (
    Keyword__Forwarded Keyword = "$Forwarded"
    Keyword__MDNSent Keyword = "$MDNSent"
    Keyword__Junk Keyword = "$Junk"
    Keyword__NotJunk Keyword = "$NotJunk"
    Keyword__Phishing Keyword = "$Phishing"
)

type UID uint32

type Message struct {
    Uid UID
    Flags []MsgFlag
    Date time.Time
    Size int
    Envelope Envelope
}
