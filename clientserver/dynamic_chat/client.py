from utils import (ClientSocket, recv_size, Message,
                   MessageSerializer)
from time import sleep
from datetime import datetime
import threading


class Client:
    def __init__(self):
        self.name = input("What is your username? ")
        self.__socket = ClientSocket()
        self.connected = False

    def establish_connection(self):
        while not self.connected:
            if self.__socket.connect_to_host() == 0:
                self.connected = True
            sleep(0.5)

    def handle_response(self):
        while self.connected:
            response = MessageSerializer.decode(self.__socket.recv(recv_size))
            if response.text == 'quit':
                self.connected = False
            print(response)

    def handle_input(self):
        while self.connected:
            message_text = input()
            message = Message(message_text, self.name, datetime.now().strftime("%H:%M:%S %d/%m/%Y"))
            self.__socket.send(MessageSerializer.encode(message))

    def activate(self):
        self.establish_connection()

        input_handler = threading.Thread(target=self.handle_input)
        response_handler = threading.Thread(target=self.handle_response)

        input_handler.start()
        response_handler.start()

        input_handler.join()
        response_handler.join()
        self.disconnect()

    def disconnect(self):
        self.__socket.terminate()

if __name__ == "__main__":
    client = Client()
    client.activate()
