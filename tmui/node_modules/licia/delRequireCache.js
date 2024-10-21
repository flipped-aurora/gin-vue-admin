const stackTrace = require('./stackTrace');
const each = require('./each');
const contain = require('./contain');

const path = require('path');

exports = function(id) {
    const filePath = findPath(id);

    if (!filePath) return;
    const mod = require.cache[filePath];
    if (!mod) return;

    const visited = {};
    function run(current) {
        visited[current.id] = true;
        each(current.children, child => {
            const { filename, id } = child;
            if (path.extname(filename) !== '.node' && !visited[id]) {
                run(child);
            }
        });
        delete require.cache[current.id];
    }
    run(mod);

    each(module.constructor._pathCache, (val, key) => {
        if (contain(key, filePath)) delete module.constructor._pathCache[key];
    });
};

function findPath(id) {
    if (id[0] === '.') {
        const stack = stackTrace();
        for (const item of stack) {
            const fileName = item.getFileName();
            if (fileName !== module.filename) {
                id = path.resolve(path.dirname(fileName), id);
                break;
            }
        }
    }

    try {
        return require.resolve(id);
    } catch (e) {}
}

module.exports = exports;
