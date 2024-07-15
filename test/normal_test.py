import socket
import time
import utils
import random

#! GET method is not tested here now.
# TODO: writing test for method GET.
# TODO: adding test bad cases such as invalid set/subset name.

st = time.time()

#? Global variables
sub_set_names = []
set_names = []
elements_value = []

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

server_address = ('localhost', 7070) #! You can change port and host for testing.
sock.connect(server_address)

def test_connect_ok():
    #* TQL command to connect to the server.
    query = "CON root super_secret_password" #! Considers you config is default.
    response = utils.make_query(query, sock)

    assert response == "OK", f"\033[91mCan't connect to server with correct info, received {response}, expected OK\033[0m"

    print(f"\033[32mReceived response: {response}, Connection test Passed.\033[32m")

def test_ping_ok():
    #* Ping database.
    response = utils.make_query("PING", sock)

    assert response == "PONG", f"\033[91mCan't ping the database, received {response}, expected PONG\033[0m"

    print(f"\033[32mReceived response: {response}, We are connected!\033[32m")

def test_new_set_ok():
    #* Creating some random sets.
    for i in range(10):
        set_names.append(utils.get_random_string_name(i+2))

        query = f"SET {set_names[i]}"
        response = utils.make_query(query, sock)

        assert response == "OK", f"\033[91mCan't create set {set_names[i]}, received {response}, expected OK\033[0m"

    print(f"\033[32mReceived response: {response}, Created {len(set_names)} sets successfully.\033[32m")

def test_new_sub_set_ok():
    #* Creating some random sub sets.
    for s in set_names:
        for i in range(7):
            sub_set_names.append(utils.get_random_string_name(i+2))

            query = f"SSET {s} {sub_set_names[i]}"
            response = utils.make_query(query, sock)

            assert response == "OK", f"\033[91mCan't create sub set {sub_set_names[i]}, received {response}, expected OK\033[0m"

    print(f"\033[32mReceived response: {response}, Created {len(sub_set_names)} sub sets successfully.\033[32m")

def test_push_element_ok():
    #* Pushing some random elements.
    set_index = 0
    for s in set_names:
        for i in range(7):
            for _ in range(1_000):
                element_value = utils.get_random_string_name(i+8)
                elements_value.append(element_value)

                query = f"PUSH {s} {sub_set_names[i]} {element_value} {int(time.mktime(time.gmtime()))}"
                response = utils.make_query(query, sock)

                assert response == "OK", f"\033[91mCan't push element {elements_value[i]}, received {response}, expected OK\033[0m"

        set_index += 7

    print(f"\033[32mReceived response: {response}, Created {len(elements_value)} elements pushed successfully.\033[32m")

def test_count_sets_ok():
    #* Test counting sets
    response = utils.make_query("CNTS", sock)

    assert response == str(len(set_names)), f"\033[91mCan't count sets, received {response}, expected {len(set_names)}\033[0m"

    print(f"\033[32mReceived response: {response}, Sets number counted successfully.\033[32m")

def test_count_sub_sets_ok():    
    #* Test Counting sub sets
    sub_sets_count = 0

    for s in set_names:
        query = f"CNTSS {s}"
        response = utils.make_query(query, sock)

        sub_sets_count += int(response)

    assert sub_sets_count == len(sub_set_names), f"\033[91mCan't count sub sets, received {sub_sets_count}, expected {len(sub_set_names)}\033[0m"

    print(f"\033[32mReceived response: {sub_sets_count}, SubSets number counted successfully.\033[32m")

def test_count_elements_ok():
    #* Test counting all elements
    set_index = 0
    elements_count = 0

    for s in set_names:
        for i in range(7):
            query = f"CNTE {s} {sub_set_names[i]}"
            response = utils.make_query(query, sock)

            elements_count += int(response)
        
        set_index += 7

    assert elements_count == len(elements_value), f"\033[91mCan't count elements, received {elements_count}, expected {len(elements_value)}\033[0m"

    print(f"\033[32mReceived response: {elements_count}, Elements number counted successfully.\033[32m")

def test_clean_sub_sets_elements_ok():
    #* Test clean sub sets
    set_index = 0

    for s in set_names:
        for i in range(7):
            query = f"CLNSS {s} {sub_set_names[i]}"
            response = utils.make_query(query, sock)

            assert response == "OK", f"\033[91mCan't clean sub sets, received {response}, expected OK\033[0m"

        set_index += 7

    print(f"\033[32mReceived response: {response}, All Subset elements cleaned successfully.\033[32m")

def test_drop_sub_sets_ok():
    #* Test drop subsets
    set_index = 0

    for s in set_names:
        for i in range(7):
            query = f"DRPSS {s} {sub_set_names[i]}"
            response = utils.make_query(query, sock)

            assert response == "OK", f"\033[91mCan't drop sub sets, received {response}, expected OK\033[0m"
        
        set_index += 7

    print(f"\033[32mReceived response: {response}, All Subsets dropped successfully.\033[32m")

def test_clean_sub_sets_ok():
    #* Test clean sets
    for s in set_names:
        query = f"CLNS {s}"
        response = utils.make_query(query, sock)

        assert response == "OK", f"\033[91mCan't clean sets, received {response}, expected OK\033[0m"

    print(f"\033[32mReceived response: {response}, All Subsets cleaned successfully.\033[32m")

def test_drop_sets_ok():
    #* Test drop sets
    for s in set_names:
        query = f"DRPS {s}"
        response = utils.make_query(query, sock)

        assert response == "OK", f"\033[91mCan't drop sets, received {response}, expected OK\033[0m"

    print(f"\033[32mReceived response: {response}, All Sets dropped successfully.\033[32m")

def test_clean_sets_ok():
    #* Clean all sets.
    response = utils.make_query("CLN", sock)

    assert response == "OK", f"\033[91mCan't clean sets, received {response}, expected OK\033[0m"

    print(f"\033[32mReceived response: {response}, All Sets cleaned successfully.\033[32m")


def main():
    test_connect_ok()
    test_ping_ok()
    test_new_set_ok()
    test_new_sub_set_ok()
    test_push_element_ok()
    test_count_sets_ok()
    test_count_sub_sets_ok()
    test_count_elements_ok()
    test_clean_sub_sets_elements_ok()
    test_drop_sub_sets_ok()
    test_clean_sub_sets_ok()
    test_drop_sets_ok()
    test_clean_sets_ok()

if __name__ == "__main__":
    main()
    sock.close()
    print('\033[34mAll test successfully passed in:\033[34m', time.time() - st, 'seconds')
