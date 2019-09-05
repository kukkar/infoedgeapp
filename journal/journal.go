package journal

import (
	"fmt"
	"time"
)

type Journal struct {
	Session string
}

type JournalData struct {
	Time      time.Time
	ToRemeber string
}

func GetInstace(sessionId string) (*Journal, error) {

	return &Journal{
		Session: "121",
	}, nil
}

func (this *Journal) ListEntries() error {

	fmt.Println("lilll")
	return nil

}

func (this *Journal) InputEntry(entry string) error {

	return nil
}
