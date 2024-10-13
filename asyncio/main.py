import asyncio


async def hello():
    await asyncio.sleep(1)
    print("hello")

async def world():
    await asyncio.sleep(1)
    print("world")

async def main():

    task_1 = asyncio.create_task(hello())
    task_2 = asyncio.create_task(world())

    await asyncio.gather(task_1, task_2)

    print(task_1.done())

    print("123123123123123")

asyncio.run(main())

# loop = asyncio.new_event_loop()
# task = loop.create_task(main())
# loop.

