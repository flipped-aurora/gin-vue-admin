exports = function(val) {
    try {
        JSON.parse(val);
        return true;
    } catch (e) {
        return false;
    }
};

module.exports = exports;
