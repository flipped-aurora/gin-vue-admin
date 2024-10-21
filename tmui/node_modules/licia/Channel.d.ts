import Emitter = require('./Emitter');

declare class Channel extends Emitter {
    send(msg: any): void;
    connect(channel: Channel): void;
    disconnect(channel: Channel): void;
    isConnected(channel: Channel): boolean;
    destroy(): void;
}

export = Channel;
