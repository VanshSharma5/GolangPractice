from fastapi import FastAPI, Body
from typing import Any

app = FastAPI()

@app.get("/get")
def serve_get_request():
    return {
        "message": "Hey it's GET Request",
        "method": "GET"
    }


@app.post("/post")
def serve_post_request(request: Any = Body()):
    return {
        "message": "Hey it's POST Request",
        "method": "POST",
        "request": str(request)
    }

@app.post("/post-form")
def serve_post_request(request: Any = Body()):
    return {
        "message": "Hey it's POST Request With Form Data",
        "method": "POST",
        "request": str(request)
    }