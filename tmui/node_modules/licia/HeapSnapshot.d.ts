import LinkedList = require('./LinkedList');

declare class HeapSnapshot {
    nodes: LinkedList;
    edges: LinkedList;
    constructor(profile: any);
}

export = HeapSnapshot;
