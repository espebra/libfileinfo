package libfileinfo

import (
	"testing"
)

func TestOpen(t *testing.T) {
	// File not found
	f, err := Open("testing/non-existing")
	if err == nil {
		t.Fatal("Error: File not found didn't trigger an error.")
	}

	// Jpg
	f, err = Open("testing/IMG_3679.jpg")
	if err != nil {
		t.Fatal("Error:", err)
	}

	if f.Bytes != 3556461 {
		t.Fatal("Wrong data: bytes")
	}

	if f.MIME != "image/jpeg" {
		t.Fatal("Wrong data: mime")
	}

	if f.MediaType != "image" {
		t.Fatal("Wrong data: mediatype")
	}

	if f.Date != "2015-02-28" {
		t.Fatal("Wrong data: date")
	}

	if f.Time != "07-18-25" {
		t.Fatal("Wrong data: date")
	}

	//if f.DateTime.String() != "2015-02-28 07:18:25 +0100 CET" {
	//	t.Fatal("Wrong data: datetime, " + f.DateTime.String())
	//}

	//if f.Mtime.String() != "2015-12-25 19:52:31 +0000 UTC" {
	//	t.Fatal("Wrong data: mtime, " + f.Mtime.String())
	//}

	if f.Checksum != "a95bd0fc493df3d89397e3a2c4f69bbc127542a5a071cdc7f0ed0828d1b9f6d8" {
		t.Fatal("Wrong data: checksum")
	}
}

func TestCalculateChecksum(t *testing.T) {
	// File not found
	f := Path{}
	f.Path = "testing/non-existing"

	err := f.CalculateChecksum()
	if err == nil {
		t.Fatal("Error: File not found didn't trigger an error.")
	}
}

func TestDetectMIME(t *testing.T) {
	// File not found
	f := Path{}
	f.Path = "testing/non-existing"

	err := f.DetectMIME()
	if err == nil {
		t.Fatal("Error: File not found didn't trigger an error.")
	}
}

func TestParseExif(t *testing.T) {
	// File not found
	f := Path{}
	f.Path = "testing/non-existing"

	err := f.ParseExif()
	if err == nil {
		t.Fatal("Error: File not found didn't trigger an error.")
	}
}

func TestIsFile(t *testing.T) {
	// File not found
	if isFile("testing/non-existing") {
		t.Fatal("Error: File not found not detected.")
	}

	if isFile("testing/") {
		t.Fatal("Error: Path is a directory not detected.")
	}

	if !isFile("testing/IMG_3679.jpg") {
		t.Fatal("Error: File not detected")
	}
}

func TestStat(t *testing.T) {
	// File not found
	f := Path{}
	f.Path = "testing/non-existing"

	err := f.Stat()
	if err == nil {
		t.Fatal("Error: File not found didn't trigger an error.")
	}
}

