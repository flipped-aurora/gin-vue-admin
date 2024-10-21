declare namespace Trace {
    interface IEvent {
        name: string;
        cat: string;
        ph: string;
        ts: number;
        pid: number;
        tid: number;
        args: any;
        [key: string]: any;
    }
    class Process {
        constructor(id);
        id(): string;
        name(): string;
        addEvent(IEvent): void;
        rmEvent(IEvent): void;
        getThread(id: number): Thread;
        rmThread(id: number): void;
        threads(): Thread[];
        toJSON(): IEvent[];
    }
    class Thread {
        constructor(id, pid);
        id(): string;
        name(): string;
        addEvent(IEvent): void;
        rmEvent(IEvent): void;
        events(): IEvent[];
        toJSON(): IEvent[];
    }
}
declare class Trace {
    constructor(events: Trace.IEvent[]);
    addEvent(event: Trace.IEvent);
    rmEvent(event: Trace.IEvent);
    getProcess(id: number): Trace.Process;
    rmProcess(id: number): void;
    processes(): Trace.Process[];
    toJSON(): Trace.IEvent[];
}

export = Trace;
