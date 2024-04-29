package pkg

import (
	"fmt"
)

const (
	DefaultChime     = "default"
	TraditionalChime = "traditional"
	MusicChime       = "music"
)

var embeddedChimes = map[string]string{
	DefaultChime:     "default.mp3",
	TraditionalChime: "traditional.mp3",
	MusicChime:       "music.mp3",
}

type DoorBell struct {
	chimes     map[string]string
	currChime  string
	volume     uint8
	onBellRing *Hook[*BellRingEvent]
}

func NewDoorBell() *DoorBell {
	return &DoorBell{
		chimes:     embeddedChimes,
		currChime:  "default",
		onBellRing: &Hook[*BellRingEvent]{},
		volume:     0,
	}
}

func (d *DoorBell) GetSound(name string) (string, error) {
	if path, ok := d.chimes[name]; !ok {
		return "", fmt.Errorf("Chime with name %s doesn't exist", name)
	} else {
		return path, nil
	}
}

func (d *DoorBell) Ring() (string, uint8, error) {
	err := d.OnBellRing().Trigger(&BellRingEvent{
		Sound:  d.currChime,
		Path:   d.chimes[d.currChime],
		Volume: d.volume,
	})
	if err != nil {
		return "", 0, err
	}
	return d.currChime, d.volume, nil
}

func (d *DoorBell) SetVolume(volume uint8) {
	d.volume = volume
}

func (d *DoorBell) AddNewChime(name string, path string) error {
	if _, ok := d.chimes[name]; ok {
		return fmt.Errorf("Chime with name %s already exists", name)
	}
	d.chimes[name] = path
	return nil
}

func (d *DoorBell) RemoveChime(name string) error {
	if _, ok := d.chimes[name]; !ok {
		return fmt.Errorf("Chime with name %s doesn't exist", name)
	}
	delete(d.chimes, name)
	return nil
}

func (d *DoorBell) SetChimeSound(chime string) error {
	if _, ok := d.chimes[chime]; !ok {
		return fmt.Errorf("Chime with name %s doesn't exist, consider adding it first", chime)
	}
	d.currChime = chime
	return nil
}

func (d *DoorBell) OnBellRing() *Hook[*BellRingEvent] {
	return d.onBellRing
}
