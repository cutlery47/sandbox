from utils import ServerSocket, MessageSerializer, recv_size
from queue import Queue
import threading


class Server:
    def __init__(self):
        self.__socket = ServerSocket()
        self.is_active = False
        self.message_q = Queue()

    def handle_connections(self):
        while self.is_active:
            client_sock, client_addr = self.__socket.accept()
            print(f"{client_addr} has connected!")

            messages_handler = threading.Thread(target=self.handle_messages, args=[client_sock])
            responses_handler = threading.Thread(target=self.handle_responses, args=[client_sock])

            messages_handler.start()
            responses_handler.start()

    def handle_messages(self, client):
        while self.is_active:
            msg = MessageSerializer.decode(client.recv(recv_size))
            self.message_q.put(msg)
            print(msg)

    def handle_responses(self, client):
        while self.is_active:
            if self.message_q.not_empty:
                msg = self.message_q.get()
                client.send(MessageSerializer.encode(msg))

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
