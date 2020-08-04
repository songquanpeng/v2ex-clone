'use strict';
const express = require('express');
const router = express.Router();
const Comment = require('../models/comment').Comment;
const getDate = require('../utils/util').getDate;
const md2html = require('../utils/util').md2html;
const checkLogin = require('../middlewares/check').checkLogin;
const checkPermission = require('../middlewares/check').checkPermission;

router.post('/', checkLogin, (req, res, next) => {
  let user_id = req.session.user.id;
  let page_id = req.body.page_id;
  let status = 1;
  let content = md2html(req.body.content);
  let post_time = getDate();
  let up_vote = 0;
  let down_vote = 0;
  let page = {
    page_id,
    user_id,
    content,
    status,
    post_time,
    up_vote,
    down_vote
  };
  Comment.add(page, (status, message) => {
    res.redirect(req.get('referer'));
  });
});

// Notice this id is page's id, not comment's id.
router.get('/parsed/:id', (req, res, next) => {
  const id = req.params.id;
  Comment.getByPageId(id, (status, message, comments) => {
    res.json({ status, message, comments: md2html(comments) });
  });
});

// Notice this id is page's id, not comment's id.
router.get('/:id', checkPermission, (req, res, next) => {
  const id = req.params.id;
  Comment.getByPageId(id, (status, message, comments) => {
    res.json({ status, message, comments });
  });
});

// Notice this id is comment's id.
router.delete('/:id', checkPermission, (req, res, next) => {
  const id = req.params.id;
  Comment.deleteById(id, (status, message) => {
    res.json({ status, message });
  });
});

module.exports = router;
