export interface Signal<T = any> {
    (): T;
    set(newValue: T): void;
    markDirty(): void;
}
export declare function signal<T>(): Signal<T | undefined>;
export declare function signal<T>(oldValue: T): Signal<T>;
