import axios from 'axios'

const BASE_URL = "http://localhost:1234"

export default class HttpClient {
  static get(path, id = null) {
    const url = new URL(BASE_URL)
    url.path = path
    if (id != null)
      url.path += id

    axios
      .get(url)
      .then(response => response.data)
  }
}
