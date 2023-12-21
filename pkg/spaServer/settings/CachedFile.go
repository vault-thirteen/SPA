package settings

import (
	"fmt"
	"path/filepath"

	"github.com/vault-thirteen/auxie/MIME"
)

const (
	MimeTypeJS  = mime.TypeApplicationJavascript
	MimeTypeCSS = mime.TypeTextCss
	MimeTypePNG = mime.TypeImagePng
)

const (
	FileExtJS  = ".js"
	FileExtCSS = ".css"
	FileExtPNG = ".png"
)

const (
	ErrFileExtensionIsUnsupported = "unsupported file extension: %v"
)

type CachedFile struct {
	Name     string
	Contents []byte
	MimeType string
}

func NewCachedFile(
	fileName string,
	contents []byte,
) (cf *CachedFile, err error) {
	cf = &CachedFile{
		Name:     fileName,
		Contents: contents,
	}

	switch filepath.Ext(filepath.Base(fileName)) {
	case FileExtJS:
		cf.MimeType = MimeTypeJS
	case FileExtCSS:
		cf.MimeType = MimeTypeCSS
	case FileExtPNG:
		cf.MimeType = MimeTypePNG
	default:
		return nil, fmt.Errorf(ErrFileExtensionIsUnsupported,
			filepath.Ext(filepath.Base(fileName)))
	}

	return cf, nil
}
