import { Kafka, logLevel } from 'kafkajs';
import * as fs from 'fs';
import { exec } from 'child_process';

const kafka = new Kafka({
  clientId: 'my-app',
  brokers: ['localhost:9092']
});

const consumer = kafka.consumer({ groupId: 'kafka' });

const cppFilePath = 'input.cpp';



async function copyStringToCppFile(inputString, filePath) {
  const deparser = inputString
      .replace(/\\n/g, '\n')       // Replace "\\n" with actual line breaks
      .replace(/^    /gm, '')      // Remove leading indentation
      .replace(/\\"/g, '"');       // Replace escaped double quotes with double quotes

  // Write the de-parsed string to the cpp file, overwriting its content
  await fs.writeFileSync(filePath, deparser);
  console.log(`String written to ${filePath}`);
}

const producer = kafka.producer();
 producer.connect()
async function compileAndExecuteCppFile(filePath: string): Promise<void> {
    // Compile the C++ file
    exec(`g++ -o ${filePath}.out ${filePath}`, async (compileError, compileStdout, compileStderr) => {
        if (compileError) {
            console.error(`Compilation failed: ${compileError.message}`);
            return;
        }
        if (compileStderr) {
            console.error(`Compilation stderr: ${compileStderr}`);
            return;
        }
        console.log(`Compilation successful: ${compileStdout}`);

        // Check if compilation was successful
        if (compileStderr) {
            console.error(`Compilation failed: ${compileStderr}`);
            return;
        }
        console.log(`Compilation successful: ${compileStdout}`);

        // Execute the compiled C++ file
        exec(`${filePath}.out`, async (execError, execStdout, execStderr) => {
            if (execError) {
                console.error(`Execution failed: ${execError.message}`);
                return;
            }
            if (execStderr) {
                console.error(`Execution stderr: ${execStderr}`);
                return;
            }
            console.log(`Execution output: ${execStdout}`);

            // Publish the execution output to Kafka topic 'code-output'
            await producer.send({
                topic: 'code-output',
                messages: [
                    { value: execStdout }
                ]
            });
        });
    });
}

// compileAndExecuteCppFile(cppFilePath);
// Example usage:
// const cppFilePath = 'path/to/your/cpp/file.cpp'; // Update with your C++ file path

async function run() {
  await consumer.connect();
  await consumer.subscribe({ topic: 'code' });

  await consumer.run({
    eachMessage: async ({ topic, partition, message }) => {
      const stringValue: string = message.value?.toString('utf8') || '';
      console.log(`Received message: ${stringValue}`);
      await copyStringToCppFile(stringValue, cppFilePath);
      await compileAndExecuteCppFile(cppFilePath);
    },
  });
}

run().catch(console.error);
