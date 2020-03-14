export function stringify(JSONdata, keyPressed) {
  if (keyPressed === null || keyPressed === "") {
    return JSON.stringify(JSONdata, null, 2)
  } else {
    const keys = Object.keys(JSONdata)
    let result = keys.filter(key => key.startsWith(keyPressed))

    return JSON.stringify(pickKeys(JSONdata, result), null, 2)
  }
}

export function arrayFilter(array, keyPressed) {
  if (keyPressed === null || keyPressed === "") {
    return array
  } else {
    let result = array.filter(obj => obj.name.startsWith(keyPressed))
    return result
  }
}

function pickKeys(obj, keys) {
  return Object.assign({}, ...keys.map(k => k in obj ? { [k]: obj[k] } : {}))
}
