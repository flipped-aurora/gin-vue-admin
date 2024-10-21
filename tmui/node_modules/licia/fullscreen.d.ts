import Emitter = require('./Emitter');

declare namespace fullscreen {
    interface IFullscreen extends Emitter {
        request(el?: Element): void;
        exit(): void;
        toggle(el?: Element): void;
        isActive(): boolean;
        getEl(): Element | null;
        isEnabled(): boolean;
    }
}
declare const fullscreen: fullscreen.IFullscreen;

export = fullscreen;
