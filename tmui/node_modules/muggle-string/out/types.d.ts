declare const NO_DATA_SYMBOL: unique symbol;
export type Segment<T = typeof NO_DATA_SYMBOL> = string | (T extends typeof NO_DATA_SYMBOL ? [
    string,
    string | undefined,
    number | [number, number]
] : [
    string,
    string | undefined,
    number | [number, number],
    T
]);
export interface StackNode {
    length: number;
    stack: string;
}
export {};
