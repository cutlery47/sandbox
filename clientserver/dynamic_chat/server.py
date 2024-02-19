from utils import ServerSocket, recv_size
from time import sleep
import threading


class Server:
    def __init__(self):
        self.__socket = ServerSocket()
        self.is_active = False
        self.connections = []

    def handle_connections(self):
        while self.is_active:
            client_sock, client_addr = self.__socket.accept()
            print(f"{client_addr} has connected!")

            messages_handler = threading.Thread(target=self.handle_messages, args=[client_sock])
            messages_handler.start()

            sleep(1)

    def handle_messages(self, client):
        while self.is_active:
            message = client.recv(recv_size).decode()
            if message == '':
                return
            print(f"{message}")
            sleep(0.5)

    def activate(self):
        self.is_active = True
        self.__socket.create_host()

        connections_handler = threading.Thread(target=self.handle_connections)
        connections_handler.start()

    def disconnect(self):
        self.__socket.terminate()

if __name__ == "__main__":
    server = Server()
    server.activate()