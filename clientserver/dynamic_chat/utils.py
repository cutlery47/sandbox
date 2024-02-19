from datetime import datetime
import socket

recv_size = 2048


class Message:
    def __init__(self, text: str = "", sender: str = "", timestamp: str = datetime.now().strftime("%H:%M:%S %d/%m/%Y")):
        try:
            MessageValidator.validate(text, sender)
        except ValueError:
            raise ValueError("Validation not passed")

        self.text = text
        self.sender = sender
        self.timestamp = timestamp

    def __str__(self):
        return f"{self.sender} sent: {self.text} at {self.timestamp}"


class MessageValidator:
    @classmethod
    def validate(cls, text, sender):
        if text.find('~~~') >= 0:
            raise ValueError("Message not allowed (\"~~~\")")
        if sender.find('~~~') >= 0:
            raise ValueError("Username not allowed (\"~~~\")")


class MessageSerializer:
    @classmethod
    def encode(cls, message: Message) -> bytes:
        text = message.text
        sender = message.sender
        date = message.timestamp

        return (text + "~~~" + sender + "~~~" + date).encode()

    @classmethod
    def decode(cls, encoded_message: bytes) -> Message:
        message = Message()

        decoded_message = encoded_message.decode()
        decoded_parts = decoded_message.split("~~~")

        message.text = decoded_parts[0]
        message.sender = decoded_parts[1]
        message.timestamp = decoded_parts[2]

        return message


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
