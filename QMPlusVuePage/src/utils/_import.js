module.exports = file => () => {
    return import ('@/' + file)
}