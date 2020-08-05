'use strict';
const express = require('express');
const router = express.Router();
const Page = require('../models/page').Page;
const checkLogin = require('../middlewares/check').checkLogin;
const checkPermission = require('../middlewares/check').checkPermission;
const getDate = require('../utils/util').getDate;
const md2html = require('../utils/util').md2html;
const Stream = require('stream');

router.post('/search', checkPermission, function(req, res, next) {
  let keyword = req.body.keyword;
  keyword = keyword ? keyword.trim() : '';
  Page.search(keyword, (status, message, pages) => {
    res.json({
      status,
      message,
      pages
    });
  });
});

// Add page
router.post('/', checkLogin, (req, res, next) => {
  req.app.locals.sitemap = undefined;
  let page_status = 1;
  let comment_status = 1;
  let title = req.body.title;
  let content = req.body.content;
  let tag = req.body.tag;
  let description = req.body.description;
  let password = req.body.password;
  let user_id = req.session.user.id;
  let post_time = getDate();
  let edit_time = post_time;
  let view = 0;
  let up_vote = 0;
  let down_vote = 0;
  if (title.trim().length > 100 || tag.trim().length > 15) {
    req.flash(
      'message',
      'Invalid parameters: the length of the title or the node name is too long.'
    );
    res.redirect('/new');
    return;
  }
  let page = {
    user_id,
    page_status,
    post_time,
    edit_time,
    comment_status,
    title,
    content,
    tag,
    description,
    password,
    view,
    up_vote,
    down_vote
  };
  page.converted_content = md2html(page.content);
  Page.add(page, (status, message, id) => {
    req.flash('message', message);
    if (status) {
      res.redirect('/');
    } else {
      res.redirect('/new');
    }
  });
});

router.get('/', checkLogin, (req, res, next) => {
  Page.all((status, message, pages) => {
    res.json({ status, message, pages });
  });
});

// Update page
router.put('/', checkPermission, (req, res, next) => {
  req.app.locals.sitemap = undefined;
  const id = req.body.id;
  let page_status = req.body.page_status;
  let comment_status = req.body.comment_status;
  let title = req.body.title;
  let content = req.body.content;
  let tag = req.body.tag;
  let description = req.body.description;
  let password = req.body.password;
  let edit_time = getDate();

  let page = {
    page_status,
    edit_time,
    comment_status,
    title,
    content,
    tag,
    password,
    description
  };
  if (page.content) {
    page.converted_content = md2html(page.content);
  } else {
    delete page.content;
  }

  Page.updateById(id, page, (status, message) => {
    res.json({ status, message });
  });
});

router.delete('/:id', checkPermission, (req, res, next) => {
  req.app.locals.sitemap = undefined;
  const id = req.params.id;
  Page.delete(id, (status, message) => {
    res.json({ status, message });
  });
});
module.exports = router;
