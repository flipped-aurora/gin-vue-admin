declare const uncaught: {
    start(): void;
    stop(): void;
    addListener(fn: (err: Error) => void): void;
    rmListener(fn: (err: Error) => void): void;
    rmAllListeners(): void;
};

export = uncaught;
