'use strict';
const express = require('express');
const router = express.Router();
const User = require('../models/user').User;
const checkLogin = require('../middlewares/check').checkLogin;
const checkPermission = require('../middlewares/check').checkPermission;

router.post('/login', function(req, res) {
  let username = req.body.username;
  let password = req.body.password;
  if (username) username = username.trim();
  if (password) password = password.trim();
  if (username === '' || password === '') {
    req.flash('message', 'Invalid parameter.');
    res.redirect('/signin');
    return;
  }
  User.check(username, password, (status, message, user) => {
    req.flash('message', message);
    if (status) {
      req.session.user = user;
      res.redirect('/');
    } else {
      res.redirect('/signin');
    }
  });
});

router.get('/logout', function(req, res, next) {
  req.session.user = undefined;
  req.flash('message', 'Logout successfully.');
  res.redirect('/');
});

router.get('/status', checkLogin, function(req, res, next) {
  res.json({
    status: true,
    user: req.session.user
  });
});

router.post('/', function(req, res) {
  const username = req.body.username;
  const password = req.body.password;
  const display_name = req.body.display_name;
  const email = req.body.email;
  const url = req.body.url;
  const status = 1;
  const avatar = req.body.avatar;

  console.log(req.body);
  if (!username.trim() || !password.trim()) {
    req.flash('message', 'Invalid parameter: username or password.');
    res.redirect('/');
  } else {
    User.register(
      {
        username,
        password,
        display_name,
        email,
        url,
        status,
        avatar
      },
      (success, message) => {
        req.flash('message', message);
        if (success) {
          res.redirect('/signin');
        } else {
          res.redirect('/signup');
        }
      }
    );
  }
});

router.get('/', checkPermission, (req, res, next) => {
  User.all((status, message, users) => {
    res.json({ status, message, users });
  });
});

router.get('/:id', checkPermission, (req, res, next) => {
  const id = req.params.id;
  User.getById(id, (status, message, user) => {
    res.json({ status, message, user });
  });
});

router.put('/', checkPermission, (req, res, next) => {
  const id = req.body.id;
  let username = req.body.username;
  let password = req.body.password;
  let display_name = req.body.display_name;
  let status = req.body.status;
  let email = req.body.email;
  let url = req.body.url;
  const avatar = req.body.avatar;

  let user = {
    username,
    password,
    display_name,
    status,
    email,
    avatar,
    url
  };

  User.updateById(id, user, (status, message) => {
    res.json({ status, message });
  });
});

router.delete('/:id', checkPermission, (req, res, next) => {
  const id = req.params.id;
  User.delete(id, (status, message) => {
    res.json({ status, message });
  });
});

module.exports = router;
