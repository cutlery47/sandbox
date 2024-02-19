from utils import ClientSocket, recv_size
from time import sleep
import threading


class Client:
    def __init__(self):
        self.name = input("What is your username? ")
        self.__socket = ClientSocket()
        self.connected = False

    def establish_connection(self):
        while not self.connected:
            status = self.__socket.connect_to_host()
            if status == 0:
                self.connected = True
            sleep(1)

    def handle_response(self):
        while self.connected:
            response = self.__socket.recv(recv_size)
            if response == '':
                self.connected = False
            else:
                print('here')
                print(response.decode())

    def handle_input(self):
        while self.connected:
            message = input()
            if message == 'quit':
                self.connected = False
            self.__socket.send(self.encode_message(message))

    def encode_message(self, message):
        message = message + "|" + self.name
        return message.encode()

    def activate(self):
        self.establish_connection()

        input_handler = threading.Thread(target=self.handle_input)
        response_handler = threading.Thread(target=self.handle_response)

        input_handler.start()
        response_handler.start()
        print('here')
        input_handler.join()
        response_handler.join()

    def disconnect(self):
        self.__socket.terminate()

if __name__ == "__main__":
    client = Client()
    client.activate()
