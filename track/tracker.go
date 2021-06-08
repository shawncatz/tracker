package track

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

type Tracker struct {
	Directory string
	Tracks    map[string]*Track
}

type Track struct {
	Name    string
	Entries []string
	Updated time.Time
	Created time.Time
}

func NewTracker(directory string) (*Tracker, error) {
	tr := &Tracker{Directory: directory, Tracks: map[string]*Track{}}
	files, err := ioutil.ReadDir(tr.Directory)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(".json$")

	for _, f := range files {
		if re.MatchString(f.Name()) {
			n := strings.ReplaceAll(f.Name(), ".json", "")

			t, err := tr.read(n)
			if err != nil {
				return tr, err
			}

			tr.Tracks[n] = t
		}
	}

	return tr, nil
}

func (tr *Tracker) NewTrack(name string) error {
	track := &Track{}

	if tr.Tracks[name] != nil {
		return fmt.Errorf("track %s already exists", name)
	}

	if _, err := os.Stat(""); !os.IsNotExist(err) {
		return fmt.Errorf("track file %s already exists", name)
	}

	if err := tr.write(name, track); err != nil {
		return err
	}

	tr.Tracks[name] = track

	return nil
}

//func (tr *Tracker) LoadTrack(name string) (*Track, error) {
//	return tr.read(name)
//}

func (tr *Tracker) Append(name, value string) error {
	t := tr.Tracks[name]
	t.Entries = append(t.Entries, value)
	return tr.write(name, t)
}

func (tr *Tracker) Save(name string) error {
	t := tr.Tracks[name]
	return tr.write(name, t)
}

func (tr *Tracker) read(name string) (*Track, error) {
	t := &Track{}

	b, err := ioutil.ReadFile(filename(tr.Directory, name))
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(b, t)
	return t, err
}

func (tr *Tracker) write(name string, t *Track) error {
	t.Name = name
	now := time.Now()
	if t.Created.IsZero() {
		t.Created = now
	}
	t.Updated = now

	b, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename(tr.Directory, name), b, 0644)
}

func filename(dir, name string) string {
	return fmt.Sprintf("%s/%s.json", dir, name)
}
