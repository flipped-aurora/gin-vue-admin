"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.forEachEmbeddedFile = exports.updateVirtualFileMaps = exports.createVirtualFiles = void 0;
const source_map_1 = require("@volar/source-map");
const sourceMaps_1 = require("./sourceMaps");
function createVirtualFiles(languages) {
    const sourceFiles = new Map();
    const virtualFiles = new Map();
    const virtualFileMaps = new WeakMap();
    const virtualFileToMirrorMap = new WeakMap();
    return {
        allSources() {
            return Array.from(sourceFiles.values());
        },
        updateSource(fileName, snapshot, languageId) {
            const key = normalizePath(fileName);
            const value = sourceFiles.get(key);
            if (value) {
                if (value.languageId !== languageId) {
                    // languageId changed
                    this.deleteSource(fileName);
                    return this.updateSource(fileName, snapshot, languageId);
                }
                else {
                    value.snapshot = snapshot;
                    deleteVirtualFiles(value);
                    value.language.updateVirtualFile(value.root, snapshot);
                    updateVirtualFiles(value);
                    return value.root; // updated
                }
            }
            for (const language of languages) {
                const virtualFile = language.createVirtualFile(fileName, snapshot, languageId);
                if (virtualFile) {
                    const source = { fileName, languageId, snapshot, root: virtualFile, language };
                    sourceFiles.set(key, source);
                    updateVirtualFiles(source);
                    return virtualFile; // created
                }
            }
        },
        deleteSource(fileName) {
            const key = normalizePath(fileName);
            const value = sourceFiles.get(key);
            if (value) {
                value.language.deleteVirtualFile?.(value.root);
                sourceFiles.delete(key); // deleted
                deleteVirtualFiles(value);
            }
        },
        getSource(fileName) {
            const key = normalizePath(fileName);
            return sourceFiles.get(key);
        },
        hasSource: (fileName) => sourceFiles.has(normalizePath(fileName)),
        getMirrorMap: getMirrorMap,
        getMaps: getMapsByVirtualFile,
        hasVirtualFile(fileName) {
            return !!virtualFiles.get(normalizePath(fileName));
        },
        getVirtualFile(fileName) {
            const sourceAndVirtual = virtualFiles.get(normalizePath(fileName));
            if (sourceAndVirtual) {
                return [sourceAndVirtual.virtualFile, sourceAndVirtual.source];
            }
            return [undefined, undefined];
        },
    };
    function deleteVirtualFiles(source) {
        forEachEmbeddedFile(source.root, file => {
            virtualFiles.delete(normalizePath(file.fileName));
        });
    }
    function updateVirtualFiles(source) {
        forEachEmbeddedFile(source.root, file => {
            virtualFiles.set(normalizePath(file.fileName), { virtualFile: file, source });
        });
    }
    function getMapsByVirtualFile(virtualFile) {
        if (!virtualFileMaps.has(virtualFile.snapshot)) {
            virtualFileMaps.set(virtualFile.snapshot, new Map());
        }
        updateVirtualFileMaps(virtualFile, sourceFileName => {
            if (sourceFileName) {
                const source = sourceFiles.get(normalizePath(sourceFileName));
                return [sourceFileName, source.snapshot];
            }
            else {
                const source = virtualFiles.get(normalizePath(virtualFile.fileName)).source;
                return [source.fileName, source.snapshot];
            }
        }, virtualFileMaps.get(virtualFile.snapshot));
        return virtualFileMaps.get(virtualFile.snapshot);
    }
    function getMirrorMap(file) {
        if (!virtualFileToMirrorMap.has(file.snapshot)) {
            virtualFileToMirrorMap.set(file.snapshot, file.mirrorBehaviorMappings ? new sourceMaps_1.MirrorMap(file.mirrorBehaviorMappings) : undefined);
        }
        return virtualFileToMirrorMap.get(file.snapshot);
    }
}
exports.createVirtualFiles = createVirtualFiles;
function updateVirtualFileMaps(virtualFile, getSourceSnapshot, map = new Map()) {
    const sources = new Set();
    for (const mapping of virtualFile.mappings) {
        if (sources.has(mapping.source))
            continue;
        sources.add(mapping.source);
        const source = getSourceSnapshot(mapping.source);
        if (!source)
            continue;
        if (!map.has(source[0]) || map.get(source[0])[0] !== source[1]) {
            map.set(source[0], [source[1], new source_map_1.SourceMap(virtualFile.mappings.filter(mapping2 => mapping2.source === mapping.source))]);
        }
    }
    return map;
}
exports.updateVirtualFileMaps = updateVirtualFileMaps;
function forEachEmbeddedFile(file, cb) {
    cb(file);
    for (const embeddedFile of file.embeddedFiles) {
        forEachEmbeddedFile(embeddedFile, cb);
    }
}
exports.forEachEmbeddedFile = forEachEmbeddedFile;
function normalizePath(fileName) {
    return fileName.replace(/\\/g, '/').toLowerCase();
}
//# sourceMappingURL=virtualFiles.js.map