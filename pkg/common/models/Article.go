package models

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/vault-thirteen/SPA/pkg/indexer/settings"
	"github.com/vault-thirteen/auxie/file"
)

const (
	TimeDateFormat = "2006-01-02 15:04" // YYYY-MM-DD HH:MM.
	DateFormat     = "2006-01-02"       // YYYY-MM-DD.
	TimeFormat     = "15:04"            //HH:MM
	AuthorNone     = "None"
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

type Article struct {
	DateTimeUTC time.Time
	Category    string
	Title       string
	Description string
	Content     string
	Icon        string
	Author      string
	CRC32Hex    string
}

func NewArticleFromFile(filePath string) (article *Article, err error) {
	var buf []byte
	buf, err = os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	rawData := new(RawArticle)
	err = json.Unmarshal(buf, rawData)
	if err != nil {
		return nil, err
	}

	article = &Article{
		//DateTimeUTC:     time.DateTimeUTC{},
		Category:    rawData.Category,
		Title:       rawData.Title,
		Description: rawData.Description,
		Content:     rawData.Content,
		Icon:        rawData.Icon,
		Author:      rawData.Author,
		CRC32Hex:    rawData.CRC32Hex,
	}

	article.DateTimeUTC, err = time.Parse(TimeDateFormat, rawData.DateUTC+" "+rawData.TimeUTC)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func NewArticle(
	category string,
	title string,
	description string,
	content string,
	icon string,
	author string,
) (article *Article, err error) {
	article = &Article{
		DateTimeUTC: time.Now(),
		Category:    category,
		Title:       title,
		Description: description,
		Content:     content,
		Icon:        icon,
		Author:      author,
	}

	if !article.HasData() {
		return nil, errors.New(ErrNoData)
	}

	article.FillCRC32()

	return article, nil
}

func (a *Article) HasData() bool {
	emptyTime := time.Time{}
	if a.DateTimeUTC == emptyTime {
		return false
	}
	if len(a.Category) == 0 {
		return false
	}
	if len(a.Title) == 0 {
		return false
	}
	if len(a.Description) == 0 {
		return false
	}
	if len(a.Content) == 0 {
		return false
	}
	if len(a.Icon) == 0 {
		return false
	}
	if len(a.Author) == 0 {
		return false
	}

	return true
}

func (a *Article) HasCRC32() bool {
	return len(a.CRC32Hex) > 0
}

func (a *Article) CalculateCRC32() (sum uint32) {
	return newRawArticle(a).calculateCRC32()
}

func (a *Article) FillCRC32() {
	a.CRC32Hex = strings.ToUpper(strconv.FormatUint(uint64(a.CalculateCRC32()), 16))
}

func (a *Article) CheckCRC32() (ok bool) {
	if !a.HasCRC32() {
		return false
	}

	return a.CRC32Hex == strings.ToUpper(strconv.FormatUint(uint64(a.CalculateCRC32()), 16))
}

func (a *Article) Serialize() (ba []byte, err error) {
	if !a.HasData() {
		return nil, errors.New(ErrNoData)
	}
	if !a.HasCRC32() {
		return nil, errors.New(ErrNoChecksum)
	}
	if !a.CheckCRC32() {
		return nil, errors.New(ErrBadChecksum)
	}

	ba, err = json.MarshalIndent(newRawArticle(a), "", "\t")
	if err != nil {
		return nil, err
	}

	ba = append(ba, []byte(NewLine)...)
	return ba, nil
}

func (a *Article) SaveAsFile(filePath string) (err error) {
	var buf []byte
	buf, err = a.Serialize()
	if err != nil {
		return err
	}

	var fileExists bool
	fileExists, err = file.FileExists(filePath)
	if err != nil {
		return err
	}

	if !fileExists {
		err = os.WriteFile(filePath, buf, settings.NewFilePerm)
		if err != nil {
			return err
		}

		return nil
	}

	// File already exists. Use a temporary file.
	tmpFileName := filePath + settings.TmpFileExt
	err = os.WriteFile(tmpFileName, buf, settings.NewFilePerm)
	if err != nil {
		return err
	}

	oldFileName := filePath + settings.OldFileExt
	err = os.Rename(filePath, oldFileName)
	if err != nil {
		return err
	}

	err = os.Rename(tmpFileName, filePath)
	if err != nil {
		return err
	}

	return nil
}
