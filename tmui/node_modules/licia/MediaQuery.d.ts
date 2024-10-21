import Emitter = require('./Emitter');

declare class MediaQuery extends Emitter {
    constructor(query: string);
    setQuery(query: string): void;
    isMatch(): boolean;
}

export = MediaQuery;
