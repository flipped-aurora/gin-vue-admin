declare class Caseless {
    constructor(obj: any);
    getKey(key: string): string | void;
    set(key: string, val: any): void;
    get(key: string): any;
    remove(key: string): void;
    has(key: string): boolean;
}

export = Caseless;
