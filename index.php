<?php

$lastContent = end(json_decode(file_get_contents("php://input"), true)['query'])['content'];

echo "event: meta\n";
echo 'data: {"content_type": "text/markdown", "linkify": false, "suggested_replies": true}' . "\n\n";
echo "event: text\n";
echo 'data: {"text": "'.stripcslashes($lastContent).'"}' . "\n\n";
echo "event: suggested_reply\n";
echo 'data: {"text": "Hi"}' . "\n\n";
echo "event: suggested_reply\n";
echo 'data: {"text": "Hello"}' . "\n\n";
echo "event: suggested_reply\n";
echo 'data: {"text": "Hey"}' . "\n\n";
echo "event: done\n";
echo 'data: {}' . "\n\n";
