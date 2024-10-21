import Emitter = require('./Emitter');

declare class Tween extends Emitter {
    constructor(target: any);
    to(props: any, duration?: number, ease?: string | Function): Tween;
    progress(): number;
    progress(progress: number): Tween;
    play(): Tween;
    pause(): Tween;
    paused(): boolean;
}

export = Tween;
