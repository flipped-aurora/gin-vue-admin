var Class = require('./Class');
var Trace = require('./Trace');
var perfNow = require('./perfNow');
var extend = require('./extend');
var isNode = require('./isNode');
var Stack = require('./Stack');
var map = require('./map');
var trim = require('./trim');
var isEmpty = require('./isEmpty');
var intersect = require('./intersect');
var convertBase = require('./convertBase');
var defPid = 0;
var defTid = 0;
var id = 0;
if (isNode) {
    defPid = process.pid;
    try {
        defTid = eval('require')('worker_threads').threadId;
    } catch (e) {}
}
exports = Class({
    initialize: function Tracing() {
        var _ref =
                arguments.length > 0 && arguments[0] !== undefined
                    ? arguments[0]
                    : {},
            _ref$pid = _ref.pid,
            pid = _ref$pid === void 0 ? defPid : _ref$pid,
            _ref$tid = _ref.tid,
            tid = _ref$tid === void 0 ? defTid : _ref$tid,
            _ref$processName = _ref.processName,
            processName =
                _ref$processName === void 0 ? 'Process' : _ref$processName,
            _ref$threadName = _ref.threadName,
            threadName =
                _ref$threadName === void 0 ? 'Thread' : _ref$threadName;
        this._pid = pid;
        this._tid = tid;
        this._processName = processName;
        this._threadName = threadName;
    },
    start: function() {
        var cat =
            arguments.length > 0 && arguments[0] !== undefined
                ? arguments[0]
                : '';
        this._targetCat = processCat(cat);
        if (!isEmpty(this._targetCat)) {
            this._targetCat.push('__metadata');
        }
        this._traceEventStack = new Stack();
        this._asyncEventMap = {};
        this._trace = new Trace();
        this.metadata(
            'process_name',
            {
                name: this._processName
            },
            {
                tid: 0,
                ts: 0
            }
        );
        this.metadata(
            'thread_name',
            {
                name: this._threadName
            },
            {
                ts: 0
            }
        );
    },
    stop: function() {
        var trace = this._trace;
        if (!trace) {
            throw Error('Need to call start first');
        }
        delete this._targetCat;
        delete this._traceEventStack;
        delete this._asyncEventMap;
        delete this._trace;
        return trace.toJSON();
    },
    metadata: function(name, args, extra) {
        this._addEvent('__metadata', name, Phase.Metadata, args, extra);
    },
    begin: function(cat, name) {
        var args =
            arguments.length > 2 && arguments[2] !== undefined
                ? arguments[2]
                : {};
        if (!this._traceEventStack) {
            return;
        }
        this._traceEventStack.push({
            cat: cat,
            name: name,
            args: args,
            ts: this._getCurTs()
        });
    },
    end: function(args) {
        if (!this._traceEventStack) {
            return;
        }
        var beginEvent = this._traceEventStack.pop();
        if (!beginEvent) {
            throw Error('Need to call begin first');
        }
        var cat = beginEvent.cat,
            name = beginEvent.name,
            ts = beginEvent.ts;
        args = extend(beginEvent.args, args);
        this._addEvent(cat, name, Phase.Complete, args, {
            dur: this._getCurTs() - ts,
            ts: ts
        });
    },
    asyncBegin: function(cat, name) {
        var id =
            arguments.length > 2 && arguments[2] !== undefined
                ? arguments[2]
                : this.id();
        var args =
            arguments.length > 3 && arguments[3] !== undefined
                ? arguments[3]
                : {};
        if (!this._asyncEventMap) {
            return id;
        }
        this._asyncEventMap[id] = {
            cat: cat,
            name: name
        };
        this._addEvent(cat, name, Phase.NestableAsyncBegin, args, {
            id: id
        });
        return id;
    },
    asyncEnd: function(id) {
        var args =
            arguments.length > 1 && arguments[1] !== undefined
                ? arguments[1]
                : {};
        if (!this._asyncEventMap) {
            return;
        }
        var asyncBeginEvent = this._asyncEventMap[id];
        if (!asyncBeginEvent) {
            throw Error('Need to call async begin first');
        }
        var cat = asyncBeginEvent.cat,
            name = asyncBeginEvent.name;
        delete this._asyncEventMap[id];
        this._addEvent(cat, name, Phase.NestableAsyncEnd, args, {
            id: id
        });
    },
    instant: function(cat, name) {
        var scope =
            arguments.length > 2 && arguments[2] !== undefined
                ? arguments[2]
                : 't';
        var args = arguments.length > 3 ? arguments[3] : undefined;
        this._addEvent(cat, name, Phase.Instant, args, {
            s: scope
        });
    },
    id: function() {
        return '0x' + convertBase(id++, 10, 16);
    },
    _addEvent: function(cat, name, ph) {
        var args =
            arguments.length > 3 && arguments[3] !== undefined
                ? arguments[3]
                : {};
        var extra =
            arguments.length > 4 && arguments[4] !== undefined
                ? arguments[4]
                : {};
        if (!this._trace) {
            return;
        }
        var targetCat = this._targetCat;
        if (!isEmpty(targetCat)) {
            var catArr = processCat(cat);
            if (isEmpty(intersect(catArr, targetCat))) {
                return;
            }
        }
        var event = extend(
            {
                name: name,
                cat: cat,
                ph: ph,
                ts: this._getCurTs(),
                pid: this._pid,
                tid: this._tid,
                args: args
            },
            extra
        );
        this._trace.addEvent(event);
    },
    _getCurTs: function() {
        return Math.round(perfNow() * 1000);
    }
});
var Phase = {
    Begin: 'B',
    End: 'E',
    Complete: 'X',
    Instant: 'I',
    NestableAsyncBegin: 'b',
    NestableAsyncEnd: 'e',
    Metadata: 'M'
};
function processCat(cat) {
    cat = trim(cat);
    if (cat === '') {
        return [];
    }
    return map(cat.split(','), trim);
}

module.exports = exports;
