declare class Wrr {
    size: number;
    set(val: any, weight: number): void;
    get(val: any): number | void;
    remove(val: any): void;
    clear(): void;
    next(): any;
}

export = Wrr;
