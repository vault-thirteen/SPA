package ind

// These tags are used by JavaScript loader script.
// They are listed here just for reference.

func (i *Indexer) htmlTagIcon(src, alt string) string {
	return "<img src='" + i.Settings.IconServerAddress + "/" + src +
		"' alt='" + alt + "' class='mcIcon' />"
}

func (i *Indexer) htmlTagPhoto(src, alt string) string {
	return "<a href='" + i.Settings.JpegServerAddress + "/" + src + "' target='_blank'>" +
		"<img src='" + i.Settings.JpegServerAddress + "/" + src + "' alt='" + alt + "' class='mcImg' /></a>"
}

func (i *Indexer) htmlTagImage(src, alt string) string {
	return "<img src='" + i.Settings.JpegServerAddress + "/" + src + "' alt='" + alt + "' class='mcImg' />"
}

func (i *Indexer) htmlTagBlock(contents string) string {
	return "<div class='mcBlock'>" + contents + "</div>"
}

func (i *Indexer) htmlTagBlockTitle(href, title string) string {
	return "<a href='" + href + "' class='mcBlockTitleLink'>" +
		"<div class='mcBlockTitle'>" + title + "</div></a>"
}

func (i *Indexer) htmlTagBlockDescription(text string) string {
	return "<div class='mcBlockDescription'>" + text + "</div>"
}

func (i *Indexer) htmlTagLinkLocal(href, text string) string {
	return "<a href='" + href + "'>" + text + "</a>"
}

func (i *Indexer) htmlTagLinkExternal(href, text string) string {
	return "<a href='" + href + "' target='_blank'>" + text + "</a>"
}

func (i *Indexer) htmlTagBold(text string) string {
	return "<b>" + text + "</b>"
}

func (i *Indexer) htmlTagBr() string {
	return "<br />"
}

func (i *Indexer) htmlTagCaption(text string) string {
	return "<div class='mcCaption'>" + text + "</div>"
}

func (i *Indexer) htmlTagH1(text string) string {
	return "<h1>" + text + "</h1>"
}

func (i *Indexer) htmlTagH2(text string) string {
	return "<h2>" + text + "</h2>"
}

func (i *Indexer) htmlTagH3(text string) string {
	return "<h3>" + text + "</h3>"
}

func (i *Indexer) htmlTagItalic(text string) string {
	return "<i>" + text + "</i>"
}

func (i *Indexer) htmlTagLi(text string) string {
	return "<li>" + text + "</li>"
}

func (i *Indexer) htmlTagDate(date string) string {
	return "<span class='mcDate'>" + date + "</span>"
}

func (i *Indexer) htmlTagUl(text string) string {
	return "<ul>" + text + "</ul>"
}

func (i *Indexer) htmlTagDescription(text string) string {
	return "<div class='mcDescription'>" + text + "</div>"
}
