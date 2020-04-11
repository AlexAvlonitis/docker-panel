class InfoData {
  constructor(obj) {
    this.containers = {
      count: obj.Containers,
      running: obj.ContainersRunning,
      paused: obj.ContainersPaused,
      stopped: obj.ContainersStopped
    }

    this.images = obj.Images
    this.id = obj.Id
  }
}

export default InfoData
