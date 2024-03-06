package helpers

type BasicSinglePerson struct {
	ID        int
	Email     string
	Name      string
	Gender    string
	Phone     string
	BageID    int
	EventID   int
	EventName string
	EventDate string
	Created   string
	Likes     string
}

type BasicAllData struct {
	AllData []BasicSinglePerson
}

type BasicEvent struct {
	ID   int
	Name string
	Date string
}

type TemplateBasicEvents struct {
	AllEvents []BasicEvent
}

type TemplateBasicEventCreate struct {
	EventName string
	EventDate string
	Text      string
	Message   string
}

type SingleEvent struct {
	ID       string
	Name     string
	Time     string
	Spot     string
	AgeGroup string
	Comment  string
	Created  string
}

type EventTemplate struct {
	AllEvents []SingleEvent
}
