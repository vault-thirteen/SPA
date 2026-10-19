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

let crc32 = {};
crc32.table = makeCrc32Table();

// This function works only for strings having ASCII symbols !
// Do not use it.
crc32.funcASCII = function (str) {
    let crc = 0 ^ (-1);
    for (let i = 0; i < str.length; i++) {
        crc = (crc >>> 8) ^ crc32.table[(crc ^ str.charCodeAt(i)) & 0xFF];
    }
    return ((crc ^ (-1)) >>> 0).toString(16).toUpperCase();
};

// This function works with Unicode symbols.
// Use it.
crc32.funcUnicode = function (unicodeStr) {
	let encoder = new TextEncoder();
	let buf = encoder.encode(unicodeStr);
	let crc = 0 ^ (-1);
	for (let i = 0; i < buf.length; i++) {
		crc = (crc >>> 8) ^ crc32.table[(crc ^ buf[i]) & 0xFF];
	}
	return ((crc ^ (-1)) >>> 0).toString(16).toUpperCase();
};
