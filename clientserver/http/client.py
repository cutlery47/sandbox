import socket

class Client:

    def __init__(self, host="localhost", port=1234):
        self.server_host = host
        self.server_port = port

        self.client_sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    def main(self):
        self.client_sock.connect((self.server_host, self.server_port))

        while True:
            data = input("Press enter to send a get request")
            recieved = self.GET(self.server_host)

            print(recieved.decode())

    # http get-request
    def GET(self, url):
        request = f"GET / HTTP/1.1\r\nHost:{url}\r\n\r\n".encode()
        self.client_sock.send(request)
        return self.client_sock.recv(2048)

client = Client(host="neetcode.io", port=80)
client.main()
