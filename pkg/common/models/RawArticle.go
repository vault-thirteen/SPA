package models

import (
	"hash/crc32"
	"strings"
)

type RawArticle struct {
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

func newRawArticle(article *Article) (rd *RawArticle) {
	return &RawArticle{
		DateUTC:     article.DateTimeUTC.UTC().Format(DateFormat),
		TimeUTC:     article.DateTimeUTC.UTC().Format(TimeFormat),
		Category:    article.Category,
		Title:       article.Title,
		Description: article.Description,
		Content:     article.Content,
		Icon:        article.Icon,
		Author:      article.Author,
		CRC32Hex:    article.CRC32Hex,
	}
}

func (ra *RawArticle) calculateCRC32() (sum uint32) {
	var sb strings.Builder
	sb.WriteString(ra.DateUTC)
	sb.WriteString(ra.TimeUTC)
	sb.WriteString(ra.Category)
	sb.WriteString(ra.Title)
	sb.WriteString(ra.Description)
	sb.WriteString(ra.Content)
	sb.WriteString(ra.Icon)
	sb.WriteString(ra.Author)

	return crc32.ChecksumIEEE([]byte(sb.String()))
}
