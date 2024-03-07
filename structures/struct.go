package structures

type Basic_SinglePerson struct {
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
	Comment   string
}

type Basic_SingleEvent struct {
	ID       int
	Name     string
	Date     string
	Spot     string
	AgeGroup string
	Comment  string
	Created  string
}

type Template_Basic_ListEvents struct {
	Slice_SingleEvent []Basic_SingleEvent
}

type Template_Basic_CreateEvent struct {
	ID                 int
	Name               string
	Date               string
	Text               string
	Message            string
	Slice_SinglePerson []Basic_SinglePerson
}

type Template_Basic_EditEvent struct {
	ID                 int
	Name               string
	Date               string
	Text               string
	Message            string
	Slice_Participants []Basic_SinglePerson
}
