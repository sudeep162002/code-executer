import * as Kafka from 'node-rdkafka';
import * as fs from 'fs';
const consumer = new Kafka.KafkaConsumer({
  'group.id': 'kafka',
  'metadata.broker.list': 'localhost:9092',
}, {});

consumer.connect();
// const inputString = "#include <iostream> int main() { std::cout << "Hello, World!" << std::endl; return 0; }";
const filePath = "./input.cpp";

function copyStringToCppFile(inputString, filePath) {
  // Append the input string to the cpp file
  fs.appendFileSync(filePath, inputString + '\n'); // Add newline for clarity
  console.log(`String appended to ${filePath}`);
}
// copyStringToCppFile(inputString,filePath);

consumer.on('ready', () => {
  console.log('consumer ready..');
  consumer.subscribe(['code']);
  consumer.consume();
}).on('data', (data: Kafka.Message) => {
  console.log(data);
  const stringValue: string = data.value?.toString('utf8') || '';
  console.log(`received message: ${stringValue}`);
  copyStringToCppFile(stringValue,filePath);
});
