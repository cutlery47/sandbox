import socket

recv_size = 2048

class ClientSocket(socket.socket):
    def __init__(self, host="localhost", port=9999):
        super().__init__(socket.AF_INET, socket.SOCK_STREAM)
        self.host = host
        self.port = port

    def connect_to_host(self):
        try:
            self.connect((self.host, self.port))
            return 0
        except socket.error as err:
            print(f"Error when connecting to server: {err}")
            return -1

    def terminate(self):
        self.close()


class ServerSocket(socket.socket):
    def __init__(self, host="localhost", port=9999):
        super().__init__(socket.AF_INET, socket.SOCK_STREAM)
        self.host = host
        self.port = port

    def create_host(self):
        self.bind((self.host, self.port))
        self.listen(5)

    def terminate(self):
        self.close()
