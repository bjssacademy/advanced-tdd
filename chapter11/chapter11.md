# 11: TDD Mistakes

1 Fragile tests
Implementation detail
Behaviour of public APIs as per consumer. Split out private stuff into more focussed public
Dipstick methods - if best we can do

2 Testing the whole elephant
Too many acts, asserts
One logical assert per test

- Explain logical assert

3 Tests with logic
Tests themselves have bugs in them!
Remove logic, split to separate tests or parameterised

4 Unclear test data
Hard to see what the connection is between the inputs and expected
Hide the irrelevant and emphasise the important: explaining variables and other refactoring steps on the test

5 Mystery guest
Test is not self-contained. We must look outside the test to understand initial conditions. Examples: global test stubs, complex fakes, test/live env databases
Make self contained. Use DIP to allow simple stubbing

6 Flaky tests
Test rely on some non deterministic behaviour: timing, sources of randomness
Remove non determinism from the test

7 Mocking what you don’t own

8 Testing the double not the component
