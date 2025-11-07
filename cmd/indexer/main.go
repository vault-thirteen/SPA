package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vault-thirteen/SPA/cmd/indexer/cla"
	"github.com/vault-thirteen/SPA/pkg/common/models"
	"github.com/vault-thirteen/SPA/pkg/indexer"
	ver "github.com/vault-thirteen/auxie/Versioneer/classes/Versioneer"
)

func main() {
	showIntro()

	var indexer *ind.Indexer
	var err error
	indexer, err = ind.NewIndexer()
	if err != nil {
		log.Println(err)
		showUsageAndExit()
	}

	if indexer.Settings.ShouldCreateCategoryFolder {
		err = indexer.CreateCategoryFolders()
		mustBeNoError(err)
	}

	if !indexer.Settings.CategoryExists(indexer.CLA.Category) {
		err = fmt.Errorf(ind.ErrCategoryIsNotFound, indexer.CLA.Category)
		log.Println(err)
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
		log.Fatal(err)
	}
}

func showIntro() {
	versioneer, err := ver.New()
	mustBeNoError(err)
	versioneer.ShowIntroText("Indexer")
	versioneer.ShowComponentsInfoText()
	fmt.Println()
}

func showUsageAndExit() {
	fmt.Println()
	cla.ShowUsage()
	os.Exit(1)
}
