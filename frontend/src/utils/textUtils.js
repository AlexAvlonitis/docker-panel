export default function stringify(JSONdata, keyPressed) {
  if (keyPressed === null || keyPressed === "") {
    return JSON.stringify(JSONdata, null, 2)
  } else {
    const keys = Object.keys(JSONdata)
    let result = keys.filter(key => key.startsWith(keyPressed))

    return JSON.stringify(pickKeys(JSONdata, result), null, 2)
  }
}

function pickKeys(obj, keys) {
  return Object.assign({}, ...keys.map(k => k in obj ? { [k]: obj[k] } : {}))
}
