class Image {
  constructor(imgObj) {
    this.name = this._repoName(imgObj.RepoTags)
    this.id = imgObj.Id
    this.size = this._bytesToGB(imgObj.Size)
  }

  _repoName(repoTags) {
    if(repoTags != null)
      return repoTags[0]
  }

  _bytesToGB(bytes) {
    return (bytes / Math.pow(10, 6)).toFixed(2) + " MB"
  }
}

export default Image
