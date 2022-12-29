package Structural

import "fmt"

type notificationResp struct {
	err           error
	destinationId string
}

type notification interface {
	send() notificationResp
}

type smsNotification struct {
	destinationId string
	msg           string
}

func (sn smsNotification) send() notificationResp {

	fmt.Println("Sending SMS notification")
	fmt.Println(sn.msg)

	if len(sn.destinationId) > 0 {
		fmt.Println("delivered to 11111111")
	}

	return simulateResponse(sn.destinationId)
}

type slackNotification struct {
	notification notification
}

func (s slackNotification) send() notificationResp {
	resp := s.notification.send()

	if len(resp.destinationId) > 0 {
		slackUser := "slack-user-eq"
		fmt.Println("sending to slack: ", slackUser)
	}

	return simulateResponse(resp.destinationId)
}

func run() {
	sms := smsNotification{
		destinationId: "1234",
		msg:           "message for sms",
	}

	slack := slackNotification{notification: sms}

	slack.send()
}

func simulateResponse(dest string) notificationResp {
	return notificationResp{
		err:           nil,
		destinationId: dest,
	}
}
