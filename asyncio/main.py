import asyncio

async def helloworld():
    print("hello")
    await helloworld()

loop = asyncio.