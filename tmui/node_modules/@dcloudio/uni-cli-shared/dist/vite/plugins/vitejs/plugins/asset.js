"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.urlToBuiltUrl = exports.getAssetHash = exports.assetFileNamesToFileName = exports.getAssetFilename = exports.fileToUrl = exports.checkPublicFile = exports.registerAssetToChunk = exports.parseAssets = exports.assetPlugin = exports.chunkToEmittedAssetsMap = exports.assetUrlRE = void 0;
const path_1 = __importDefault(require("path"));
const url_1 = require("url");
const lite_1 = __importDefault(require("mime/lite"));
const fs_extra_1 = __importStar(require("fs-extra"));
const magic_string_1 = __importDefault(require("magic-string"));
const crypto_1 = require("crypto");
const shared_1 = require("@vue/shared");
const utils_1 = require("../utils");
const utils_2 = require("../../../../vite/utils/utils");
const utils_3 = require("../../../../utils");
const static_1 = require("./static");
exports.assetUrlRE = /__VITE_ASSET__([a-z\d]{8})__(?:\$_(.*?)__)?/g;
const rawRE = /(\?|&)raw(?:&|$)/;
const urlRE = /(\?|&)url(?:&|$)/;
exports.chunkToEmittedAssetsMap = new WeakMap();
const assetCache = new WeakMap();
const assetHashToFilenameMap = new WeakMap();
// save hashes of the files that has been emitted in build watch
const emittedHashMap = new WeakMap();
/**
 * Also supports loading plain strings with import text from './foo.txt?raw'
 */
function assetPlugin(config, options) {
    // assetHashToFilenameMap initialization in buildStart causes getAssetFilename to return undefined
    assetHashToFilenameMap.set(config, new Map());
    return {
        name: 'vite:asset',
        buildStart() {
            assetCache.set(config, new Map());
            emittedHashMap.set(config, new Set());
        },
        resolveId(id) {
            if (!config.assetsInclude((0, utils_1.cleanUrl)(id))) {
                return;
            }
            // imports to absolute urls pointing to files in /public
            // will fail to resolve in the main resolver. handle them here.
            const publicFile = checkPublicFile(id, config);
            if (publicFile) {
                return id;
            }
        },
        async load(id) {
            if (id.startsWith('\0')) {
                // Rollup convention, this id should be handled by the
                // plugin that marked it with \0
                return;
            }
            // raw requests, read from disk
            if (rawRE.test(id)) {
                const file = checkPublicFile(id, config) || (0, utils_1.cleanUrl)(id);
                // raw query, read file and return as string
                return `export default ${JSON.stringify(await fs_extra_1.promises.readFile(file, 'utf-8'))}`;
            }
            if (!config.assetsInclude((0, utils_1.cleanUrl)(id)) && !urlRE.test(id)) {
                return;
            }
            id = id.replace(urlRE, '$1').replace(/[\?&]$/, '');
            const url = await fileToUrl(id, config, options?.isAndroidX
                ? {
                    emitFile(emittedFile) {
                        // 直接写入目标目录
                        fs_extra_1.default.outputFileSync(path_1.default.resolve(process.env.UNI_OUTPUT_DIR, emittedFile.fileName), emittedFile.source);
                    },
                }
                : this, false, (0, static_1.getIsStaticFile)());
            if (options?.isAndroidX) {
                this.emitFile({
                    type: 'asset',
                    fileName: (0, utils_3.normalizeEmitAssetFileName)((0, utils_3.normalizeNodeModules)(path_1.default.relative(process.env.UNI_INPUT_DIR, id))),
                    source: `export default ${JSON.stringify(parseAssets(config, url))}`,
                });
            }
            return `export default ${JSON.stringify(url)}`;
        },
        renderChunk(code, chunk) {
            let match;
            let s;
            exports.assetUrlRE.lastIndex = 0;
            // Urls added with JS using e.g.
            // imgElement.src = "__VITE_ASSET__5aa0ddc0__" are using quotes
            // Urls added in CSS that is imported in JS end up like
            // var inlined = ".inlined{color:green;background:url(__VITE_ASSET__5aa0ddc0__)}\n";
            // In both cases, the wrapping should already be fine
            while ((match = exports.assetUrlRE.exec(code))) {
                s = s || (s = new magic_string_1.default(code));
                const [full, hash, postfix = ''] = match;
                // some internal plugins may still need to emit chunks (e.g. worker) so
                // fallback to this.getFileName for that.
                const file = getAssetFilename(hash, config) || this.getFileName(hash);
                registerAssetToChunk(chunk, file);
                const outputFilepath = config.base + file + postfix;
                s.overwrite(match.index, match.index + full.length, outputFilepath);
            }
            if (s) {
                return {
                    code: s.toString(),
                    map: (0, utils_2.withSourcemap)(config) ? s.generateMap({ hires: true }) : null,
                };
            }
            else {
                return null;
            }
        },
    };
}
exports.assetPlugin = assetPlugin;
function parseAssets(config, code) {
    let match;
    let s;
    exports.assetUrlRE.lastIndex = 0;
    while ((match = exports.assetUrlRE.exec(code))) {
        s = s || (s = new magic_string_1.default(code));
        const [full, hash, postfix = ''] = match;
        // some internal plugins may still need to emit chunks (e.g. worker) so
        // fallback to this.getFileName for that.
        const file = getAssetFilename(hash, config);
        const outputFilepath = config.base + file + postfix;
        s.overwrite(match.index, match.index + full.length, outputFilepath);
    }
    if (s) {
        return s.toString();
    }
    return code;
}
exports.parseAssets = parseAssets;
function registerAssetToChunk(chunk, file) {
    let emitted = exports.chunkToEmittedAssetsMap.get(chunk);
    if (!emitted) {
        emitted = new Set();
        exports.chunkToEmittedAssetsMap.set(chunk, emitted);
    }
    emitted.add((0, utils_1.cleanUrl)(file));
}
exports.registerAssetToChunk = registerAssetToChunk;
function checkPublicFile(url, { publicDir }) {
    // note if the file is in /public, the resolver would have returned it
    // as-is so it's not going to be a fully resolved path.
    if (!publicDir || !url.startsWith('/')) {
        return;
    }
    const publicFile = path_1.default.join(publicDir, (0, utils_1.cleanUrl)(url));
    if (fs_extra_1.default.existsSync(publicFile)) {
        return publicFile;
    }
    else {
        return;
    }
}
exports.checkPublicFile = checkPublicFile;
function fileToUrl(id, config, ctx, canInline = false, isStaticFile) {
    return fileToBuiltUrl(id, config, ctx, false, canInline, isStaticFile);
}
exports.fileToUrl = fileToUrl;
function getAssetFilename(hash, config) {
    return assetHashToFilenameMap.get(config)?.get(hash);
}
exports.getAssetFilename = getAssetFilename;
/**
 * converts the source filepath of the asset to the output filename based on the assetFileNames option. \
 * this function imitates the behavior of rollup.js. \
 * https://rollupjs.org/guide/en/#outputassetfilenames
 *
 * @example
 * ```ts
 * const content = Buffer.from('text');
 * const fileName = assetFileNamesToFileName(
 *   'assets/[name].[hash][extname]',
 *   '/path/to/file.txt',
 *   getAssetHash(content),
 *   content
 * )
 * // fileName: 'assets/file.982d9e3e.txt'
 * ```
 *
 * @param assetFileNames filename pattern. e.g. `'assets/[name].[hash][extname]'`
 * @param file filepath of the asset
 * @param contentHash hash of the asset. used for `'[hash]'` placeholder
 * @param content content of the asset. passed to `assetFileNames` if `assetFileNames` is a function
 * @returns output filename
 */
function assetFileNamesToFileName(assetFileNames, file, contentHash, content) {
    const basename = path_1.default.basename(file);
    // placeholders for `assetFileNames`
    // `hash` is slightly different from the rollup's one
    const extname = path_1.default.extname(basename);
    const ext = extname.substring(1);
    const name = basename.slice(0, -extname.length);
    const hash = contentHash;
    if ((0, shared_1.isFunction)(assetFileNames)) {
        assetFileNames = assetFileNames({
            name: file,
            source: content,
            type: 'asset',
        });
        if (!(0, shared_1.isString)(assetFileNames)) {
            throw new TypeError('assetFileNames must return a string');
        }
    }
    else if (!(0, shared_1.isString)(assetFileNames)) {
        throw new TypeError('assetFileNames must be a string or a function');
    }
    const fileName = assetFileNames.replace(/\[\w+\]/g, (placeholder) => {
        switch (placeholder) {
            case '[ext]':
                return ext;
            case '[extname]':
                return extname;
            case '[hash]':
                return hash;
            case '[name]':
                return sanitizeFileName(name);
        }
        throw new Error(`invalid placeholder ${placeholder} in assetFileNames "${assetFileNames}"`);
    });
    return fileName;
}
exports.assetFileNamesToFileName = assetFileNamesToFileName;
// taken from https://github.com/rollup/rollup/blob/a8647dac0fe46c86183be8596ef7de25bc5b4e4b/src/utils/sanitizeFileName.ts
// https://datatracker.ietf.org/doc/html/rfc2396
// eslint-disable-next-line no-control-regex
const INVALID_CHAR_REGEX = /[\x00-\x1F\x7F<>*#"{}|^[\]`;?:&=+$,]/g;
const DRIVE_LETTER_REGEX = /^[a-z]:/i;
function sanitizeFileName(name) {
    const match = DRIVE_LETTER_REGEX.exec(name);
    const driveLetter = match ? match[0] : '';
    // A `:` is only allowed as part of a windows drive letter (ex: C:\foo)
    // Otherwise, avoid them because they can refer to NTFS alternate data streams.
    return (driveLetter +
        name.substring(driveLetter.length).replace(INVALID_CHAR_REGEX, '_'));
}
/**
 * Register an asset to be emitted as part of the bundle (if necessary)
 * and returns the resolved public URL
 */
function fileToBuiltUrl(id, config, pluginContext, skipPublicCheck = false, canInline = false, isStaticFile) {
    if (!skipPublicCheck && checkPublicFile(id, config)) {
        return config.base + id.slice(1);
    }
    const cache = assetCache.get(config);
    const cached = cache.get(id);
    if (cached) {
        return cached;
    }
    const file = (0, utils_1.cleanUrl)(id);
    const content = fs_extra_1.default.readFileSync(file);
    let url;
    if (canInline && content.length < Number(config.build.assetsInlineLimit)) {
        // base64 inlined as a string
        url = `data:${lite_1.default.getType(file)};base64,${content.toString('base64')}`;
    }
    else {
        const map = assetHashToFilenameMap.get(config);
        const contentHash = getAssetHash(content);
        const { search, hash } = (0, url_1.parse)(id);
        const postfix = (search || '') + (hash || '');
        const output = config.build?.rollupOptions?.output;
        const defaultAssetFileNames = path_1.default.posix.join(config.build.assetsDir, '[name].[hash][extname]');
        // Steps to determine which assetFileNames will be actually used.
        // First, if output is an object or string, use assetFileNames in it.
        // And a default assetFileNames as fallback.
        let assetFileNames = (output && !Array.isArray(output) ? output.assetFileNames : undefined) ??
            defaultAssetFileNames;
        if (output && Array.isArray(output)) {
            // Second, if output is an array, adopt assetFileNames in the first object.
            assetFileNames = output[0].assetFileNames ?? assetFileNames;
        }
        const inputDir = (0, utils_1.normalizePath)(process.env.UNI_INPUT_DIR);
        const isStatic = isStaticFile(file);
        let fileName = file.startsWith(inputDir) && isStatic
            ? // 需要处理 HBuilderX 项目中的 node_modules 目录
                (0, utils_3.normalizeNodeModules)(path_1.default.posix.relative(inputDir, file))
            : assetFileNamesToFileName(assetFileNames, file, contentHash, content);
        if (!map.has(contentHash)) {
            map.set(contentHash, fileName);
        }
        if (!isStatic) {
            const emittedSet = emittedHashMap.get(config);
            if (!emittedSet.has(contentHash)) {
                pluginContext.emitFile({
                    name: fileName,
                    fileName,
                    type: 'asset',
                    source: content,
                });
                emittedSet.add(contentHash);
            }
        }
        url = `__VITE_ASSET__${contentHash}__${postfix ? `$_${postfix}__` : ``}`;
    }
    cache.set(id, url);
    return url;
}
function getAssetHash(content) {
    return (0, crypto_1.createHash)('sha256').update(content).digest('hex').slice(0, 8);
}
exports.getAssetHash = getAssetHash;
function urlToBuiltUrl(url, importer, config, pluginContext, isStaticFile) {
    if (checkPublicFile(url, config)) {
        return config.base + url.slice(1);
    }
    const file = url.startsWith('/')
        ? path_1.default.join(config.root, url)
        : path_1.default.join(path_1.default.dirname(importer), url);
    return fileToBuiltUrl(file, config, pluginContext, 
    // skip public check since we just did it above
    true, false, isStaticFile);
}
exports.urlToBuiltUrl = urlToBuiltUrl;
