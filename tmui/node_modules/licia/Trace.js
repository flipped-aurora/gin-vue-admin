var Class = require('./Class');
var each = require('./each');
var map = require('./map');
exports = Class({
    initialize: function Trace() {
        var _this = this;
        var events =
            arguments.length > 0 && arguments[0] !== undefined
                ? arguments[0]
                : [];
        this._processes = {};
        each(events, function(event) {
            return _this.addEvent(event);
        });
    },
    addEvent: function(event) {
        var process = this.getProcess(event.pid);
        process.addEvent(event);
    },
    rmEvent: function(event) {
        var process = this.getProcess(event.pid);
        process.rmEvent(event);
    },
    getProcess: function(id) {
        var process = this._processes[id];
        if (!process) {
            process = new Process(id);
            this._processes[id] = process;
        }
        return process;
    },
    rmProcess: function(id) {
        delete this._processes[id];
    },
    processes: function() {
        return map(this._processes);
    },
    toJSON: function() {
        var events = [];
        each(this.processes(), function(process) {
            events.push.apply(events, process.toJSON());
        });
        return events;
    }
});
var Process = Class({
    initialize: function Process(id) {
        this._id = id;
        this._name = '';
        this._threads = {};
        this._metadata = {};
    },
    id: function() {
        return this._id;
    },
    name: function() {
        return this._name;
    },
    addEvent: function(event) {
        if (event.cat === '__metadata') {
            if (event.name === 'process_name') {
                this._name = event.args.name;
            }
            if (event.tid === 0) {
                this._metadata[event.name] = event.args;
                return;
            }
        }
        var thread = this.getThread(event.tid);
        thread.addEvent(event);
    },
    rmEvent: function(event) {
        var thread = this.getThread(event.tid);
        thread.rmEvent(event);
    },
    getThread: function(id) {
        var thread = this._threads[id];
        if (!thread) {
            thread = new Thread(id, this.id());
            this._threads[id] = thread;
        }
        return thread;
    },
    rmThread: function(id) {
        delete this._threads[id];
    },
    threads: function() {
        return map(this._threads);
    },
    toJSON: function() {
        var _this2 = this;
        var events = [];
        each(this._metadata, function(args, name) {
            events.push(createMetaEvent(_this2._id, 0, name, args));
        });
        each(this.threads(), function(thread) {
            events.push.apply(events, thread.toJSON());
        });
        return events;
    }
});
var Thread = Class({
    initialize: function Thread(id, pid) {
        this._id = id;
        this._pid = pid;
        this._name = '';
        this._events = [];
        this._metadata = {};
    },
    id: function() {
        return this._id;
    },
    name: function() {
        return this._name;
    },
    addEvent: function(event) {
        if (event.cat === '__metadata') {
            if (event.name === 'thread_name') {
                this._name = event.args.name;
            }
            this._metadata[event.name] = event.args;
            return;
        }
        this._events.push(event);
    },
    rmEvent: function(event) {
        var events = this._events;
        events.splice(events.indexOf(event), 1);
    },
    events: function() {
        return map(this._events);
    },
    toJSON: function() {
        var _this3 = this;
        var events = [];
        each(this._metadata, function(args, name) {
            events.push(createMetaEvent(_this3._pid, _this3._id, name, args));
        });
        each(this.events(), function(event) {
            events.push(event);
        });
        return events;
    }
});
function createMetaEvent(pid, tid, name, args) {
    return {
        args: args,
        cat: '__metadata',
        name: name,
        ph: 'M',
        pid: pid,
        tid: tid,
        ts: 0
    };
}

module.exports = exports;
