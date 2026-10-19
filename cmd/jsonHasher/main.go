package main

import (
	"os"

	ver "github.com/vault-thirteen/auxie/Versioneer/classes/Versioneer"
	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"

	"github.com/vault-thirteen/SPA/pkg/common/models"
	"github.com/vault-thirteen/SPA/pkg/jsonHasher/cla"
)

const (
	Msg_Crc32SumIsAlreadyCorrect = "CRC32 check sum is already correct."
	Msg_Crc32SumUpdated          = "CRC32 check sum has been updated."
	Msg_Crc32SumIsCorrect        = "CRC32 check sum is correct."
	Msg_Crc32SumIsBad            = "CRC32 check sum is bad."
)

const AppName = "Hasher"

func mustBeNoError(err error) {
	if err != nil {
		errors.LogError(err)
		os.Exit(1)
	}
}

func main() {
	showIntro()

	args, err := cla.ReadCLA()
	if err != nil {
		errors.LogError(err)
		cla.ShowUsage()
		os.Exit(1)
	}

	var ok bool
	switch args.ActionId {
	case cla.ActionIdHash:
		ok, err = fillFile(args.FileToProcess)
		mustBeNoError(err)
		if ok {
			cio.PrintMessage(Msg_Crc32SumIsAlreadyCorrect)
		} else {
			cio.PrintMessage(Msg_Crc32SumUpdated)
		}

	case cla.ActionIdCheck:
		ok, err = checkFile(args.FileToProcess)
		mustBeNoError(err)
		if ok {
			cio.PrintMessage(Msg_Crc32SumIsCorrect)
		} else {
			cio.PrintMessage(Msg_Crc32SumIsBad)
		}

	default:
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
