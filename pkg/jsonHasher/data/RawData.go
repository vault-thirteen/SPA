package data

import (
	"hash/crc32"
	"strings"
)

type RawData struct {
	DateUTC     string `json:"Date"`
	TimeUTC     string `json:"Time"`
	Category    string `json:"Category"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
	Icon        string `json:"Icon"`
	Author      string `json:"Author"`
	CRC32Hex    string `json:"CRC32"`
}

func newRawData(data *Data) (rd *RawData) {
	return &RawData{
		DateUTC:     data.DateTimeUTC.UTC().Format(DateFormat),
		TimeUTC:     data.DateTimeUTC.UTC().Format(TimeFormat),
		Category:    data.Category,
		Title:       data.Title,
		Description: data.Description,
		Content:     data.Content,
		Icon:        data.Icon,
		Author:      data.Author,
		CRC32Hex:    data.CRC32Hex,
	}
}

func (rd *RawData) calculateCRC32() (sum uint32) {
	var sb strings.Builder
	sb.WriteString(rd.DateUTC)
	sb.WriteString(rd.TimeUTC)
	sb.WriteString(rd.Category)
	sb.WriteString(rd.Title)
	sb.WriteString(rd.Description)
	sb.WriteString(rd.Content)
	sb.WriteString(rd.Icon)
	sb.WriteString(rd.Author)

	return crc32.ChecksumIEEE([]byte(sb.String()))
}
