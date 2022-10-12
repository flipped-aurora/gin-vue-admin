const { app, BrowserWindow } = require('electron')
const path = require('path')

function createWindow() {
  const mainWindow = new BrowserWindow({
    webPreferences: {
      enableRemoteModule: true,
      webSecurity: false,
    }
  })
  if (app.isPackaged) {
    mainWindow.loadFile(path.join(__dirname, './dist/index.html'))
  } else {
    mainWindow.loadURL('http://127.0.0.1:8080')
    mainWindow.webContents.openDevTools()
  }
}

app.whenReady().then(() => {
  createWindow()
  app.on('activate', function() {
    if (BrowserWindow.getAllWindows().length === 0) createWindow()
  })
})

app.on('window-all-closed', function() {
  if (process.platform !== 'darwin') app.quit()
})
