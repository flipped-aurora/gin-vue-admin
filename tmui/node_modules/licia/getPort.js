const toArr = require('./toArr');

const net = require('net');

exports = function(ports, host) {
    ports = toArr(ports);
    ports.push(0);

    return ports.reduce((seq, port) => {
        return seq.catch(() => isAvailable(port, host));
    }, Promise.reject());
};

function isAvailable(port, host) {
    return new Promise((resolve, reject) => {
        const server = net.createServer();

        server.unref();
        server.on('error', reject);
        const options = {};
        options.port = port;
        if (host) options.host = host;
        server.listen(options, () => {
            const { port } = server.address();
            server.close(() => {
                resolve(port);
            });
        });
    });
}

module.exports = exports;
