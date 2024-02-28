# 07: FIRST Tests

Now we've completed one full Red-Green-Refactor TDD cycle, let's look into what makes for a good test.

We know tests are executable specifications. But what can we do to make them most effective?

Five rules helps us, neatly summarised by an acronym FIRST:

![FIRST Acronym](/chapter07/images/first.png)

## What FIRST means

- **Fast** Tests must run quickly. Milliseconds or less per test
- **Isolated** Tests can be run in isolation - any order or individually
- **Repeatable** Test results must be repeatable if a test is run multiple times
- **Self-validating** Tests must check their own results giving a PASS/FAIL
- **Timely** Tests must be written at the same time as production code

Let's look at what this means in practice.

### Fast tests are critical to success

Fast tests make TDD viable.

If we have a suite of tests that takes an hour to run, those tests will be run infrequently. Chances are they will never be run by developers at all. It's just human nature to do this.

That leads to test-first being abandoned in favour of code-and-fix-then-test. It is obvious, but bears repeating: if you don't use TDD you won't get any of its benefits.

> Having a bunch of tests you never run is pointless

The goal in TDD is to have tests that run quickly enough that we will run them often. As we saw in the previous chapter, this is often several times _as we write a single test/production code pair_.

A single test should be so trivially fast that ity could be run by the IDE automatically.

How fast are we talking? Milliseconds, typically.

We often encounter a slow (1-2 second) 'warm up' time when we run a test suite. That's not ideal, but so long as the rest of the tests run in milliseconds we will be ok.

> Tests run in milliseconds. This is a game changer

The whole test suite must be run before each commit to main trunk. This will only happen if the test suite as a whole is fast enough to not be annoying. This can take several seconds to low minutes. Anything slower again means commits to main will happen untested. Whenever we try to fight human nature, we will lose

> Don't fight human nature - Fix the system

We must design our code and tests to run quickly.

### Isolated tests enable fast, accurate feedback

Isolated tests mean that the test result will be the same no matter what order we run it in.

This allows us to focus on getting a single test to pass. We can also be confident that the test will not suddenly break once we run all the tests.

If we find that we can only run tests in a certain order, this is telling us something about our production code. It remembers state and the order of operations. If we need this, we must make that clear in our tests. It will be a very important part of the programming interface.

### Repeatable

Test results must be consistent to be useful

If we run a test multiple times, without changing any code, it must return the same result every time.

When this does not happen, it is called a **flaky test**.

Flaky tests are worse than useless. They have two major issues:

- They mask true faults
- They train developers to ignore test results

Avoiding flaky tests needs design work in both test and production code.

### Self-validating

A test which does not check the results of the Act step is **useless**.

It is not even a test.

### Timely

When we write a test is important.

Think about this timeline and the effect of when we write the test. The vertical arrow is the time when we are writing the production code for a single behaviour, perhaps one function or method.

![Timeliness of tests](/chapter06/images/timely.png)

#### Writing all tests before any production code

This results in highly speculative waterfall code.

- How do we know those are the behaviours we need?
- Why are we using unproven programming interface design decisions?
- Can we even build that system?
- When will we learn of errors and unknowns?
- How much delay and cost will rework involve?

Writing all the tests before writing any code feels like a 'modern' way of doing a waterfall project. It is a recipe for failure.

#### Writing all tests after all production code declared complete

The opposite end of the timeline results in _test theatre_.

It is far too late by then for the tests to have any real impact. They clearly cannot be executable specifications. They cannot be exploratory design aids. They are a way to smush together the words 'test automation' and 'test phase'.

By this time, all the testing will have been done by hand.

Tests as a separate step will lock the tests to whatever behaviour is there and whatever programming interface exists. They are not really acting as doubly-entry checks, nor design assistants. They will verify logic, but that's about all.

If you even remember to add all the tests needed for every edge case.

> You won't remember to add all the tests needed for every edge case

#### Writing one test just after the production code

Known as POUTing (Plain Old Unit Testing) this works better than writing all the tests in one go.

It denies us the benefit of thinking about required behaviour and implementation separately.

POUTing often locks-in suboptimal programming interfaces by mistake. It can seem too costly to go back and clean up the production code interface once we have a test in place and the code 'works'.

The design feedback we hope to get from TDD is weaker, and often ignored.

#### Writing one test as we write the production code

This is the test-first TDD approach we have been learning.

By writing the test in small micro-iterations, following the Red-Green-Refactor cycle, we maximise the value of writing the test. It serves as an executable specification. It provides feedback on ease of use of our chosen programming interface.

Significantly, feedback from a test written at this point happens at the point of _lowest cost of change_ of the production code. We can afford to make changes and experiment. We can afford to improve the design of our code to reduce the _total cost of ownership (TCO)_ of this code over its lifetime.

### What are the causes behind non-FIRST tests?

Given that FIRST tests make sense, they are surprisingly rare in our industry.

Why?

Let's look at pitfalls for each characteristic.

## Fast: What causes slow tests?

Top causes of slower tests:

- Code involving external systems such as databases, web services, cloud services
- End to end testing through the UI with complex page navigation

The common aspect here is that our production code contains logic that is fast to run, but it is coupled to a slow dependency:

![Slow dependency slows test](/chapter07/images/slow-dependency.png)

In our test Act step, we execute the fast logic in the production code. It is delayed by the slow response from the dependency. That ripples through, making the test itself slow.

We can avoid that by replacing the dependency with a faster simulation. We'll cover how to do that in later chapters.

## Isoloated: What prevents tests running in isolation?

Tests cannot be run in isolation if:

- Production code stores state (has memory of what happened before)
- Each test is written to depend on stored state from the previous test

Ways to avoid this include:

- Use stateless production code wherever sensible
- Require all test state to be in the Arrange section _only_
- Avoid "Before Each test" or similar styles of common Arrange code
- Have one test group per distinct setup - good for configurable code

## Reliable: What causes unreliable tests?

Non determinism in either the test harness or the production code will cause a different output for the same test conditions on each run.

Common causes are:

- Random number generators
- System clock (ie the time now)
- External systems with memory - databases, web services and more
- Data races - failures to synchronise concurrent programs correctly
- Selenium of old, which has the world's flakiets browser driver

We can solve non-determinism due to random sources, the system clock and externally stored data by using Test Doubles. These need a bit of software design magic, so we will describe these in a later chapter.

Failures such as incorrect concurrent programming simply have to be identified and then corrected.

> Testing concurrent code is _hard_

## Self-validating: When causes false positives?

You will see tests that either have their Assert section either missing or broken.

- Assert missing - nothing is checked
- Assert carelessly copy-pasted and not modified to the new expectation
- Assert is checking the wrong thing
- Act step had code in it, so the test was in effect testing itself
- Act step ran some dead code

Having a missing assert seems weird until you work on a dysfunctional project. Political forces pressurise you into "getting all tests green by close of play" and making the tests lie can be an attractive option. Needless to see, we don;t recommend this. But we have seen it done by other (lower-cost) consultancies. Let's be better.

> Don't worry, dear reader: you _will_ work on a dysfunctional project or two in this industry ...

There is an exception to "missing assert" that is valid. A compile failure is a valid "assertion fail". We can write a test using specific syntax that will fail to compile under the wrong version of language. That's ok, but a little weird and definitely needs a comment.

## Timely: When is it the wrong time to write a test?

We looked into the four options earlier. Writing a test just before - and in step with - production code provides the most valuable design feedback.

Other options exist, and provide more limited feedback.

## Next

If we have tests which are not FIRST, we don't have to stay stuck with them.

In the next chapter, we'll look into design patterns to get us back on the FIRST track.
