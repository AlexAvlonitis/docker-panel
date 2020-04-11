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
    this.version = obj.ServerVersion
    this.os = obj.OperatingSystem
  }
}

export default InfoData
