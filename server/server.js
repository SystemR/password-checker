const spawn = require('child_process').spawn;
const express = require('express');
const fs = require('fs');

const app = express();

let args = process.argv.slice(2);
let passFolderTest = `${args[0]}/a.txt`;

if (fs.existsSync(passFolderTest)) {
  app.use(express.static(__dirname + '/public'));

  app.get('/api', (req, res) => {
    let q = req.query.q;
    let start = new Date();
    if (q && q.length > 2) {
      let char = q[0];
      let filename = 'other';
      if (/[a-z0-9]/i.test(char)) {
        filename = char.toLowerCase();
      }

      let stdout = '';
      let cmd = spawn('$GOBIN/finder', [`'pass/${filename}.txt'`, `'${q}'`]);
      cmd.stdout.on('data', function(data) {
        stdout += data;
      });

      cmd.on('close', function() {
        let output = stdout.split('\n');
        output.pop(); // remove last empty line
        let len = output.length;
        res.send({
          data: output,
          count: len
        });
      });
    } else {
      res.send({
        data: [],
        count: 0
      });
    }
  });

  app.listen(1337, () => console.log('Password checker listening on port 1337!'));
} else {
  console.error('Please specify the password folder when running the script.');
}
