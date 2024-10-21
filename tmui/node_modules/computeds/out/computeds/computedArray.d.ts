export declare function computedArray<I, O>(arr: () => I[], computedItem: (item: () => I, index: number) => () => O): readonly Readonly<O>[];
