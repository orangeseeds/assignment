package pkg

import "testing"

func TestBellChimeTraditional(t *testing.T) {
	expected := TraditionalChime

	doorBell := NewDoorBell()
	err := doorBell.SetChimeSound(expected)
	if err != nil {
		t.Fatalf(err.Error())
	}

	got, _, err := doorBell.Ring()
	if err != nil {
		t.Fatal("error ringing bell", err.Error())
	}

	if expected != got {
		t.Fatalf("expected sound: %s but got %s", expected, got)
	}
}

func TestBellChimeCustom(t *testing.T) {
	expected := "voice clip"
	path := "voice_clip.mp3"

	doorBell := NewDoorBell()

	err := doorBell.AddNewChime(expected, path)
	if err != nil {
		t.Fatal("error adding chime sound", err.Error())
	}

	err = doorBell.SetChimeSound(expected)
	if err != nil {
		t.Fatal("error setting chime sound", err.Error())
	}

	got, _, err := doorBell.Ring()
	if err != nil {
		t.Fatal("error ringing bell", err.Error())
	}

	if expected != got {
		t.Fatalf("expected sound: %s but got %s", expected, got)
	}
}

func TestBellVolume(t *testing.T) {
	var expected uint8 = 100

	doorBell := NewDoorBell()
	doorBell.SetVolume(expected)

	_, got, err := doorBell.Ring()
	if err != nil {
		t.Fatal("error ringing bell", err.Error())
	}

	if expected != got {
		t.Fatalf("expected volume: %d but got %d", expected, got)
	}
}

func TestChimeAndVolume(t *testing.T) {
	var chime string = TraditionalChime
	var vol uint8 = 28

	doorBell := NewDoorBell()
	err := doorBell.SetChimeSound(chime)
	if err != nil {
		t.Fatal("error setting chime sound", err.Error())
	}
	doorBell.SetVolume(vol)

	gotChime, gotVol, err := doorBell.Ring()
	if err != nil {
		t.Fatal("error ringing bell", err.Error())
	}

	if gotChime != chime {
		t.Fatalf("expected sound: %s but got %s", chime, gotChime)
	}
	if gotVol != vol {
		t.Fatalf("expected volume: %d but got %d", vol, gotVol)
	}
}

func TestRingTriggers(t *testing.T) {
	var cases = []struct {
		name              string
		shouldBeTriggered bool
		isTriggered       bool
	}{
		{"Porch lights", true, false},
		{"Front Gate", true, false},
		{"Watering Hose", false, false},
	}

	doorBell := NewDoorBell()

	for i := range cases {
		if cases[i].shouldBeTriggered {
			doorBell.OnBellRing().Add(func(e *BellRingEvent) error {
				cases[i].isTriggered = true
				return nil
			})
		}
	}
	_, _, err := doorBell.Ring()
	if err != nil {
		t.Fatal("Error ringing bell", err.Error())
	}

	for i := range cases {
		if cases[i].shouldBeTriggered != cases[i].isTriggered {
			t.Fatalf("component: %s toBeTriggered: %v but isTriggered: %v", cases[i].name, cases[i].shouldBeTriggered, cases[i].isTriggered)
		}
	}

}
