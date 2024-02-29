const Kafka = require('node-rdkafka');

var consumer = new Kafka.KafkaConsumer({
  'group.id': 'kafka',
  'metadata.broker.list': 'localhost:9092',
}, {});

consumer.connect();

consumer.on('ready', () => {
  console.log('consumer ready..')
  consumer.subscribe(['code']);
  consumer.consume();
}).on('data', function(data) {
  console.log(data)
  const stringValue = data.value.toString('utf8'); 
  console.log(`received message: ${stringValue}`);
  
});
