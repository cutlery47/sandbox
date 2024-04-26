from pika import exceptions
import pika
import os


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
    def callback(ch, method, properties, body):
        print(" [x] Received %r" % body)

    channel.basic_consume(queue=name,
                          auto_ack=True,
                          on_message_callback=callback)

    print(' [*] Waiting for messages. To exit press CTRL+C')
    channel.start_consuming()

if __name__ == '__main__':
    main()
