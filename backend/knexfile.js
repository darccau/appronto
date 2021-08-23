module.exports = {
  development: {
    client: 'postgresql',
    connection: {
      database: 'appronto',
      user: 'appronto',
      password: 'appronto'
    },
    pool: {
      min: 2,
      max: 10
    },
    migrations: {
      directory: 'src/database/migrations'
    }
  }
}
