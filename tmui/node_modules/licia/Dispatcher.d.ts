import types = require('./types');

declare class Dispatcher {
    dispatch(payload: any);
    register(cb: types.AnyFn): void;
    waitFor(ids: string[]): void;
    unregister(id: string): void;
    isDispatching(): boolean;
}

export = Dispatcher;
