import socket
import threading


class Server:
    def __init__(self, host="localhost", port=1234):
        # creting TCP internet socket
        self.host = host
        self.port = port

        self.server_sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.server_sock.bind((self.host, self.port))
        self.server_sock.listen(5)

    def main(self):
        while True:
            client_sock, client_addr = self.server_sock.accept()

            worker = threading.Thread(target=self.work, daemon=True, args=(client_sock, client_addr))
            worker.start()

    def work(self, client_sock: socket.socket, client_addr):
        while True:
            data = client_sock.recv(2048).decode()

            if not data:
                return

            print("data recieved from", client_addr)
            print("the data:", data)

            method, path, protocol, headers, body = self.handleHttp(data)
            response = self.processResponse("hi client")

            client_sock.send(response.encode())

    def handleHttp(self, http):
        request, *headers, _, body = http.split("\r\n")
        method, path, protocol = request.split(" ")
        headers = dict(
            line.split(":", maxsplit=1) for line in headers
        )
        return method, path, protocol, headers, body

    def processResponse(self, response):
        return (
            "HTTP/1.1 200 OK\r\n"
            f"Content-Length: {len(response)}\r\n"
            "Content-Type: text/html\r\n"
            f"\r\n{response}\r\n"
        )

server = Server()
server.main()