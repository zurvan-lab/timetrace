import socket
import time
import utils
import random

#! GET method is not tested here now.
# TODO: writing test for method GET.
# TODO: adding test bad cases such as invalid set/subset name.

st = time.time()

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

server_address = ('localhost', 7070) # You can change port and host for testing.
sock.connect(server_address)

#? TQL command to connect to the server.
query = "CON root super_secret_password" #! Considers you config is default.
response = utils.make_query(query, sock)

assert response == "OK"
print(f"Received response: {response}, Connection test Passed!")

#? Ping database.
assert utils.make_query("PING", sock) == "PONG"
print(f"Received response: {response}, We are connected!")

#* Creating some random sets.
set_names = []
for i in range(10):
    set_names.append(utils.get_random_string(i+2))
    query = f"SET {set_names[i]}"
    
    response = utils.make_query(query, sock)
    assert response == "OK"

print(f"Received response: {response}, Created {len(set_names)} sets successfully.")

#* Creating some random sub sets.
sub_set_names = []
for s in set_names:
    for i in range(7):
        sub_set_names.append(utils.get_random_string(i+2))

        query = f"SSET {s} {sub_set_names[i]}"
        response = utils.make_query(query, sock)
        assert response == "OK"

print(f"Received response: {response}, Created {len(sub_set_names)} sub sets successfully.")

#* Pushing some random elements.
set_index = 0
elements_value = []
for s in set_names:
    for i in range(7):
        for _ in range(1_000):
            element_value = utils.get_random_string(i+8)
            elements_value.append(element_value)

            query = f"PUSH {s} {sub_set_names[i]} {element_value} {int(time.mktime(time.gmtime()))}"
        
            response = utils.make_query(query, sock)
            assert response == "OK"
    
    set_index += 7

print(f"Received response: {response}, Created {len(elements_value)} elements pushed successfully.")

#* Test count
response = utils.make_query("CNTS", sock)
assert response == str(len(set_names))
print(f"Received response: {response}, Sets number counted successfully.")

sub_sets_count = 0
for s in set_names:
    query = f"CNTSS {s}"
    response = utils.make_query(query, sock)
    sub_sets_count += int(response)

assert sub_sets_count == len(sub_set_names)
print(f"Received response: {sub_sets_count}, SubSets number counted successfully.")

set_index = 0
elements_count = 0
for s in set_names:
    for i in range(7):
        query = f"CNTE {s} {sub_set_names[i]}"
        
        response = utils.make_query(query, sock)
        elements_count += int(response)
    
    set_index += 7

assert elements_count == len(elements_value)
print(f"Received response: {elements_count}, Elements number counted successfully.")

#* Test clean sub sets
set_index = 0
for s in set_names:
    for i in range(7):
        query = f"CLNSS {s} {sub_set_names[i]}"
        
        response = utils.make_query(query, sock)
        assert response == "OK"
    
    set_index += 7

print(f"Received response: {response}, All Subset elements cleaned successfully.")


#* Test drop subsets
set_index = 0
for s in set_names:
    for i in range(7):
        query = f"DRPSS {s} {sub_set_names[i]}"
        
        response = utils.make_query(query, sock)
        assert response == "OK"
    
    set_index += 7

print(f"Received response: {response}, All Subset dropped successfully.")

#* Test clean sets
for s in set_names:
    query = f"CLNS {s}"

    response = utils.make_query(query, sock)
    assert response == "OK"    

print(f"Received response: {response}, All Subsets cleaned successfully.")

#* Test drop sets
for s in set_names:
    query = f"DRPS {s}"

    response = utils.make_query(query, sock)
    assert response == "OK"    

print(f"Received response: {response}, All Sets dropped successfully.")

#* Clean all sets.
response = utils.make_query("CLN", sock)
assert response == "OK"
print(f"Received response: {response}, All Sets cleaned successfully.")

sock.close()
print('All test successfully passed in:', time.time() - st, 'seconds')
