from fastapi import APIRouter
import json
import httpx

router = APIRouter()
config = json.loads(open('config.json', 'r').read())

@router.get("/service_1/")
def service_1():
    res = httpx.get(f"http://localhost:{config['service_port_1']}/endpoint")
    return res.text

@router.get("/service_2/")
def service_2():
    res = httpx.get(f"http://localhost:{config['service_port_2']}/endpoint")
    return res.text
