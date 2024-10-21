declare class Semaphore {
    constructor(counter?: number);
    wait(fn: () => void): void;
    signal(): void;
}

export = Semaphore;
