const fs = require('fs')

const ls = fs.readdirSync('.').filter((item) => /\.txt$/.test(item))

for (let i = 0; i < ls.length; i++) {
  let file = ls[i]

  let data = fs.readFileSync(file).toString()
  let parsed = data
    .split('\n')
    .filter((line) => line.length > 3)
    .map((line) => {
      const pos = line.indexOf(' ')
      const emoji = line.substring(0, pos).trim()
      const desc = line.substring(pos + 1).trim()
      return [emoji, desc]
    })
  let nfile = file.replace(/\.txt/, '.json').toLowerCase().substring(1)

  fs.writeFileSync(nfile, JSON.stringify(parsed))
}
