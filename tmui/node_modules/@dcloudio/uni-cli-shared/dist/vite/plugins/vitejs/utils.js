"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.combineSourcemaps = exports.processSrcSet = exports.generateCodeFrame = exports.locToStartAndEnd = exports.posToNumber = exports.offsetToLineColumnByLines = exports.offsetToLineColumn = exports.offsetToStartAndEnd = exports.pad = exports.isObject = exports.asyncReplace = exports.multilineCommentsRE = exports.isDataUrl = exports.dataUrlRE = exports.isExternalUrl = exports.externalRE = exports.cleanUrl = exports.hashRE = exports.queryRE = exports.normalizePath = exports.isWindows = exports.deepImportRE = exports.bareImportRE = exports.slash = void 0;
/**
 * https://github.com/vitejs/vite/blob/main/packages/vite/src/node/utils.ts
 */
const os_1 = __importDefault(require("os"));
const path_1 = __importDefault(require("path"));
const remapping_1 = __importDefault(require("@ampproject/remapping"));
const lines_and_columns_1 = require("lines-and-columns");
function slash(p) {
    return p.replace(/\\/g, '/');
}
exports.slash = slash;
exports.bareImportRE = /^[\w@](?!.*:\/\/)/;
exports.deepImportRE = /^([^@][^/]*)\/|^(@[^/]+\/[^/]+)\//;
exports.isWindows = os_1.default.platform() === 'win32';
function normalizePath(id) {
    return path_1.default.posix.normalize(exports.isWindows ? slash(id) : id);
}
exports.normalizePath = normalizePath;
exports.queryRE = /\?.*$/s;
exports.hashRE = /#.*$/s;
const cleanUrl = (url) => url.replace(exports.hashRE, '').replace(exports.queryRE, '');
exports.cleanUrl = cleanUrl;
exports.externalRE = /^(https?:)?\/\//;
const isExternalUrl = (url) => exports.externalRE.test(url);
exports.isExternalUrl = isExternalUrl;
exports.dataUrlRE = /^\s*data:/i;
const isDataUrl = (url) => exports.dataUrlRE.test(url);
exports.isDataUrl = isDataUrl;
exports.multilineCommentsRE = /\/\*(.|[\r\n])*?\*\//gm;
async function asyncReplace(input, re, replacer) {
    let match;
    let remaining = input;
    let rewritten = '';
    while ((match = re.exec(remaining))) {
        rewritten += remaining.slice(0, match.index);
        rewritten += await replacer(match);
        remaining = remaining.slice(match.index + match[0].length);
    }
    rewritten += remaining;
    return rewritten;
}
exports.asyncReplace = asyncReplace;
function isObject(value) {
    return Object.prototype.toString.call(value) === '[object Object]';
}
exports.isObject = isObject;
const splitRE = /\r?\n/;
const range = 2;
function pad(source, n = 2) {
    const lines = source.split(splitRE);
    return lines.map((l) => ` `.repeat(n) + l).join(`\n`);
}
exports.pad = pad;
function offsetToStartAndEnd(source, startOffset, endOffset) {
    const lines = new lines_and_columns_1.LinesAndColumns(source);
    return {
        start: offsetToLineColumnByLines(lines, startOffset),
        end: offsetToLineColumnByLines(lines, endOffset),
        source: '',
    };
}
exports.offsetToStartAndEnd = offsetToStartAndEnd;
function offsetToLineColumn(source, offset) {
    return offsetToLineColumnByLines(new lines_and_columns_1.LinesAndColumns(source), offset);
}
exports.offsetToLineColumn = offsetToLineColumn;
function offsetToLineColumnByLines(lines, offset) {
    let location = lines.locationForIndex(offset);
    if (!location) {
        location = lines.locationForIndex(offset);
    }
    // lines-and-columns is 0-indexed while positions are 1-indexed
    return { line: location.line + 1, column: location.column, offset: 0 };
}
exports.offsetToLineColumnByLines = offsetToLineColumnByLines;
function posToNumber(source, pos) {
    if (typeof pos === 'number')
        return pos;
    return posToNumberByLines(new lines_and_columns_1.LinesAndColumns(source), pos.line, pos.column);
}
exports.posToNumber = posToNumber;
function posToNumberByLines(lines, line, column) {
    // lines-and-columns is 0-indexed while positions are 1-indexed
    return lines.indexForLocation({ line: line - 1, column }) || 0;
}
function locToStartAndEnd(source, loc) {
    const lines = new lines_and_columns_1.LinesAndColumns(source);
    const start = posToNumberByLines(lines, loc.start.line, loc.start.column);
    const end = loc.end
        ? posToNumberByLines(lines, loc.end.line, loc.end.column)
        : undefined;
    return { start, end };
}
exports.locToStartAndEnd = locToStartAndEnd;
function generateCodeFrame(source, start = 0, end) {
    start = posToNumber(source, start);
    end = end || start;
    // Split the content into individual lines but capture the newline sequence
    // that separated each line. This is important because the actual sequence is
    // needed to properly take into account the full line length for offset
    // comparison
    let lines = source.split(/(\r?\n)/);
    // Separate the lines and newline sequences into separate arrays for easier referencing
    const newlineSequences = lines.filter((_, idx) => idx % 2 === 1);
    lines = lines.filter((_, idx) => idx % 2 === 0);
    let count = 0;
    const res = [];
    for (let i = 0; i < lines.length; i++) {
        count +=
            lines[i].length +
                ((newlineSequences[i] && newlineSequences[i].length) || 0);
        if (count >= start) {
            for (let j = i - range; j <= i + range || end > count; j++) {
                if (j < 0 || j >= lines.length)
                    continue;
                const line = j + 1;
                res.push(`${line}${' '.repeat(Math.max(3 - String(line).length, 0))}|  ${lines[j]}`);
                const lineLength = lines[j].length;
                const newLineSeqLength = (newlineSequences[j] && newlineSequences[j].length) || 0;
                if (j === i) {
                    // push underline
                    const pad = start - (count - (lineLength + newLineSeqLength));
                    const length = Math.max(1, end > count ? lineLength - pad : end - start);
                    res.push(`   |  ` + ' '.repeat(pad) + '^'.repeat(length));
                }
                else if (j > i) {
                    if (end > count) {
                        const length = Math.max(Math.min(end - count, lineLength), 1);
                        res.push(`   |  ` + '^'.repeat(length));
                    }
                    count += lineLength + newLineSeqLength;
                }
            }
            break;
        }
    }
    return res.join('\n');
}
exports.generateCodeFrame = generateCodeFrame;
const escapedSpaceCharacters = /( |\\t|\\n|\\f|\\r)+/g;
const imageSetUrlRE = /^(?:[\w\-]+\(.*?\)|'.*?'|".*?"|\S*)/;
async function processSrcSet(srcs, replacer) {
    const imageCandidates = srcs
        .split(',')
        .map((s) => {
        const src = s.replace(escapedSpaceCharacters, ' ').trim();
        const [url] = imageSetUrlRE.exec(src) || [];
        return {
            url: url,
            descriptor: src?.slice(url.length).trim(),
        };
    })
        .filter(({ url }) => !!url);
    const ret = await Promise.all(imageCandidates.map(async ({ url, descriptor }) => {
        return {
            url: await replacer({ url, descriptor }),
            descriptor,
        };
    }));
    const url = ret.reduce((prev, { url, descriptor }, index) => {
        descriptor = descriptor || '';
        return (prev +=
            url + ` ${descriptor}${index === ret.length - 1 ? '' : ', '}`);
    }, '');
    return url;
}
exports.processSrcSet = processSrcSet;
function escapeToLinuxLikePath(path) {
    if (/^[A-Z]:/.test(path)) {
        return path.replace(/^([A-Z]):\//, '/windows/$1/');
    }
    if (/^\/[^/]/.test(path)) {
        return `/linux${path}`;
    }
    return path;
}
function unescapeToLinuxLikePath(path) {
    if (path.startsWith('/linux/')) {
        return path.slice('/linux'.length);
    }
    if (path.startsWith('/windows/')) {
        return path.replace(/^\/windows\/([A-Z])\//, '$1:/');
    }
    return path;
}
// based on https://github.com/sveltejs/svelte/blob/abf11bb02b2afbd3e4cac509a0f70e318c306364/src/compiler/utils/mapped_code.ts#L221
const nullSourceMap = {
    names: [],
    sources: [],
    mappings: '',
    version: 3,
};
function combineSourcemaps(filename, sourcemapList, excludeContent = true) {
    if (sourcemapList.length === 0 ||
        sourcemapList.every((m) => m.sources.length === 0)) {
        return { ...nullSourceMap };
    }
    // hack for parse broken with normalized absolute paths on windows (C:/path/to/something).
    // escape them to linux like paths
    // also avoid mutation here to prevent breaking plugin's using cache to generate sourcemaps like vue (see #7442)
    sourcemapList = sourcemapList.map((sourcemap) => {
        const newSourcemaps = { ...sourcemap };
        newSourcemaps.sources = sourcemap.sources.map((source) => source ? escapeToLinuxLikePath(source) : null);
        if (sourcemap.sourceRoot) {
            newSourcemaps.sourceRoot = escapeToLinuxLikePath(sourcemap.sourceRoot);
        }
        return newSourcemaps;
    });
    const escapedFilename = escapeToLinuxLikePath(filename);
    // We don't declare type here so we can convert/fake/map as RawSourceMap
    let map; //: SourceMap
    let mapIndex = 1;
    const useArrayInterface = sourcemapList.slice(0, -1).find((m) => m.sources.length !== 1) === undefined;
    if (useArrayInterface) {
        map = (0, remapping_1.default)(sourcemapList, () => null, excludeContent);
    }
    else {
        map = (0, remapping_1.default)(sourcemapList[0], function loader(sourcefile) {
            if (sourcefile === escapedFilename && sourcemapList[mapIndex]) {
                return sourcemapList[mapIndex++];
            }
            else {
                return null;
            }
        }, excludeContent);
    }
    if (!map.file) {
        delete map.file;
    }
    // unescape the previous hack
    map.sources = map.sources.map((source) => source ? unescapeToLinuxLikePath(source) : source);
    map.file = filename;
    return map;
}
exports.combineSourcemaps = combineSourcemaps;
