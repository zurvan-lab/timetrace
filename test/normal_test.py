import socket
import time
import utils
import random

#! GET method is not tested here now.
# TODO: writing test for method GET.
# TODO: adding test bad cases such as invalid set/subset name.
# TODO: popping items from global variables while dropping or cleaning them(?)

st = time.time()

#? Global variables
sub_set_names = []
set_names = []
elements_value = []

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

server_address = ('localhost', 7070) #! You can change port and host for testing.
sock.connect(server_address)

def test_connect_ok():
    """
    Connecting to database and creating a session.
    Testing: CON TQL command
    Error: null
    """
    query = "CON root super_secret_password" #! Considers you config is default.
    response = utils.make_query(query, sock)

    assert response == "OK", f"\033[91mCan't connect to server with default info, received: {response}, expected: OK\033[0m"

    print(f"\033[32mReceived response: {response}, Connection test Passed.\033[32m")

def test_ping_ok():
    """
    Ping the database to check session health.
    Testing: PING TQL command
    Error: null
    """
    response = utils.make_query("PING", sock)

    assert response == "PONG", f"\033[91mCan't ping the database, received: {response}, expected: PONG\033[0m"

    print(f"\033[32mReceived response: {response}, we are connected!\033[32m")

def test_new_set_ok():
    """
    Creating random sets.
    Testing: SET TQL command
    Error: null
    """
    for i in range(10):
        set_names.append(utils.get_random_string_name(i+2))

        query = f"SET {set_names[i]}"
        response = utils.make_query(query, sock)

        assert response == "OK", f"\033[91mCan't create set: {set_names[i]}, received: {response}, expected: OK\033[0m"

    print(f"\033[32mReceived response: {response}, {len(set_names)} sets created successfully.\nSets: {set_names}\033[32m")

def test_new_sub_set_ok():
    """
    Creating random subsets for sets.
    Testing: SSET command
    Error: null
    """
    for s in set_names:
        for i in range(7):
            sub_set_names.append(utils.get_random_string_name(i+2))

            query = f"SSET {s} {sub_set_names[i]}"
            response = utils.make_query(query, sock)

            assert response == "OK", f"\033[91mCan't create subset: {sub_set_names[i]} in set {s}, received: {response}, expected: OK\033[0m"

    print(f"\033[32mReceived response: {response}, {len(sub_set_names)} subsets created successfully.\nSubsets: {sub_set_names}\033[32m")

def test_push_element_ok():
    """
    Pushing randomly generated elements in all created subsets.
    Testing: PUSH TQL command
    Error: null
    """
    set_index = 0

    for s in set_names:
        for i in range(7):
            for _ in range(1_000):
                element_value = utils.get_random_string_name(i+8)
                elements_value.append(element_value)

                element_time = int(time.mktime(time.gmtime()))
                query = f"PUSH {s} {sub_set_names[i]} {element_value} {element_time}"
                response = utils.make_query(query, sock)

                assert response == "OK", f"\033[91mCan't push element with value of: {elements_value[i]} and time of: {element_time}, received: {response}, expected: OK\033[0m"

        set_index += 7

    #? Change `elements_value[:10]` to get more or less elements.
    print(f"\033[32mReceived response: {response}, {len(elements_value)} elements pushed successfully.\nElements: {elements_value[:10]}\033[32m")

def test_count_sets_ok():
    """
    Counting all sets.
    Testing: CNTS TQL command
    Error: null
    """
    response = utils.make_query("CNTS", sock)

    assert response == str(len(set_names)), f"\033[91mCan't count sets, received: {response}, expected: {len(set_names)}\033[0m"

    print(f"\033[32mReceived response: {response}, sets number counted successfully.\033[32m")

def test_count_sub_sets_ok():    
    """
    Counting all subsets.
    Testing: CNTSS TQL command
    Error: null
    """
    sub_sets_count = 0

    for s in set_names:
        query = f"CNTSS {s}"
        response = utils.make_query(query, sock)

        sub_sets_count += int(response)

    assert sub_sets_count == len(sub_set_names), f"\033[91mCan't count subsets, received: {sub_sets_count}, expected: {len(sub_set_names)}\033[0m"

    print(f"\033[32mReceived response: {sub_sets_count}, subsets counted successfully.\033[32m")

def test_count_elements_ok():
    """
    Counting all elements in all subsets.
    Testing: CNTE TQL command
    Error: null
    """
    set_index = 0
    elements_count = 0

    for s in set_names:
        for i in range(7):
            query = f"CNTE {s} {sub_set_names[i]}"
            response = utils.make_query(query, sock)

            elements_count += int(response)
        
        set_index += 7

    assert elements_count == len(elements_value), f"\033[91mCan't count elements, received: {elements_count}, expected: {len(elements_value)}\033[0m"

    print(f"\033[32mReceived response: {elements_count}, elements counted successfully.\033[32m")

def test_clean_sub_sets_elements_ok():
    """
    Cleaning all elements in all subsets.
    Testing: CLNSS TQL command
    Error: null
    """
    set_index = 0

    for s in set_names:
        for i in range(7):
            query = f"CLNSS {s} {sub_set_names[i]}"
            response = utils.make_query(query, sock)

            assert response == "OK", f"\033[91mCan't clean subset: {sub_set_names[i]} of set {s}, received: {response}, expected: OK\033[0m"

        set_index += 7

    print(f"\033[32mReceived response: {response}, subset elements cleaned successfully.\033[32m")

def test_drop_sub_sets_ok():
    """
    Dropping all subsets.
    Testing: DRPSS TQL command
    Error: null
    """
    set_index = 0

    for s in set_names:
        for i in range(7):
            query = f"DRPSS {s} {sub_set_names[i]}"
            response = utils.make_query(query, sock)

            assert response == "OK", f"\033[91mCan't drop subset: {sub_set_names[i]} from set: {s}, received: {response}, expected: OK\033[0m"
        
        set_index += 7

    print(f"\033[32mReceived response: {response}, subsets dropped successfully.\033[32m")

def test_clean_sub_sets_ok():
    """
    Cleaning all subsets in all sets.
    Testing: CLNS TQL command
    Error: null
    """
    for s in set_names:
        query = f"CLNS {s}"
        response = utils.make_query(query, sock)

        assert response == "OK", f"\033[91mCan't clean set: {s}, received: {response}, expected: OK\033[0m"

    print(f"\033[32mReceived response: {response}, sets cleaned successfully.\033[32m")

def test_drop_sets_ok():
    """
    Dropping all sets.
    Testing: DRPS TQL command
    Error: null
    """
    for s in set_names:
        query = f"DRPS {s}"
        response = utils.make_query(query, sock)

        assert response == "OK", f"\033[91mCan't drop set: {s}, received: {response}, expected: OK\033[0m"

    print(f"\033[32mReceived response: {response}, sets dropped successfully.\033[32m")

def test_clean_sets_ok():
    """
    Cleaning all sets.
    Testing: CLN TQL command
    Error: null
    """
    response = utils.make_query("CLN", sock)

    assert response == "OK", f"\033[91mCan't clean sets, received: {response}, expected: OK\033[0m"

    print(f"\033[32mReceived response: {response},sets cleaned successfully.\033[32m")


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
    print('\033[34mAll tests successfully passed in:\033[34m', time.time() - st, 'seconds')
