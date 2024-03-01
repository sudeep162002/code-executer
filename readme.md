code deparser

const deparser = inputString
      .replace(/\\n/g, '\n')       // Replace "\\n" with actual line breaks
      .replace(/^    /gm, '')      // Remove leading indentation
      .replace(/\\"/g, '"');       // Replace escaped double quotes with double quotes



code parser

const code = normalCode
    .replace(/"/g, '\\"') // Replace all double quotes with escaped double quotes
    .split('\n')          // Split the code into lines
    .map(line => `    ${line}`) // Indent each line
    .join('\\n\\n');      // Join the lines with "\n\n"

console.log(`const code = "${code}";`);



command for running kafka container

docker compose up



command for input data to a topic from cmd 

kafka-console-producer.sh --broker-list localhost:9092 --topic code


command for checking what is inside a tpic

kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic code-output --from-beginning





sampel input
{"id":"69888","userName":"sudeep162002","output":"#include <iostream>\\n\\n    int main() {\\n\\n        std::cout << \"hello, World!\" << std::endl;\\n\\n        return 0;\\n    }"}