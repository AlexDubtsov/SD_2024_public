package sorting

import (
	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
)

// Bubble sorting forever!
func Basic_SortingGender() {

}

func Basic_SortingAll(templateData *structures.Template_Basic_EditEvent) {
	var sorted int = 1
	for sorted > 0 {
		sorted = 0
		for i := 1; i < len(templateData.Slice_Participants); i++ {
			var tempBasic_SinglePerson structures.Basic_SinglePerson
			if templateData.Slice_Participants[i].BageID < templateData.Slice_Participants[i-1].BageID {
				tempBasic_SinglePerson = templateData.Slice_Participants[i]
				templateData.Slice_Participants[i] = templateData.Slice_Participants[i-1]
				templateData.Slice_Participants[i-1] = tempBasic_SinglePerson

				sorted++
			}
		}
	}
}

func Basic_SortingMale() {

}

func Basic_SortingFemale() {

}
