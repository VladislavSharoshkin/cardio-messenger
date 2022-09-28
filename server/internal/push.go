package internal

import (
	"awesomeProject/database"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
	"log"
)

var fcmClient *messaging.Client


func SendPushNotifications(tokens []string, title string, body string) error {

	var messages []*messaging.Message
	for _, token := range tokens {
		messages = append(messages, &messaging.Message {
			Notification: &messaging.Notification {
				Title: title,
				Body:  body,
			},
			APNS: &messaging.APNSConfig{Payload: &messaging.APNSPayload{Aps: &messaging.Aps{Sound: "default"}}},
			Token: token, // a token that you received from a client
		})
	}

	_, err := fcmClient.SendAll(database.Ctx, messages)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	opts := []option.ClientOption{option.WithCredentialsFile("cardiomessenger-firebase-adminsdk-k00j0-036a6227ba.json")}

	app, err := firebase.NewApp(database.Ctx, nil, opts...)
	if err != nil {
		log.Fatalf("new firebase app: %s", err)
	}

	fcmClient, err = app.Messaging(context.TODO())
	IfLogError(err, 217941)
}
