package guestbook

type GuestBookEntry struct {
	Id      int
	Email   string
	Title   string
	Content string
}

type GuestBook struct {
	guestBookData []*GuestBookEntry
}
