import types = require('./types');

declare class Readiness {
    signal(tasks: string | string[]): void;
    isReady(tasks: string | string[]): boolean;
    ready(tasks: string | string[], fn?: types.AnyFn): Promise<void>;
}

export = Readiness;
