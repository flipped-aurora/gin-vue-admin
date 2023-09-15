const child_process = require('child_process')

export default function GvaPositionServer() {
  return {
    name: 'gva-position-server',
    apply: 'serve',
    configureServer(server) {
      server.middlewares.use((req, _, next) => {
        if (req._parsedUrl.pathname === '/gvaPositionCode') {
          const path =
            req._parsedUrl.query && req._parsedUrl.query.split('=')[1]
          if (path && path !== 'null') {
            if (process.env.VITE_EDITOR === 'webstorm') {
              const linePath = path.split(':')[1]
              const filePath = path.split(':')[0]
              const platform = os()
              if (platform === 'win32') {
                child_process.exec(
                  `webstorm64.exe  --line ${linePath} ${filePath}`
                )
              } else {
                child_process.exec(
                  `webstorm --line ${linePath} ${filePath}`
                )
              }
            } else {
              child_process.exec('code -r -g ' + path)
            }
          }
        }
        next()
      })
    },
  }
}

function os() {
  'use strict'
  const os = require('os')
  const platform = os.platform()
  return platform
}
