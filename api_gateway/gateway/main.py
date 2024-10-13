from fastapi import FastAPI
from router import router, config
import uvicorn

app = FastAPI()
app.include_router(router)

if __name__ == "__main__":
    uvicorn.run(app, port=int(config['self_port']), host=config['self_host'])

