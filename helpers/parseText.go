package helpers

import (
	"strings"
)

func ParseInputText(tempText string) ([]BasicSinglePerson, string) {
	var arrayBasicPersons []BasicSinglePerson
	var err bool
	var errorMessage string
	var bageID int

	maleNr := 1
	femaleNr := 2

	// Split input into lines
	lines := strings.Split(tempText, "\n")

	// Iterate over each line and parse the values
	for _, line := range lines {
		// Split line into fields
		fields := strings.Split(line, "\t")

		// Check if there are enough fields
		if len(fields) != 4 || !(fields[2] == "Male" || fields[2] == "Female") {
			err = true
			errorMessage = "invalid input line \"" + line + "\"\n"
		} else {
			switch fields[2] {
			case "Male":
				bageID = maleNr
				maleNr += 2
			case "Female":
				bageID = femaleNr
				femaleNr += 2
			}
			// Create a UserData structure and append it to the array
			tempBasicSinglePerson := BasicSinglePerson{
				Email:     fields[0],
				Name:      fields[1],
				Gender:    fields[2],
				Phone:     fields[3],
				BageID:    bageID,
				EventID:   0,
				EventName: "",
				EventDate: "",
				Created:   "",
				Likes:     "",
			}

			arrayBasicPersons = append(arrayBasicPersons, tempBasicSinglePerson)
		}

	}
	if err {
		return nil, errorMessage
	} else {
		return arrayBasicPersons, errorMessage
	}
}
