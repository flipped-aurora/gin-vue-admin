declare class BloomFilter {
    constructor(size?: number, k?: number);
    add(val: string): void;
    test(val: string): boolean;
}

export = BloomFilter;
