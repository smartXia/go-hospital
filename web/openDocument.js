var child_process = require('child_process')
var url = 'www.baidu.com'
var cmd = 'www.baotao.com'
console.log(process.platform)
switch (process.platform) {
  case 'win32':
    cmd = 'start'
    child_process.exec(cmd + ' ' + url)
    break

  case 'darwin':
    cmd = 'open'
    child_process.exec(cmd + ' ' + url)
    break
}
