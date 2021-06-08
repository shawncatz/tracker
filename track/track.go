package track

import (
	"time"
)

type Track struct {
	Name    string
	Entries []Entry
	Updated time.Time
	Created time.Time
}

type Entry struct {
	Value string
}
