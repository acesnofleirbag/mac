package imap

type Mailbox struct {
    Uidnext UID
    Messages []Message
}

func (self *Mailbox) AddMessage(msg Message) {

}

func (self *Mailbox) RmMessage(uid UID) {

}
