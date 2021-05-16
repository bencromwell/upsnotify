package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mailgun/mailgun-go/v4"
)

type upsNotification struct {
	notificationType    string
	upsName             string
	notificationMessage string
}

type emailNotifier struct {
	mailgun   *mailgun.MailgunImpl
	recipient string
	from      string
}

func (n emailNotifier) sendNotification(u upsNotification) {
	subject := fmt.Sprintf("UPS: %s (%s)", u.notificationType, u.upsName)
	body := u.notificationMessage

	message := n.mailgun.NewMessage(n.from, subject, body, n.recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := n.mailgun.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}

func main() {
	godotenv.Load("/etc/default/upsnotify")

	upsNotification := upsNotification{
		notificationType:    os.Getenv("NOTIFYTYPE"),
		upsName:             os.Getenv("UPSNAME"),
		notificationMessage: os.Args[1],
	}

	mailgun := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_PRIVATE_KEY"))

	notifier := emailNotifier{
		mailgun,
		os.Getenv("ALERT_EMAIL"),
		os.Getenv("ALERT_FROM"),
	}

	notifier.sendNotification(upsNotification)
}
