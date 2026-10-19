package ind

import (
	"github.com/vault-thirteen/auxie/cio"

	"github.com/vault-thirteen/SPA/cmd/indexer/cla"
	cm "github.com/vault-thirteen/SPA/pkg/common/message"
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
		cio.LogMessage(cm.Msg_UsingDefaultConfigurationFile)
	}

	indexer.Settings, err = settings.NewSettingsFromFile(indexer.CLA.ConfigurationFilePath)
	if err != nil {
		return nil, err
	}

	return indexer, nil
}
