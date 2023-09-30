package structure

type Farm struct {
	AntsNumber int
	StartRoom  string
	EndRoom    string
	AllPaths   []string
	AllRooms   []string
}

type Room struct {
	RoomName   string
	RoomStatus bool
}

type AntPath struct {
	PathIndex       int
	TotalOfAntEroom int
}

type Coordinates struct {
	X string
	Y string
}

type Context struct {
	Data string
}

// Define a struct to hold the data you want to pass to the HTML template.
type PageData struct {
	AntsData []Ant
}

// Define a struct to represent an ant.
type Ant struct {
	Name      string
	Movements []Movement
}

// Define a struct to represent a movement.
type Movement struct {
	X    int
	Y    int
	Step int
}
