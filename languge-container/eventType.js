const avro = require('avsc');

const schema = avro.Type.forSchema({
  type: 'record',
  fields: [
    {
      name: 'code',
      type: 'string',
    }
  ]
});

module.exports = schema;
