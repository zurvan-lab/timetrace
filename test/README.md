# Tests

This directory contains a set of python scripts for testing time-trace database. Each change or pull request must pass this test otherwise it won't be merged.

## Checklist

Here is a list of different tests which must be added here and their status:

| Test | Status |
| -------- | ------- |
| Normal test | ✅ |
| Concurrent test | ⚒️ |

## Test writing guide

All of test must be in form of functions. The functions name must start with keyword `test` and finish with the result of test such as `ok`, `ssnf`, `snf`, `invalid` and so on.

Example of function name:

Correct: `test_new_sub_set_ok`
Incorrect: `new_sub_set_test`
Incorrect: `new_subset_test`
Incorrect: `test_new_sub_set`

> [!NOTE]
> Consider following python naming convention and code style.

Since in we use a connection a server with once instance, consider that sometimes order of commands can affect the result of test so make sure you do them in correct order and restart database each time. For example you can't clean a set when you dropped it!

> [!IMPORTANT]
> Codes in this directory are not about benchmark, these are only for testing functionality of project, for benchmark please check benchmark directory.
