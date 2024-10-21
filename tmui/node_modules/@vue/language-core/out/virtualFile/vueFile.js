"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.VueFile = void 0;
const language_core_1 = require("@volar/language-core");
const computedFiles_1 = require("./computedFiles");
const computedMappings_1 = require("./computedMappings");
const computedSfc_1 = require("./computedSfc");
const computedVueSfc_1 = require("./computedVueSfc");
const computeds_1 = require("computeds");
const jsxReg = /^\.(js|ts)x?$/;
class VueFile {
    get embeddedFiles() {
        return this.getEmbeddedFiles();
    }
    get mainScriptName() {
        let res = '';
        (0, language_core_1.forEachEmbeddedFile)(this, file => {
            if (file.kind === language_core_1.FileKind.TypeScriptHostFile && file.fileName.replace(this.fileName, '').match(jsxReg)) {
                res = file.fileName;
            }
        });
        return res;
    }
    get snapshot() {
        return this._snapshot();
    }
    get mappings() {
        return this.getMappings();
    }
    constructor(fileName, initSnapshot, vueCompilerOptions, plugins, ts, codegenStack) {
        this.fileName = fileName;
        this.initSnapshot = initSnapshot;
        this.vueCompilerOptions = vueCompilerOptions;
        this.plugins = plugins;
        this.ts = ts;
        this.codegenStack = codegenStack;
        // computeds
        this.getVueSfc = (0, computedVueSfc_1.computedVueSfc)(this.plugins, this.fileName, () => this._snapshot());
        this.sfc = (0, computedSfc_1.computedSfc)(this.ts, this.plugins, this.fileName, () => this._snapshot(), this.getVueSfc);
        this.getMappings = (0, computedMappings_1.computedMappings)(() => this._snapshot(), this.sfc);
        this.getEmbeddedFiles = (0, computedFiles_1.computedFiles)(this.plugins, this.fileName, this.sfc, this.codegenStack);
        // others
        this.capabilities = language_core_1.FileCapabilities.full;
        this.kind = language_core_1.FileKind.TextFile;
        this.codegenStacks = [];
        this._snapshot = (0, computeds_1.signal)(initSnapshot);
    }
    update(newSnapshot) {
        this._snapshot.set(newSnapshot);
    }
}
exports.VueFile = VueFile;
//# sourceMappingURL=vueFile.js.map