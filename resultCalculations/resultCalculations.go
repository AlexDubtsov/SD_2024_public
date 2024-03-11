package resultCalculations

import (
	"fmt"
	"strings"

	"github.com/AlexDubtsov/SD_2024_public/m/v2/structures"
)

func ResultPrint(templateData *structures.Template_Basic_EditEvent) string {

	all_Members_Result := allMembersResultSliceCollection(templateData)

	likesConvertToSlice(&all_Members_Result)

	likesFilterMutual(&all_Members_Result)

	resultString := collectPrintString(&all_Members_Result)
	return resultString
}

func allMembersResultSliceCollection(templateData *structures.Template_Basic_EditEvent) []structures.Result_Basic {
	var result []structures.Result_Basic
	for i := range templateData.Slice_Participants {
		var tempMemberResult structures.Result_Basic
		tempMemberResult.ID = templateData.Slice_Participants[i].ID
		tempMemberResult.Email = templateData.Slice_Participants[i].Email
		tempMemberResult.Name = templateData.Slice_Participants[i].Name
		tempMemberResult.Phone = templateData.Slice_Participants[i].Phone
		tempMemberResult.BageID = templateData.Slice_Participants[i].BageID
		tempMemberResult.Likes = templateData.Slice_Participants[i].Likes
		tempMemberResult.Comment = templateData.Slice_Participants[i].Comment

		result = append(result, tempMemberResult)
	}
	return result
}

func likesConvertToSlice(all_Members_Result *[]structures.Result_Basic) {
	// Check if the pointer is not nil
	if all_Members_Result == nil {
		fmt.Println("Pointer == nil for some reason #0")
		return
	}

	// Dereference the pointer to get the slice
	members := *all_Members_Result

	// Clean Likes string out of non-number characters
	// Split likes string into slice of strings
	for i := range members {
		var separated string
		for j := range members[i].Likes {
			if members[i].Likes[j] >= '0' && members[i].Likes[j] <= '9' {
				separated = separated + string(members[i].Likes[j])
			} else {
				separated = separated + " "
			}

		}
		// Split the string into individual numbers
		members[i].Likes_Slice = strings.Fields(separated)

	}
}

func collectPrintString(all_Members_Result *[]structures.Result_Basic) string {
	var resultString string
	// Check if the pointer is not nil
	if all_Members_Result == nil {
		fmt.Println("Pointer == nil for some reason #1")
		return ""
	}

	// Dereference the pointer to get the slice
	members := *all_Members_Result

	// Collect result
	for i := range members {
		if i > 0 {
			resultString = resultString + "\n\n"
		}

		// Collect to result Member Name and Email
		resultString = resultString + "**************\n" + fmt.Sprint(members[i].BageID) + " " + members[i].Name + "\n" + members[i].Email + "\n\nLikes:\n"
		if len(members[i].Likes_Slice) == 0 {
			resultString = resultString + "Match number today = 0 \n\n"
		} else { // Collect to result Member[Like Number] Name and Email
			for j := range members[i].Likes_Slice {
				for k := range members {
					tempBageID := fmt.Sprint(members[k].BageID)
					if members[i].Likes_Slice[j] == tempBageID {
						resultString = resultString + tempBageID + " " + members[k].Name + "\n" + members[k].Email + "\n\n"
					}
				}
			}
		}
	}
	return resultString
}

func likesFilterMutual(all_Members_Result *[]structures.Result_Basic) {
	// Check if the pointer is not nil
	if all_Members_Result == nil {
		fmt.Println("Pointer == nil for some reason #2")
		return
	}

	// Dereference the pointer to get the slice
	members := *all_Members_Result

	for i := range members {
		requestMember := fmt.Sprint(members[i].BageID)
		var tempLikeSlice []string
		for j := range members[i].Likes_Slice {
			for k := range members {
				responseMember := fmt.Sprint(members[k].BageID)
				if responseMember == members[i].Likes_Slice[j] {
					for l := range members[k].Likes_Slice {
						if members[k].Likes_Slice[l] == requestMember {
							tempLikeSlice = append(tempLikeSlice, responseMember)
						}
					}
				}
			}
		}
		members[i].Likes_Slice = tempLikeSlice
	}
}
