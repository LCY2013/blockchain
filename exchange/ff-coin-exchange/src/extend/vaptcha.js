export function vaptcha () {
  return new Promise(function (resolve, reject) {
    const tag = document.getElementsByTagName('script')
    for (let i of tag) {
      if (i.src === 'https://v.vaptcha.com/v3.js') {
        resolve(window.vaptcha)
        return
      }
    }
    const script = document.createElement('script')
    script.type = 'text/javascript'
    script.src = 'https://v.vaptcha.com/v3.js'
    script.onerror = reject
    document.body.appendChild(script)
    script.onload = () => {
      resolve(window.vaptcha)
    }
  })
}