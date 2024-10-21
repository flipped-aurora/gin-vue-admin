const net = require('net');

exports = function(port, host) {
    return new Promise(resolve => {
        const server = net.createServer();

        server.unref();
        server.on('error', () => resolve(false));
        const options = {
            port
        };
        if (host) options.host = host;
        server.listen(options, () => {
            server.close(() => {
                resolve(true);
            });
        });
    });
};

module.exports = exports;
