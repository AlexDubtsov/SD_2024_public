package textParsing

import (
	"regexp"
	"strings"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
)

func isValidEmail(email string) bool {
	// Define a regular expression for a basic email validation
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	re := regexp.MustCompile(emailRegex)

	// Use the MatchString method to check if the email matches the pattern
	return re.MatchString(email)
}

func ParseFormText(formText string) ([]structures.Basic_SinglePerson, string) {
	var slice_SinglePerson []structures.Basic_SinglePerson
	var errStr string
	var err bool

	// Split input text into lines
	lines := strings.Split(formText, "\n")

	// Iterate over each line and parse the values
	for _, line := range lines {
		// Split line into fields
		fields := strings.Split(line, "\t")

		// Check if there are enough fields
		if len(fields) != 4 {
			err = true
			errStr = "Participants info should be in format \"EMAIL NAME GENDER PHONE\". Issue with line: \"" + line + "\"\n"
		} else if !isValidEmail(fields[0]) {
			err = true
			errStr = "Invalid Mail address in line \"" + line + "\"\n"
		} else if fields[2] != "Male" && fields[2] != "Female" {
			err = true
			errStr = "Invalid Gender in line \"" + line + "\"\n"
		} else {
			// Create a UserData structure and append it to the array
			tempBasicSinglePerson := structures.Basic_SinglePerson{
				ID:        0,
				Email:     fields[0],
				Name:      fields[1],
				Gender:    fields[2],
				Phone:     fields[3],
				BageID:    0,
				EventID:   0,
				EventName: "",
				EventDate: "",
				Created:   "",
				Likes:     "",
				Comment:   "",
			}

			slice_SinglePerson = append(slice_SinglePerson, tempBasicSinglePerson)
		}

	}
	if err {
		return nil, errStr
	} else {
		return slice_SinglePerson, errStr
	}
}
