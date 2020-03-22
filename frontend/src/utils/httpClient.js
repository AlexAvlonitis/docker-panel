import axios from 'axios'

const BASE_URL = "http://localhost:1234"

const HttpClient = {
  get(path, id = null) {
    const url = new URL(BASE_URL)
    url.pathname = path
    if (id != null)
      url.pathname += `/${id}`

    return(
      axios
        .get(url)
        .then(response => response.data)
    )
  }
}

export default HttpClient
