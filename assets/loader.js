/* Version 0.0.0 */

class Category {
    constructor(name, path, cssClass) {
        this.name = name;
        this.path = path;
        this.cssClass = cssClass;
    }
}

let crc32 = {};
let mca = {};

// makePage is the main function which is an entry point to the SPA.
async function initPage() {
    initServers();
    initCrc();
    initClicks();
    initCategories();
    initYear();
    initArticleNumber();
    initFooter();
    await loadPage();
}

async function loadPage() {
    if (!checkCurrentServer()) {
        console.error("Domain Error.");
        return
    }

    checkCategory();
    checkYear();
    checkArticleNumber();
    checkPageURI();
    makeTopMenu();

    let obj = await getJsonObject(mca.page.uri.path);
    if (obj == null) {
        console.error("Object is either damaged or not available.");
        return
    }
    fillPageWithContent(obj);
}

// initCrc initializes the CRC32 calculator.
function initCrc() {
    crc32.table = makeCrc32Table();
    crc32.func = function (unicodeStr) {
        let encoder = new TextEncoder();
        let buf = encoder.encode(unicodeStr);
        let crc = 0 ^ (-1);
        for (let i = 0; i < buf.length; i++) {
            crc = (crc >>> 8) ^ crc32.table[(crc ^ buf[i]) & 0xFF];
        }
        return ((crc ^ (-1)) >>> 0).toString(16).toUpperCase();
    };
}

function makeCrc32Table() {
    let c;
    let crcTable = [];
    for (let n = 0; n < 256; n++) {
        c = n;
        for (let k = 0; k < 8; k++) {
            c = ((c & 1) ? (0xEDB88320 ^ (c >>> 1)) : (c >>> 1));
        }
        crcTable[n] = c;
    }
    return crcTable;
}

function initClicks() {
    document.addEventListener('click', (event) => {
        documentMouseClick(event).then();
    }, true);
}

// documentMouseClick is a mouse click handler for the whole document.
async function documentMouseClick(event) {
    event.preventDefault();

    // Try to find the caller.
    let target = event.target;
    while (target.href === undefined) {
        target = target.parentNode;
        if (target === document) {
            return;
        }
    }

    // Enable changes of page URL for clicks on links leading to external addresses.
    let targetUrl = target.href;
    if (isExternalUrl(targetUrl)) {
        if (target.target === "_blank") {
            window.open(targetUrl, '_blank').focus();
        } else {
            window.location = targetUrl;
        }
        return;
    }

    // Local link click.
    window.history.pushState({}, "", targetUrl);
    await loadPage();
}

function isExternalUrl(url) {
    let host = window.location.hostname;

    let linkHost = host;
    if (/^https?:\/\//.test(url)) {
        let parser = document.createElement('a');
        parser.href = url;
        linkHost = parser.hostname;
        parser = null;
    }

    return host !== linkHost;
}

// composeURL composes a URL.
// path must start with a slash.
function composeURL3P(protocol, hostPort, path) {
    return protocol + "://" + hostPort + path;
}

function composeURL2P(protocol, hostPort) {
    return protocol + "://" + hostPort;
}

// initServers creates a setting with addresses of servers.
// Each kind of resources is fetched from its own server,
// e.g. JSON files are requested from JSON server.
// The exception is only for start page data – HTML code of
// the index page, loader script in JS and CSS style sheet.
function initServers() {
    mca.server = {}

    // Base (common) setting.
    mca.server.common = {
        protocol: "http",
        host: "localhost",
    };

    // Main server which serves the index page.
    mca.server.main = {
        protocol: mca.server.common.protocol,
        hostPort: mca.server.common.host + ":8000",
    };
    mca.server.main.address = composeURL2P(mca.server.main.protocol, mca.server.main.hostPort)

    // Server for icons.
    mca.server.icon = {
        protocol: mca.server.common.protocol,
        hostPort: mca.server.common.host + ":8001",
    }
    mca.server.icon.address = composeURL2P(mca.server.icon.protocol, mca.server.icon.hostPort)

    // Server for JPEG images.
    mca.server.jpeg = {
        protocol: mca.server.common.protocol,
        hostPort: mca.server.common.host + ":8002",
    }
    mca.server.jpeg.address = composeURL2P(mca.server.jpeg.protocol, mca.server.jpeg.hostPort)

    // Server for JSON files.
    mca.server.json = {
        protocol: mca.server.common.protocol,
        hostPort: mca.server.common.host + ":8003",
    }
    mca.server.json.address = composeURL2P(mca.server.json.protocol, mca.server.json.hostPort)
}

function checkCurrentServer() {
    let curPageHostPort = window.location.host;
    return curPageHostPort === mca.server.main.hostPort;
}

// initCategories creates a setting with parameters of categories.
// Categories of pages are hard-coded here in JS and in the HTML.
// This function also parses current URL into parts (called paths).
// Category is used in the first URL part (path).
function initCategories() {
    /* Information about categories in alphabetical order */
    let categoriesRD = [
        ["Events", "event", "tmbEvents"],
        ["Games", "game", "tmbGames"],
        ["Hardware", "hard", "tmbHardware"],
        ["Life", "life", "tmbLife"],
        ["Multimedia", "media", "tmbMultimedia"],
        ["Motorsport", "motor", "tmbMotorsport"],
        ["News", "news", "tmbNews"],
        ["Reviews", "review", "tmbReviews"],
        ["Software", "soft", "tmbSoftware"],
        ["Technology", "tech", "tmbTechnology"]
    ];
    mca.category = {};
    mca.category.indexDefault = 6; // News.
    mca.category.list = [];
    for (let i = 0; i < categoriesRD.length; i++)
        mca.category.list.push(
            new Category(categoriesRD[i][0], categoriesRD[i][1], categoriesRD[i][2]));
}

function checkCategory() {
    // Current URL path parts.
    mca.paths = window.location.pathname.split("/").filter(Boolean);

    // Index of the current category.
    mca.category.curIndex = mca.category.indexDefault;
    if (mca.paths.length === 0) {
        mca.paths.push(mca.category.list[mca.category.indexDefault].path);
    }
    for (let i = 0; i < mca.category.list.length; i++) {
        if (mca.category.list[i].path === mca.paths[0].toLowerCase()) {
            mca.category.curIndex = i;
            break;
        }
    }
    mca.paths[0] = mca.category.list[mca.category.curIndex].path;

    // Current category.
    mca.category.current = mca.category.list[mca.category.curIndex];
}

// initYears stores current year and the minimum year usable.
// Year is used in the second URL part (path).
function initYear() {
    mca.year = {};
    mca.year.first = 2023;
    mca.year.now = new Date().getFullYear();

    // List of possible years.
    mca.year.list = [];
    for (let i = mca.year.first; i <= mca.year.now; i++) {
        mca.year.list.push(i);
    }
}

function checkYear() {
    mca.year.selected = 0;
    if (mca.paths.length >= 2) {
        mca.year.selected = parseInt(mca.paths[1]);
        if (isNaN(mca.year.selected)) {
            mca.year.selected = mca.year.now;
        } else if ((mca.year.selected < mca.year.first) || (mca.year.selected > mca.year.now)) {
            mca.year.selected = mca.year.now;
        }
    }
}

function initArticleNumber() {
    mca.article = {};
}

function makeTopMenu() {
    let topMenuButtonClass = "tmButton";
    let el = document.getElementById("topMenuCell");
    el.innerHTML = "";

    let categoriesShuffled = shuffledCopy(mca.category.list);

    for (let i = 0; i < categoriesShuffled.length; i++) {
        let cat = categoriesShuffled[i];

        let tagA = document.createElement('a');
        tagA.href = "/" + cat.path;

        let tagDiv = document.createElement('div');
        let classAttr = topMenuButtonClass + " " + cat.cssClass;
        tagDiv.setAttribute("class", classAttr);
        tagDiv.innerHTML = cat.name;

        tagA.appendChild(tagDiv);
        el.appendChild(tagA);
    }
}

// randomInt generates a random integer value from interval [min;max].
function randomInt(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

// shuffle shuffles an array in place.
function shuffle(array) {
    let currentIndex = array.length, randomIndex;

    while (currentIndex !== 0) {
        randomIndex = Math.floor(Math.random() * currentIndex);
        currentIndex--;
        [array[currentIndex], array[randomIndex]] = [
            array[randomIndex], array[currentIndex]];
    }

    return array;
}

// shuffledCopy returns a shuffled copy of an array.
function shuffledCopy(array) {
    let arrayCopy = array.slice();
    shuffle(arrayCopy);
    return arrayCopy;
}

function initFooter() {
    let el = document.getElementById("footerCell");
    el.innerHTML = "Information about this site";
}

// checkArticleNumber extracts the article number from current URL.
// Article number is used in the third URL part (path).
function checkArticleNumber() {
    mca.article.selected = 0;
    if (mca.paths.length >= 3) {
        mca.article.selected = parseInt(mca.paths[2]);
        if (isNaN(mca.article.selected)) {
            mca.article.selected = 0;
        }
    }
}

// checkPageURI counts on which level of URL we are and composes a usable page URI.
// This is needed because a client may enter a fake URL.
// Here we normalize the real URI of the site's page.
function checkPageURI() {
    let pathSegmentsCount = 3;
    if (mca.article.selected === 0) {
        pathSegmentsCount--;
    }
    if (mca.year.selected === 0) {
        pathSegmentsCount--;
    }

    mca.page = {};
    mca.page.uri = {};
    mca.page.uri.segments = [];
    if (pathSegmentsCount >= 1) {
        mca.page.uri.segments.push(mca.category.current.path);
    }
    if ((pathSegmentsCount >= 2) && (mca.year.selected !== 0)) {
        mca.page.uri.segments.push(mca.year.selected);
    }
    if ((pathSegmentsCount >= 3) && (mca.article.selected !== 0)) {
        mca.page.uri.segments.push(mca.article.selected);
    }
    mca.page.uri.path = '/' + mca.page.uri.segments.join('/');

    window.history.pushState({}, "", mca.page.uri.path);
}

// getJsonObject gets the JSON content.
// It also checks the integrity of data received.
async function getJsonObject(pagePath) {
    if (pagePath.length > 0) {
        let firstSymbol = Array.from(pagePath)[0];
        if (firstSymbol !== '/') {
            return null;
        }
    }

    let url = composeURL3P(mca.server.json.protocol, mca.server.json.hostPort, pagePath);
    let resp = await fetch(url);
    if (resp.status !== 200) {
        return null;
    }
    let obj = await resp.json();
    if (!checkJsonObjectsIntegrity(obj)) {
        return null;
    }

    return obj;
}

function checkJsonObjectsIntegrity(obj) {
    if (obj.CRC32.length === 0) return false;
    let hashedContent = obj.Date + obj.Time + obj.Category + obj.Title +
        obj.Description + obj.Content + obj.Icon + obj.Author;
    let hashSum = crc32.func(hashedContent);
    return hashSum === obj.CRC32;
}

const sleep = ms => new Promise(r => setTimeout(r, ms));

// fillPageWithContent tries to fill the page from data taken from
// the provided JSON object. For pages of the first and second
// levels some fields are not shown, among them are date, time,
// author.
// True is returned on success, and false – otherwise.
function fillPageWithContent(obj) {
    // Page Title.
    document.title = obj.Title;

    let tag = {};
    let pl = pageLevel();

    // Date.
    if (pl >= 3) {
        let dateParts = obj.Date.split("-");
        if (dateParts.length !== 3) return false;
        tag = document.getElementById("mcDate");
        tag.innerHTML = dateParts[2] + "." + dateParts[1] + "." + dateParts[0];
    } else {
        tag = document.getElementById("mcDate");
        tag.innerHTML = "";
    }


    // Time.
    if (pl >= 3) {
        let timeParts = obj.Time.split(":");
        if (timeParts.length !== 2) return false;
        tag = document.getElementById("mcTime");
        tag.innerHTML = "[" + timeParts[0] + ":" + timeParts[1] + "]";
    } else {
        tag = document.getElementById("mcTime");
        tag.innerHTML = "";
    }

    // Category.
    let cat = findCategoryByPath(obj.Category);
    if (cat == null) return false;
    tag = document.getElementById("mcCategoryPrefix");
    tag.innerHTML = "Category: ";
    tag = document.getElementById("mcCategoryName");
    tag.innerHTML = tagLinkLocal("/" + cat.path, cat.name);

    // Main Content.
    let mainContent = revealAllowedHtmlTags(hideHtmlTags(decodeURIComponent(obj.Content)));
    tag = document.getElementById("mainContentCell");
    let description = preProcessDescription(obj.Description);
    tag.innerHTML = tagH1(obj.Title) + "\r\n" +
        tagBlock(tagIcon(obj.Icon, obj.Title) + tagDescription(description)) + "\r\n" +
        mainContent;

    // Author.
    if (pl >= 3) {
        tag = document.getElementById("mcAuthorPrefix");
        tag.innerHTML = "Author: ";
        tag = document.getElementById("mcAuthorName");
        tag.innerHTML = obj.Author;
    } else {
        tag = document.getElementById("mcAuthorPrefix");
        tag.innerHTML = "";
        tag = document.getElementById("mcAuthorName");
        tag.innerHTML = "";
    }
}

// findCategoryByPath tries to find a category by its URL part (path).
// Null is returned if no category is found.
function findCategoryByPath(categoryPath) {
    for (let i = 0; i < mca.category.list.length; i++) {
        if (mca.category.list[i].path === categoryPath) {
            return mca.category.list[i];
        }
    }

    return null;
}

// hideHtmlTags replaces <> with [].
function hideHtmlTags(text) {
    text = text.replaceAll("<", "[");
    text = text.replaceAll(">", "]");
    return text;
}

// revealAllowedHtmlTags converts BBCode tags into allowed HTML tags.
function revealAllowedHtmlTags(text) {
    // A normal link. Opens in the same page.
    // [url=www.yahoo.com]text[/url] -> <a href='www.yahoo.com'>text</a>.
    text = text.replaceAll(/\[url=(.*?)](.*?)\[\/url]/gim, tagLinkLocal("$1", "$2"));

    // An external link. Opens in a new page.
    // [ext=www.yahoo.com]text[/ext] -> <a href='www.yahoo.com' target='_blank'>text</a>.
    text = text.replaceAll(/\[ext=(.*?)](.*?)\[\/ext]/gim, tagLinkExternal("$1", "$2"));

    // [b]Abc[/b] -> <b>Abc</b>.
    text = text.replaceAll(/\[b](.*?)\[\/b]/gim, tagBold("$1"));

    // [br] -> <br>.
    text = text.replaceAll("[br]", tagBr());

    // Block for news description with icon.
    // [block]aaa[/block] -> <div>aaa</div>.
    text = text.replaceAll(/\[block](.*?)\[\/block]/gim, tagBlock("$1"));

    // Block title.
    // [bt=123]Title[/bt] -> <div><a href='json_server/123'>Title</a></div>.
    text = text.replaceAll(/\[bt=(.*?)](.*?)\[\/bt]/gim, tagBlockTitle("$1", "$2"));

    // Block description.
    // [bd]aaa[/bd] -> <div>aaa</div>.
    text = text.replaceAll(/\[bd](.*?)\[\/bd]/gim, tagBlockDescription("$1"));

    // Image caption with padding below it.
    // [caption]abc[/caption] -> <div>abc</div>.
    text = text.replaceAll(/\[caption](.*?)\[\/caption]/gim, tagCaption("$1"));

    // [h2]Abc[/h2] -> <h2>Abc</h2>.
    text = text.replaceAll(/\[h2](.*?)\[\/h2]/gim, tagH2("$1"));

    // [h3]Abc[/h3] -> <h3>Abc</h3>.
    text = text.replaceAll(/\[h3](.*?)\[\/h3]/gim, tagH3("$1"));

    // [i]Abc[/i] -> <i>Abc</i>.
    text = text.replaceAll(/\[i](.*?)\[\/i]/gim, tagItalic("$1"));

    // Local image.
    // [img=123]alt[/img] -> <img src='jpeg_server/123' alt='alt' />.
    text = text.replaceAll(/\[img=(.*?)](.*?)\[\/img]/gim, tagImage("$1", "$2"));

    // Clickable local image that opens a full size image in a new page.
    // [photo=123]alt[/photo] ->
    // <a href='jpeg_server/123' target='_blank'><img src='jpeg_server/123' alt='alt'/></a>.
    text = text.replaceAll(/\[photo=(.*?)](.*?)\[\/photo]/gim, tagPhoto("$1", "$2"));

    // Icon for a list of articles.
    // [icon=123]alt[/icon] -> <img src='icon_server/123' alt='alt' />.
    text = text.replaceAll(/\[icon=(.*?)](.*?)\[\/icon]/gim, tagIcon("$1", "$2"));

    // [li]Abc[/li] -> <li>Abc</li>.
    text = text.replaceAll(/\[li](.*?)\[\/li]/gim, tagLi("$1"));

    // [date]01.01.2001[/date] -> <span>01.01.2001</span>.
    text = text.replaceAll(/\[date](.*?)\[\/date]/gim, tagDate("$1"));

    // [ul]Abc[/ul] -> <ul>Abc</ul>.
    text = text.replaceAll(/\[ul](.*?)\[\/ul]/gim, tagUl("$1"));

    return text;
}

function preProcessDescription(text) {
    return text.replaceAll("[br]", tagBr());
}

function pageLevel() {
    return mca.paths.length;
}

// Following several functions return a piece of HTML, an HTML tag for something.
// If it is used from the 'revealAllowedHtmlTags' function, placeholders ($1, $2, ...) must be provided.
// If it is used from the 'fillPageWithContent' function, real values (123) must be provided.

function tagIcon(src, alt) {
    return "<img src='" + mca.server.icon.address + "/" + src + "' alt='" + alt + "' class='mcIcon' />";
}

function tagPhoto(src, alt) {
    return "<a href='" + mca.server.jpeg.address + "/" + src + "' target='_blank'>" +
        "<img src='" + mca.server.jpeg.address + "/" + src + "' alt='" + alt + "' class='mcImg' /></a>";
}

function tagImage(src, alt) {
    return "<img src='" + mca.server.jpeg.address + "/" + src + "' alt='" + alt + "' class='mcImg' />";
}

function tagDescription(text) {
    return "<div class='mcDescription'>" + text + "</div>";
}

function tagBlock(contents) {
    return "<div class='mcBlock'>" + contents + "</div>";
}

function tagBlockTitle(href, title) {
    return "<a href='" + href + "' class='mcBlockTitleLink'>" +
        "<div class='mcBlockTitle'>" + title + "</div></a>";
}

function tagBlockDescription(text) {
    return "<div class='mcBlockDescription'>" + text + "</div>";
}

function tagLinkLocal(href, text) {
    return "<a href='" + href + "'>" + text + "</a>";
}

function tagLinkExternal(href, text) {
    return "<a href='" + href + "' target='_blank'>" + text + "</a>";
}

function tagBold(text) {
    return "<b>" + text + "</b>";
}

function tagBr() {
    return "<br />"
}

function tagCaption(text) {
    return "<div class='mcCaption'>" + text + "</div>";
}

function tagH1(text) {
    return "<h1>" + text + "</h1>";
}

function tagH2(text) {
    return "<h2>" + text + "</h2>";
}

function tagH3(text) {
    return "<h3>" + text + "</h3>";
}

function tagItalic(text) {
    return "<i>" + text + "</i>";
}

function tagLi(text) {
    return "<li>" + text + "</li>";
}

function tagDate(date) {
    return "<span class='mcDate'>" + date + "</span>";
}

function tagUl(text) {
    return "<ul>" + text + "</ul>";
}
