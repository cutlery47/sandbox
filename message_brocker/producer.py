import pika

connection = pika.BlockingConnection(pika.ConnectionParameters(host="localhost"))
channel = connection.channel()

channel.queue_declare(queue="example")
channel.basic_publish(exchange="",
                      routing_key="example",
                      body="example message")

connection.close()

