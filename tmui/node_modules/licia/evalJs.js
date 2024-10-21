var root = require('./root');
exports = function(js, ctx) {
    ctx = ctx || root;

    try {
        return new Function('return (' + js + ');').call(ctx);
    } catch (e) {
        try {
            return new Function('return ' + js).call(ctx);
        } catch (e) {
            return new Function(js).call(ctx);
        }
    }
};

module.exports = exports;
