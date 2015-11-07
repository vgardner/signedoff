var express = require('express');
var router = express.Router();
var request = require("request");

/* GET users listing. */
router.get('/', function(req, res, next) {
  res.send('respond with a resource');
});

/* GET users listing. */
router.get('/vgardner/go-lights', function(req, res, next) {

  request("http://localhost:3001/api/releases/vgardner/go-lights", function(error, response, body) {
    res.render('index', { title: 'Express', releaseData: body });
   });
});

router.get('/dev', function(req, res, next) {
  res.render('releases', { title: 'Releases for Go Lights' });
});

module.exports = router;
