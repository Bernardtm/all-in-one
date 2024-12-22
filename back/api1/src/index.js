const express = require('express')
const cors = require('cors');

const app = express()

// Enable CORS with default settings
app.use(cors());

const port = 3000

app.get('/', (req, res) => {
  res.send('UP!')
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})