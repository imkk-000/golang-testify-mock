package speak

import "testify-mock/greeting"

func Speak(g greeting.GreetingInterface) string {
	g.Handshake()
	return g.Say("Hi there!")
}
