class Network {
  constructor(imgObj) {
    this.name = imgObj.Name
    this.ipAddr = this._ipAddr(imgObj)
    this.id = imgObj.Id
  }

  _ipAddr(imgObj) {
    let conf = imgObj.IPAM.Config[0]
    return conf ? conf.Subnet : null
  }
}

export default Network
