const defaults = require('./defaults');
const md5 = require('./md5');
const safeGet = require('./safeGet');
const FileStore = require('./FileStore');
const FileBlobStore = require('./FileBlobStore');
const mkdir = require('./mkdir');

let isEnabled = false;

exports = function(options = {}) {
    if (isEnabled) {
        return;
    }
    isEnabled = true;
    defaults(options, defOptions);

    const Module = require('module');
    const fs = require('fs');
    const path = require('path');
    const vm = require('vm');

    const cacheDir = options.dir || getCacheDir();
    if (!fs.existsSync(cacheDir)) {
        mkdir.sync(cacheDir);
    }

    if (options.requirePath) {
        const _resolveFilename = Module._resolveFilename;
        const requirePathMapStore = new FileStore(
            path.resolve(cacheDir, 'require-path.json')
        );

        Module._resolveFilename = function(request, parent, isMain, options) {
            let currentModuleCache = requirePathMapStore.get(parent.filename);
            if (!currentModuleCache) {
                currentModuleCache = {};
                requirePathMapStore.set(parent.filename, currentModuleCache);
            }
            if (currentModuleCache[request]) {
                return currentModuleCache[request];
            }

            const pathToLoad = _resolveFilename(
                request,
                parent,
                isMain,
                options
            );
            currentModuleCache[request] = pathToLoad;

            return pathToLoad;
        };
    }

    if (options.code) {
        const fileMapStore = new FileStore(
            path.resolve(cacheDir, 'require-file.json')
        );

        Module._extensions['.js'] = function(module, filename) {
            let content;
            if (fileMapStore.get(filename)) {
                content = fileMapStore.get(filename);
            } else {
                content = fs.readFileSync(filename, 'utf8');
                fileMapStore.set(filename, content);
            }
            return module._compile(content, filename);
        };
    }

    if (options.compileCache) {
        const compileCacheBlobStore = new FileBlobStore(
            path.resolve(cacheDir, 'compile-cache')
        );
        process.once('exit', () => compileCacheBlobStore.save());
        Module.prototype._compile = function(content, filename) {
            const mod = this;

            function require(id) {
                return mod.require(id);
            }

            function resolve(request, options) {
                return Module._resolveFilename(request, mod, false, options);
            }
            require.resolve = resolve;

            resolve.paths = function paths(request) {
                return Module._resolveLookupPaths(request, mod, true);
            };

            require.main = process.mainModule;

            require.extensions = Module._extensions;
            require.cache = Module._cache;

            const dirname = path.dirname(filename);

            const compiledWrapper = moduleCompile(filename, content);

            const args = [
                mod.exports,
                require,
                mod,
                filename,
                dirname,
                process,
                global,
                Buffer
            ];
            return compiledWrapper.apply(mod.exports, args);
        };

        function moduleCompile(filename, content) {
            const wrapper = Module.wrap(content);

            const buffer = compileCacheBlobStore.get(filename);

            const script = new vm.Script(wrapper, {
                filename,
                lineOffset: 0,
                displayErrors: true,
                cachedData: buffer,
                produceCachedData: true
            });

            if (script.cachedDataProduced) {
                compileCacheBlobStore.set(filename, script.cachedData);
            } else if (script.cachedDataRejected) {
                compileCacheBlobStore.delete(filename);
            }

            const compiledWrapper = script.runInThisContext({
                filename,
                lineOffset: 0,
                columnOffset: 0,
                displayErrors: true
            });

            return compiledWrapper;
        }
    }
};

const defOptions = {
    requirePath: true,
    code: false,
    compileCache: true
};

function getCacheDir() {
    const os = require('os');

    return (
        os.tmpdir() +
        '/' +
        md5(
            `${get(process, 'versions.v8')}${get(require, 'main.filename') ||
                get(module, 'parent.filename')}`
        )
    );
}

function get(obj, path) {
    return safeGet(obj, path) || '';
}

module.exports = exports;
