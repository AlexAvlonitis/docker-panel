class Image {
  constructor(imgObj) {
    this.id = this._trimID(imgObj.Id)
    this.nameTag = this._repoName(imgObj.RepoTags)
    this.size = this._bytesToGB(imgObj.Size)
    this.meta = imgObj
  }

  _trimID(id) {
    return id.slice(0, 20) + '...'
  }

  _repoName(repoTags) {
    return repoTags[0]
  }

  _bytesToGB(bytes) {
    return (bytes / Math.pow(10, 6)).toFixed(2) + " MB"
  }
}

export default Image
