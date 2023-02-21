package data

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	TimeDateFormat = "2006-01-02 15:04" // YYYY-MM-DD HH:MM.
	DateFormat     = "2006-01-02"       // YYYY-MM-DD.
	TimeFormat     = "15:04"            //HH:MM
)

const (
	ErrNoData      = "no data"
	ErrNoChecksum  = "no checksum"
	ErrBadChecksum = "bad checksum"
)

const (
	CR = "\r"
	LF = "\n"

	// NewLine is a new line.
	// This is the correct new line, as it used to be in the history of mankind.
	// If you are a user of Unix, Linux, Macintosh or something else, check the
	// Wikipedia: https://en.wikipedia.org/wiki/Newline
	NewLine = CR + LF
)

type Data struct {
	DateTimeUTC time.Time
	Category    string
	Title       string
	Description string
	Content     string
	Icon        string
	Author      string
	CRC32Hex    string
}

func NewFromFile(filePath string) (data *Data, err error) {
	var buf []byte
	buf, err = os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	rawData := new(RawData)
	err = json.Unmarshal(buf, rawData)
	if err != nil {
		return nil, err
	}

	data = &Data{
		//DateTimeUTC:     time.DateTimeUTC{},
		Category:    rawData.Category,
		Title:       rawData.Title,
		Description: rawData.Description,
		Content:     rawData.Content,
		Icon:        rawData.Icon,
		Author:      rawData.Author,
		CRC32Hex:    rawData.CRC32Hex,
	}

	data.DateTimeUTC, err = time.Parse(TimeDateFormat, rawData.DateUTC+" "+rawData.TimeUTC)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *Data) HasData() bool {
	emptyTime := time.Time{}
	if d.DateTimeUTC == emptyTime {
		return false
	}
	if len(d.Category) == 0 {
		return false
	}
	if len(d.Title) == 0 {
		return false
	}
	if len(d.Description) == 0 {
		return false
	}
	if len(d.Content) == 0 {
		return false
	}
	if len(d.Icon) == 0 {
		return false
	}
	if len(d.Author) == 0 {
		return false
	}

	return true
}

func (d *Data) HasCRC32() bool {
	return len(d.CRC32Hex) > 0
}

func (d *Data) CalculateCRC32() (sum uint32) {
	return newRawData(d).calculateCRC32()
}

func (d *Data) FillCRC32() {
	d.CRC32Hex = strings.ToUpper(strconv.FormatUint(uint64(d.CalculateCRC32()), 16))
}

func (d *Data) CheckCRC32() (ok bool) {
	if !d.HasCRC32() {
		return false
	}

	return d.CRC32Hex == strings.ToUpper(strconv.FormatUint(uint64(d.CalculateCRC32()), 16))
}

func (d *Data) Serialize() (ba []byte, err error) {
	if !d.HasData() {
		return nil, errors.New(ErrNoData)
	}
	if !d.HasCRC32() {
		return nil, errors.New(ErrNoChecksum)
	}
	if !d.CheckCRC32() {
		return nil, errors.New(ErrBadChecksum)
	}

	ba, err = json.MarshalIndent(newRawData(d), "", "\t")
	if err != nil {
		return nil, err
	}

	ba = append(ba, []byte(NewLine)...)
	return ba, nil
}
