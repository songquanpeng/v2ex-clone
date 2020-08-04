const db = require('../utils/database').db;
const uuid = require('uuid/v1');

class Comment {
  add(comment, callback) {
    comment.id = uuid();
    db('comments')
      .insert(comment)
      .asCallback(error => {
        if (error) {
          console.error(error.message);
          callback(false, error.message);
        } else {
          callback(true, '');
        }
      });
  }

  getByPageId(id, callback) {
    db('comments')
      .select([
        'comments.id as id',
        'page_id',
        'users.id as user_id',
        'users.username as username',
        'users.avatar as user_avatar',
        'post_time',
        'comments.status as status',
        'content'
      ])
      .innerJoin('users', 'users.id', 'comments.user_id')
      .where('page_id', id)
      .asCallback((error, data) => {
        if (error) {
          console.error(error.message);
          callback(false, error.message, undefined);
        } else {
          callback(true, '', data);
        }
      });
  }

  deleteById(id, callback) {
    db('comments')
      .where('id', id)
      .del()
      .asCallback(error => {
        if (error) {
          console.error(error.message);
          callback(false, error.message);
        } else {
          callback(true, '');
        }
      });
  }
}

let comment = new Comment();
module.exports.Comment = comment;
