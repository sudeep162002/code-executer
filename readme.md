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