package ind

import (
	"log"

	"github.com/vault-thirteen/SPA/cmd/indexer/cla"
	"github.com/vault-thirteen/SPA/pkg/indexer/settings"
)

type Indexer struct {
	CLA      *cla.CommandLineArguments
	Settings *settings.Settings
}

func NewIndexer() (indexer *Indexer, err error) {
	indexer = &Indexer{}

	indexer.CLA, err = cla.ReadCLA()
	if err != nil {
		return nil, err
	}

	if indexer.CLA.IsDefaultFile() {
		log.Println("Using the default configuration file.")
	}

	indexer.Settings, err = settings.NewSettingsFromFile(indexer.CLA.ConfigurationFilePath)
	if err != nil {
		return nil, err
	}

	return indexer, nil
}
