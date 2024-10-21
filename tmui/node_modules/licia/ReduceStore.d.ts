import types = require('./types');

declare class ReduceStore {
    constructor(reducer: types.AnyFn, initialState: any);
    subscribe(listener: types.AnyFn): types.AnyFn;
    dispatch(action: any): any;
    getState(): any;
}

export = ReduceStore;
