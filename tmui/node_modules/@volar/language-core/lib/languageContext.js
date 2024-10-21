"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createLanguageContext = void 0;
const virtualFiles_1 = require("./virtualFiles");
function createLanguageContext(rawHost, languages) {
    let host = rawHost;
    let lastRootFiles = new Map();
    let lastProjectVersion;
    const virtualFiles = (0, virtualFiles_1.createVirtualFiles)(languages);
    for (const language of languages.reverse()) {
        if (language.resolveHost) {
            const pastHost = host;
            let proxyHost = language.resolveHost(host);
            if (proxyHost === pastHost) {
                console.warn(`[volar] language.resolveHost() should not return the same host instance.`);
                proxyHost = { ...proxyHost };
            }
            host = new Proxy(proxyHost, {
                get(target, p) {
                    if (p in target) {
                        return target[p];
                    }
                    return pastHost[p];
                }
            });
        }
    }
    return {
        rawHost,
        host,
        virtualFiles: new Proxy(virtualFiles, {
            get: (target, property) => {
                syncVirtualFiles();
                return target[property];
            },
        }),
    };
    function syncVirtualFiles() {
        const newProjectVersion = host.getProjectVersion();
        const shouldUpdate = newProjectVersion !== lastProjectVersion;
        if (!shouldUpdate)
            return;
        const nowRootFiles = new Map();
        const remainRootFiles = new Set(lastRootFiles.keys());
        for (const rootFileName of host.getScriptFileNames()) {
            nowRootFiles.set(rootFileName, host.getScriptSnapshot(rootFileName));
        }
        for (const [fileName, snapshot] of nowRootFiles) {
            remainRootFiles.delete(fileName);
            if (lastRootFiles.get(fileName) !== nowRootFiles.get(fileName)) {
                if (snapshot) {
                    virtualFiles.updateSource(fileName, snapshot, host.getLanguageId?.(fileName));
                }
                else {
                    virtualFiles.deleteSource(fileName);
                }
            }
        }
        for (const fileName of remainRootFiles) {
            virtualFiles.deleteSource(fileName);
        }
        lastRootFiles = nowRootFiles;
        lastProjectVersion = newProjectVersion;
    }
}
exports.createLanguageContext = createLanguageContext;
//# sourceMappingURL=languageContext.js.map