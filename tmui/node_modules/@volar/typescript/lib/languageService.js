"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.decorateLanguageService = void 0;
const language_core_1 = require("@volar/language-core");
function decorateLanguageService(virtualFiles, languageService, isTsPlugin) {
    const _organizeImports = languageService.organizeImports.bind(languageService);
    const _getDefinitionAtPosition = languageService.getDefinitionAtPosition.bind(languageService);
    const _getDefinitionAndBoundSpan = languageService.getDefinitionAndBoundSpan.bind(languageService);
    const _getTypeDefinitionAtPosition = languageService.getTypeDefinitionAtPosition.bind(languageService);
    const _getImplementationAtPosition = languageService.getImplementationAtPosition.bind(languageService);
    const _getFileReferences = languageService.getFileReferences.bind(languageService);
    const _findRenameLocations = languageService.findRenameLocations.bind(languageService);
    const _getReferencesAtPosition = languageService.getReferencesAtPosition.bind(languageService);
    const _findReferences = languageService.findReferences.bind(languageService);
    languageService.organizeImports = organizeImports;
    languageService.getDefinitionAtPosition = getDefinitionAtPosition;
    languageService.getDefinitionAndBoundSpan = getDefinitionAndBoundSpan;
    languageService.getTypeDefinitionAtPosition = getTypeDefinitionAtPosition;
    languageService.getImplementationAtPosition = getImplementationAtPosition;
    languageService.findRenameLocations = findRenameLocations;
    languageService.getReferencesAtPosition = getReferencesAtPosition;
    languageService.getFileReferences = getFileReferences;
    languageService.findReferences = findReferences;
    // apis
    function organizeImports(args, formatOptions, preferences) {
        let edits = [];
        const file = virtualFiles.getSource(args.fileName)?.root;
        if (file) {
            (0, language_core_1.forEachEmbeddedFile)(file, embeddedFile => {
                if (embeddedFile.kind === language_core_1.FileKind.TypeScriptHostFile && embeddedFile.capabilities.codeAction) {
                    edits = edits.concat(_organizeImports({
                        ...args,
                        fileName: embeddedFile.fileName,
                    }, formatOptions, preferences));
                }
            });
        }
        else {
            return _organizeImports(args, formatOptions, preferences);
        }
        return edits.map(transformFileTextChanges).filter(notEmpty);
    }
    function getReferencesAtPosition(fileName, position) {
        return findLocations(fileName, position, 'references');
    }
    function getFileReferences(fileName) {
        return findLocations(fileName, -1, 'fileReferences');
    }
    function getDefinitionAtPosition(fileName, position) {
        return findLocations(fileName, position, 'definition');
    }
    function getTypeDefinitionAtPosition(fileName, position) {
        return findLocations(fileName, position, 'typeDefinition');
    }
    function getImplementationAtPosition(fileName, position) {
        return findLocations(fileName, position, 'implementation');
    }
    function findRenameLocations(fileName, position, findInStrings, findInComments, preferences) {
        return findLocations(fileName, position, 'rename', findInStrings, findInComments, preferences);
    }
    function findLocations(fileName, position, mode, findInStrings = false, findInComments = false, preferences) {
        const loopChecker = new Set();
        let symbols = [];
        withMirrors(fileName, position);
        return symbols.map(s => transformDocumentSpanLike(s, mode === 'definition')).filter(notEmpty);
        function withMirrors(fileName, position) {
            if (loopChecker.has(fileName + ':' + position))
                return;
            loopChecker.add(fileName + ':' + position);
            const _symbols = mode === 'definition' ? _getDefinitionAtPosition(fileName, position)
                : mode === 'typeDefinition' ? _getTypeDefinitionAtPosition(fileName, position)
                    : mode === 'references' ? _getReferencesAtPosition(fileName, position)
                        : mode === 'fileReferences' ? _getFileReferences(fileName)
                            : mode === 'implementation' ? _getImplementationAtPosition(fileName, position)
                                : mode === 'rename' && preferences ? _findRenameLocations(fileName, position, findInStrings, findInComments, preferences)
                                    : undefined;
            if (!_symbols)
                return;
            symbols = symbols.concat(_symbols);
            for (const ref of _symbols) {
                loopChecker.add(ref.fileName + ':' + ref.textSpan.start);
                const [virtualFile] = getVirtualFile(ref.fileName);
                if (!virtualFile)
                    continue;
                const mirrorMap = virtualFiles.getMirrorMap(virtualFile);
                if (!mirrorMap)
                    continue;
                for (const [mirrorOffset, data] of mirrorMap.findMirrorOffsets(ref.textSpan.start)) {
                    if ((mode === 'definition' || mode === 'typeDefinition' || mode === 'implementation') && !data.definition)
                        continue;
                    if ((mode === 'references') && !data.references)
                        continue;
                    if ((mode === 'rename') && !data.rename)
                        continue;
                    if (loopChecker.has(ref.fileName + ':' + mirrorOffset))
                        continue;
                    withMirrors(ref.fileName, mirrorOffset);
                }
            }
        }
    }
    function getDefinitionAndBoundSpan(fileName, position) {
        const loopChecker = new Set();
        let textSpan;
        let symbols = [];
        withMirrors(fileName, position);
        if (!textSpan)
            return;
        return {
            textSpan: textSpan,
            definitions: symbols?.map(s => transformDocumentSpanLike(s, true)).filter(notEmpty),
        };
        function withMirrors(fileName, position) {
            if (loopChecker.has(fileName + ':' + position))
                return;
            loopChecker.add(fileName + ':' + position);
            const _symbols = _getDefinitionAndBoundSpan(fileName, position);
            if (!_symbols)
                return;
            if (!textSpan) {
                textSpan = _symbols.textSpan;
            }
            if (!_symbols.definitions)
                return;
            symbols = symbols.concat(_symbols.definitions);
            for (const ref of _symbols.definitions) {
                loopChecker.add(ref.fileName + ':' + ref.textSpan.start);
                const [virtualFile] = getVirtualFile(ref.fileName);
                if (!virtualFile)
                    continue;
                const mirrorMap = virtualFiles.getMirrorMap(virtualFile);
                if (!mirrorMap)
                    continue;
                for (const [mirrorOffset, data] of mirrorMap.findMirrorOffsets(ref.textSpan.start)) {
                    if (!data.definition)
                        continue;
                    if (loopChecker.has(ref.fileName + ':' + mirrorOffset))
                        continue;
                    withMirrors(ref.fileName, mirrorOffset);
                }
            }
        }
    }
    function findReferences(fileName, position) {
        const loopChecker = new Set();
        let symbols = [];
        withMirrors(fileName, position);
        return symbols.map(s => transformReferencedSymbol(s)).filter(notEmpty);
        function withMirrors(fileName, position) {
            if (loopChecker.has(fileName + ':' + position))
                return;
            loopChecker.add(fileName + ':' + position);
            const _symbols = _findReferences(fileName, position);
            if (!_symbols)
                return;
            symbols = symbols.concat(_symbols);
            for (const symbol of _symbols) {
                for (const ref of symbol.references) {
                    loopChecker.add(ref.fileName + ':' + ref.textSpan.start);
                    const [virtualFile] = getVirtualFile(ref.fileName);
                    if (!virtualFile)
                        continue;
                    const mirrorMap = virtualFiles.getMirrorMap(virtualFile);
                    if (!mirrorMap)
                        continue;
                    for (const [mirrorOffset, data] of mirrorMap.findMirrorOffsets(ref.textSpan.start)) {
                        if (!data.references)
                            continue;
                        if (loopChecker.has(ref.fileName + ':' + mirrorOffset))
                            continue;
                        withMirrors(ref.fileName, mirrorOffset);
                    }
                }
            }
        }
    }
    // transforms
    function transformFileTextChanges(changes) {
        const [_, source] = getVirtualFile(changes.fileName);
        if (source) {
            return {
                ...changes,
                fileName: source.fileName,
                textChanges: changes.textChanges.map(c => {
                    const span = transformSpan(changes.fileName, c.span);
                    if (span) {
                        return {
                            ...c,
                            span: span.textSpan,
                        };
                    }
                }).filter(notEmpty),
            };
        }
        else {
            return changes;
        }
    }
    function transformReferencedSymbol(symbol) {
        const definition = transformDocumentSpanLike(symbol.definition, false);
        const references = symbol.references.map(r => transformDocumentSpanLike(r, false)).filter(notEmpty);
        if (definition) {
            return {
                definition,
                references,
            };
        }
        else if (references.length) { // TODO: remove patching
            return {
                definition: {
                    ...symbol.definition,
                    fileName: references[0].fileName,
                    textSpan: references[0].textSpan,
                },
                references,
            };
        }
    }
    function transformDocumentSpanLike(documentSpan, isDefinition) {
        let textSpan = transformSpan(documentSpan.fileName, documentSpan.textSpan);
        if (isDefinition && !textSpan) {
            const [virtualFile, source] = getVirtualFile(documentSpan.fileName);
            if (virtualFile && source) {
                textSpan = {
                    fileName: source.fileName,
                    textSpan: { start: 0, length: 0 },
                };
            }
        }
        if (!textSpan)
            return;
        const contextSpan = transformSpan(documentSpan.fileName, documentSpan.contextSpan);
        const originalTextSpan = transformSpan(documentSpan.originalFileName, documentSpan.originalTextSpan);
        const originalContextSpan = transformSpan(documentSpan.originalFileName, documentSpan.originalContextSpan);
        return {
            ...documentSpan,
            fileName: textSpan.fileName,
            textSpan: textSpan.textSpan,
            contextSpan: contextSpan?.textSpan,
            originalFileName: originalTextSpan?.fileName,
            originalTextSpan: originalTextSpan?.textSpan,
            originalContextSpan: originalContextSpan?.textSpan,
        };
    }
    function transformSpan(fileName, textSpan) {
        if (!fileName)
            return;
        if (!textSpan)
            return;
        const [virtualFile, source] = getVirtualFile(fileName);
        if (virtualFile && source) {
            if (isTsPlugin) {
                textSpan = {
                    start: textSpan.start - source.snapshot.getLength(),
                    length: textSpan.length,
                };
            }
            for (const [_, [sourceSnapshot, map]] of virtualFiles.getMaps(virtualFile)) {
                if (source.snapshot !== sourceSnapshot)
                    continue;
                const sourceLoc = map.toSourceOffset(textSpan.start);
                if (sourceLoc) {
                    return {
                        fileName: source.fileName,
                        textSpan: {
                            start: sourceLoc[0],
                            length: textSpan.length,
                        },
                    };
                }
            }
        }
        else {
            return {
                fileName,
                textSpan,
            };
        }
    }
    function getVirtualFile(fileName) {
        if (isTsPlugin) {
            let result;
            const source = virtualFiles.getSource(fileName);
            if (source) {
                (0, language_core_1.forEachEmbeddedFile)(source.root, file => {
                    const ext = file.fileName.replace(fileName, '');
                    if (file.kind === language_core_1.FileKind.TypeScriptHostFile && (ext === '.d.ts' || ext.match(/^\.(js|ts)x?$/))) {
                        result = file;
                    }
                });
            }
            return [result, source];
        }
        else {
            return virtualFiles.getVirtualFile(fileName);
        }
    }
}
exports.decorateLanguageService = decorateLanguageService;
function notEmpty(value) {
    return value !== null && value !== undefined;
}
//# sourceMappingURL=languageService.js.map