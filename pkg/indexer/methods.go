package ind

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/vault-thirteen/SPA/pkg/common/helper"
	"github.com/vault-thirteen/SPA/pkg/common/models"
	"github.com/vault-thirteen/SPA/pkg/indexer/settings"
	"github.com/vault-thirteen/auxie/file"
)

const (
	// CategoryNews is the name of the news category[*].
	CategoryNews = "news"

	NewsBlockDateFormat  = "02.01.2006"
	EmptyCategoryContent = "There are no articles in this category at the moment."
)

const (
	ErrCategoryIsNotFound             = "category is not found: %v"
	ErrNewsCategoryIsASpecialCategory = "news category is a special category"
	ErrArticleIsDamaged               = "article is damaged: %v"
)

// CreateCategoryFolders creates category[*] folders if necessary.
func (i *Indexer) CreateCategoryFolders() (err error) {
	var fullPath string
	var ok bool
	for _, catPath := range i.Settings.CategoryPaths {
		if catPath == CategoryNews {
			continue
		}

		fullPath = filepath.Join(i.Settings.JsonDataFolder, catPath)
		ok, err = file.FolderExists(fullPath)
		if err != nil {
			return err
		}
		if ok {
			continue
		}

		log.Println(fmt.Sprintf("Creating a category folder: %v", catPath))
		err = os.MkdirAll(fullPath, settings.NewFolderPerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetArticleDateTime extracts the DateTime from an article's file.
func (i *Indexer) GetArticleDateTime(filePath string) (dateTime time.Time, err error) {
	var article *models.Article
	article, err = models.NewArticleFromFile(filePath)
	if err != nil {
		return dateTime, err
	}

	ok := article.CheckCRC32()
	if !ok {
		return dateTime, fmt.Errorf(ErrArticleIsDamaged, filePath)
	}

	return article.DateTimeUTC, nil
}

// GetDateTimeFromFolder gets the DateTime from all the files in the specified
// folder, including its sub-folders. A returned array's items have two fields:
// date-time and a file path relative to the JSON data folder.
func (i *Indexer) GetDateTimeFromFolder(folderPath string) (unsortedFiles models.SortedFiles, err error) {
	unsortedFiles = make(models.SortedFiles, 0)
	var dateTime time.Time
	var relPath string

	err = filepath.Walk(
		folderPath,
		func(path string, fi os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if fi.IsDir() {
				return nil
			}
			if filepath.Ext(path) != settings.JsonFileExt {
				return nil
			}

			dateTime, err = i.GetArticleDateTime(path)
			if err != nil {
				return err
			}

			relPath = helper.TrimSlashPrefix(strings.TrimPrefix(path, i.Settings.JsonDataFolder))

			unsortedFiles = append(unsortedFiles,
				&models.SortedFile{
					FilePath:        relPath,
					ArticleDateTime: dateTime,
				},
			)

			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return unsortedFiles, nil
}

// GetSortedFilesForCategory gets all files from a single category[*] and sorts
// them. The 'category' parameter is the category's path.
func (i *Indexer) GetSortedFilesForCategory(category string) (sortedFiles models.SortedFiles, err error) {
	if category == CategoryNews {
		return nil, errors.New(ErrNewsCategoryIsASpecialCategory)
	}

	catFullPath := filepath.Join(i.Settings.JsonDataFolder, category)
	sortedFiles, err = i.GetDateTimeFromFolder(catFullPath)
	if err != nil {
		return nil, err
	}

	sort.Sort(sortedFiles)

	return sortedFiles, nil
}

// GetAllSortedFiles gets all the files from all the categories[*] and sorts them.
func (i *Indexer) GetAllSortedFiles() (sortedFiles models.SortedFiles, err error) {
	sortedFiles = make(models.SortedFiles, 0)
	var files models.SortedFiles

	for _, catPath := range i.Settings.CategoryPaths {
		if catPath == CategoryNews {
			continue
		}

		catFullPath := filepath.Join(i.Settings.JsonDataFolder, catPath)

		files, err = i.GetDateTimeFromFolder(catFullPath)
		if err != nil {
			return nil, err
		}

		sortedFiles = append(sortedFiles, files...)
	}

	// Sort the records.
	sort.Sort(sortedFiles)

	return sortedFiles, nil
}

// createIndexForCategory creates an index file for any category.
func (i *Indexer) createIndexForCategory(category string, files models.SortedFiles) (err error) {
	topArticles := files.GetTopItems(i.Settings.TopNewsCount)
	nonTopArticles := files.GetNonTopItems(i.Settings.TopNewsCount)

	content := new(strings.Builder)
	var block string

	// Collect blocks of top articles.
	for _, topArticle := range topArticles {
		block, err = i.GetLongNewsBlock(topArticle.FilePath)
		if err != nil {
			return err
		}

		content.WriteString(block)
	}

	// Collect blocks of other (non-top) articles.
	for _, nonTopArticle := range nonTopArticles {
		block, err = i.GetShortNewsBlock(nonTopArticle.FilePath)
		if err != nil {
			return err
		}

		content.WriteString(block)
	}

	if len(topArticles) == 0 {
		content.WriteString(EmptyCategoryContent)
	}

	var article *models.Article
	article, err = models.NewArticle(
		category,
		models.Categories[category].Name,
		models.Categories[category].Description,
		content.String(),
		models.Categories[category].Icon,
		models.AuthorNone,
	)
	if err != nil {
		return err
	}

	fileName := filepath.Join(i.Settings.JsonDataFolder, category+settings.JsonFileExt)
	err = article.SaveAsFile(fileName)
	if err != nil {
		return err
	}

	return nil
}

// CreateIndexForCategory creates an index file for a single category[*].
func (i *Indexer) CreateIndexForCategory(category string, files models.SortedFiles) (err error) {
	if category == CategoryNews {
		return errors.New(ErrNewsCategoryIsASpecialCategory)
	}

	return i.createIndexForCategory(category, files)
}

// CreateIndexForNewsCategory creates an index file for the news category[*].
func (i *Indexer) CreateIndexForNewsCategory(files models.SortedFiles) (err error) {
	return i.createIndexForCategory(CategoryNews, files)
}

// GetShortNewsBlock gets the text of a short news block from an article.
// Short news block consists of date and title.
// fileRelPath is a path relative to the JSON data folder.
func (i *Indexer) GetShortNewsBlock(fileRelPath string) (block string, err error) {
	filePath := filepath.Join(i.Settings.JsonDataFolder, fileRelPath)

	var article *models.Article
	article, err = models.NewArticleFromFile(filePath)
	if err != nil {
		return block, err
	}

	return i.bbTagBlock(
		i.bbTagDate(article.DateTimeUTC.Format(NewsBlockDateFormat)) +
			i.bbTagLinkLocal(
				filepath.ToSlash(strings.TrimSuffix(fileRelPath, filepath.Ext(fileRelPath))),
				article.Title,
			),
	), nil
}

// GetLongNewsBlock gets the text of a long news block from an article.
// Long news block consists of icon, title and description.
// fileRelPath is a path relative to the JSON data folder.
func (i *Indexer) GetLongNewsBlock(fileRelPath string) (block string, err error) {
	filePath := filepath.Join(i.Settings.JsonDataFolder, fileRelPath)

	var article *models.Article
	article, err = models.NewArticleFromFile(filePath)
	if err != nil {
		return block, err
	}

	return i.bbTagBlock(
		i.bbTagIcon(article.Icon, "") +
			i.bbTagBlockTitle(
				filepath.ToSlash(strings.TrimSuffix(fileRelPath, filepath.Ext(fileRelPath))),
				article.Title,
			) +
			i.bbTagBlockDescription(article.Description),
	), nil
}

// Notes:
//	[*]	News category is a separate special category which virtually consists
//		of all other categories, i.e. it is the main index of the website. It
//		does not have a folder, and thus it has no own articles. It only has an
//		index which provides links to articles of all other categories.
