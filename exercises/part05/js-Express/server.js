const express = require('express');
const bodyParser = require('body-parser');

const app = express();
const PORT = 9000

app.get('/health-check', (req, res) => {
    res.status(200).json({
	"status" : "ok"
    })
});

const server = app.listen("9000", () => {
    console.log('Starting server on port: '+ PORT);
});

module.exports = server;
