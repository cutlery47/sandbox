from fastapi import FastAPI
import uvicorn
import json

config = json.loads(open('config.json', 'r').read())

app = FastAPI()

@app.get("/endpoint")
def foo():
    return "service 1 data"

