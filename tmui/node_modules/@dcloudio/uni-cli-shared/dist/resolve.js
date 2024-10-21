"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.resolveComponentsLibPath = exports.resolveVueI18nRuntime = exports.resolveBuiltIn = exports.getBuiltInPaths = exports.resolveMainPathOnce = exports.relativeFile = exports.requireResolve = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const debug_1 = __importDefault(require("debug"));
const resolve_1 = __importDefault(require("resolve"));
const uni_shared_1 = require("@dcloudio/uni-shared");
const utils_1 = require("./utils");
const env_1 = require("./hbx/env");
const constants_1 = require("./constants");
function requireResolve(filename, basedir) {
    return resolveWithSymlinks(filename, basedir);
}
exports.requireResolve = requireResolve;
function resolveWithSymlinks(id, basedir) {
    return resolve_1.default.sync(id, {
        basedir,
        extensions: process.env.UNI_APP_X === 'true' ? constants_1.uni_app_x_extensions : constants_1.extensions,
        // necessary to work with pnpm
        preserveSymlinks: true,
    });
}
function relativeFile(from, to) {
    const relativePath = (0, utils_1.normalizePath)(path_1.default.relative(path_1.default.dirname(from), to));
    return relativePath.startsWith('.') ? relativePath : './' + relativePath;
}
exports.relativeFile = relativeFile;
exports.resolveMainPathOnce = (0, uni_shared_1.once)((inputDir) => {
    if (process.env.UNI_APP_X === 'true') {
        const mainUTSPath = path_1.default.resolve(inputDir, 'main.uts');
        if (fs_1.default.existsSync(mainUTSPath)) {
            return (0, utils_1.normalizePath)(mainUTSPath);
        }
    }
    const mainTsPath = path_1.default.resolve(inputDir, 'main.ts');
    if (fs_1.default.existsSync(mainTsPath)) {
        return (0, utils_1.normalizePath)(mainTsPath);
    }
    return (0, utils_1.normalizePath)(path_1.default.resolve(inputDir, 'main.js'));
});
const ownerModules = [
    '@dcloudio/uni-app',
    '@dcloudio/vite-plugin-uni',
    '@dcloudio/uni-cli-shared',
];
const paths = [];
function resolveNodeModulePath(modulePath) {
    const nodeModulesPaths = [];
    const nodeModulesPath = path_1.default.join(modulePath, 'node_modules');
    if (fs_1.default.existsSync(nodeModulesPath)) {
        nodeModulesPaths.push(nodeModulesPath);
    }
    const index = modulePath.lastIndexOf('node_modules');
    if (index > -1) {
        nodeModulesPaths.push(path_1.default.join(modulePath.slice(0, index), 'node_modules'));
    }
    return nodeModulesPaths;
}
function initPaths() {
    const cliContext = process.env.UNI_CLI_CONTEXT || process.cwd();
    if (cliContext) {
        const pathSet = new Set();
        pathSet.add(path_1.default.join(cliContext, 'node_modules'));
        if (!(0, env_1.isInHBuilderX)()) {
            ;
            [`@dcloudio/uni-` + process.env.UNI_PLATFORM, ...ownerModules].forEach((ownerModule) => {
                let pkgPath = '';
                try {
                    pkgPath = require.resolve(ownerModule + '/package.json', {
                        paths: [cliContext],
                    });
                }
                catch (e) { }
                if (pkgPath) {
                    resolveNodeModulePath(path_1.default.dirname(pkgPath)).forEach((nodeModulePath) => {
                        pathSet.add(nodeModulePath);
                    });
                }
            });
        }
        paths.push(...pathSet);
        (0, debug_1.default)('uni-paths')(paths);
    }
}
function getBuiltInPaths() {
    if (!paths.length) {
        initPaths();
    }
    return paths;
}
exports.getBuiltInPaths = getBuiltInPaths;
function resolveBuiltIn(path) {
    return require.resolve(path, { paths: getBuiltInPaths() });
}
exports.resolveBuiltIn = resolveBuiltIn;
function resolveVueI18nRuntime() {
    return path_1.default.resolve(__dirname, '../lib/vue-i18n/dist/vue-i18n.runtime.esm-bundler.js');
}
exports.resolveVueI18nRuntime = resolveVueI18nRuntime;
let componentsLibPath = '';
function resolveComponentsLibPath() {
    if (!componentsLibPath) {
        const dir = process.env.UNI_APP_X === 'true' ? '../lib-x' : '../lib';
        if ((0, env_1.isInHBuilderX)()) {
            componentsLibPath = path_1.default.join(resolveBuiltIn('@dcloudio/uni-components/package.json'), dir);
        }
        else {
            try {
                componentsLibPath = path_1.default.join(resolveWithSymlinks('@dcloudio/uni-components/package.json', process.env.UNI_INPUT_DIR), dir);
            }
            catch (e) {
                try {
                    componentsLibPath = path_1.default.join(resolveWithSymlinks('@dcloudio/uni-components/package.json', process.cwd()), dir);
                }
                catch (e) {
                    console.log(e);
                }
            }
        }
    }
    return componentsLibPath;
}
exports.resolveComponentsLibPath = resolveComponentsLibPath;
