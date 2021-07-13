exports.up = function(knex) {
  return knex.schema.createTable('report', (table) => {
    table.increments('id')
    table.string('first_name').notNullable
    table.string('last_name').notNullable
    table.text('report').notNullable
    table.timestamps(true, true);
  })  
};

exports.down = function(knex) {
  return knex.schema.dropTable('report');
};

