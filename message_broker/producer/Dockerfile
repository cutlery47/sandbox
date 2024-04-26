FROM python:3.10-slim
WORKDIR /app
COPY . .

RUN python3 -m pip install pika

CMD ["python3", "producer.py"]