"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createLanguageServiceHost = void 0;
const path = require("path-browserify");
const utilities_1 = require("./typescript/utilities");
const fileVersions = new Map();
function createLanguageServiceHost(ctx, ts, sys) {
    let lastProjectVersion;
    let tsProjectVersion = 0;
    let tsFileNames = [];
    let tsDirectories = new Set();
    const _tsHost = {
        ...sys,
        getCurrentDirectory: () => ctx.host.workspacePath,
        getCompilationSettings: () => ctx.host.getCompilationSettings(),
        getCancellationToken: ctx.host.getCancellationToken ? () => ctx.host.getCancellationToken() : undefined,
        getLocalizedDiagnosticMessages: ctx.host.getLocalizedDiagnosticMessages ? () => ctx.host.getLocalizedDiagnosticMessages() : undefined,
        getProjectReferences: ctx.host.getProjectReferences ? () => ctx.host.getProjectReferences() : undefined,
        getDefaultLibFileName: (options) => {
            try {
                return ts.getDefaultLibFilePath(options);
            }
            catch {
                // web
                return `/node_modules/typescript/lib/${ts.getDefaultLibFileName(options)}`;
            }
        },
        useCaseSensitiveFileNames: () => sys.useCaseSensitiveFileNames,
        getNewLine: () => sys.newLine,
        readFile: fileName => {
            const snapshot = getScriptSnapshot(fileName);
            if (snapshot) {
                return snapshot.getText(0, snapshot.getLength());
            }
        },
        readDirectory,
        getDirectories,
        directoryExists,
        fileExists,
        getProjectVersion: () => {
            return tsProjectVersion + ':' + sys.version;
        },
        getTypeRootsVersion: () => {
            return sys.version ?? -1; // TODO: only update for /node_modules changes?
        },
        getScriptFileNames: () => tsFileNames,
        getScriptVersion,
        getScriptSnapshot,
        getScriptKind(fileName) {
            if (ts) {
                if (ctx.virtualFiles.hasSource(fileName))
                    return ts.ScriptKind.Deferred;
                switch (path.extname(fileName)) {
                    case '.js': return ts.ScriptKind.JS;
                    case '.cjs': return ts.ScriptKind.JS;
                    case '.mjs': return ts.ScriptKind.JS;
                    case '.jsx': return ts.ScriptKind.JSX;
                    case '.ts': return ts.ScriptKind.TS;
                    case '.cts': return ts.ScriptKind.TS;
                    case '.mts': return ts.ScriptKind.TS;
                    case '.tsx': return ts.ScriptKind.TSX;
                    case '.json': return ts.ScriptKind.JSON;
                    default: return ts.ScriptKind.Unknown;
                }
            }
            return 0;
        },
    };
    const fsFileSnapshots = new Map();
    if (ctx.host.resolveModuleName) {
        // TODO: can this share between monorepo packages?
        const moduleCache = ts.createModuleResolutionCache(_tsHost.getCurrentDirectory(), _tsHost.useCaseSensitiveFileNames ? s => s : s => s.toLowerCase(), _tsHost.getCompilationSettings());
        let lastSysVersion = sys.version;
        _tsHost.resolveModuleNameLiterals = (moduleLiterals, containingFile, redirectedReference, options, sourceFile) => {
            if (lastSysVersion !== sys.version) {
                lastSysVersion = sys.version;
                moduleCache.clear();
            }
            return moduleLiterals.map((moduleLiteral) => {
                let moduleName = moduleLiteral.text;
                moduleName = ctx.host.resolveModuleName(moduleName, sourceFile.impliedNodeFormat);
                return ts.resolveModuleName(moduleName, containingFile, options, _tsHost, moduleCache, redirectedReference, sourceFile.impliedNodeFormat);
            });
        };
        _tsHost.resolveModuleNames = (moduleNames, containingFile, _reusedNames, redirectedReference, options, sourceFile) => {
            if (lastSysVersion !== sys.version) {
                lastSysVersion = sys.version;
                moduleCache.clear();
            }
            return moduleNames.map((moduleName) => {
                moduleName = ctx.host.resolveModuleName(moduleName, sourceFile?.impliedNodeFormat);
                return ts.resolveModuleName(moduleName, containingFile, options, _tsHost, moduleCache, redirectedReference, sourceFile?.impliedNodeFormat).resolvedModule;
            });
        };
    }
    let oldTsVirtualFileSnapshots = new Set();
    let oldOtherVirtualFileSnapshots = new Set();
    return new Proxy(_tsHost, {
        get: (target, property) => {
            sync();
            return target[property];
        },
    });
    function sync() {
        const newProjectVersion = ctx.host.getProjectVersion();
        const shouldUpdate = newProjectVersion !== lastProjectVersion;
        if (!shouldUpdate)
            return;
        lastProjectVersion = newProjectVersion;
        const newTsVirtualFileSnapshots = new Set();
        const newOtherVirtualFileSnapshots = new Set();
        for (const { root } of ctx.virtualFiles.allSources()) {
            forEachEmbeddedFile(root, embedded => {
                if (embedded.kind === 1) {
                    newTsVirtualFileSnapshots.add(embedded.snapshot);
                }
                else {
                    newOtherVirtualFileSnapshots.add(embedded.snapshot);
                }
            });
        }
        if (!setEquals(oldTsVirtualFileSnapshots, newTsVirtualFileSnapshots)) {
            tsProjectVersion++;
        }
        else if (setEquals(oldOtherVirtualFileSnapshots, newOtherVirtualFileSnapshots)) {
            // no any meta language files update, it mean project version was update by source files this time
            tsProjectVersion++;
        }
        oldTsVirtualFileSnapshots = newTsVirtualFileSnapshots;
        oldOtherVirtualFileSnapshots = newOtherVirtualFileSnapshots;
        const tsFileNamesSet = new Set();
        for (const { root } of ctx.virtualFiles.allSources()) {
            forEachEmbeddedFile(root, embedded => {
                if (embedded.kind === 1) {
                    tsFileNamesSet.add(embedded.fileName); // virtual .ts
                }
            });
        }
        for (const fileName of ctx.host.getScriptFileNames()) {
            if (!ctx.virtualFiles.hasSource(fileName)) {
                tsFileNamesSet.add(fileName); // .ts
            }
        }
        tsFileNames = [...tsFileNamesSet];
        // Update tsDirectories for `directoryExists()`
        tsDirectories.clear();
        for (const fileName of tsFileNames) {
            tsDirectories.add(path.dirname(normalizePath(fileName)));
        }
    }
    function readDirectory(dirName, extensions, excludes, includes, depth) {
        let matches = (0, utilities_1.matchFiles)(dirName, extensions, excludes, includes, sys?.useCaseSensitiveFileNames ?? false, ctx.host.workspacePath, depth, (dirPath) => {
            const files = [];
            for (const fileName of tsFileNames) {
                if (fileName.toLowerCase().startsWith(dirPath.toLowerCase())) {
                    const baseName = fileName.substring(dirPath.length);
                    if (baseName.indexOf('/') === -1) {
                        files.push(baseName);
                    }
                }
            }
            return {
                files,
                directories: getVirtualFileDirectories(dirPath),
            };
        }, sys?.realpath ? (path => sys.realpath(path)) : (path => path));
        matches = matches.map(match => {
            const [_, source] = ctx.virtualFiles.getVirtualFile(match);
            if (source) {
                return source.fileName;
            }
            return match;
        });
        return [...new Set([
                ...matches,
                ...sys.readDirectory(dirName, extensions, excludes, includes, depth),
            ])];
    }
    function getDirectories(dirName) {
        return [...new Set([
                ...getVirtualFileDirectories(dirName),
                ...sys.getDirectories(dirName),
            ])];
    }
    function getVirtualFileDirectories(dirName) {
        const names = new Set();
        for (const fileName of tsFileNames) {
            if (fileName.toLowerCase().startsWith(dirName.toLowerCase())) {
                const path = fileName.substring(dirName.length);
                if (path.indexOf('/') >= 0) {
                    names.add(path.split('/')[0]);
                }
            }
        }
        return [...names];
    }
    function getScriptSnapshot(fileName) {
        // virtual files
        const [virtualFile] = ctx.virtualFiles.getVirtualFile(fileName);
        if (virtualFile) {
            return virtualFile.snapshot;
        }
        // root files / opened files
        const tsScript = ctx.host.getScriptSnapshot(fileName);
        if (tsScript) {
            return tsScript;
        }
        // fs files
        const cache = fsFileSnapshots.get(fileName);
        const modifiedTime = sys.getModifiedTime?.(fileName)?.valueOf();
        if (!cache || cache[0] !== modifiedTime) {
            if (sys.fileExists(fileName)) {
                const text = sys.readFile(fileName);
                const snapshot = text !== undefined ? ts.ScriptSnapshot.fromString(text) : undefined;
                fsFileSnapshots.set(fileName, [modifiedTime, snapshot]);
            }
            else {
                fsFileSnapshots.set(fileName, [modifiedTime, undefined]);
            }
        }
        return fsFileSnapshots.get(fileName)?.[1];
    }
    function getScriptVersion(fileName) {
        // virtual files / root files / opened files
        const [virtualFile] = ctx.virtualFiles.getVirtualFile(fileName);
        const snapshot = virtualFile?.snapshot ?? ctx.host.getScriptSnapshot(fileName);
        if (snapshot) {
            if (!fileVersions.has(fileName)) {
                fileVersions.set(fileName, { lastVersion: 0, snapshotVersions: new WeakMap() });
            }
            const version = fileVersions.get(fileName);
            if (!version.snapshotVersions.has(snapshot)) {
                version.snapshotVersions.set(snapshot, version.lastVersion++);
            }
            return version.snapshotVersions.get(snapshot).toString();
        }
        // fs files
        return sys.getModifiedTime?.(fileName)?.valueOf().toString() ?? '';
    }
    function directoryExists(dirName) {
        return tsDirectories.has(normalizePath(dirName)) || sys.directoryExists(dirName);
    }
    function fileExists(fileName) {
        // fill external virtual files
        const ext = fileName.substring(fileName.lastIndexOf('.'));
        if (ext === '.js'
            || ext === '.ts'
            || ext === '.jsx'
            || ext === '.tsx') {
            /**
             * If try to access a external .vue file that outside of the project,
             * the file will not process by language service host,
             * so virtual file will not be created.
             *
             * We try to create virtual file here.
             */
            const sourceFileName = fileName.substring(0, fileName.lastIndexOf('.'));
            if (!ctx.virtualFiles.hasSource(sourceFileName)) {
                const scriptSnapshot = getScriptSnapshot(sourceFileName);
                if (scriptSnapshot) {
                    ctx.virtualFiles.updateSource(sourceFileName, scriptSnapshot, ctx.host.getLanguageId?.(sourceFileName));
                }
            }
        }
        // virtual files
        if (ctx.virtualFiles.hasVirtualFile(fileName)) {
            return true;
        }
        // root files
        if (ctx.host.getScriptSnapshot(fileName)) {
            return true;
        }
        // fs files
        return !!sys.fileExists(fileName);
    }
}
exports.createLanguageServiceHost = createLanguageServiceHost;
function setEquals(a, b) {
    if (a.size !== b.size)
        return false;
    for (const item of a) {
        if (!b.has(item))
            return false;
    }
    return true;
}
function forEachEmbeddedFile(file, cb) {
    cb(file);
    for (const embeddedFile of file.embeddedFiles) {
        forEachEmbeddedFile(embeddedFile, cb);
    }
}
function normalizePath(fileName) {
    return fileName.replace(/\\/g, '/').toLowerCase();
}
//# sourceMappingURL=languageServiceHost.js.map