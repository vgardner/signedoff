var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Express' });
});

router.get('/helloworld', function(req, res, next) {
  res.render('helloworld', { title: 'Hello World' });
});

router.get('/testing', function(req, res, next) {
  res.send("that's all folks");
});

module.exports = router;
