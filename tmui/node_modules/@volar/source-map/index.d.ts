import { Segment, StackNode } from 'muggle-string';
export * from 'muggle-string';
export interface Mapping<T = any> {
    source?: string;
    sourceRange: [number, number];
    generatedRange: [number, number];
    data: T;
}
export interface Stack {
    source: string;
    range: [number, number];
}
export declare class SourceMap<Data = any> {
    readonly mappings: Mapping<Data>[];
    private _memo;
    private get memo();
    constructor(mappings: Mapping<Data>[]);
    toSourceOffset(start: number, baseOnRight?: boolean): readonly [number, Mapping<Data>] | undefined;
    toGeneratedOffset(start: number, baseOnRight?: boolean): readonly [number, Mapping<Data>] | undefined;
    toSourceOffsets(start: number, baseOnRight?: boolean): Generator<readonly [number, Mapping<Data>], void, unknown>;
    toGeneratedOffsets(start: number, baseOnRight?: boolean): Generator<readonly [number, Mapping<Data>], void, unknown>;
    matching(startOffset: number, from: 'sourceRange' | 'generatedRange', to: 'sourceRange' | 'generatedRange', baseOnRight: boolean): Generator<readonly [number, Mapping<Data>], void, unknown>;
    matchOffset(start: number, mappedFromRange: [number, number], mappedToRange: [number, number], baseOnRight: boolean): number | undefined;
    private binarySearchMemo;
}
export declare function buildMappings<T>(chunks: Segment<T>[]): Mapping<T>[];
export declare function buildStacks<T>(chunks: Segment<T>[], stacks: StackNode[]): Stack[];
//# sourceMappingURL=index.d.ts.map