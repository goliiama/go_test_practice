package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
	text := fmt.Sprintf("%v", msg.getMessage())
	cost := len(text) * 3
	return text, cost
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func main() {
	msg := birthdayMessage{time.Now(), "Naz"}
	text, cost := sendMessage(msg)
	fmt.Println(text, cost)
}
