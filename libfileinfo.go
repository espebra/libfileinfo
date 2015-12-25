package libfileinfo

import (
	"os"
	"io"
	"time"
	"errors"
	"regexp"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/rwcarlsen/goexif/exif"

        //"github.com/dustin/go-humanize"
	//"github.com/disintegration/imaging"
)

type Path struct {
        Path			string		`json:"path"`
        Bytes			int64		`json:"bytes"`
        Mtime			time.Time	`json:"mtime"`
        MIME			string		`json:"mime"`
        MediaType		string		`json:"mediatype"`
        Checksum		string		`json:"checksum"`
        Algorithm		string		`json:"algorithm"`
        Model			string		`json:"model,omitempty"`
        DateTime		time.Time	`json:"datetime,omitempty"`
        Date			string		`json:"date,omitempty"`
        Time			string		`json:"time,omitempty"`
        Latitude		float64		`json:"latitude,omitempty"`
        Longitude		float64		`json:"longitude,omitempty"`
        Exif			*exif.Exif	`json:"-"`
}

func Open(fpath string) (Path, error) {
	var err error
	f := Path {}
	if !isFile(fpath) {
		err = errors.New("The specified file does not exist.")
		return f, err
	}
	f.Path = fpath
	f.CalculateChecksum()
	f.DetectMIME()
	f.Stat()
	f.ParseExif()
	
	return f, nil
}

func (f *Path) CalculateChecksum() error {
	var err error
	var result []byte
	fp, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer fp.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, fp)
	if err != nil {
		return err
	}
	f.Checksum = hex.EncodeToString(hash.Sum(result))
	f.Algorithm = "SHA256"
	return nil
}

func (f *Path) DetectMIME() error {
	var err error
	fp, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer fp.Close()
	buffer := make([]byte, 512)
	_, err = fp.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = fp.Read(buffer)
	if err != nil {
		return err
	}
	f.MIME = http.DetectContentType(buffer)

	s := regexp.MustCompile("/").Split(f.MIME, 2)
	if len(s) > 0 {
		f.MediaType = s[0]
	}

	return nil
}

func (f *Path) ParseExif() error {
	var err error
	fp, err := os.Open(f.Path)
	defer fp.Close()
	if err != nil {
		return err
	}
	
	f.Exif, err = exif.Decode(fp)
	if err != nil {
		return err
	}

	f.DateTime, err = f.Exif.DateTime()
	if err != nil {
		return err
	}
	f.Date = f.DateTime.Format("2006-01-02")
	f.Time = f.DateTime.Format("15-04-05")

	f.Latitude, f.Longitude, err = f.Exif.LatLong()
	if err != nil {
		return err
	}

	v, err := f.Exif.Get(exif.Model)
	if err == nil {
		f.Model, _ = v.StringVal()
	}
	return nil
}

func isFile(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return false
	} else {
		return true
	}
}

func (f *Path) Stat() error {
	i, err := os.Lstat(f.Path)
	if err != nil {
		return err
	}
	f.Bytes = i.Size()
	f.Mtime = i.ModTime().UTC()
	return nil
}
