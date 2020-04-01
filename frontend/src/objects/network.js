class Network {
  constructor(obj) {
    this.name = obj.Name
    this.ipAddr = this._ipAddr(obj)
    this.id = obj.Id
  }

  _ipAddr(obj) {
    let conf = obj.IPAM.Config[0]
    return conf ? conf.Subnet : 'N/A'
  }
}

export default Network
