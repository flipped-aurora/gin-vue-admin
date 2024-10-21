"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const language_core_1 = require("@volar/language-core");
const styleReg = /^(.*)\.style_(\d+)\.([^.]+)$/;
const plugin = () => {
    return {
        version: 1,
        getEmbeddedFileNames(fileName, sfc) {
            const names = [];
            for (let i = 0; i < sfc.styles.length; i++) {
                const style = sfc.styles[i];
                names.push(fileName + '.style_' + i + '.' + style.lang);
            }
            return names;
        },
        resolveEmbeddedFile(_fileName, sfc, embeddedFile) {
            const match = embeddedFile.fileName.match(styleReg);
            if (match) {
                const index = parseInt(match[2]);
                const style = sfc.styles[index];
                embeddedFile.capabilities = language_core_1.FileCapabilities.full;
                embeddedFile.content.push([
                    style.content,
                    style.name,
                    0,
                    language_core_1.FileRangeCapabilities.full,
                ]);
            }
        },
    };
};
exports.default = plugin;
//# sourceMappingURL=vue-sfc-styles.js.map