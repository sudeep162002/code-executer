-- Create a table named 'output'
CREATE TABLE output (
    id INT AUTO_INCREMENT PRIMARY KEY,
    reqid INT NOT NULL,
    username VARCHAR(255) NOT NULL,
    code_output TEXT
);
