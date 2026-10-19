package main

import (
	"os"

	ver "github.com/vault-thirteen/auxie/Versioneer/classes/Versioneer"
	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"

	"github.com/vault-thirteen/SPA/cmd/indexer/cla"
	"github.com/vault-thirteen/SPA/pkg/common/models"
	"github.com/vault-thirteen/SPA/pkg/indexer"
)

const AppName = "Indexer"

func main() {
	showIntro()

	var indexer *ind.Indexer
	var err error
	indexer, err = ind.NewIndexer()
	if err != nil {
		errors.LogError(err)
		showUsageAndExit()
	}

	if indexer.Settings.ShouldCreateCategoryFolder {
		err = indexer.CreateCategoryFolders()
		mustBeNoError(err)
	}

	if !indexer.Settings.CategoryExists(indexer.CLA.Category) {
		err = errors.ErrorFromFormattedText(ind.Errf_CategoryIsNotFound, indexer.CLA.Category)
		errors.LogError(err)
		showUsageAndExit()
	}

	var files models.SortedFiles
	if indexer.CLA.Category != ind.CategoryNews {
		// Create an index for a single category (except 'news').
		files, err = indexer.GetSortedFilesForCategory(indexer.CLA.Category)
		if err != nil {
			mustBeNoError(err)
		}

		err = indexer.CreateIndexForCategory(indexer.CLA.Category, files)
		if err != nil {
			mustBeNoError(err)
		}

		return
	}

	// Create an index for the 'news' category.
	files, err = indexer.GetAllSortedFiles()
	mustBeNoError(err)

	err = indexer.CreateIndexForNewsCategory(files)
	if err != nil {
		mustBeNoError(err)
	}

	return
}

func mustBeNoError(err error) {
	if err != nil {
		errors.LogError(err)
		os.Exit(1)
	}
}

func showIntro() {
	versioneer, err := ver.New(false)
	mustBeNoError(err)
	versioneer.ShowIntroText(AppName)
	versioneer.ShowComponentsInfoText()
	cio.PrintMessage()
}

func showUsageAndExit() {
	cio.PrintMessage()
	cla.ShowUsage()
	os.Exit(1)
}
