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
exports.rewriteScssReadFileSync = exports.hoistAtRules = exports.minifyCSS = exports.importCssRE = exports.cssDataUriRE = exports.cssUrlRE = exports.formatPostcssSourceMap = exports.cssPostPlugin = exports.cssPlugin = exports.isDirectCSSRequest = exports.isCSSRequest = exports.commonjsProxyRE = exports.cssLangRE = void 0;
const fs_1 = __importDefault(require("fs"));
const fs_extra_1 = __importDefault(require("fs-extra"));
const path_1 = __importDefault(require("path"));
const fast_glob_1 = __importDefault(require("fast-glob"));
const picocolors_1 = __importDefault(require("picocolors"));
const postcss_load_config_1 = __importDefault(require("postcss-load-config"));
const pluginutils_1 = require("@rollup/pluginutils");
const utils_1 = require("../utils");
const asset_1 = require("./asset");
const magic_string_1 = __importDefault(require("magic-string"));
const Postcss = __importStar(require("postcss"));
const shared_1 = require("@vue/shared");
const preprocess_1 = require("../../../../preprocess");
const uniapp_1 = require("../../../../postcss/plugins/uniapp");
const cleanString_1 = require("../cleanString");
const constants_1 = require("../../../../constants");
const utils_2 = require("../../../utils/utils");
const compiler_core_1 = require("@vue/compiler-core");
const utils_3 = require("../../../../utils");
const utils_4 = require("../../../utils");
const static_1 = require("./static");
const cssLangs = `\\.(css|less|sass|scss|styl|stylus|pcss|postcss)($|\\?)`;
exports.cssLangRE = new RegExp(cssLangs);
const cssModuleRE = new RegExp(`\\.module${cssLangs}`);
const directRequestRE = /(\?|&)direct\b/;
exports.commonjsProxyRE = /\?commonjs-proxy/;
const varRE = /^var\(/i;
const isCSSRequest = (request) => exports.cssLangRE.test(request) && !directRequestRE.test(request);
exports.isCSSRequest = isCSSRequest;
const isDirectCSSRequest = (request) => exports.cssLangRE.test(request) && directRequestRE.test(request);
exports.isDirectCSSRequest = isDirectCSSRequest;
const cssModulesCache = new WeakMap();
const postcssConfigCache = new WeakMap();
function wrapResolve(resolve, code, source, getDescriptor) {
    if (!source) {
        return resolve;
    }
    return async (id, importer, aliasOnly, ssr) => {
        try {
            return await resolve(id, importer, aliasOnly, ssr);
        }
        catch (e) {
            if (importer && getDescriptor) {
                const { filename, query } = (0, utils_4.parseVueRequest)(importer);
                // 仅处理 vue | uvue
                if (query.vue &&
                    query.type === 'style' &&
                    'index' in query &&
                    /\.[u]vue/.test(filename)) {
                    const descriptor = getDescriptor(importer.split('?')[0]);
                    if (descriptor) {
                        const styleBlock = descriptor.styles[query.index];
                        if (styleBlock) {
                            code = descriptor.source;
                            const offsetLine = styleBlock.loc.start.line - 1;
                            source.start.line = source.start.line + offsetLine;
                            source.end.line = source.end.line + offsetLine;
                        }
                    }
                }
            }
            const error = (0, utils_2.createRollupError)('', importer || '', (0, compiler_core_1.createCompilerError)(0, { start: source.start, end: source.end, source: '' }, { 0: (0, utils_3.createResolveErrorMsg)(id, importer) }), code);
            error.line = source.start.line;
            error.column = source.start.column;
            throw error;
        }
        return;
    };
}
/**
 * Plugin applied before user plugins
 */
function cssPlugin(config, options = { isAndroidX: false }) {
    let server;
    let moduleCache;
    const resolveUrl = config.createResolver({
        preferRelative: true,
        tryIndex: false,
        extensions: [],
    });
    const atImportResolvers = createCSSResolvers(config);
    return {
        name: 'vite:css',
        configureServer(_server) {
            server = _server;
        },
        buildStart() {
            // Ensure a new cache for every build (i.e. rebuilding in watch mode)
            moduleCache = new Map();
            cssModulesCache.set(config, moduleCache);
        },
        async transform(raw, id) {
            if (!exports.cssLangRE.test(id) || exports.commonjsProxyRE.test(id)) {
                return;
            }
            const urlReplacer = (options.createUrlReplacer && options.createUrlReplacer(resolveUrl)) ||
                (async (url, importer, source) => {
                    if (url.startsWith('/') && !url.startsWith('//')) {
                        // /static/logo.png => @/static/logo.png
                        url = '@' + url;
                    }
                    const resolved = await wrapResolve(resolveUrl, source?.input.css || raw, source, options.getDescriptor)(url, importer);
                    if (resolved) {
                        return (0, asset_1.fileToUrl)(resolved, config, options?.isAndroidX
                            ? {
                                emitFile(emittedFile) {
                                    const fileName = path_1.default.resolve(process.env.UNI_OUTPUT_DIR, emittedFile.fileName);
                                    // 忽略static（可能有只读文件，写入覆盖只读会报错权限）
                                    if ((0, static_1.getIsStaticFile)()((0, utils_1.normalizePath)(emittedFile.fileName))) {
                                        return;
                                    }
                                    // 直接写入目标目录
                                    fs_extra_1.default.outputFileSync(fileName, emittedFile.source);
                                },
                            }
                            : this, true, (0, static_1.getIsStaticFile)());
                    }
                    return url;
                });
            const { code: css, modules, deps, } = await compileCSS(id, raw, config, urlReplacer, atImportResolvers, server);
            if (modules) {
                moduleCache.set(id, modules);
            }
            // track deps for build watch mode
            if (config.command === 'build' && config.build.watch && deps) {
                for (const file of deps) {
                    this.addWatchFile(file);
                }
            }
            return {
                code: css,
                // TODO CSS source map
                map: { mappings: '' },
            };
        },
    };
}
exports.cssPlugin = cssPlugin;
function findCssModuleIds(moduleId, includeComponentCss = true, cssModuleIds, seen) {
    if (!cssModuleIds) {
        cssModuleIds = new Set();
    }
    if (!seen) {
        seen = new Set();
    }
    if (seen.has(moduleId)) {
        return cssModuleIds;
    }
    seen.add(moduleId);
    const moduleInfo = this.getModuleInfo(moduleId);
    if (moduleInfo) {
        moduleInfo.importedIds.forEach((id) => {
            if (id.includes(constants_1.PAGES_JSON_JS) || id.includes(constants_1.PAGES_JSON_UTS)) {
                // 查询main.js时，需要忽略pages.json.js，否则会把所有页面样式加进来
                return;
            }
            if (exports.cssLangRE.test(id) && !exports.commonjsProxyRE.test(id)) {
                cssModuleIds.add(id);
            }
            else {
                if (!includeComponentCss &&
                    (id.includes('.vue') || id.includes('.uvue') || id.includes('.nvue'))) {
                    // 不包含组件样式，不需要继续查找，uni x中不需要包含
                    return;
                }
                findCssModuleIds.call(this, id, includeComponentCss, cssModuleIds, seen);
            }
        });
    }
    return cssModuleIds;
}
/**
 * Plugin applied after user plugins
 */
function cssPostPlugin(config, { platform, isJsCode, preserveModules, chunkCssFilename, chunkCssCode, includeComponentCss, }) {
    // styles initialization in buildStart causes a styling loss in watch
    const styles = new Map();
    let cssChunks;
    return {
        name: 'vite:css-post',
        buildStart() {
            cssChunks = new Map();
        },
        async transform(css, id) {
            if (!exports.cssLangRE.test(id) || exports.commonjsProxyRE.test(id)) {
                return;
            }
            const modules = cssModulesCache.get(config).get(id);
            const modulesCode = modules && (0, pluginutils_1.dataToEsm)(modules, { namedExports: true, preferConst: true });
            // build CSS handling ----------------------------------------------------
            styles.set(id, css);
            return {
                code: modulesCode || '',
                map: { mappings: '' },
                // avoid the css module from being tree-shaken so that we can retrieve
                // it in renderChunk()
                moduleSideEffects: 'no-treeshake',
            };
        },
        async renderChunk(_code, chunk, _opts) {
            const id = chunk.facadeModuleId;
            if (id) {
                const filename = chunkCssFilename(id);
                if (filename) {
                    if (platform === 'app' &&
                        (filename === 'app.css' || filename.startsWith('App.style'))) {
                        // 获取 unocss 的样式文件信息
                        const ids = Object.keys(chunk.modules).filter((id) => styles.has(id) &&
                            (id.includes('__uno.css') || id.includes('-unocss-')));
                        cssChunks.set(filename, ids);
                    }
                    else {
                        let ids = Object.keys(chunk.modules).filter((id) => styles.has(id));
                        // 当页面作为组件使用时，上一步找不到依赖的css，需要再次查找
                        // renderChunk会执行两次，一次是页面chunk，一次是组件chunk，两者生成的css文件名和内容都是一样的
                        if (!ids.length) {
                            ids = [...findCssModuleIds.call(this, id, includeComponentCss)];
                        }
                        cssChunks.set(filename, ids);
                    }
                }
            }
            return null;
        },
        async generateBundle() {
            // app 平台页面并未 chunk，所以 renderChunk 并不会处理页面的 css，需要这里再补充查找
            if (platform === 'app' || preserveModules) {
                const moduleIds = Array.from(this.getModuleIds());
                moduleIds.forEach((id) => {
                    const filename = chunkCssFilename(id);
                    if (filename) {
                        const ids = findCssModuleIds.call(this, id, includeComponentCss);
                        if (cssChunks.has(filename)) {
                            cssChunks.get(filename).forEach((id) => {
                                ids.add(id);
                            });
                        }
                        cssChunks.set(filename, [...ids]);
                    }
                });
            }
            if (!cssChunks.size) {
                return;
            }
            // resolve asset URL placeholders to their built file URLs and perform
            // minification if necessary
            const processChunkCSS = async (css, { filename, inlined, minify, }) => {
                // replace asset url references with resolved url.
                css = css.replace(asset_1.assetUrlRE, (_, fileHash, postfix = '') => {
                    return (0, utils_1.normalizePath)(path_1.default.relative(path_1.default.dirname(filename), (0, asset_1.getAssetFilename)(fileHash, config) + postfix));
                });
                if (isJsCode) {
                    return chunkCssCode(filename, css);
                }
                // only external @imports and @charset should exist at this point
                // hoist them to the top of the CSS chunk per spec (#1845 and #6333)
                if (css.includes('@import') || css.includes('@charset')) {
                    css = await hoistAtRules(css);
                }
                if (minify && config.build.minify) {
                    css = await minifyCSS(css, config);
                }
                // 压缩后再处理，小程序平台会补充 @import nvue 代码，esbuild 的压缩会把 `@import "./nvue.css";` 的空格移除，变成 `@import"./nvue.css";` 在支付宝小程序中不支持
                return chunkCssCode(filename, css);
            };
            const genCssCode = (fileName) => {
                return [...cssChunks.get(fileName)]
                    .map((id) => styles.get(id) || '')
                    .join('\n');
            };
            for (const filename of cssChunks.keys()) {
                const cssCode = genCssCode(filename);
                let source = await processChunkCSS(cssCode, {
                    filename: filename,
                    inlined: false,
                    minify: true,
                });
                this.emitFile({
                    fileName: filename,
                    type: 'asset',
                    source,
                });
            }
        },
    };
}
exports.cssPostPlugin = cssPostPlugin;
function createCSSResolvers(config) {
    let cssResolve;
    let sassResolve;
    let lessResolve;
    return {
        get css() {
            return (cssResolve ||
                (cssResolve = config.createResolver({
                    extensions: ['.css'],
                    mainFields: ['style'],
                    tryIndex: false,
                    preferRelative: true,
                })));
        },
        get sass() {
            return (sassResolve ||
                (sassResolve = config.createResolver({
                    extensions: ['.scss', '.sass', '.css'],
                    mainFields: ['sass', 'style'],
                    tryIndex: true,
                    tryPrefix: '_',
                    preferRelative: true,
                })));
        },
        get less() {
            return (lessResolve ||
                (lessResolve = config.createResolver({
                    extensions: ['.less', '.css'],
                    mainFields: ['less', 'style'],
                    tryIndex: false,
                    preferRelative: true,
                })));
        },
    };
}
function getCssResolversKeys(resolvers) {
    return Object.keys(resolvers);
}
async function compileCSS(id, code, config, urlReplacer, atImportResolvers, server) {
    const { modules: modulesOptions, preprocessorOptions, devSourcemap, } = config.css || {};
    const isModule = modulesOptions !== false && cssModuleRE.test(id);
    // although at serve time it can work without processing, we do need to
    // crawl them in order to register watch dependencies.
    const needInlineImport = code.includes('@import');
    const hasUrl = exports.cssUrlRE.test(code) || cssImageSetRE.test(code);
    const postcssConfig = await resolvePostcssConfig(config);
    const lang = id.match(exports.cssLangRE)?.[1];
    // 1. plain css that needs no processing
    if (lang === 'css' &&
        !postcssConfig &&
        !isModule &&
        !needInlineImport &&
        !hasUrl) {
        return { code, map: null };
    }
    let preprocessorMap;
    let modules;
    const deps = new Set();
    // 2. pre-processors: sass etc.
    if (isPreProcessor(lang)) {
        const preProcessor = preProcessors[lang];
        let opts = (preprocessorOptions && preprocessorOptions[lang]) || {};
        // support @import from node dependencies by default
        switch (lang) {
            case "scss" /* PreprocessLang.scss */:
            case "sass" /* PreprocessLang.sass */:
                opts = {
                    includePaths: ['node_modules'],
                    alias: config.resolve.alias,
                    ...opts,
                };
                break;
            case "less" /* PreprocessLang.less */:
            case "styl" /* PreprocessLang.styl */:
            case "stylus" /* PreprocessLang.stylus */:
                opts = {
                    paths: ['node_modules'],
                    alias: config.resolve.alias,
                    ...opts,
                };
        }
        // important: set this for relative import resolving
        opts.filename = (0, utils_1.cleanUrl)(id);
        opts.enableSourcemap = devSourcemap ?? false;
        const preprocessResult = await preProcessor(code, config.root, opts, atImportResolvers, !!config.nvue);
        if (preprocessResult.errors.length) {
            throw preprocessResult.errors[0];
        }
        // TODO 升级
        // if (preprocessResult.error) {
        //   throw preprocessResult.error
        // }
        code = preprocessResult.code;
        preprocessorMap = combineSourcemapsIfExists(opts.filename, preprocessResult.map, preprocessResult.additionalMap);
        if (preprocessResult.deps) {
            preprocessResult.deps.forEach((dep) => {
                // sometimes sass registers the file itself as a dep
                if ((0, utils_1.normalizePath)(dep) !== (0, utils_1.normalizePath)(opts.filename)) {
                    deps.add(dep);
                }
            });
        }
    }
    // 3. postcss
    const postcssOptions = (postcssConfig && postcssConfig.options) || {};
    const postcssPlugins = postcssConfig && postcssConfig.plugins ? postcssConfig.plugins.slice() : [];
    if (needInlineImport) {
        postcssPlugins.unshift((await Promise.resolve().then(() => __importStar(require('postcss-import')))).default({
            async resolve(id, basedir) {
                // const publicFile = checkPublicFile(id, config)
                // if (publicFile) {
                //   return publicFile
                // }
                const resolved = await atImportResolvers.css(id, path_1.default.join(basedir, '*'));
                if (resolved) {
                    return path_1.default.resolve(resolved);
                }
                return id;
            },
        }));
    }
    postcssPlugins.push(UrlRewritePostcssPlugin({
        replacer: urlReplacer,
    }));
    if (isModule) {
        postcssPlugins.unshift((await Promise.resolve().then(() => __importStar(require('postcss-modules')))).default({
            ...modulesOptions,
            getJSON(cssFileName, _modules, outputFileName) {
                modules = _modules;
                if (modulesOptions && (0, shared_1.isFunction)(modulesOptions.getJSON)) {
                    modulesOptions.getJSON(cssFileName, _modules, outputFileName);
                }
            },
            async resolve(id) {
                for (const key of getCssResolversKeys(atImportResolvers)) {
                    const resolved = await atImportResolvers[key](id);
                    if (resolved) {
                        return path_1.default.resolve(resolved);
                    }
                }
                return id;
            },
        }));
    }
    if (!postcssPlugins.length) {
        return {
            code,
            map: preprocessorMap,
        };
    }
    let postcssResult;
    try {
        // postcss is an unbundled dep and should be lazy imported
        postcssResult = await (await Promise.resolve().then(() => __importStar(require('postcss'))))
            .default(postcssPlugins)
            .process(code, {
            ...postcssOptions,
            to: id,
            from: id,
            ...(devSourcemap
                ? {
                    map: {
                        inline: false,
                        annotation: false,
                        // postcss may return virtual files
                        // we cannot obtain content of them, so this needs to be enabled
                        sourcesContent: true,
                        // when "prev: preprocessorMap", the result map may include duplicate filename in `postcssResult.map.sources`
                        // prev: preprocessorMap,
                    },
                }
                : {}),
        });
        // record CSS dependencies from @imports
        for (const message of postcssResult.messages) {
            if (message.type === 'dependency') {
                deps.add((0, utils_1.normalizePath)(message.file));
            }
            else if (message.type === 'dir-dependency') {
                // https://github.com/postcss/postcss/blob/main/docs/guidelines/plugin.md#3-dependencies
                const { dir, glob: globPattern = '**' } = message;
                const pattern = (0, utils_1.normalizePath)(path_1.default.resolve(path_1.default.dirname(id), dir)) + `/` + globPattern;
                const files = fast_glob_1.default.sync(pattern, {
                    ignore: ['**/node_modules/**'],
                });
                for (let i = 0; i < files.length; i++) {
                    deps.add(files[i]);
                }
                if (server) {
                    // register glob importers so we can trigger updates on file add/remove
                    if (!(id in server._globImporters)) {
                        ;
                        server._globImporters[id] = {
                            module: server.moduleGraph.getModuleById(id),
                            importGlobs: [],
                        };
                    }
                    ;
                    server._globImporters[id].importGlobs.push({
                        base: config.root,
                        pattern,
                    });
                }
            }
            else if (message.type === 'warning') {
                let msg = `[vite:css] ${message.text}`;
                if (message.line && message.column) {
                    msg += `\n${(0, utils_1.generateCodeFrame)(code, {
                        line: message.line,
                        column: message.column,
                    })}`;
                }
                config.logger.warn(picocolors_1.default.yellow(msg));
            }
        }
    }
    catch (e) {
        e.message = `[postcss] ${e.message}`;
        e.code = code;
        e.loc = {
            column: e.column,
            line: e.line,
        };
        throw e;
    }
    if (!devSourcemap) {
        return {
            ast: postcssResult,
            code: postcssResult.css,
            map: { mappings: '' },
            modules,
            deps,
        };
    }
    const rawPostcssMap = postcssResult.map.toJSON();
    const postcssMap = formatPostcssSourceMap(
    // version property of rawPostcssMap is declared as string
    // but actually it is a number
    rawPostcssMap, (0, utils_1.cleanUrl)(id));
    return {
        ast: postcssResult,
        code: postcssResult.css,
        map: combineSourcemapsIfExists((0, utils_1.cleanUrl)(id), postcssMap, preprocessorMap),
        modules,
        deps,
    };
}
function formatPostcssSourceMap(rawMap, file) {
    const inputFileDir = path_1.default.dirname(file);
    const sources = rawMap.sources.map((source) => {
        const cleanSource = (0, utils_1.cleanUrl)(decodeURIComponent(source));
        // postcss returns virtual files
        if (/^<.+>$/.test(cleanSource)) {
            return `\0${cleanSource}`;
        }
        return (0, utils_1.normalizePath)(path_1.default.resolve(inputFileDir, cleanSource));
    });
    return {
        file,
        mappings: rawMap.mappings,
        names: rawMap.names,
        sources,
        sourcesContent: rawMap.sourcesContent,
        version: rawMap.version,
    };
}
exports.formatPostcssSourceMap = formatPostcssSourceMap;
function combineSourcemapsIfExists(filename, map1, map2) {
    return map1 && map2
        ? (0, utils_1.combineSourcemaps)(filename, [
            // type of version property of ExistingRawSourceMap is number
            // but it is always 3
            map1,
            map2,
        ])
        : map1;
}
async function resolvePostcssConfig(config) {
    let result = postcssConfigCache.get(config);
    if (result !== undefined) {
        return result;
    }
    // inline postcss config via vite config
    const inlineOptions = config.css?.postcss;
    if ((0, utils_1.isObject)(inlineOptions)) {
        const options = { ...inlineOptions };
        delete options.plugins;
        result = {
            options,
            plugins: inlineOptions.plugins || [],
        };
    }
    else {
        const searchPath = (0, shared_1.isString)(inlineOptions) ? inlineOptions : config.root;
        try {
            result = await (0, postcss_load_config_1.default)({}, searchPath);
        }
        catch (e) {
            if (!/No PostCSS Config found/.test(e.message)) {
                if (e instanceof Error) {
                    const { name, message, stack } = e;
                    e.name = 'Failed to load PostCSS config';
                    e.message = `Failed to load PostCSS config (searchPath: ${searchPath}): [${name}] ${message}\n${stack}`;
                    e.stack = ''; // add stack to message to retain stack
                    throw e;
                }
                else {
                    throw new Error(`Failed to load PostCSS config: ${e}`);
                }
            }
            result = null;
        }
    }
    postcssConfigCache.set(config, result);
    return result;
}
// https://drafts.csswg.org/css-syntax-3/#identifier-code-point
exports.cssUrlRE = /(?<=^|[^\w\-\u0080-\uffff])url\(\s*('[^']+'|"[^"]+"|[^'")]+)\s*\)/;
exports.cssDataUriRE = /(?<=^|[^\w\-\u0080-\uffff])data-uri\(\s*('[^']+'|"[^"]+"|[^'")]+)\s*\)/;
exports.importCssRE = /@import ('[^']+\.css'|"[^"]+\.css"|[^'")]+\.css)/;
const cssImageSetRE = /(?<=image-set\()((?:[\w\-]+\([^\)]*\)|[^)])*)(?=\))/;
const UrlRewritePostcssPlugin = (opts) => {
    if (!opts) {
        throw new Error('base or replace is required');
    }
    return {
        postcssPlugin: 'vite-url-rewrite',
        Once(root) {
            const promises = [];
            root.walkDecls((declaration) => {
                const isCssUrl = exports.cssUrlRE.test(declaration.value);
                const isCssImageSet = cssImageSetRE.test(declaration.value);
                if (isCssUrl || isCssImageSet) {
                    const replacerForDeclaration = (rawUrl) => {
                        const importer = declaration.source?.input.file;
                        return opts.replacer(rawUrl, importer, declaration.source);
                    };
                    const rewriterToUse = isCssImageSet
                        ? rewriteCssImageSet
                        : rewriteCssUrls;
                    promises.push(rewriterToUse(declaration.value, replacerForDeclaration).then((url) => {
                        declaration.value = url;
                    }));
                }
            });
            if (promises.length) {
                return Promise.all(promises);
            }
        },
    };
};
UrlRewritePostcssPlugin.postcss = true;
function rewriteCssUrls(css, replacer) {
    return (0, utils_1.asyncReplace)(css, exports.cssUrlRE, async (match) => {
        const [matched, rawUrl] = match;
        return await doUrlReplace(rawUrl, matched, replacer);
    });
}
function rewriteCssDataUris(css, replacer) {
    return (0, utils_1.asyncReplace)(css, exports.cssDataUriRE, async (match) => {
        const [matched, rawUrl] = match;
        return await doUrlReplace(rawUrl, matched, replacer, 'data-uri');
    });
}
function rewriteImportCss(css, replacer) {
    return (0, utils_1.asyncReplace)(css, exports.importCssRE, async (match) => {
        const [matched, rawUrl] = match;
        return await doImportCSSReplace(rawUrl, matched, replacer);
    });
}
// TODO: image and cross-fade could contain a "url" that needs to be processed
// https://drafts.csswg.org/css-images-4/#image-notation
// https://drafts.csswg.org/css-images-4/#cross-fade-function
const cssNotProcessedRE = /(gradient|element|cross-fade|image)\(/;
async function rewriteCssImageSet(css, replacer) {
    return await (0, utils_1.asyncReplace)(css, cssImageSetRE, async (match) => {
        const [, rawUrl] = match;
        const url = await (0, utils_1.processSrcSet)(rawUrl, async ({ url }) => {
            // the url maybe url(...)
            if (exports.cssUrlRE.test(url)) {
                return await rewriteCssUrls(url, replacer);
            }
            if (!cssNotProcessedRE.test(url)) {
                return await doUrlReplace(url, url, replacer);
            }
            return url;
        });
        return url;
    });
}
async function doUrlReplace(rawUrl, matched, replacer, funcName = 'url') {
    let wrap = '';
    const first = rawUrl[0];
    if (first === `"` || first === `'`) {
        wrap = first;
        rawUrl = rawUrl.slice(1, -1);
    }
    if ((0, utils_1.isExternalUrl)(rawUrl) ||
        (0, utils_1.isDataUrl)(rawUrl) ||
        rawUrl.startsWith('#') ||
        varRE.test(rawUrl)) {
        return matched;
    }
    const newUrl = await replacer(rawUrl);
    if (wrap === '' && newUrl !== encodeURI(newUrl)) {
        // The new url might need wrapping even if the original did not have it, e.g. if a space was added during replacement
        wrap = "'";
    }
    return `${funcName}(${wrap}${newUrl}${wrap})`;
}
async function doImportCSSReplace(rawUrl, matched, replacer) {
    let wrap = '';
    const first = rawUrl[0];
    if (first === `"` || first === `'`) {
        wrap = first;
        rawUrl = rawUrl.slice(1, -1);
    }
    if ((0, utils_1.isExternalUrl)(rawUrl) || (0, utils_1.isDataUrl)(rawUrl) || rawUrl.startsWith('#')) {
        return matched;
    }
    return `@import ${wrap}${await replacer(rawUrl)}${wrap}`;
}
async function minifyCSS(css, config) {
    try {
        const { code, warnings } = await Promise.resolve().then(() => __importStar(require('esbuild'))).then(({ transform }) => {
            if (process.env.VUE_APP_DARK_MODE !== 'true') {
                const postcssRoot = Postcss.parse(css);
                (0, uniapp_1.filterPrefersColorScheme)(postcssRoot, true);
                css = postcssRoot.toResult().css;
            }
            return transform(css, {
                loader: 'css',
                minify: true,
                target: config.build.cssTarget || undefined,
            });
        });
        if (warnings.length) {
            const msgs = await Promise.resolve().then(() => __importStar(require('esbuild'))).then(({ formatMessages }) => {
                return formatMessages(warnings, { kind: 'warning' });
            });
            config.logger.warn(picocolors_1.default.yellow(`warnings when minifying css:\n${msgs.join('\n')}`));
        }
        return code;
    }
    catch (e) {
        if (e.errors) {
            const msgs = await Promise.resolve().then(() => __importStar(require('esbuild'))).then(({ formatMessages }) => {
                return formatMessages(e.errors, { kind: 'error' });
            });
            e.frame = '\n' + msgs.join('\n');
            e.loc = e.errors[0].location;
        }
        throw e;
    }
}
exports.minifyCSS = minifyCSS;
async function hoistAtRules(css) {
    const s = new magic_string_1.default(css);
    const cleanCss = (0, cleanString_1.emptyCssComments)(css);
    let match;
    // #1845
    // CSS @import can only appear at top of the file. We need to hoist all @import
    // to top when multiple files are concatenated.
    // match until semicolon that's not in quotes
    const atImportRE = /@import\s*(?:url\([^\)]*\)|"([^"]|(?<=\\)")*"|'([^']|(?<=\\)')*'|[^;]*).*?;/gm;
    while ((match = atImportRE.exec(cleanCss))) {
        s.remove(match.index, match.index + match[0].length);
        // Use `appendLeft` instead of `prepend` to preserve original @import order
        s.appendLeft(0, match[0]);
    }
    // #6333
    // CSS @charset must be the top-first in the file, hoist the first to top
    const atCharsetRE = /@charset\s*(?:"([^"]|(?<=\\)")*"|'([^']|(?<=\\)')*'|[^;]*).*?;/gm;
    let foundCharset = false;
    while ((match = atCharsetRE.exec(cleanCss))) {
        s.remove(match.index, match.index + match[0].length);
        if (!foundCharset) {
            s.prepend(match[0]);
            foundCharset = true;
        }
    }
    return s.toString();
}
exports.hoistAtRules = hoistAtRules;
const loadedPreprocessors = {};
function loadPreprocessor(lang, root) {
    if (lang in loadedPreprocessors) {
        return loadedPreprocessors[lang];
    }
    try {
        // Search for the preprocessor in the root directory first, and fall back
        // to the default require paths.
        const fallbackPaths = require.resolve.paths?.(lang) || [];
        const resolved = require.resolve(lang, { paths: [root, ...fallbackPaths] });
        return (loadedPreprocessors[lang] = require(resolved));
    }
    catch (e) {
        throw new Error(`Preprocessor dependency "${lang}" not found. Did you install it?`);
    }
}
// .scss/.sass processor
const scss = async (source, root, options, resolvers, isNVue) => {
    const render = loadPreprocessor("sass" /* PreprocessLang.sass */, root).render;
    // NOTE: `sass` always runs it's own importer first, and only falls back to
    // the `importer` option when it can't resolve a path
    const internalImporter = (url, importer, done) => {
        resolvers.sass(url, importer).then((resolved) => {
            if (resolved) {
                rebaseUrls(resolved, options.filename, options.alias, isNVue)
                    .then((data) => done?.(data))
                    .catch((data) => done?.(data));
            }
            else {
                done?.(null);
            }
        });
    };
    const importer = [internalImporter];
    if (options.importer) {
        (0, shared_1.isArray)(options.importer)
            ? importer.push(...options.importer)
            : importer.push(options.importer);
    }
    const { content: data, map: additionalMap } = await getSource(source, options.filename, options.additionalData, options.enableSourcemap);
    const finalOptions = {
        ...options,
        data,
        file: options.filename,
        outFile: options.filename,
        importer,
        ...(options.enableSourcemap
            ? {
                sourceMap: true,
                omitSourceMapUrl: true,
                sourceMapRoot: path_1.default.dirname(options.filename),
            }
            : {}),
    };
    try {
        const result = await new Promise((resolve, reject) => {
            render(finalOptions, (err, res) => {
                if (err) {
                    reject(err);
                }
                else {
                    resolve(res);
                }
            });
        });
        const deps = result.stats.includedFiles;
        const map = result.map
            ? JSON.parse(result.map.toString())
            : undefined;
        return {
            code: result.css.toString(),
            map,
            additionalMap,
            errors: [],
            deps,
        };
    }
    catch (e) {
        // normalize SASS error
        e.id = e.file;
        e.frame = e.formatted;
        return { code: '', errors: [e], deps: [] };
    }
};
const sass = (source, root, options, aliasResolver, isNVue) => scss(source, root, {
    ...options,
    indentedSyntax: true,
}, aliasResolver, isNVue);
function preprocessCss(content, isNVue = false) {
    if (content.includes('#endif')) {
        return isNVue ? (0, preprocess_1.preNVueCss)(content) : (0, preprocess_1.preCss)(content);
    }
    return content;
}
/**
 * relative url() inside \@imported sass and less files must be rebased to use
 * root file as base.
 */
async function rebaseUrls(file, rootFile, alias, isNVue = false) {
    file = path_1.default.resolve(file); // ensure os-specific flashes
    // fixed by xxxxxx 条件编译
    let contents = preprocessCss(fs_extra_1.default.readFileSync(file, 'utf-8'), isNVue);
    // in the same dir, no need to rebase
    const fileDir = path_1.default.dirname(file);
    const rootDir = path_1.default.dirname(rootFile);
    if (fileDir === rootDir) {
        return { file, contents };
    }
    // no url()
    const hasUrls = exports.cssUrlRE.test(contents);
    // data-uri() calls
    const hasDataUris = exports.cssDataUriRE.test(contents);
    // no @import xxx.css
    const hasImportCss = exports.importCssRE.test(contents);
    if (!hasUrls && !hasDataUris && !hasImportCss) {
        return { file, contents };
    }
    let rebased;
    const rebaseFn = (url) => {
        if (url.startsWith('/'))
            return url;
        // match alias, no need to rewrite
        for (const { find } of alias) {
            const matches = (0, shared_1.isString)(find) ? url.startsWith(find) : find.test(url);
            if (matches) {
                return url;
            }
        }
        const absolute = path_1.default.resolve(fileDir, url);
        const relative = path_1.default.relative(rootDir, absolute);
        return (0, utils_1.normalizePath)(relative);
    };
    // fix css imports in less such as `@import "foo.css"`
    if (hasImportCss) {
        contents = await rewriteImportCss(contents, rebaseFn);
    }
    if (hasUrls) {
        contents = await rewriteCssUrls(rebased || contents, rebaseFn);
    }
    if (hasDataUris) {
        contents = await rewriteCssDataUris(rebased || contents, rebaseFn);
    }
    return {
        file,
        contents,
    };
}
// .less
const less = async (source, root, options, resolvers, isNVue) => {
    const nodeLess = loadPreprocessor("less" /* PreprocessLang.less */, root);
    const viteResolverPlugin = createViteLessPlugin(nodeLess, options.filename, options.alias, resolvers, isNVue);
    const { content, map: additionalMap } = await getSource(source, options.filename, options.additionalData, options.enableSourcemap);
    let result;
    try {
        result = await nodeLess.render(content, {
            ...options,
            plugins: [viteResolverPlugin, ...(options.plugins || [])],
            ...(options.enableSourcemap
                ? {
                    sourceMap: {
                        outputSourceFiles: true,
                        sourceMapFileInline: false,
                    },
                }
                : {}),
        });
    }
    catch (e) {
        const error = e;
        // normalize error info
        const normalizedError = new Error(error.message || error.type);
        normalizedError.loc = {
            file: error.filename || options.filename,
            line: error.line,
            column: error.column,
        };
        return { code: '', errors: [normalizedError], deps: [] };
    }
    const map = result.map && JSON.parse(result.map);
    if (map) {
        delete map.sourcesContent;
    }
    return {
        code: result.css.toString(),
        map,
        additionalMap,
        deps: result.imports,
        errors: [],
    };
};
/**
 * Less manager, lazy initialized
 */
let ViteLessManager;
function createViteLessPlugin(less, rootFile, alias, resolvers, isNVue) {
    if (!ViteLessManager) {
        ViteLessManager = class ViteManager extends less.FileManager {
            constructor(rootFile, resolvers, alias) {
                super();
                this.rootFile = rootFile;
                this.resolvers = resolvers;
                this.alias = alias;
            }
            supports() {
                return true;
            }
            supportsSync() {
                return false;
            }
            async loadFile(filename, dir, opts, env) {
                const resolved = await this.resolvers.less(filename, path_1.default.join(dir, '*'));
                if (resolved) {
                    const result = await rebaseUrls(resolved, this.rootFile, this.alias, isNVue);
                    let contents;
                    if (result && 'contents' in result) {
                        contents = result.contents;
                    }
                    else {
                        contents = fs_extra_1.default.readFileSync(resolved, 'utf-8');
                    }
                    return {
                        filename: path_1.default.resolve(resolved),
                        contents,
                    };
                }
                else {
                    return super.loadFile(filename, dir, opts, env);
                }
            }
        };
    }
    return {
        install(_, pluginManager) {
            pluginManager.addFileManager(new ViteLessManager(rootFile, resolvers, alias));
        },
        minVersion: [3, 0, 0],
    };
}
// .styl
const styl = async (source, root, options) => {
    const nodeStylus = loadPreprocessor("stylus" /* PreprocessLang.stylus */, root);
    // Get source with preprocessor options.additionalData. Make sure a new line separator
    // is added to avoid any render error, as added stylus content may not have semi-colon separators
    const { content, map: additionalMap } = await getSource(source, options.filename, options.additionalData, options.enableSourcemap, '\n');
    // Get preprocessor options.imports dependencies as stylus
    // does not return them with its builtin `.deps()` method
    const importsDeps = (options.imports ?? []).map((dep) => path_1.default.resolve(dep));
    try {
        const ref = nodeStylus(content, options);
        if (options.enableSourcemap) {
            ref.set('sourcemap', {
                comment: false,
                inline: false,
                basePath: root,
            });
        }
        const result = ref.render();
        // Concat imports deps with computed deps
        const deps = [...ref.deps(), ...importsDeps];
        // @ts-expect-error sourcemap exists
        const map = ref.sourcemap;
        return {
            code: result,
            map: formatStylusSourceMap(map, root),
            additionalMap,
            errors: [],
            deps,
        };
    }
    catch (e) {
        return { code: '', errors: [e], deps: [] };
    }
};
function formatStylusSourceMap(mapBefore, root) {
    if (!mapBefore)
        return undefined;
    const map = { ...mapBefore };
    const resolveFromRoot = (p) => (0, utils_1.normalizePath)(path_1.default.resolve(root, p));
    if (map.file) {
        map.file = resolveFromRoot(map.file);
    }
    map.sources = map.sources.map(resolveFromRoot);
    return map;
}
async function getSource(source, filename, additionalData, enableSourcemap, sep = '') {
    if (!additionalData)
        return { content: source };
    if ((0, shared_1.isFunction)(additionalData)) {
        const newContent = await additionalData(source, filename);
        if ((0, shared_1.isString)(newContent)) {
            return { content: newContent };
        }
        return newContent;
    }
    if (!enableSourcemap) {
        return { content: additionalData + sep + source };
    }
    const ms = new magic_string_1.default(source);
    ms.appendLeft(0, sep);
    ms.appendLeft(0, additionalData);
    const map = ms.generateMap({ hires: true });
    map.file = filename;
    map.sources = [filename];
    return {
        content: ms.toString(),
        map,
    };
}
const preProcessors = Object.freeze({
    ["less" /* PreprocessLang.less */]: less,
    ["sass" /* PreprocessLang.sass */]: sass,
    ["scss" /* PreprocessLang.scss */]: scss,
    ["styl" /* PreprocessLang.styl */]: styl,
    ["stylus" /* PreprocessLang.stylus */]: styl,
});
function isPreProcessor(lang) {
    return lang && lang in preProcessors;
}
const preCssExtNames = ['.scss', '.sass', '.styl', '.stylus'];
/**
 * 重写 readFileSync
 * 目前主要解决 scss 文件被 @import 的条件编译
 */
function rewriteScssReadFileSync() {
    // 目前 1.0 App 端，只要包含了APP-NVUE条件编译，就不pre，因为区分不出来APP-NVUE
    const ignoreAppNVue = process.env.UNI_APP_X !== 'true' &&
        (process.env.UNI_PLATFORM === 'app' ||
            process.env.UNI_PLATFORM === 'app-plus');
    const { readFileSync } = fs_1.default;
    fs_1.default.readFileSync = ((filepath, options) => {
        const content = readFileSync(filepath, options);
        if ((0, shared_1.isString)(filepath) &&
            (0, shared_1.isString)(content) &&
            preCssExtNames.includes(path_1.default.extname(filepath)) &&
            content.includes('#endif')
        // 目前无法区分app-nvue
        ) {
            if (ignoreAppNVue && content.includes('APP-NVUE')) {
                return content;
            }
            return (0, preprocess_1.preCss)(content);
        }
        return content;
    });
    // TODO 在 web 端，目前使用的是 vite 内置的 css 预处理
    // 最新的 vite less 使用的是 readFile，https://github.com/vitejs/vite/blob/main/packages/vite/src/node/plugins/css.ts#L2336
    // 拦截 readFile 的话，要考虑的细节较多，暂不实现了
    // 稍后 web 端，也直接内部定制 css 的预处理
}
exports.rewriteScssReadFileSync = rewriteScssReadFileSync;
