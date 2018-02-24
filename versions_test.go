package versions

import (
	"testing"
)

func Test_String_encodesCorrectly(t *testing.T) {
	expected := "1.2.3.4-rc"
	version := NewVersion(1, 2, 3, 4, "rc")

	actual := version.String()

	if expected != actual {
		t.Fatalf("%s != %s", expected, actual)
	}
}

func Test_ParseString_invalidFormats(t *testing.T) {
	formats := []string{
		"",
		"1.1.1.1.1.1",
		"1-ssuff-ffu",
		"1--ssuff",
		"123456789123456789123456789",
	}
	for _, format := range formats {
		if _, err := ParseString(format); err != ErrInvalidVersionFormat {
			t.Fatalf("expected invalid version format error when parsing '%s', got: %s", format, err)
		}
	}
}

func Test_ParseString_validFormats(t *testing.T) {
	cases := map[string]Version{
		"1.1.1.1": NewVersion(1, 1, 1, 1, ""),
	}

	for toParse, expected := range cases {
		parsed, err := ParseString(toParse)
		if err != nil {
			t.Fatalf("expected valid version format when parsing '%s', got: %s", toParse, err)
		} else if !Equal(expected, parsed) {
			t.Fatalf("expected %v, but got %v", expected, parsed)
		}
	}

}
