package models

import "time"

type SortedFile struct {
	FilePath        string
	ArticleDateTime time.Time
}

// SortedFiles are files are sorted in a reverse way:
// first file is the newest one.
type SortedFiles []*SortedFile

func (sfs SortedFiles) Len() int {
	return len(sfs)
}

func (sfs SortedFiles) Less(i, j int) bool {
	return sfs[i].ArticleDateTime.After(sfs[j].ArticleDateTime)
}

func (sfs SortedFiles) Swap(i, j int) {
	sfs[i], sfs[j] = sfs[j], sfs[i]
}

// GetTopItems tries to return N top items from the list.
// If there are less than N top items available, they are returned.
func (sfs SortedFiles) GetTopItems(n int) SortedFiles {
	if len(sfs) >= n {
		return sfs[:n]
	} else {
		return sfs
	}
}

// GetNonTopItems tries to return non-top items excluding the N top items.
func (sfs SortedFiles) GetNonTopItems(n int) SortedFiles {
	if len(sfs) >= n {
		return sfs[n:]
	} else {
		return make(SortedFiles, 0)
	}
}
