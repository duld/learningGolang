package doug

// IsYourName checks if you know the name of the individual you are
// speaking to. Returns 'Yes' if you are correct, but will correct
// you if you are wrong. You should be more considerate.
func IsYourName(name string) string {
	if name != "Doug" {
		return "No, my name is Doug."
	}
	return "Yes"
}

// DoYouKnow do you know this person? If so they owe me money.
func DoYouKnow(name string) bool {
	peopleIKnow := []string{
		"Susan McLady",
		"Bill BillsPerson",
		"Rogo Rogoralaloie pie",
		"Jeff",
	}

	// loop over the array checking if name is included in peopleIKnow
	for _, v := range peopleIKnow {
		if v == name {
			return true
		}
	}
	return false // I don't know that person.
}
