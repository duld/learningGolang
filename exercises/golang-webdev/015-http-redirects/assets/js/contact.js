console.log("Hello!")
console.log("new script")

// select button
let patch_btn = document.querySelector("#patch-btn")
let response_section = document.querySelector("#http-response")

patch_btn.addEventListener('click', e => {
  e.preventDefault()

  let n_xhr = new XMLHttpRequest();

  n_xhr.onreadystatechange = () => {
    if (n_xhr.readyState === 4) {
      // create a new header
      console.log(n_xhr.responseURL)
      buildPatchResponseElement(n_xhr)
    }
  }

  n_xhr.open("patch", "/users")
  n_xhr.method = 'PATCH'
  n_xhr.reqURL = "/users"
  n_xhr.send(null)
})

function buildPatchResponseElement(xhr_res) {
  // build out our elements
  n_h3 = document.createElement("h4")
  headerText = document.createTextNode(`HTTP Method: ${xhr_res.method} -- Request URL: ${xhr_res.reqURL} -- Response URL: ${xhr_res.responseURL}`)
  n_h3.appendChild(headerText)

  // add new header to responses
  response_section.appendChild(n_h3)
}