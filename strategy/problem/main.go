package main

import "fmt"

type NotificationService struct {
	notificationType string
}

func (ns *NotificationService) SendNotification() {
	if ns.notificationType == "SMS" {
		fmt.Println("Send SMS")
	} else if ns.notificationType == "Email" {
		fmt.Println("Send Email")
	}
}

func main() {
	ns := NotificationService{notificationType: "SMS"}
	ns.SendNotification()
}
