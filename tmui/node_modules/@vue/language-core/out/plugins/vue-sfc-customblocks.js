"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const language_core_1 = require("@volar/language-core");
const customBlockReg = /^(.*)\.customBlock_([^_]+)_(\d+)\.([^.]+)$/;
const plugin = () => {
    return {
        version: 1,
        getEmbeddedFileNames(fileName, sfc) {
            const names = [];
            for (let i = 0; i < sfc.customBlocks.length; i++) {
                const customBlock = sfc.customBlocks[i];
                names.push(fileName + '.customBlock_' + customBlock.type + '_' + i + '.' + customBlock.lang);
            }
            return names;
        },
        resolveEmbeddedFile(_fileName, sfc, embeddedFile) {
            const match = embeddedFile.fileName.match(customBlockReg);
            if (match) {
                const index = parseInt(match[3]);
                const customBlock = sfc.customBlocks[index];
                embeddedFile.capabilities = language_core_1.FileCapabilities.full;
                embeddedFile.content.push([
                    customBlock.content,
                    customBlock.name,
                    0,
                    language_core_1.FileRangeCapabilities.full,
                ]);
            }
        },
    };
};
exports.default = plugin;
//# sourceMappingURL=vue-sfc-customblocks.js.map