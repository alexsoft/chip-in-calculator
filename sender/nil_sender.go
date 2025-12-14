package sender

import "fmt"

type NilSender struct {
}

func (s *NilSender) Send(message string) error {
	fmt.Println("Sending message via NilSender:", message)

	return nil
}
