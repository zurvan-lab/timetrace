import random
import string
import socket

def get_random_string_name(length: int) -> str:
    return ''.join(random.choice(string.ascii_lowercase) for i in range(length))

def make_query(query: str, sock: socket.socket) -> str:
    sock.sendall(query.encode())
    return sock.recv(1024).decode()
