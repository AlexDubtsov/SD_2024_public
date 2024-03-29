package textPrepare

import (
	"fmt"
	"strconv"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
)

func MalePrint(templateData *structures.Template_Basic_EditEvent) string {
	var males []structures.Basic_SinglePerson
	numberParticipants := 0
	var numberParticipantsString string

	for i := range templateData.Slice_Participants {
		if templateData.Slice_Participants[i].Gender == "Male" {
			males = append(males, templateData.Slice_Participants[i])
			numberParticipants++
		}
	}

	numberParticipantsString = strconv.Itoa(numberParticipants)

	initString := templateData.Name + "\t" + templateData.Date + "\t" + "MALE" + "\t" + "Total " + numberParticipantsString + "\n\n"

	stringForPrint := initString + printTablePrepare(males)
	return stringForPrint
}

func FemalePrint(templateData *structures.Template_Basic_EditEvent) string {
	var females []structures.Basic_SinglePerson
	numberParticipants := 0
	var numberParticipantsString string

	for i := range templateData.Slice_Participants {
		if templateData.Slice_Participants[i].Gender == "Female" {
			females = append(females, templateData.Slice_Participants[i])
			numberParticipants++
		}
	}

	numberParticipantsString = strconv.Itoa(numberParticipants)

	initString := templateData.Name + "\t" + templateData.Date + "\t" + "FEMALE" + "\t" + "Total " + numberParticipantsString + "\n\n"

	stringForPrint := initString + printTablePrepare(females)
	return stringForPrint
}

func printTablePrepare(inputStruct []structures.Basic_SinglePerson) string {
	var resultString string
	for i := range inputStruct {
		resultString = resultString + fmt.Sprint(inputStruct[i].BageID) + "\t" + inputStruct[i].Name + "\t" + inputStruct[i].Email + "\t" + inputStruct[i].Phone
		if len(inputStruct[i].Comment) > 0 {
			resultString = resultString + "\t" + inputStruct[i].Comment
		}
		if i < len(inputStruct)-1 {
			resultString = resultString + "\n"
		}
	}
	return resultString
}
