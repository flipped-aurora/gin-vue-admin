declare class QuickLru {
    constructor(max: number);
    has(key: string): boolean;
    remove(key: string): void;
    get(key: string): any;
    set(key: string, val: any): void;
    clear(): void;
}

export = QuickLru;
