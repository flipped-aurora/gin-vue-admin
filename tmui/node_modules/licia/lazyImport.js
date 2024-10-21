const stackTrace = require('./stackTrace');
const splitPath = require('./splitPath');
const startWith = require('./startWith');
const endWith = require('./endWith');
const defineProp = require('./defineProp');
const isStr = require('./isStr');
const has = require('./has');
const objToStr = require('./objToStr');
const unique = require('./unique');
const concat = require('./concat');
const keys = require('./keys');
const isArr = require('./isArr');
const toBool = require('./toBool');

const path = require('path');

exports = function(importFn, dirname) {
    return function(moduleId) {
        if (isRelative(moduleId)) {
            if (!dirname) {
                dirname = findDirName();
            }
            moduleId = path.join(dirname, moduleId);
        }

        const { cache } = importFn;
        if (cache) {
            if (cache[moduleId]) {
                return importFn(moduleId);
            }
            if (!endWith(moduleId, '.js') && cache[`${moduleId}.js`]) {
                return importFn(`${moduleId}.js`);
            }
        }

        return proxyExports(importFn, moduleId);
    };
};

function proxyExports(importFn, moduleId) {
    const fakeExports = function() {};
    let cache;

    function doImport() {
        if (cache) {
            return;
        }
        const module = importFn(moduleId);
        cache = Object(module);

        const valueOfDescriptor = createDescriptor(0, 0, 1);
        if (isStr(module)) {
            valueOfDescriptor.value = () => String(module.valueOf());
        } else {
            valueOfDescriptor.value = () => Number(module.valueOf());
        }
        defineProp(cache, 'valueOf', valueOfDescriptor);

        defineProp(
            cache,
            'toString',
            createDescriptor(0, 0, 1, () => String(module.toString()))
        );

        if (!has(cache, Symbol.toStringTag)) {
            const realType = objToStr(module).slice(8, -1);
            Object.defineProperty(cache, Symbol.toStringTag, {
                configurable: true,
                get() {
                    return realType;
                }
            });
        }
    }

    return new Proxy(fakeExports, {
        get(target, property) {
            doImport();
            return cache[property];
        },
        set(target, property, value) {
            doImport();
            cache[property] = value;
            return true;
        },

        has(target, prop) {
            doImport();
            return prop in cache;
        },

        construct(target, argumentsList) {
            doImport();

            return new cache(...argumentsList);
        },

        apply(target, thisArg, argumentsList) {
            doImport();
            return cache.apply(thisArg, argumentsList);
        },

        ownKeys() {
            doImport();

            const descriptors = Object.getOwnPropertyDescriptors(cache);
            delete descriptors.valueOf;
            delete descriptors.toString;

            return unique(
                concat(
                    [
                        'arguments',
                        'caller',
                        'prototype',
                        'name',
                        'length',
                        Symbol.toStringTag
                    ],
                    keys(descriptors)
                )
            );
        },
        getOwnPropertyDescriptor(target, prop) {
            if (has(cache, prop)) {
                if (isArr(cache) && prop === 'length') {
                    return {
                        configurable: true,
                        enumerable: false,
                        writable: true
                    };
                } else {
                    const descriptor = Object.getOwnPropertyDescriptor(
                        cache,
                        prop
                    );
                    if (descriptor.configurable) {
                        return descriptor;
                    }
                    if (!fakeExports.prop) {
                        defineProp(fakeExports, prop, descriptor);
                    }
                    return descriptor;
                }
            } else {
                switch (prop) {
                    case 'arguments':
                        return createDescriptor(0, 0, 0, null);
                    case 'caller':
                        return createDescriptor(0, 0, 0, null);
                    case 'prototype':
                        return createDescriptor(0, 0, 1, null);
                    case 'length':
                        return createDescriptor(1, 0, 0, 0);
                    case 'name':
                        return createDescriptor(1, 0, 0, '');
                    default:
                        return {
                            configurable: true,
                            enumerable: true,
                            writable: true
                        };
                }
            }
        }
    });
}

function createDescriptor(configurable, enumerable, writable, value) {
    return {
        configurable: toBool(configurable),
        enumerable: toBool(enumerable),
        writable: toBool(writable),
        value
    };
}

function findDirName() {
    const stack = stackTrace();
    for (const item of stack) {
        const fileName = item.getFileName();
        if (fileName !== module.filename) {
            return splitPath(fileName).dir;
        }
    }

    return '';
}

function isRelative(moduleId) {
    return startWith(moduleId, './') || startWith(moduleId, '../');
}

module.exports = exports;
