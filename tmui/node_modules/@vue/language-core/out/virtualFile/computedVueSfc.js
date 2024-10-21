"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.computedVueSfc = void 0;
const computeds_1 = require("computeds");
function computedVueSfc(plugins, fileName, snapshot) {
    let cache;
    return (0, computeds_1.computed)(() => {
        // incremental update
        if (cache?.plugin.updateSFC) {
            const change = snapshot().getChangeRange(cache.snapshot);
            if (change) {
                const newSfc = cache.plugin.updateSFC(cache.sfc, {
                    start: change.span.start,
                    end: change.span.start + change.span.length,
                    newText: snapshot().getText(change.span.start, change.span.start + change.newLength),
                });
                if (newSfc) {
                    cache.snapshot = snapshot();
                    // force dirty
                    cache.sfc = JSON.parse(JSON.stringify(newSfc));
                    return cache.sfc;
                }
            }
        }
        for (const plugin of plugins) {
            const sfc = plugin.parseSFC?.(fileName, snapshot().getText(0, snapshot().getLength()));
            if (sfc) {
                if (!sfc.errors.length) {
                    cache = {
                        snapshot: snapshot(),
                        sfc,
                        plugin,
                    };
                }
                return sfc;
            }
        }
    });
}
exports.computedVueSfc = computedVueSfc;
//# sourceMappingURL=computedVueSfc.js.map