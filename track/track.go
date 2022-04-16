package track

import "time"

type Track struct {
	Name    string
	Entries []string
	Updated time.Time
	Created time.Time
}

func (t *Track) Count() int {
	return len(t.Entries)
}
