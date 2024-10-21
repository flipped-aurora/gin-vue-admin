import Emitter = require('./Emitter');

declare class State extends Emitter {
    constructor(initial: string, events: any);
    is(state: string): boolean;
    [event: string]: any;
}

export = State;
