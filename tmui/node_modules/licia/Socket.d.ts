import Emitter = require('./Emitter');

declare class Socket extends Emitter {
    constructor(
        url: string,
        options?: {
            protocols?: string | string[];
            reconnect?: boolean;
        }
    );
    send(message: any): void;
    close(code?: number, reason?: string): void;
    connect(): void;
}

export = Socket;
