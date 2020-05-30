package speak_test

import (
	"testing"

	. "testify-mock/speak"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type greetingMock struct {
	mock.Mock
}

func (o *greetingMock) Say(message string) string {
	args := o.Called()
	return args.String(0)
}

func (o *greetingMock) Handshake() {
	o.Called()
}

func TestSay(t *testing.T) {
	mockTest := new(greetingMock)
	mockTest.On("Handshake").Return()
	mockTest.On("Say").Return("Hi there!")

	resultMessage := Speak(mockTest)

	mockTest.AssertExpectations(t)

	assert.NotEmpty(t, resultMessage)
	assert.Equal(t, "Hi there!", resultMessage)
}
