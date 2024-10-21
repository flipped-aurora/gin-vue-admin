import Emitter = require('./Emitter');

declare namespace notify {
    class Notification extends Emitter {
        constructor(title: string, options?: object);
        show(): void;
    }
}
declare function notify(title: string, options?: object): void;

export = notify;
