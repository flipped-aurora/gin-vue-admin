"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const language_core_1 = require("@volar/language-core");
const scriptFormatReg = /^(.*)\.script_format\.([^.]+)$/;
const scriptSetupFormatReg = /^(.*)\.scriptSetup_format\.([^.]+)$/;
const plugin = () => {
    return {
        version: 1,
        getEmbeddedFileNames(fileName, sfc) {
            const names = [];
            if (sfc.script) {
                names.push(fileName + '.script_format.' + sfc.script.lang);
            }
            if (sfc.scriptSetup) {
                names.push(fileName + '.scriptSetup_format.' + sfc.scriptSetup.lang);
            }
            return names;
        },
        resolveEmbeddedFile(_fileName, sfc, embeddedFile) {
            const scriptMatch = embeddedFile.fileName.match(scriptFormatReg);
            const scriptSetupMatch = embeddedFile.fileName.match(scriptSetupFormatReg);
            const script = scriptMatch ? sfc.script : scriptSetupMatch ? sfc.scriptSetup : undefined;
            if (script) {
                embeddedFile.kind = language_core_1.FileKind.TextFile;
                embeddedFile.capabilities = {
                    ...language_core_1.FileCapabilities.full,
                    diagnostic: false,
                    codeAction: false,
                    inlayHint: false,
                };
                embeddedFile.content.push([
                    script.content,
                    script.name,
                    0,
                    {},
                ]);
            }
        },
    };
};
exports.default = plugin;
//# sourceMappingURL=vue-sfc-scripts.js.map