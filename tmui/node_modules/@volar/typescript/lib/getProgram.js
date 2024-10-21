"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getProgram = void 0;
function getProgram(ts, core, ls, sys) {
    const proxy = {
        getRootFileNames,
        emit,
        getSyntacticDiagnostics,
        getSemanticDiagnostics,
        getGlobalDiagnostics,
        // @ts-expect-error
        getBindAndCheckDiagnostics,
    };
    return new Proxy({}, {
        get: (target, property) => {
            if (property in proxy) {
                return proxy[property];
            }
            const program = getProgram();
            if (property in program) {
                return program[property];
            }
            return target[property];
        },
        // #17
        // notice: https://github.com/vuejs/language-tools/issues/2403
        set: (target, property, newValue) => {
            const program = getProgram();
            target[property] = program[property] = newValue;
            return true;
        },
    });
    function getProgram() {
        return ls.getProgram();
    }
    function getRootFileNames() {
        return getProgram().getRootFileNames().filter(fileName => sys.fileExists?.(fileName));
    }
    // for vue-tsc --noEmit --watch
    function getBindAndCheckDiagnostics(sourceFile, cancellationToken) {
        return getSourceFileDiagnosticsWorker(sourceFile, cancellationToken, 'getBindAndCheckDiagnostics');
    }
    // for vue-tsc --noEmit
    function getSyntacticDiagnostics(sourceFile, cancellationToken) {
        return getSourceFileDiagnosticsWorker(sourceFile, cancellationToken, 'getSyntacticDiagnostics');
    }
    function getSemanticDiagnostics(sourceFile, cancellationToken) {
        return getSourceFileDiagnosticsWorker(sourceFile, cancellationToken, 'getSemanticDiagnostics');
    }
    function getSourceFileDiagnosticsWorker(sourceFile, cancellationToken, api) {
        if (sourceFile) {
            const [virtualFile, source] = core.virtualFiles.getVirtualFile(sourceFile.fileName);
            if (virtualFile && source) {
                if (!virtualFile.capabilities.diagnostic)
                    return [];
                const errors = transformDiagnostics(ls.getProgram()?.[api](sourceFile, cancellationToken) ?? []);
                return errors;
            }
        }
        return transformDiagnostics(getProgram()[api](sourceFile, cancellationToken) ?? []);
    }
    function getGlobalDiagnostics(cancellationToken) {
        return transformDiagnostics(getProgram().getGlobalDiagnostics(cancellationToken) ?? []);
    }
    function emit(targetSourceFile, _writeFile, cancellationToken, emitOnlyDtsFiles, customTransformers) {
        const scriptResult = getProgram().emit(targetSourceFile, (sys.writeFile ?? ts.sys.writeFile), cancellationToken, emitOnlyDtsFiles, customTransformers);
        return {
            emitSkipped: scriptResult.emitSkipped,
            emittedFiles: scriptResult.emittedFiles,
            diagnostics: transformDiagnostics(scriptResult.diagnostics),
        };
    }
    // transform
    function transformDiagnostics(diagnostics) {
        const result = [];
        for (const diagnostic of diagnostics) {
            if (diagnostic.file !== undefined
                && diagnostic.start !== undefined
                && diagnostic.length !== undefined) {
                const [virtualFile, source] = core.virtualFiles.getVirtualFile(diagnostic.file.fileName);
                if (virtualFile && source) {
                    if (sys.fileExists?.(source.fileName) === false)
                        continue;
                    for (const [_, [sourceSnapshot, map]] of core.virtualFiles.getMaps(virtualFile)) {
                        if (sourceSnapshot !== source.snapshot)
                            continue;
                        for (const start of map.toSourceOffsets(diagnostic.start)) {
                            const reportStart = typeof start[1].data.diagnostic === 'object' ? start[1].data.diagnostic.shouldReport() : !!start[1].data.diagnostic;
                            if (!reportStart)
                                continue;
                            for (const end of map.toSourceOffsets(diagnostic.start + diagnostic.length, true)) {
                                const reportEnd = typeof end[1].data.diagnostic === 'object' ? end[1].data.diagnostic.shouldReport() : !!end[1].data.diagnostic;
                                if (!reportEnd)
                                    continue;
                                onMapping(diagnostic, source.fileName, start[0], end[0], source.snapshot.getText(0, source.snapshot.getLength()));
                                break;
                            }
                            break;
                        }
                    }
                }
                else {
                    if (sys.fileExists?.(diagnostic.file.fileName) === false)
                        continue;
                    onMapping(diagnostic, diagnostic.file.fileName, diagnostic.start, diagnostic.start + diagnostic.length, diagnostic.file.text);
                }
            }
            else if (diagnostic.file === undefined) {
                result.push(diagnostic);
            }
        }
        return result;
        function onMapping(diagnostic, fileName, start, end, docText) {
            let file = fileName === diagnostic.file?.fileName
                ? diagnostic.file
                : undefined;
            if (!file) {
                if (docText === undefined) {
                    const snapshot = core.host.getScriptSnapshot(fileName);
                    if (snapshot) {
                        docText = snapshot.getText(0, snapshot.getLength());
                    }
                }
                else {
                    file = ts.createSourceFile(fileName, docText, ts.ScriptTarget.Latest, undefined, ts.ScriptKind.Deferred);
                    // fix https://github.com/vuejs/language-tools/issues/2622 for TS 5.0
                    file.originalFileName = fileName;
                    file.path = fileName.toLowerCase();
                    file.resolvedPath = fileName.toLowerCase();
                }
            }
            const newDiagnostic = {
                ...diagnostic,
                file,
                start: start,
                length: end - start,
            };
            const relatedInformation = diagnostic.relatedInformation;
            if (relatedInformation) {
                newDiagnostic.relatedInformation = transformDiagnostics(relatedInformation);
            }
            result.push(newDiagnostic);
        }
    }
}
exports.getProgram = getProgram;
//# sourceMappingURL=getProgram.js.map