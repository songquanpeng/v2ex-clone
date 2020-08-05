'use strict';
const express = require('express');
const router = express.Router();
const Page = require('../models/page').Page;
const Comment = require('../models/comment').Comment;
const User = require('../models/user').User;
const checkLogin = require('../middlewares/check').checkLogin;
const checkPermission = require('../middlewares/check').checkPermission;

router.get('/', function(req, res, next) {
  let start = 0;
  let end = 9;
  Page.getByRange(start, end, pages => {
    res.render('index', {
      pages: pages,
      prev: ``,
      next: `10-19`,
      message: req.flash('message')
    });
  });
});

router.get('/t/:link', function(req, res, next) {
  const link = req.params.link;
  Page.getByLink(link, (success, message, page, links) => {
    if (success && page !== undefined) {
      page.view++;
      Comment.getByPageId(page.id, (status, message, comments) => {
        res.locals.comments = comments;
        res.locals.links = links;
        res.render('post', { page });
      });
      Page.updateViewCounter(page.id);
    } else {
      res.render('error', { title: 'Error!', message });
    }
  });
});

router.get('/member/:username', (req, res, next) => {
  const username = req.params.username;
  User.getByUsername(username, (success, message, user) => {
    if (success && user !== undefined) {
      res.render('member', user);
    }
  });
});

router.get('/new', checkLogin, (req, res, next) => {
  res.render('editor', {
    message: req.flash('message')
  });
});

router.get('/signin', (req, res, next) => {
  res.render('signin', {
    message: req.flash('message')
  });
});

router.get('/signup', (req, res, next) => {
  res.render('signup', {
    message: req.flash('message')
  });
});

module.exports = router;
