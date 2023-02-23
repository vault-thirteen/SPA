package ind

// These tags are used in JSON data files.
// They remind so called BBCode tags used in forum engines.

func (i *Indexer) bbTagIcon(src, alt string) string {
	return "[icon=" + src + "]" + alt + "[/icon]"
}

func (i *Indexer) bbTagPhoto(src, alt string) string {
	return "[photo=" + src + "]" + alt + "[/photo]"
}

func (i *Indexer) bbTagImage(src, alt string) string {
	return "[img=" + src + "]" + alt + "[/img]"
}

func (i *Indexer) bbTagBlock(contents string) string {
	return "[block]" + contents + "[/block]"
}

func (i *Indexer) bbTagBlockTitle(href, title string) string {
	return "[bt=" + href + "]" + title + "[/bt]"
}

func (i *Indexer) bbTagBlockDescription(text string) string {
	return "[bd]" + text + "[/bd]"
}

func (i *Indexer) bbTagLinkLocal(href, text string) string {
	return "[url=" + href + "]" + text + "[/url]"
}

func (i *Indexer) bbTagLinkExternal(href, text string) string {
	return "[ext=" + href + "]" + text + "[/ext]"
}

func (i *Indexer) bbTagBold(text string) string {
	return "[b]" + text + "[/b]"
}

func (i *Indexer) bbTagBr() string {
	return "[br]"
}

func (i *Indexer) bbTagCaption(text string) string {
	return "[caption]" + text + "[/caption]"
}

func (i *Indexer) bbTagH1(text string) string {
	return "[h1]" + text + "[/h1]"
}

func (i *Indexer) bbTagH2(text string) string {
	return "[h2]" + text + "[/h2]"
}

func (i *Indexer) bbTagH3(text string) string {
	return "[h3]" + text + "[/h3]"
}

func (i *Indexer) bbTagItalic(text string) string {
	return "[i]" + text + "[/i]"
}

func (i *Indexer) bbTagLi(text string) string {
	return "[li]" + text + "[/li]"
}

func (i *Indexer) bbTagDate(date string) string {
	return "[date]" + date + "[/date]"
}

func (i *Indexer) bbTagUl(text string) string {
	return "[ul]" + text + "[/ul]"
}
