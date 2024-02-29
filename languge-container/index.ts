import * as Kafka from 'node-rdkafka';

const consumer = new Kafka.KafkaConsumer({
  'group.id': 'kafka',
  'metadata.broker.list': 'localhost:9092',
}, {});

consumer.connect();

consumer.on('ready', () => {
  console.log('consumer ready..');
  consumer.subscribe(['code']);
  consumer.consume();
}).on('data', (data: Kafka.Message) => {
  console.log(data);
  const stringValue: string = data.value?.toString('utf8') || '';
  console.log(`received message: ${stringValue}`);
});
