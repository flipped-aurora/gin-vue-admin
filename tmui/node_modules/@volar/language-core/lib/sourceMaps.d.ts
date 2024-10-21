import * as SourceMaps from '@volar/source-map';
import { MirrorBehaviorCapabilities } from './types';
export declare class MirrorMap extends SourceMaps.SourceMap<[MirrorBehaviorCapabilities, MirrorBehaviorCapabilities]> {
    findMirrorOffsets(start: number): Generator<readonly [number, MirrorBehaviorCapabilities], void, unknown>;
}
//# sourceMappingURL=sourceMaps.d.ts.map