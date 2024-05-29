# 11: TDD Anti-Patterns

If you've been in the industry for a while, you will have heard that using TDD "slows you down and makes the code twice as hard to work with".

When this happens, it's generally through using some anti-patterns.

Here are ten common problems that will Torpedo your TDD.

## 1 Fragile tests

"Dude - every single change we make causes loads of tests to break!"

- Reason: Testing implementation details instead of observable behaivour
- Symptom: It's not possible to change the implementation of the code under test without breaking tests.
- Solution: Test only the public behaviour of the component
- Workround: If we can;t do that for some reason, consider using a _dipstick method_

## 2 Testing the whole elephant

"These tests are huuuuuge - hundreds of lines!"

- Reason: A single test is testing too much.
- Symptom: Complex Arrange section, Multiple Act steps per test,Many Asserts per test
- Solution: Split tests to test only one thing
- Solution: Consider parameterised tests or extracting helper functions to reduce duplicated code in the tests. But DAMP tests are generally better.

Testing one thing _can_ use a custom assert to check several related results. A good example might be to check that all five fields of a User record are correct.

## 3 Tests with embedded logic

"I can't believe I'm having to debug the test code!"

- Reason: Tests have conditional/loop logic inside them
- Symptom: Difficult to follow what the test is doing
- Solution: Split tests into individual, more specific cases - remove the need for logic

## 4 Unclear test data

"Bro, I have no idea why this test expects 22.6 as the answer!"

- Reason: The actual data and logic under test are either implicit, or specified outside the test
- Symptom: Hard to see what the connection is between the inputs and expected
- Solution: Make tests explicit
- Consider: using explaining variables in the tests, making test data explicit, avoiding global test data setups

## 5 Mystery guest

"wtf???"

- Reason: The test relies on the state of something not controlled by the test
- Symptom: It is not possible to work out what the expected result is from the test case alone
- Solution: Make test self contained
- Consider: Use Dependency Inversion to allow simple stubbing, controlled inside the test itself

## 6 Flaky tests

"Oh, the team just ignores that test when it fails - it's usually ok"

- Reason: Test relies on some non-deterministic behaviour, like a race condition not happening, or a specific random value being generated
- Symptom: Test result is not repeatable on consecutive runs
- Solution: Remove the source of non-determinism either form the test or the component under test

## 7 Mocking what you don’t own

"Production is down, but all the tests in the pipeline are passing!"

- Reason: Tests rely on a mock of a third-party service. The third-party has changed their API. The mock has not been updated.
- Symptom: Tests pass, but using the code for real fails.
- Solution:

## 8 Testing the mock

"Why is that test so big when the component has no code in it?"

- Reason: A test is verifying behaviour of a mock object, and not the logic of the component under test
- Symptom: Component has basic logic (or none) but there are many complex tests
- Solution: Scope the test to only the logic inside the component
- Consider: Write a separate test for any complex Mock object or Fake you require

## 9 Tests only pass if they run in a specific order

"You have to use the script to run the tests, else they fail!"

- Reason: The tests are coupled by some shared state
- Symptom: Test pass when run in order, but fail when run out of order
- Solution: either (a) Remove that shared state, or (b) Make each test explicitly build up any required initial state

## 10 It's not possible to write a test

"Dude TDD sucks because we can;t even write tests for this code"

- Reason: Excess coupling in the component design
- Symptom: Tests become huge, spinning up and configuring the entire system
- Solution: Refactor to decouple the system
- Consider: applying ideas from the book _Working with Legacy Code_ by Michael Feathers
