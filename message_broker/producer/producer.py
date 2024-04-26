from pika import exceptions
import pika
import os
import time



def main():
    host = os.environ.get('RABBITMQ_HOST')
    name = os.environ.get('RABBITMQ_NAME')

    try:
        connection = pika.BlockingConnection(pika.ConnectionParameters(host=host))
    except exceptions.AMQPConnectionError as err:
        print(host)
        print(err.args)
        return

    channel = connection.channel()
    channel.queue_declare(queue=name)

    while True:
        channel.basic_publish(exchange="",
                              routing_key="example",
                              body="example message")
        time.sleep(5)

    connection.close()

if __name__ == '__main__':
    main()
