package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vault-thirteen/SPA/pkg/common/models"
	"github.com/vault-thirteen/SPA/pkg/jsonHasher/cla"
	ver "github.com/vault-thirteen/auxie/Versioneer"
)

func mustBeNoError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	showIntro()

	args, err := cla.ReadCLA()
	if err != nil {
		log.Println(err)
		cla.ShowUsage()
		os.Exit(1)
	}

	var ok bool
	switch args.ActionId {
	case cla.ActionIdHash:
		ok, err = fillFile(args.FileToProcess)
		mustBeNoError(err)
		if ok {
			fmt.Println("CRC32 check sum is already correct.")
		} else {
			fmt.Println("CRC32 check sum has been updated.")
		}

	case cla.ActionIdCheck:
		ok, err = checkFile(args.FileToProcess)
		mustBeNoError(err)
		if ok {
			fmt.Println("CRC32 check sum is correct.")
		} else {
			fmt.Println("CRC32 check sum is bad.")
		}

	default:
		os.Exit(1)
	}
}

func showIntro() {
	versioneer, err := ver.New()
	mustBeNoError(err)
	versioneer.ShowIntroText("Hasher")
	versioneer.ShowComponentsInfoText()
	fmt.Println()
}

// checkFile checks the CRC32 field of the JSON file.
// 'true' is returned if the checksum is good, otherwise – 'false'.
func checkFile(filePath string) (ok bool, err error) {
	var dat *models.Article
	dat, err = models.NewArticleFromFile(filePath)
	if err != nil {
		return false, err
	}

	if !dat.CheckCRC32() {
		return false, nil
	}

	return true, nil
}

// fillFile updates the CRC32 field of the JSON file.
// 'true' is returned if the checksum is already correct and no update is
// needed. If the checksum was updated, 'false' is returned.
//
// If the file needs an update, it is updated using a temporary file. This
// means that a new file receives updates and the old file is renamed with an
// old file postfix. This method ensures that original data is not damaged if
// something goes wrong. It is the user's duty to clean all the old files after
// the update. This function does not delete the old file for safety reasons.
func fillFile(filePath string) (alreadyOk bool, err error) {
	var article *models.Article
	article, err = models.NewArticleFromFile(filePath)
	if err != nil {
		return false, err
	}

	if article.CheckCRC32() {
		return true, nil
	}

	article.FillCRC32()

	err = article.SaveAsFile(filePath)
	if err != nil {
		return false, err
	}

	return false, nil
}
