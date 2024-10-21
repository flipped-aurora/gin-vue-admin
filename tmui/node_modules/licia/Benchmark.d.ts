import Promise = require('./Promise');
import delay = require('./delay');
import types = require('./types');

declare namespace Benchmark {
    interface IOptions {
        minTime?: number;
        maxTime?: number;
        minSamples?: number;
        delay?: number;
        name?: string;
    }
    interface IResult {
        name: string;
        mean: number;
        variance: number;
        deviation: number;
        sem: number;
        moe: number;
        rme: number;
        hz: number;
        sample: number[];
    }
}
declare class Benchmark {
    constructor(fn: types.AnyFn, options?: Benchmark.IOptions);
    run(): Promise<Benchmark.IResult>;
    static all(
        benches: Array<types.AnyFn | Benchmark>,
        options?: Benchmark.IOptions
    ): Promise<Benchmark.IResult[]>;
}

export = Benchmark;
