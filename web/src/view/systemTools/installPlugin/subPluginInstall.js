export const canInstallSubPlugin = (plugin) => {
  return Boolean(plugin?.pluginName) && plugin.pluginType === 'full'
}

export const createSubPluginUploadData = (plugin) => {
  if (!canInstallSubPlugin(plugin)) {
    return null
  }

  return { parentPlugin: plugin.pluginName }
}
