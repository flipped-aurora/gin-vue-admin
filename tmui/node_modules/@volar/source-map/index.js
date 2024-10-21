"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __exportStar = (this && this.__exportStar) || function(m, exports) {
    for (var p in m) if (p !== "default" && !Object.prototype.hasOwnProperty.call(exports, p)) __createBinding(exports, m, p);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildStacks = exports.buildMappings = exports.SourceMap = void 0;
__exportStar(require("muggle-string"), exports);
class SourceMap {
    get memo() {
        if (!this._memo) {
            const self = this;
            this._memo = {
                sourceRange: createMemo('sourceRange'),
                generatedRange: createMemo('generatedRange'),
            };
            function createMemo(key) {
                const offsets = new Set();
                for (const mapping of self.mappings) {
                    offsets.add(mapping[key][0]);
                    offsets.add(mapping[key][1]);
                }
                const arr = [...offsets].sort((a, b) => a - b).map(offset => ({ offset, mappings: new Set() }));
                for (const mapping of self.mappings) {
                    const startIndex = binarySearch(mapping[key][0]);
                    const endIndex = binarySearch(mapping[key][1]);
                    for (let i = startIndex; i <= endIndex; i++) {
                        arr[i].mappings.add(mapping);
                    }
                }
                return arr;
                function binarySearch(start) {
                    let low = 0;
                    let high = arr.length - 1;
                    while (low <= high) {
                        const mid = Math.floor((low + high) / 2);
                        const midValue = arr[mid];
                        if (midValue.offset < start) {
                            low = mid + 1;
                        }
                        else if (midValue.offset > start) {
                            high = mid - 1;
                        }
                        else {
                            return mid;
                        }
                    }
                }
            }
        }
        return this._memo;
    }
    constructor(mappings) {
        this.mappings = mappings;
    }
    toSourceOffset(start, baseOnRight = false) {
        for (const mapped of this.matching(start, 'generatedRange', 'sourceRange', baseOnRight)) {
            return mapped;
        }
    }
    toGeneratedOffset(start, baseOnRight = false) {
        for (const mapped of this.matching(start, 'sourceRange', 'generatedRange', baseOnRight)) {
            return mapped;
        }
    }
    toSourceOffsets(start, baseOnRight = false) {
        return this.matching(start, 'generatedRange', 'sourceRange', baseOnRight);
    }
    toGeneratedOffsets(start, baseOnRight = false) {
        return this.matching(start, 'sourceRange', 'generatedRange', baseOnRight);
    }
    *matching(startOffset, from, to, baseOnRight) {
        const memo = this.memo[from];
        if (memo.length === 0)
            return;
        const { low: start, high: end, } = this.binarySearchMemo(memo, startOffset);
        const skip = new Set();
        for (let i = start; i <= end; i++) {
            for (const mapping of memo[i].mappings) {
                if (skip.has(mapping)) {
                    continue;
                }
                skip.add(mapping);
                const mapped = this.matchOffset(startOffset, mapping[from], mapping[to], baseOnRight);
                if (mapped !== undefined) {
                    yield [mapped, mapping];
                }
            }
        }
    }
    matchOffset(start, mappedFromRange, mappedToRange, baseOnRight) {
        if (start >= mappedFromRange[0] && start <= mappedFromRange[1]) {
            let offset = mappedToRange[0] + start - mappedFromRange[0];
            if (baseOnRight) {
                offset += (mappedToRange[1] - mappedToRange[0]) - (mappedFromRange[1] - mappedFromRange[0]);
            }
            if (offset >= mappedToRange[0] && offset <= mappedToRange[1]) {
                return offset;
            }
        }
    }
    binarySearchMemo(array, start) {
        let low = 0;
        let high = array.length - 1;
        while (low <= high) {
            const mid = Math.floor((low + high) / 2);
            const midValue = array[mid];
            if (midValue.offset < start) {
                low = mid + 1;
            }
            else if (midValue.offset > start) {
                high = mid - 1;
            }
            else {
                low = mid;
                high = mid;
                break;
            }
        }
        return {
            low: Math.max(Math.min(low, high, array.length - 1), 0),
            high: Math.min(Math.max(low, high, 0), array.length - 1),
        };
    }
}
exports.SourceMap = SourceMap;
function buildMappings(chunks) {
    let length = 0;
    const mappings = [];
    for (const segment of chunks) {
        if (typeof segment === 'string') {
            length += segment.length;
        }
        else {
            mappings.push({
                generatedRange: [length, length + segment[0].length],
                source: segment[1],
                sourceRange: typeof segment[2] === 'number' ? [segment[2], segment[2] + segment[0].length] : segment[2],
                // @ts-ignore
                data: segment[3],
            });
            length += segment[0].length;
        }
    }
    return mappings;
}
exports.buildMappings = buildMappings;
function buildStacks(chunks, stacks) {
    let offset = 0;
    let index = 0;
    const result = [];
    for (const stack of stacks) {
        const start = offset;
        for (let i = 0; i < stack.length; i++) {
            const segment = chunks[index + i];
            if (typeof segment === 'string') {
                offset += segment.length;
            }
            else {
                offset += segment[0].length;
            }
        }
        index += stack.length;
        result.push({
            range: [start, offset],
            source: stack.stack,
        });
    }
    return result;
}
exports.buildStacks = buildStacks;
//# sourceMappingURL=index.js.map