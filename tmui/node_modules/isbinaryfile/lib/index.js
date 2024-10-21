"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.isBinaryFileSync = exports.isBinaryFile = void 0;
const fs = require("fs");
const util_1 = require("util");
const statAsync = (0, util_1.promisify)(fs.stat);
const openAsync = (0, util_1.promisify)(fs.open);
const closeAsync = (0, util_1.promisify)(fs.close);
const MAX_BYTES = 512;
// A very basic non-exception raising reader. Read bytes and
// at the end use hasError() to check whether this worked.
class Reader {
    fileBuffer;
    size;
    offset;
    error;
    constructor(fileBuffer, size) {
        this.fileBuffer = fileBuffer;
        this.size = size;
        this.offset = 0;
        this.error = false;
    }
    hasError() {
        return this.error;
    }
    nextByte() {
        if (this.offset === this.size || this.hasError()) {
            this.error = true;
            return 0xff;
        }
        return this.fileBuffer[this.offset++];
    }
    next(len) {
        const n = new Array();
        for (let i = 0; i < len; i++) {
            n[i] = this.nextByte();
        }
        return n;
    }
}
// Read a Google Protobuf var(iable)int from the buffer.
function readProtoVarInt(reader) {
    let idx = 0;
    let varInt = 0;
    while (!reader.hasError()) {
        const b = reader.nextByte();
        varInt = varInt | ((b & 0x7f) << (7 * idx));
        if ((b & 0x80) === 0) {
            break;
        }
        idx++;
    }
    return varInt;
}
// Attempt to taste a full Google Protobuf message.
function readProtoMessage(reader) {
    const varInt = readProtoVarInt(reader);
    const wireType = varInt & 0x7;
    switch (wireType) {
        case 0:
            readProtoVarInt(reader);
            return true;
        case 1:
            reader.next(8);
            return true;
        case 2:
            const len = readProtoVarInt(reader);
            reader.next(len);
            return true;
        case 5:
            reader.next(4);
            return true;
    }
    return false;
}
// Check whether this seems to be a valid protobuf file.
function isBinaryProto(fileBuffer, totalBytes) {
    const reader = new Reader(fileBuffer, totalBytes);
    let numMessages = 0;
    while (true) {
        // Definitely not a valid protobuf
        if (!readProtoMessage(reader) && !reader.hasError()) {
            return false;
        }
        // Short read?
        if (reader.hasError()) {
            break;
        }
        numMessages++;
    }
    return numMessages > 0;
}
async function isBinaryFile(file, size) {
    if (isString(file)) {
        const stat = await statAsync(file);
        isStatFile(stat);
        const fileDescriptor = await openAsync(file, 'r');
        const allocBuffer = Buffer.alloc(MAX_BYTES);
        // Read the file with no encoding for raw buffer access.
        // NB: something is severely wrong with promisify, had to construct my own Promise
        return new Promise((fulfill, reject) => {
            fs.read(fileDescriptor, allocBuffer, 0, MAX_BYTES, 0, (err, bytesRead, _) => {
                closeAsync(fileDescriptor);
                if (err) {
                    reject(err);
                }
                else {
                    fulfill(isBinaryCheck(allocBuffer, bytesRead));
                }
            });
        });
    }
    else {
        if (size === undefined) {
            size = file.length;
        }
        return isBinaryCheck(file, size);
    }
}
exports.isBinaryFile = isBinaryFile;
function isBinaryFileSync(file, size) {
    if (isString(file)) {
        const stat = fs.statSync(file);
        isStatFile(stat);
        const fileDescriptor = fs.openSync(file, 'r');
        const allocBuffer = Buffer.alloc(MAX_BYTES);
        const bytesRead = fs.readSync(fileDescriptor, allocBuffer, 0, MAX_BYTES, 0);
        fs.closeSync(fileDescriptor);
        return isBinaryCheck(allocBuffer, bytesRead);
    }
    else {
        if (size === undefined) {
            size = file.length;
        }
        return isBinaryCheck(file, size);
    }
}
exports.isBinaryFileSync = isBinaryFileSync;
function isBinaryCheck(fileBuffer, bytesRead) {
    // empty file. no clue what it is.
    if (bytesRead === 0) {
        return false;
    }
    let suspiciousBytes = 0;
    const totalBytes = Math.min(bytesRead, MAX_BYTES);
    // UTF-8 BOM
    if (bytesRead >= 3 && fileBuffer[0] === 0xef && fileBuffer[1] === 0xbb && fileBuffer[2] === 0xbf) {
        return false;
    }
    // UTF-32 BOM
    if (bytesRead >= 4 &&
        fileBuffer[0] === 0x00 &&
        fileBuffer[1] === 0x00 &&
        fileBuffer[2] === 0xfe &&
        fileBuffer[3] === 0xff) {
        return false;
    }
    // UTF-32 LE BOM
    if (bytesRead >= 4 &&
        fileBuffer[0] === 0xff &&
        fileBuffer[1] === 0xfe &&
        fileBuffer[2] === 0x00 &&
        fileBuffer[3] === 0x00) {
        return false;
    }
    // GB BOM
    if (bytesRead >= 4 &&
        fileBuffer[0] === 0x84 &&
        fileBuffer[1] === 0x31 &&
        fileBuffer[2] === 0x95 &&
        fileBuffer[3] === 0x33) {
        return false;
    }
    if (totalBytes >= 5 && fileBuffer.slice(0, 5).toString() === '%PDF-') {
        /* PDF. This is binary. */
        return true;
    }
    // UTF-16 BE BOM
    if (bytesRead >= 2 && fileBuffer[0] === 0xfe && fileBuffer[1] === 0xff) {
        return false;
    }
    // UTF-16 LE BOM
    if (bytesRead >= 2 && fileBuffer[0] === 0xff && fileBuffer[1] === 0xfe) {
        return false;
    }
    for (let i = 0; i < totalBytes; i++) {
        if (fileBuffer[i] === 0) {
            // NULL byte--it's binary!
            return true;
        }
        else if ((fileBuffer[i] < 7 || fileBuffer[i] > 14) && (fileBuffer[i] < 32 || fileBuffer[i] > 127)) {
            // UTF-8 detection
            if (fileBuffer[i] > 193 && fileBuffer[i] < 224 && i + 1 < totalBytes) {
                i++;
                if (fileBuffer[i] > 127 && fileBuffer[i] < 192) {
                    continue;
                }
            }
            else if (fileBuffer[i] > 223 && fileBuffer[i] < 240 && i + 2 < totalBytes) {
                i++;
                if (fileBuffer[i] > 127 && fileBuffer[i] < 192 && fileBuffer[i + 1] > 127 && fileBuffer[i + 1] < 192) {
                    i++;
                    continue;
                }
            }
            suspiciousBytes++;
            // Read at least 32 fileBuffer before making a decision
            if (i >= 32 && (suspiciousBytes * 100) / totalBytes > 10) {
                return true;
            }
        }
    }
    if ((suspiciousBytes * 100) / totalBytes > 10) {
        return true;
    }
    if (suspiciousBytes > 1 && isBinaryProto(fileBuffer, totalBytes)) {
        return true;
    }
    return false;
}
function isString(x) {
    return typeof x === 'string';
}
function isStatFile(stat) {
    if (!stat.isFile()) {
        throw new Error(`Path provided was not a file!`);
    }
}
