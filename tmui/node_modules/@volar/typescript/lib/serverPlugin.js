"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getExternalFiles = exports.searchExternalFiles = exports.decorateLanguageServiceHost = void 0;
const language_core_1 = require("@volar/language-core");
function decorateLanguageServiceHost(virtualFiles, languageServiceHost, ts, exts) {
    let extraProjectVersion = 0;
    const scripts = new Map();
    const readDirectory = languageServiceHost.readDirectory?.bind(languageServiceHost);
    const resolveModuleNameLiterals = languageServiceHost.resolveModuleNameLiterals?.bind(languageServiceHost);
    const resolveModuleNames = languageServiceHost.resolveModuleNames?.bind(languageServiceHost);
    const getProjectVersion = languageServiceHost.getProjectVersion?.bind(languageServiceHost);
    const getScriptSnapshot = languageServiceHost.getScriptSnapshot.bind(languageServiceHost);
    const getScriptKind = languageServiceHost.getScriptKind?.bind(languageServiceHost);
    // path completion
    if (readDirectory) {
        languageServiceHost.readDirectory = (path, extensions, exclude, include, depth) => {
            if (extensions) {
                for (const ext of exts) {
                    if (!extensions.includes(ext)) {
                        extensions = [...extensions, ...ext];
                    }
                }
            }
            return readDirectory(path, extensions, exclude, include, depth);
        };
    }
    if (resolveModuleNameLiterals) {
        languageServiceHost.resolveModuleNameLiterals = (moduleNames, containingFile, redirectedReference, options, ...rest) => {
            const resolvedModules = resolveModuleNameLiterals(moduleNames, containingFile, redirectedReference, options, ...rest);
            return moduleNames.map((name, i) => {
                if (exts.some(ext => name.text.endsWith(ext))) {
                    const resolved = resolveModuleName(name.text, containingFile, options, redirectedReference);
                    if (resolved.resolvedModule) {
                        return resolved;
                    }
                }
                return resolvedModules[i];
            });
        };
    }
    else if (resolveModuleNames) {
        languageServiceHost.resolveModuleNames = (moduleNames, containingFile, reusedNames, redirectedReference, options, containingSourceFile) => {
            const resolvedModules = resolveModuleNames(moduleNames, containingFile, reusedNames, redirectedReference, options, containingSourceFile);
            return moduleNames.map((name, i) => {
                if (exts.some(ext => name.endsWith(ext))) {
                    const resolved = resolveModuleName(name, containingFile, options, redirectedReference);
                    if (resolved.resolvedModule) {
                        return resolved.resolvedModule;
                    }
                }
                return resolvedModules[i];
            });
        };
    }
    if (getProjectVersion) {
        languageServiceHost.getProjectVersion = () => {
            return getProjectVersion() + ':' + extraProjectVersion;
        };
    }
    languageServiceHost.getScriptSnapshot = (fileName) => {
        if (exts.some(ext => fileName.endsWith(ext))) {
            updateScript(fileName);
            return scripts.get(fileName)?.snapshot;
        }
        return getScriptSnapshot(fileName);
    };
    if (getScriptKind) {
        languageServiceHost.getScriptKind = (fileName) => {
            if (exts.some(ext => fileName.endsWith(ext))) {
                updateScript(fileName);
                const script = scripts.get(fileName);
                if (script) {
                    if (script.extension.endsWith('.js')) {
                        return ts.ScriptKind.JS;
                    }
                    if (script.extension.endsWith('.jsx')) {
                        return ts.ScriptKind.JSX;
                    }
                    if (script.extension.endsWith('.ts')) {
                        return ts.ScriptKind.TS;
                    }
                    if (script.extension.endsWith('.tsx')) {
                        return ts.ScriptKind.TSX;
                    }
                }
                return ts.ScriptKind.Deferred;
            }
            return getScriptKind(fileName);
        };
    }
    function resolveModuleName(name, containingFile, options, redirectedReference) {
        const resolved = ts.resolveModuleName(name, containingFile, options, {
            readFile(fileName) {
                return languageServiceHost.readFile(fileName);
            },
            fileExists(fileName) {
                if (exts.some(ext => fileName.endsWith(ext + '.d.ts'))) {
                    return fileExists(fileName.slice(0, -'.d.ts'.length));
                }
                return languageServiceHost.fileExists(fileName);
            },
        }, undefined, redirectedReference);
        if (resolved.resolvedModule) {
            resolved.resolvedModule.resolvedFileName = resolved.resolvedModule.resolvedFileName.slice(0, -'.d.ts'.length);
            const script = updateScript(resolved.resolvedModule.resolvedFileName);
            if (script) {
                resolved.resolvedModule.extension = script.extension;
            }
        }
        return resolved;
    }
    // fix https://github.com/vuejs/language-tools/issues/3332
    function fileExists(fileName) {
        if (languageServiceHost.fileExists(fileName)) {
            const fileSize = ts.sys.getFileSize?.(fileName) ?? languageServiceHost.readFile(fileName)?.length ?? 0;
            return fileSize < 4 * 1024 * 1024;
        }
        return false;
    }
    function updateScript(fileName) {
        const version = languageServiceHost.getScriptVersion(fileName);
        if (version !== scripts.get(fileName)?.version) {
            const text = languageServiceHost.readFile(fileName);
            let snapshot;
            let extension = '.ts';
            if (text !== undefined) {
                extraProjectVersion++;
                const virtualFile = virtualFiles.updateSource(fileName, ts.ScriptSnapshot.fromString(text), undefined);
                if (virtualFile) {
                    let patchedText = text.split('\n').map(line => ' '.repeat(line.length)).join('\n');
                    (0, language_core_1.forEachEmbeddedFile)(virtualFile, file => {
                        const ext = file.fileName.substring(fileName.length);
                        if (file.kind === language_core_1.FileKind.TypeScriptHostFile && (ext === '.d.ts' || ext.match(/^\.(js|ts)x?$/))) {
                            extension = ext;
                            patchedText += file.snapshot.getText(0, file.snapshot.getLength());
                        }
                    });
                    snapshot = ts.ScriptSnapshot.fromString(patchedText);
                }
            }
            else if (virtualFiles.hasSource(fileName)) {
                extraProjectVersion++;
                virtualFiles.deleteSource(fileName);
            }
            scripts.set(fileName, {
                version,
                snapshot,
                extension,
            });
        }
        return scripts.get(fileName);
    }
}
exports.decorateLanguageServiceHost = decorateLanguageServiceHost;
function searchExternalFiles(ts, project, exts) {
    if (project.projectKind !== ts.server.ProjectKind.Configured) {
        return [];
    }
    const configFile = project.getProjectName();
    const config = ts.readJsonConfigFile(configFile, project.readFile.bind(project));
    const parseHost = {
        useCaseSensitiveFileNames: project.useCaseSensitiveFileNames(),
        fileExists: project.fileExists.bind(project),
        readFile: project.readFile.bind(project),
        readDirectory: (...args) => {
            args[1] = exts;
            return project.readDirectory(...args);
        },
    };
    const parsed = ts.parseJsonSourceFileConfigFileContent(config, parseHost, project.getCurrentDirectory());
    return parsed.fileNames;
}
exports.searchExternalFiles = searchExternalFiles;
/**
 * @deprecated use `searchExternalFiles` instead
 */
exports.getExternalFiles = searchExternalFiles;
//# sourceMappingURL=serverPlugin.js.map