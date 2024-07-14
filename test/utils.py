import random
import string

def get_random_string(length):
    return ''.join(random.choice(string.ascii_lowercase) for i in range(length))

def make_query(query, sock):
    sock.sendall(query.encode())
    return sock.recv(1024).decode()
