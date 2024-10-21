import SingleEmitter = require('./SingleEmitter');
import extend = require('./extend');

declare class ResizeSensor extends SingleEmitter {
    constructor(el: HTMLElement);
    destroy(): void;
}

export = ResizeSensor;
