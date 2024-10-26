
 export default function trimAddress(address: string, len: number) {
    return address.substring(0, len) + '.....' + address.slice(-len)
}