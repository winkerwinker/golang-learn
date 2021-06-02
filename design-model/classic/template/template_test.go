package template

import (
	"testing"
)

func TestNewTelecomSms(t *testing.T) {

	tel := NewTelecomSms()
	tel.Send("test", 1239999)
}
