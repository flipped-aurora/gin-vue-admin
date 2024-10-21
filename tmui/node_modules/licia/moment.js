var Class = require('./Class');
var toDate = require('./toDate');
var dateFormat = require('./dateFormat');
var isLeapYear = require('./isLeapYear');
var extend = require('./extend');
var toStr = require('./toStr');
var isNil = require('./isNil');
var ms = require('./ms');
exports = function(val) {
    return new Moment(val);
};
var Moment = Class({
    initialize: function Moment(val) {
        this._d = toDate(val);
        this._init();
    },
    _init: function() {
        var d = this._d;
        extend(this, {
            _year: d.getFullYear(),
            _month: d.getMonth(),
            _date: d.getDate(),
            _hour: d.getHours(),
            _minute: d.getMinutes(),
            _second: d.getSeconds(),
            _millisecond: d.getMilliseconds()
        });
        return this;
    },
    format: function(mask) {
        return dateFormat(this._d, mask);
    },
    isValid: function() {
        return !(this._d.toString() === 'Invalid Date');
    },
    isLeapYear: function() {
        return isLeapYear(this._year);
    },
    isSame: function(that) {
        return this.valueOf() === that.valueOf();
    },
    valueOf: function() {
        return this._d.getTime();
    },
    isBefore: function(that) {
        return this.valueOf() < that.valueOf();
    },
    isAfter: function(that) {
        return this.valueOf() > that.valueOf();
    },
    year: makeGetSet('year'),
    month: makeGetSet('month'),
    date: makeGetSet('date'),
    hour: makeGetSet('hour'),
    minute: makeGetSet('minute'),
    second: makeGetSet('second'),
    millisecond: makeGetSet('millisecond'),
    unix: function() {
        return floor(this.valueOf() / 1000);
    },
    clone: function() {
        return new Moment(this);
    },
    toDate: function() {
        return new Date(this._d);
    },
    toArray: function() {
        return [
            this._year,
            this._month,
            this._date,
            this._hour,
            this._minute,
            this._second,
            this._millisecond
        ];
    },
    toJSON: function() {
        return this.toISOString();
    },
    toISOString: function() {
        return this.toDate().toISOString();
    },
    toObject: function() {
        return {
            years: this._year,
            months: this._month,
            date: this._date,
            hours: this._hour,
            minutes: this._minute,
            seconds: this._second,
            milliseconds: this._millisecond
        };
    },
    toString: function() {
        return this._d.toUTCString();
    },
    set: function(unit, num) {
        var d = this._d;
        unit = normalizeUnit(unit);
        switch (unit) {
            case 'year':
                d.setFullYear(num);
                break;
            case 'month':
                d.setMonth(num);
                break;
            case 'date':
                d.setDate(num);
                break;
            case 'hour':
                d.setHours(num);
                break;
            case 'minute':
                d.setMinutes(num);
                break;
            case 'second':
                d.setSeconds(num);
                break;
            case 'millisecond':
                d.setMilliseconds(num);
                break;
        }
        return this._init();
    },
    startOf: function(unit) {
        unit = normalizeUnit(unit);

        /* eslint-disable no-fallthrough */
        switch (unit) {
            case 'year':
                this.month(0);
            case 'month':
                this.date(1);
            case 'day':
            case 'date':
                this.hour(0);
            case 'hour':
                this.minute(0);
            case 'minute':
                this.second(0);
            case 'second':
                this.millisecond(0);
        }
        return this;
    },
    endOf: function(unit) {
        return this.startOf(unit)
            .add(1, unit)
            .subtract(1, 'ms');
    },
    daysInMonth: function() {
        return this.clone()
            .endOf('month')
            .date();
    },
    add: createAdder(1),
    subtract: createAdder(-1),
    diff: function(input, unit, asFloat) {
        var that = input instanceof Moment ? input : new Moment(input);
        var ret;
        unit = normalizeUnit(unit);
        var diff = this - that;
        switch (unit) {
            case 'year':
                ret = monthDiff(this, that) / 12;
                break;
            case 'month':
                ret = monthDiff(this, that);
                break;
            case 'second':
                ret = diff / 1e3;
                break;
            // 1000
            case 'minute':
                ret = diff / 6e4;
                break;

            case 'hour':
                ret = diff / 36e5;
                break;
            // 1000 * 60 * 60
            case 'day':
                ret = diff / 864e5;
                break;

            default:
                ret = diff;
        }
        return asFloat ? ret : absFloor(ret);
    }
});
var floor = Math.floor;
var ceil = Math.ceil;
function absFloor(num) {
    return num < 0 ? ceil(num) || 0 : floor(num);
}
var unitShorthandMap = {
    y: 'year',
    M: 'month',
    D: 'date',
    d: 'day',
    h: 'hour',
    m: 'minute',
    s: 'second',
    ms: 'millisecond'
};
var regEndS = /s$/;

function normalizeUnit(unit) {
    unit = toStr(unit);
    if (unitShorthandMap[unit]) return unitShorthandMap[unit];
    return unit.toLowerCase().replace(regEndS, '');
}
function makeGetSet(unit) {
    return function(num) {
        return isNil(num) ? this['_' + unit] : this.set(unit, num);
    };
}
function createAdder(dir) {
    return function(num, unit) {
        unit = normalizeUnit(unit);
        if (unit === 'month') return this.month(this._month + dir * num);
        if (unit === 'year') return this.year(this._year + dir * num);
        var duration = createDuration(num, unit);
        this._d = new Date(this.valueOf() + dir * duration);
        return this._init();
    };
}
var msMap = {
    day: 'd',
    hour: 'h',
    minute: 'm',
    second: 's',
    millisecond: ''
};
function createDuration(num, unit) {
    return ms(num + msMap[unit]);
}

function monthDiff(a, b) {
    var wholeMonthDiff = (b.year() - a.year()) * 12 + (b.month() - a.month());
    var anchor = a.clone().add(wholeMonthDiff, 'months');
    var anchor2;
    var adjust;
    if (b - anchor < 0) {
        anchor2 = a.clone().add(wholeMonthDiff - 1, 'months');
        adjust = (b - anchor) / (anchor - anchor2);
    } else {
        anchor2 = a.clone().add(wholeMonthDiff + 1, 'months');
        adjust = (b - anchor) / (anchor2 - anchor);
    }
    return -(wholeMonthDiff + adjust) || 0;
}

module.exports = exports;
