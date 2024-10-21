var Emitter = require('./Emitter');
var Enum = require('./Enum');
var toArr = require('./toArr');
var isUndef = require('./isUndef');
var clone = require('./clone');
var isStr = require('./isStr');
var isNum = require('./isNum');
exports = Emitter.extend(
    {
        initialize: function Logger(name, level) {
            this.name = name;
            this.setLevel(isUndef(level) ? exports.level.DEBUG : level);
            this.callSuper(Emitter, 'initialize', arguments);
        },
        setLevel: function(level) {
            if (isStr(level)) {
                level = exports.level[level.toUpperCase()];
                if (level) this._level = level;
                return this;
            }
            if (isNum(level)) this._level = level;
            return this;
        },
        getLevel: function() {
            return this._level;
        },
        formatter: function(type, argList) {
            return argList;
        },
        trace: function() {
            return this._log('trace', arguments);
        },
        debug: function() {
            return this._log('debug', arguments);
        },
        info: function() {
            return this._log('info', arguments);
        },
        warn: function() {
            return this._log('warn', arguments);
        },
        error: function() {
            return this._log('error', arguments);
        },
        _log: function(type, argList) {
            argList = toArr(argList);
            if (argList.length === 0) return this;
            this.emit('all', type, clone(argList));
            if (exports.level[type.toUpperCase()] < this._level) return this;
            this.emit(type, clone(argList));

            var consoleMethod = type === 'debug' ? console.log : console[type];
            consoleMethod.apply(console, this.formatter(type, argList));
            return this;
        }
    },
    {
        level: new Enum({
            TRACE: 0,
            DEBUG: 1,
            INFO: 2,
            WARN: 3,
            ERROR: 4,
            SILENT: 5
        })
    }
);

module.exports = exports;
