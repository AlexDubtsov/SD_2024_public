package helpers

// Function looks through all data and returns slice of unique events
func BasicEventsList() []BasicEvent {
	var eventsList []BasicEvent
	allData, _, _ := BasicReadoutDB()

	for _, allDataElement := range allData {
		duplicate := false
		for _, eventsListElement := range eventsList {
			if allDataElement.EventID == eventsListElement.ID {
				duplicate = true
			}
		}
		if !duplicate {
			var tempEvent BasicEvent
			tempEvent.ID = allDataElement.EventID
			tempEvent.Name = allDataElement.EventName
			tempEvent.Date = allDataElement.EventDate
			eventsList = append(eventsList, tempEvent)
		}
	}
	return eventsList
}
