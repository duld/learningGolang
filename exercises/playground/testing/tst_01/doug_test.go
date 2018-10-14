package doug

import (
	"testing"
)

func TestIsYourName(t *testing.T) {
	badCase := "Susan"
	passCase := "Doug"

	var v string
	v = IsYourName(passCase)
	if v != "Yes" {
		t.Error("Expected, 'yes', got: ", v)
	}

	v = IsYourName(badCase)
	if v == "Yes" {
		t.Error("Expected, 'No, my name is Doug' got: ", v)
	}
}

func TestDoYouKnow(t *testing.T) {
	friend := "Jeff"
	foe := "Susan McNotLady"

	var v bool
	v = DoYouKnow(friend)
	if v != true {
		t.Error("You are a liar, you definately know: ", v)
	}

	v = DoYouKnow(foe)
	if v != false {
		t.Error("You do not know this person. You are a living a lie: ", v)
	}
}
