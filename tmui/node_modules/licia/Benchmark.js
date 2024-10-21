var Class = require('./Class');
var defaults = require('./defaults');
var Promise = require('./Promise');
var perfNow = require('./perfNow');
var delay = require('./delay');
var average = require('./average');
var reduce = require('./reduce');
var each = require('./each');
var map = require('./map');
var table = require('./table');
var toStr = require('./toStr');
exports = Class(
    {
        initialize: function Benchmark(fn) {
            var options =
                arguments.length > 1 && arguments[1] !== undefined
                    ? arguments[1]
                    : {};
            defaults(options, defOpts);
            this._fn = fn;
            this._isRunning = false;
            this._options = options;
        },
        run: function() {
            var _this = this;
            if (this._isRunning) {
                return this._pendingPromise;
            }
            this._reset();
            this._isRunning = true;
            var options = this._options;
            var pendingPromise = new Promise(function(resolve, reject) {
                var runSample = function() {
                    var initCount =
                        arguments.length > 0 && arguments[0] !== undefined
                            ? arguments[0]
                            : 1;
                    delay(function() {
                        _this
                            ._runSample(initCount)
                            .then(function(_ref) {
                                var period = _ref.period,
                                    count = _ref.count;
                                var sample = _this._sample;
                                sample.push(period);
                                if (
                                    perfNow() - _this._timeStamp <
                                        options.maxTime ||
                                    sample.length < options.minSamples
                                ) {
                                    runSample(count);
                                } else {
                                    resolve(_this._calcResult());
                                }
                            })
                            .catch(function(err) {
                                reject(err);
                            });
                    }, options.delay);
                };
                runSample();
            });
            function complete() {
                this._isRunning = false;
                delete this._pendingPromise;
            }
            pendingPromise.then(complete).catch(complete);
            this._pendingPromise = pendingPromise;
            return pendingPromise;
        },
        _reset: function() {
            this._timeStamp = perfNow();
            this._sample = [];
        },
        _calcResult: function() {
            var sample = this._sample;
            var result = {
                sample: sample,
                toString: function() {
                    var hz = this.hz,
                        rme = this.rme,
                        name = this.name;
                    var size = this.sample.length;
                    return ''
                        .concat(name, ' x ')
                        .concat(
                            formatNumber(hz.toFixed(hz < 100 ? 2 : 0)),
                            ' ops/sec \xB1'
                        )
                        .concat(rme.toFixed(2), '% (')
                        .concat(size, ' run')
                        .concat(size === 1 ? '' : 's', ' sampled)');
                }
            };
            var size = sample.length;
            result.name = this._options.name || this._fn.name || 'anonymous';
            result.mean = average.apply(null, sample);
            function varOf(sum, x) {
                return sum + Math.pow(x - result.mean, 2);
            }
            result.variance = reduce(sample, varOf, 0) / (size - 1) || 0;
            result.deviation = Math.sqrt(result.variance);
            result.sem = result.deviation / Math.sqrt(size);
            var critical = tTable[Math.round(size - 1) || 1] || tTable.infinity;
            result.moe = result.sem * critical;
            result.rme = (result.moe / result.mean) * 100 || 0;
            result.hz = 1000 / result.mean;
            return result;
        },
        _runSample: function(count) {
            var _this2 = this;
            var options = this._options;
            var minTime = options.minTime;
            return new Promise(function(resolve, reject) {
                var runCycle = function(count) {
                    delay(function() {
                        var elapsed = 0;
                        try {
                            elapsed = _this2._runCycle(count);
                        } catch (e) {
                            return reject(e);
                        }
                        var period = elapsed / count;
                        if (elapsed < minTime) {
                            if (elapsed === 0) {
                                count *= 100;
                            } else {
                                count += Math.ceil(
                                    (minTime - elapsed) / period
                                );
                            }
                            runCycle(count);
                        } else {
                            resolve({
                                count: count,
                                period: period
                            });
                        }
                    }, options.delay);
                };
                runCycle(count);
            });
        },
        _runCycle: function(count) {
            var fn = this._fn;
            var now = perfNow();
            while (count--) {
                fn();
            }
            return perfNow() - now;
        }
    },
    {
        all: function(benches, options) {
            var promises = [];
            each(benches, function(bench) {
                if (!(bench instanceof exports)) {
                    bench = new exports(bench, options);
                }
                promises.push(bench.run());
            });
            return Promise.all(promises).then(function(results) {
                results.toString = function() {
                    var data = map(results, function(_ref2, idx) {
                        var name = _ref2.name,
                            sample = _ref2.sample,
                            hz = _ref2.hz,
                            rme = _ref2.rme;
                        var columns = [];
                        var size = sample.length;
                        columns.push(
                            toStr(idx + 1),
                            name || 'anonymous',
                            formatNumber(hz.toFixed(hz < 100 ? 2 : 0)),
                            '\xB1'.concat(rme.toFixed(2), '%'),
                            ''
                                .concat(size, ' run')
                                .concat(size === 1 ? '' : 's')
                        );
                        return columns;
                    });
                    data.unshift([
                        'index',
                        'name',
                        'ops/sec',
                        'rme',
                        'sampled'
                    ]);
                    return table(data);
                };
                return results;
            });
        }
    }
);
var defOpts = {
    minTime: 50,
    maxTime: 5000,
    minSamples: 5,
    delay: 5,
    name: ''
};
var tTable = {
    '1': 12.706,
    '2': 4.303,
    '3': 3.182,
    '4': 2.776,
    '5': 2.571,
    '6': 2.447,
    '7': 2.365,
    '8': 2.306,
    '9': 2.262,
    '10': 2.228,
    '11': 2.201,
    '12': 2.179,
    '13': 2.16,
    '14': 2.145,
    '15': 2.131,
    '16': 2.12,
    '17': 2.11,
    '18': 2.101,
    '19': 2.093,
    '20': 2.086,
    '21': 2.08,
    '22': 2.074,
    '23': 2.069,
    '24': 2.064,
    '25': 2.06,
    '26': 2.056,
    '27': 2.052,
    '28': 2.048,
    '29': 2.045,
    '30': 2.042,
    infinity: 1.96
};
function formatNumber(number) {
    number = String(number).split('.');
    return (
        number[0].replace(/(?=(?:\d{3})+$)(?!\b)/g, ',') +
        (number[1] ? '.' + number[1] : '')
    );
}

module.exports = exports;
