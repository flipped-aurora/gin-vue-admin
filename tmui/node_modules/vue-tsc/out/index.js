"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createProgram = void 0;
const vue = require("@vue/language-core");
const volarTs = require("@volar/typescript");
const shared_1 = require("./shared");
const windowsPathReg = /\\/g;
function createProgram(options) {
    if (!options.options.noEmit && !options.options.emitDeclarationOnly)
        throw toThrow('js emit is not supported');
    if (!options.options.noEmit && options.options.noEmitOnError)
        throw toThrow('noEmitOnError is not supported');
    if (options.options.extendedDiagnostics || options.options.generateTrace)
        throw toThrow('--extendedDiagnostics / --generateTrace is not supported, please run `Write Virtual Files` in VSCode to write virtual files and use `--extendedDiagnostics` / `--generateTrace` via tsc instead of vue-tsc to debug.');
    if (!options.host)
        throw toThrow('!options.host');
    const ts = require('typescript');
    let program = options.oldProgram;
    if (shared_1.state.hook) {
        program = shared_1.state.hook.program;
        program.__vue.options = options;
    }
    else if (!program) {
        const ctx = {
            projectVersion: 0,
            options,
            get languageHost() {
                return languageHost;
            },
            get vueCompilerOptions() {
                return vueCompilerOptions;
            },
            get languageService() {
                return vueTsLs;
            },
            get langaugeContext() {
                return languageContext;
            },
        };
        const vueCompilerOptions = getVueCompilerOptions();
        const scripts = new Map();
        const languageHost = {
            workspacePath: ctx.options.host.getCurrentDirectory().replace(windowsPathReg, '/'),
            rootPath: ctx.options.host.getCurrentDirectory().replace(windowsPathReg, '/'),
            getCompilationSettings: () => ctx.options.options,
            getScriptFileNames: () => {
                return ctx.options.rootNames;
            },
            getScriptSnapshot,
            getProjectVersion: () => {
                return ctx.projectVersion.toString();
            },
            getProjectReferences: () => ctx.options.projectReferences,
            getCancellationToken: ctx.options.host.getCancellationToken ? () => ctx.options.host.getCancellationToken() : undefined,
        };
        const languageContext = vue.createLanguageContext(languageHost, vue.createLanguages(ts, languageHost.getCompilationSettings(), vueCompilerOptions));
        const languageServiceHost = volarTs.createLanguageServiceHost(languageContext, ts, ts.sys);
        const vueTsLs = ts.createLanguageService(languageServiceHost, volarTs.getDocumentRegistry(ts, ts.sys.useCaseSensitiveFileNames, languageHost.workspacePath));
        volarTs.decorateLanguageService(languageContext.virtualFiles, vueTsLs, false);
        program = volarTs.getProgram(ts, languageContext, vueTsLs, ts.sys);
        program.__vue = ctx;
        function getVueCompilerOptions() {
            const tsConfig = ctx.options.options.configFilePath;
            if (typeof tsConfig === 'string') {
                return vue.createParsedCommandLine(ts, ts.sys, tsConfig.replace(windowsPathReg, '/')).vueOptions;
            }
            return {};
        }
        function getScriptSnapshot(fileName) {
            return getScript(fileName)?.scriptSnapshot;
        }
        function getScript(fileName) {
            const script = scripts.get(fileName);
            if (script?.projectVersion === ctx.projectVersion) {
                return script;
            }
            const modifiedTime = ts.sys.getModifiedTime?.(fileName)?.valueOf() ?? 0;
            if (script?.modifiedTime === modifiedTime) {
                return script;
            }
            if (ctx.options.host.fileExists(fileName)) {
                const fileContent = ctx.options.host.readFile(fileName);
                if (fileContent !== undefined) {
                    const script = {
                        projectVersion: ctx.projectVersion,
                        modifiedTime,
                        scriptSnapshot: ts.ScriptSnapshot.fromString(fileContent),
                        version: ctx.options.host.createHash?.(fileContent) ?? fileContent,
                    };
                    scripts.set(fileName, script);
                    return script;
                }
            }
        }
    }
    else {
        const ctx = program.__vue;
        ctx.options = options;
        ctx.projectVersion++;
    }
    const vueCompilerOptions = program.__vue.vueCompilerOptions;
    if (vueCompilerOptions?.hooks) {
        const index = (shared_1.state.hook?.index ?? -1) + 1;
        if (index < vueCompilerOptions.hooks.length) {
            const hookPath = vueCompilerOptions.hooks[index];
            const hook = require(hookPath);
            shared_1.state.hook = {
                program,
                index,
                worker: (async () => await hook(program))(),
            };
            throw 'hook';
        }
    }
    for (const rootName of options.rootNames) {
        // register file watchers
        options.host.getSourceFile(rootName, ts.ScriptTarget.ESNext);
    }
    return program;
}
exports.createProgram = createProgram;
function toThrow(msg) {
    console.error(msg);
    return msg;
}
//# sourceMappingURL=index.js.map