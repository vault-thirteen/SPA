package settings

import (
	"errors"
	"os"
	"strings"

	"github.com/vault-thirteen/auxie/boolean"
	ae "github.com/vault-thirteen/auxie/errors"
	"github.com/vault-thirteen/auxie/number"
	"github.com/vault-thirteen/auxie/reader"
)

const (
	NewFilePerm   = 0644
	NewFolderPerm = 0644
	TmpFileExt    = ".tmp"
	OldFileExt    = ".old"
	JsonFileExt   = ".json"
)

const (
	ErrFileIsNotSet           = "file is not set"
	ErrJsonDataFolderIsNotSet = "JSON data folder is not set"
	ErrCategoryPathsAreNotSet = "category paths are not set"
	ErrTopNewsCountIsNotSet   = "top news count is not set"
	ErrServerAddressIsNotSet  = "server address is not set"
)

// Settings is Server's settings.
type Settings struct {
	// Path to the File with these Settings.
	File string

	// JSON Data Folder.
	JsonDataFolder string

	// Category Paths.
	CategoryPaths []string

	// Should we create category folders if they do not exist ?
	ShouldCreateCategoryFolder bool

	// Number of top items to be shown on the index page with full detail (i.e.
	// with an icon and a description text). All other items are shown briefly
	// (only date-time and a title are shown).
	TopNewsCount int

	// Addresses of servers.
	MainServerAddress string
	IconServerAddress string
	JpegServerAddress string
	JsonServerAddress string
}

func NewSettingsFromFile(filePath string) (stn *Settings, err error) {
	stn = &Settings{
		File: filePath,
	}

	var file *os.File
	file, err = os.Open(stn.File)
	if err != nil {
		return stn, err
	}
	defer func() {
		derr := file.Close()
		if derr != nil {
			err = ae.Combine(err, derr)
		}
	}()

	rdr := reader.New(file)
	var buf = make([][]byte, 8)

	for i := range buf {
		buf[i], err = rdr.ReadLineEndingWithCRLF()
		if err != nil {
			return stn, err
		}
	}

	stn.JsonDataFolder = strings.TrimSpace(string(buf[0]))

	var parts = strings.Split(string(buf[1]), ",")
	paths := make([]string, 0, len(parts))
	for _, part := range parts {
		paths = append(paths, strings.TrimSpace(part))
	}
	stn.CategoryPaths = paths

	stn.ShouldCreateCategoryFolder, err = boolean.FromString(
		strings.TrimSpace(string(buf[2])),
	)
	if err != nil {
		return stn, err
	}

	stn.TopNewsCount, err = number.ParseInt(strings.TrimSpace(string(buf[3])))
	if err != nil {
		return stn, err
	}

	stn.MainServerAddress = strings.TrimSpace(string(buf[4]))
	stn.JsonServerAddress = strings.TrimSpace(string(buf[5]))
	stn.IconServerAddress = strings.TrimSpace(string(buf[6]))
	stn.JpegServerAddress = strings.TrimSpace(string(buf[7]))

	return stn, nil
}

func (stn *Settings) Check() (err error) {
	if len(stn.File) == 0 {
		return errors.New(ErrFileIsNotSet)
	}

	if len(stn.JsonDataFolder) == 0 {
		return errors.New(ErrJsonDataFolderIsNotSet)
	}

	if len(stn.CategoryPaths) == 0 {
		return errors.New(ErrCategoryPathsAreNotSet)
	}

	if stn.TopNewsCount < 0 {
		return errors.New(ErrTopNewsCountIsNotSet)
	}

	if len(stn.MainServerAddress) == 0 {
		return errors.New(ErrServerAddressIsNotSet)
	}

	if len(stn.JsonServerAddress) == 0 {
		return errors.New(ErrServerAddressIsNotSet)
	}

	if len(stn.IconServerAddress) == 0 {
		return errors.New(ErrServerAddressIsNotSet)
	}

	if len(stn.JpegServerAddress) == 0 {
		return errors.New(ErrServerAddressIsNotSet)
	}

	return nil
}

func (stn *Settings) CategoryExists(category string) bool {
	for _, catPath := range stn.CategoryPaths {
		if catPath == category {
			return true
		}
	}

	return false
}
