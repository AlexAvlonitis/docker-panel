class Container {
  constructor(obj) {
    this.name = this._getFirstName(obj.Names)
    this.image = obj.Image
    this.id = obj.Id
  }

  _getFirstName(names = []) {
    return names[0].substring(1)
  }
}

export default Container
