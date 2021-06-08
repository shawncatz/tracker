package track

import "os"

func NewTrack(name string) error {
	if _, err := os.Stat(""); os.IsNotExist(err) {

	}
	return nil
}
