package helpers

type BasicSinglePerson struct {
	Email     string
	Name      string
	Gender    string
	Phone     string
	EventName string
	Created   string
	Likes     string
}

type ProcessTemplate struct {
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
