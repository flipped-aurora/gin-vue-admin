const Store = require('./Store');
const isEmpty = require('./isEmpty');
const debounce = require('./debounce');
const stringify = require('./stringify');
const isObj = require('./isObj');
const defaults = require('./defaults');
const noop = require('./noop');

const fs = require('fs');

exports = Store.extend({
    initialize: function FileStore(path, data) {
        this._path = path;
        data = data || {};

        let fileData;
        if (fs.existsSync(path)) {
            try {
                fileData = JSON.parse(fs.readFileSync(path, 'utf8'));
            } catch (e) {
                fileData = {};
            }
        }
        if (!isObj(fileData)) {
            fileData = {};
        }
        data = defaults(fileData, data);

        this.save = debounce(data => {
            if (isEmpty(data)) {
                fs.unlink(this._path, noop);
                return;
            }
            fs.writeFile(this._path, stringify(data), 'utf8', noop);
        }, 300);

        this.callSuper(Store, 'initialize', [data]);
    }
});

module.exports = exports;
