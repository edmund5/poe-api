const http = require('http');
const fs = require('fs');

const server = http.createServer(async (req, res) => {
  let rawData = '';
  req.on('data', (chunk) => { rawData += chunk; });
  req.on('end', async () => {
    const lastContent = JSON.parse(rawData)['query'].pop()['content'];

    res.writeHead(200, {
      'Content-Type': 'text/event-stream',
    });

    res.write("event: meta\n");
    res.write('data: {"content_type": "text/markdown", "linkify": false, "suggested_replies": true}\n\n');

    res.write("event: text\n");
    res.write('data: {"text": "' + lastContent.replace(/\\/g, '') + '"}\n\n');

    res.write("event: suggested_reply\n");
    res.write('data: {"text": "Hi"}\n\n');

    res.write("event: suggested_reply\n");
    res.write('data: {"text": "Hello"}\n\n');

    res.write("event: suggested_reply\n");
    res.write('data: {"text": "Hey"}\n\n');

    res.write("event: done\n");
    res.write('data: {}\n\n');
    
    res.end();
  });
});

server.listen(3000, () => {
  console.log('Server running at http://localhost:3000/');
});