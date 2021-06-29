const express = require('express')
const app = express()
const port = process.env.PORT || 3000

app.get('/',(req,res) => {

    res.sendFile(__dirname+'/view/homepage.html')
    return
})



app.get('/*',(req,res) => {

    res.status(404).sendFile(__dirname+'/view/404.html')
    return
})




app.listen(port, () => {
    console.log(`app is listening on port ${port}`)
})