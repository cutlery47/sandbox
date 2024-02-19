from utils import ServerSocket, Message, recv_size
from time import sleep
import threading


class Server:
    def __init__(self):
        self.__socket = ServerSocket()
        self.is_active = False

    def handle_connections(self):
        while self.is_active:
            client_sock, client_addr = self.__socket.accept()
            print(f"{client_addr} has connected!")

            messages_handler = threading.Thread(target=self.handle_messages, args=[client_sock])
            messages_handler.start()

    def handle_messages(self, client):
        while self.is_active:
            msg = client.recv(recv_size).decode()
            print(f"{message}")
            self.respond(message, client)

    def respond(self, message):
        pass

    def activate(self):
        self.is_active = True
        self.__socket.create_host()

        connections_handler = threading.Thread(target=self.handle_connections)
        connections_handler.start()

        connections_handler.join()
        self.disconnect()

    def disconnect(self):
        self.__socket.terminate()

if __name__ == "__main__":
    server = Server()
    server.activate()