import Trace = require('./Trace');

declare class Tracing {
    constructor(options?: {
        pid?: number;
        tid?: number;
        processName?: string;
        threadName?: string;
    });
    start(cat?: string): void;
    stop(): Trace.IEvent[];
    metadata(name: string, args: any): void;
    begin(cat: string, name: string, args?: any): void;
    end(args?: any): void;
    asyncBegin(cat: string, name: string, id?: string, args?: any): string;
    asyncEnd(id: string, args?: any): void;
    instant(
        cat: string,
        name: string,
        scope?: 'g' | 'p' | 't',
        args?: any
    ): void;
    id(): string;
}

export = Tracing;
