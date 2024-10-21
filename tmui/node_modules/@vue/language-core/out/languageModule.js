"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createLanguages = exports.createVueLanguage = void 0;
const path = require("path-browserify");
const plugins_1 = require("./plugins");
const vueFile_1 = require("./virtualFile/vueFile");
const sharedTypes = require("./utils/globalTypes");
const ts_1 = require("./utils/ts");
const fileRegistries = [];
function getVueFileRegistry(key, plugins) {
    let fileRegistry = fileRegistries.find(r => r.key === key
        && r.plugins.length === plugins.length
        && r.plugins.every(plugin => plugins.includes(plugin)))?.files;
    if (!fileRegistry) {
        fileRegistry = new Map();
        fileRegistries.push({
            key: key,
            plugins: plugins,
            files: fileRegistry,
        });
    }
    return fileRegistry;
}
function createVueLanguage(ts, compilerOptions = {}, _vueCompilerOptions = {}, codegenStack = false) {
    const vueCompilerOptions = (0, ts_1.resolveVueCompilerOptions)(_vueCompilerOptions);
    const plugins = (0, plugins_1.getDefaultVueLanguagePlugins)(ts, compilerOptions, vueCompilerOptions, codegenStack);
    const keys = [
        ...Object.keys(vueCompilerOptions)
            .sort()
            .filter(key => key !== 'plugins')
            .map(key => [key, vueCompilerOptions[key]]),
        [...new Set(plugins.map(plugin => plugin.requiredCompilerOptions ?? []).flat())]
            .sort()
            .map(key => [key, compilerOptions[key]]),
    ];
    const fileRegistry = getVueFileRegistry(JSON.stringify(keys), _vueCompilerOptions.plugins ?? []);
    const allowLanguageIds = new Set(['vue']);
    if (vueCompilerOptions.extensions.includes('.md')) {
        allowLanguageIds.add('markdown');
    }
    if (vueCompilerOptions.extensions.includes('.html')) {
        allowLanguageIds.add('html');
    }
    return {
        createVirtualFile(fileName, snapshot, languageId) {
            if ((languageId && allowLanguageIds.has(languageId))
                || (!languageId && vueCompilerOptions.extensions.some(ext => fileName.endsWith(ext)))) {
                if (fileRegistry.has(fileName)) {
                    const reusedVueFile = fileRegistry.get(fileName);
                    reusedVueFile.update(snapshot);
                    return reusedVueFile;
                }
                const vueFile = new vueFile_1.VueFile(fileName, snapshot, vueCompilerOptions, plugins, ts, codegenStack);
                fileRegistry.set(fileName, vueFile);
                return vueFile;
            }
        },
        updateVirtualFile(sourceFile, snapshot) {
            sourceFile.update(snapshot);
        },
        resolveHost(host) {
            const sharedTypesSnapshot = ts.ScriptSnapshot.fromString(sharedTypes.getTypesCode(vueCompilerOptions));
            const sharedTypesFileName = path.join(host.rootPath, sharedTypes.baseName);
            return {
                ...host,
                resolveModuleName(moduleName, impliedNodeFormat) {
                    if (impliedNodeFormat === ts.ModuleKind.ESNext && vueCompilerOptions.extensions.some(ext => moduleName.endsWith(ext))) {
                        return `${moduleName}.js`;
                    }
                    return host.resolveModuleName?.(moduleName, impliedNodeFormat) ?? moduleName;
                },
                getScriptFileNames() {
                    return [
                        sharedTypesFileName,
                        ...host.getScriptFileNames(),
                    ];
                },
                getScriptSnapshot(fileName) {
                    if (fileName === sharedTypesFileName) {
                        return sharedTypesSnapshot;
                    }
                    return host.getScriptSnapshot(fileName);
                },
            };
        },
    };
}
exports.createVueLanguage = createVueLanguage;
/**
 * @deprecated planed to remove in 2.0, please use createVueLanguage instead of
 */
function createLanguages(ts, compilerOptions = {}, vueCompilerOptions = {}, codegenStack = false) {
    return [
        createVueLanguage(ts, compilerOptions, vueCompilerOptions, codegenStack),
        ...vueCompilerOptions.experimentalAdditionalLanguageModules?.map(module => require(module)) ?? [],
    ];
}
exports.createLanguages = createLanguages;
//# sourceMappingURL=languageModule.js.map