"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const language_core_1 = require("@volar/language-core");
const templateReg = /^(.*)\.template\.([^.]+)$/;
const plugin = () => {
    return {
        version: 1,
        getEmbeddedFileNames(fileName, sfc) {
            if (sfc.template) {
                return [fileName + '.template.' + sfc.template.lang];
            }
            return [];
        },
        resolveEmbeddedFile(_fileName, sfc, embeddedFile) {
            const match = embeddedFile.fileName.match(templateReg);
            if (match && sfc.template) {
                embeddedFile.capabilities = language_core_1.FileCapabilities.full;
                embeddedFile.content.push([
                    sfc.template.content,
                    sfc.template.name,
                    0,
                    language_core_1.FileRangeCapabilities.full,
                ]);
            }
        },
    };
};
exports.default = plugin;
//# sourceMappingURL=vue-sfc-template.js.map