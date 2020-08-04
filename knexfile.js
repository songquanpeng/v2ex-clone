module.exports = {
  development: {
    client: 'sqlite3',
    connection: {
      filename: './data.db'
    },
    useNullAsDefault: true
  },

  staging: {
    client: 'sqlite3',
    connection: {
      filename: './data.db'
    },
    useNullAsDefault: true
  },

  production: {
    client: 'sqlite3',
    connection: {
      filename: './data.db'
    },
    useNullAsDefault: true
  }
};

/*
MySQL example:
production: {
  client: 'mysql',
  connection: {
    host: '127.0.0.1',
    user: 'root',
    password: 'root',
    database: 'blog'
  }
}

SQLite example:
production: {
  client: 'sqlite3',
  connection: {
    filename: './data.db'
  }
}
 */
