package controllers

import (
	"gopkg.in/maddevsio/fcm.v1"
	"fmt"
)

func SendNotification() {
	data := map[string]string{
		"msg": "Test",
		"sum": "Test",
	}
	c := fcm.NewFCM("serverKey")
	token := "token"
	response, err := c.Send(fcm.Message{
		Data:             data,
		RegistrationIDs:  []string{token},
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
		Notification: fcm.Notification{
			Title: "Test",
			Body:  "Test",
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Status Code   :", response.StatusCode)
	fmt.Println("Success       :", response.Success)
}
