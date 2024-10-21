import Emitter = require('./Emitter');
import Enum = require('./Enum');

declare class Logger extends Emitter {
    name: string;
    formatter(type: string, argList: any[]): any[];
    constructor(name: string, level?: string | number);
    setLevel(level: string | number): Logger;
    getLevel(): number;
    trace(...args: any[]): Logger;
    debug(...args: any[]): Logger;
    info(...args: any[]): Logger;
    warn(...args: any[]): Logger;
    error(...args: any[]): Logger;
    static level: Enum;
}

export = Logger;
