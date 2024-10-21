"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.computedMappings = void 0;
const language_core_1 = require("@volar/language-core");
const muggle = require("muggle-string");
const computeds_1 = require("computeds");
function computedMappings(snapshot, sfc) {
    return (0, computeds_1.computed)(() => {
        const str = [[snapshot().getText(0, snapshot().getLength()), undefined, 0, language_core_1.FileRangeCapabilities.full]];
        for (const block of [
            sfc.script,
            sfc.scriptSetup,
            sfc.template,
            ...sfc.styles,
            ...sfc.customBlocks,
        ]) {
            if (block) {
                muggle.replaceSourceRange(str, undefined, block.startTagEnd, block.endTagStart, [
                    block.content,
                    undefined,
                    block.startTagEnd,
                    {},
                ]);
            }
        }
        return str.map((m) => {
            const text = m[0];
            const start = m[2];
            const end = start + text.length;
            return {
                sourceRange: [start, end],
                generatedRange: [start, end],
                data: m[3],
            };
        });
    });
}
exports.computedMappings = computedMappings;
//# sourceMappingURL=computedMappings.js.map