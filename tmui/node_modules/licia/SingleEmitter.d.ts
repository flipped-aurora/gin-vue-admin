import types = require('./types');

declare class SingleEmitter {
    addListener(listener: types.AnyFn): void;
    rmListener(listener: types.AnyFn): void;
    emit(...args: any[]): void;
    rmAllListeners(): void;
    static mixin(obj: any): void;
}

export = SingleEmitter;
