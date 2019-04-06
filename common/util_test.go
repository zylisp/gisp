package common

import (
	"testing"
)

func TestRemoveExtension(t *testing.T) {
	wanted := "thing"
	filePart := RemoveExtension("thing.zsp")
	if filePart != wanted {
		t.Errorf("RemoveExtension filePart incorrect, got: %s, want: %s", filePart, wanted)
	}
}

func TestRemoveExtensionWithoutExtension(t *testing.T) {
	wanted := "thing"
	filePart := RemoveExtension("thing.")
	if filePart != wanted {
		t.Errorf("RemoveExtension filePart incorrect, got: %s, want: %s", filePart, wanted)
	}
}

func TestRemoveExtensionWithoutDot(t *testing.T) {
	wanted := "thing"
	filePart := RemoveExtension("thing")
	if filePart != wanted {
		t.Errorf("RemoveExtension filePart incorrect, got: %s, want: %s", filePart, wanted)
	}
	wanted = "thingzsp"
	filePart = RemoveExtension("thingzsp")
	if filePart != wanted {
		t.Errorf("RemoveExtension filePart incorrect, got: %s, want: %s", filePart, wanted)
	}
}
