"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.MirrorMap = void 0;
const SourceMaps = require("@volar/source-map");
class MirrorMap extends SourceMaps.SourceMap {
    *findMirrorOffsets(start) {
        for (const mapped of this.toGeneratedOffsets(start)) {
            yield [mapped[0], mapped[1].data[1]];
        }
        for (const mapped of this.toSourceOffsets(start)) {
            yield [mapped[0], mapped[1].data[0]];
        }
    }
}
exports.MirrorMap = MirrorMap;
//# sourceMappingURL=sourceMaps.js.map