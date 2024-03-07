import { Kafka, logLevel } from 'kafkajs';
import * as fs from 'fs';
import { exec } from 'child_process';


//for local development
// const kafka = new Kafka({
//   clientId: 'my-app',
//   brokers: ['localhost:9092']
// });


//for container development
const kafka = new Kafka({
  clientId: 'my-app-java',
  brokers: ['kafka:29092']
});


const consumer = kafka.consumer({ groupId: 'kafka-java' });

const javaFilePath = 'HelloWorld';

async function runWorker(id: string, username:string,paylode:any){
        await copyStringToJavaFile(paylode, javaFilePath);
        await compileAndExecuteJavaFile(id, javaFilePath,username);
}


async function copyStringToJavaFile(inputString, filePath) {
  const deparser = inputString
      .replace(/\\n/g, '\n')       // Replace "\\n" with actual line breaks
      .replace(/^    /gm, '')      // Remove leading indentation
      .replace(/\\"/g, '"');       // Replace escaped double quotes with double quotes

  // Write the de-parsed string to the java file, overwriting its content
  await fs.writeFileSync(`${filePath}.java`, deparser);
  console.log(`String written to ${filePath}.java`);
}

const producer = kafka.producer();
 producer.connect()
 async function compileAndExecuteJavaFile(id: string, filePath: string, username: string): Promise<void> {
  // Compile the Java file
  exec(`javac ${filePath}.java`, async (compileError, compileStdout, compileStderr) => {
      if (compileError) {
          console.error(`Compilation failed: ${compileError.message}`);
          return;
      }
      if (compileStderr) {
          console.error(`Compilation stderr: ${compileStderr}`);
          return;
      }
      console.log(`Compilation successful: ${compileStdout}`);

      // Execute the compiled Java file
      exec(`java ${filePath}`, async (execError, execStdout, execStderr) => {
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
          const jsonOutput = JSON.stringify({
              "id": id,
              "userName": username,
              "output": execStdout
          });

          // Assuming 'producer' is defined outside this function
          await producer.send({
              topic: 'code-output',
              messages: [
                  { value: jsonOutput }
              ]
          });
      });
  });
}

async function run() {
  await consumer.connect();
  await consumer.subscribe({ topic: 'java-code' });

  await consumer.run({
    eachMessage: async ({ topic, partition, message }) => {
      const stringValue: string = message.value?.toString('utf8') || '';
      console.log(`Received message: ${stringValue}`);
      
      try {
        const parsedValue = JSON.parse(stringValue);
        const paylode= parsedValue.output;
        const id = parsedValue.id;
        const userName=parsedValue.userName;
       await runWorker(id,userName,paylode);
        
      } catch (error) {
        console.error('Error parsing JSON:', error);
      }
    },
  });
}

run().catch(console.error);
