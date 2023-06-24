from fastapi import FastAPI, Request, Response
from fastapi.responses import StreamingResponse
import json

app = FastAPI()

@app.post("/")
async def root(request: Request):
    body = await request.json()
    last_content = body['query'][-1]['content']

    def event_stream():
        yield "event: meta\n"
        yield 'data: {"content_type": "text/markdown", "linkify": false, "suggested_replies": true}\n\n'

        yield "event: text\n"
        yield f'data: {{"text": "{last_content.replace("\\\\", "")}"}}\n\n'

        yield "event: suggested_reply\n"
        yield 'data: {"text": "Hi"}\n\n'

        yield "event: suggested_reply\n"
        yield 'data: {"text": "Hello"}\n\n'

        yield "event: suggested_reply\n"
        yield 'data: {"text": "Hey"}\n\n'

        yield "event: done\n"
        yield 'data: {}\n\n'

    return StreamingResponse(event_stream(), media_type="text/event-stream")