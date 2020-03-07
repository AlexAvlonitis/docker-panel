class Image {
  constructor(imgObj) {
    this.id = this._trimID(imgObj.Id)
    this.containers = imgObj.Containers
  }

  toJSON() {
    return {
      id: this.id,
      containers: this.containers
    };
  }

  _trimID(id) {
    return id
  }
}

export default Image
