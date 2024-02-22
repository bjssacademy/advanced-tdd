# 07: FIRST Tests

Now we've completed one full Red-Green-Refactor TDD cycle, let's look into what makes for a good test.

We know tests are executable specificationsd. But what can we do to make them execute well?

Five rules helps us, neatly summarised by an acronym FIRST:

- **Fast** Tests must run quickly. Milliseconds or less per test
- **Isolated** Tests can be run in isolation - any order or individually
- **Repeatable** Test results must be repeatable if a test is run multiple times
- **Self-validating** Tests must check their own results giving a PASS/FAIL
- **Timely** Tests must be written at the same time as production code

## What FIRST means

Let's look at what this means for developers and the TDD workflow.

### Fast tests are critical to success

Fast tests make TDD viable.

If we have a suite of tests that takes an hour to run, those tests will be run very infrequently. That leads to test-first being abandoned in favour of code-and-fix-then-test. It's just human nature to do this.

The goal in TDD is to have tests that run so quickly, we can run them multiple times as we develop the code.

A single test should be so trivially fast that ity could be run by the IDE automatically.

The whole test suite must be run before each commit to main trunk. This will only happen if the test suite as a whole is fast enough to not be annoying. This can take several seconds to low minutes. Anything slower again means commits to main will happen untested. Whenever we try to fight human nature, we will lose

> Don't fight human nature - Fix the system

We must design our code and tests to run quickly.

### Isolated tests enable fast, accurate feedback

Isolated tests mean that the test result will be the same no matter what order we run it in.

This enables focus on a single test as we write production code in one RGR cycle. Once we get to green on that test, we can run the whole suite and confirm we haven't broken anything by mistake.

We need isolated tests to do this. If we can only run tests in a particular order, we will not be able to work on making a single test passed during RGR.

Inaccurate results often happen when tests have a hidden 'correct' order - technically, tests are coupled to each other by memorised program state.

Running a test in these conditions can generate a false pass. For example, some setup code is not run in isolation, but only when all tests are run. Production code will then be driven out for this incorrect setup. It will pass the test in isolation _and_ be incorrect for the system as a whole.

Tests should be independent of coupling. This will drive out production code to handle state correctly.

### Repeatable

Test results must be consistent to be useful

If we run a test multiple times, without changing any code, it must return the same result every time.

When this does not happen, it is called a **flaky test**.

Flaky tests are worse than useless. They have two major issues:

- They mask true fault
- They train developers to ignore tests

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

It is far too late by then for the tests to have any real impact. they clearly cannot be executable specifications. They cannot be exploratory design aids. They are a way to smuch together the words 'test automation' and 'test phase'.

By this time, all the testing will have been done by hand.

Having a phase of writing automated tests without feedback into the production code ends up tying the tests to whatever behaviour is there and whatever programming interface exists. They are not really acting as doubly-entry checks, nor design assistants.

#### Writing one test just after the production code

Known as POUTing (Plain Old Unit Testing) this works better than writing all the tests in one go.

It denies us the benefit of thinking about required behaviour and implementation separately.

POUTing often locks-in suboptimal programming interfaces by mistake. It can seem too costly to go back and clean up the production code interface once we have a test in place and the code 'works'.

The design feedback we hope to get from TDD is weaker, and often ignored.

#### Writing one test as we write the production code

This is the test-first TDD approach we have been learning.

By writing the test in small micro-iterations, following the Red-Green-Refactor cycle, we maximise the value of writing the test. It serves as an executable specification. It provides feedback on ease of use of our chosen programming interface.

Significantly, feedback from a test written atthis point happens at the point of _lowest cost of change_ of the production code. We can afford to make changes and experiment. We can afford to improve the design of our code to reduce the _total cost of ownership (TCO)_ of this code over its lifetime.

### What are the causes behind non-FIRST tests?

The next section presents some common technical and business drivers that frustrate writing FIRST tests. We will cover solutions in the following chapters.

## Fast: What causes slow tests?

Given that FIRST tests make sense, they are surprisingly rare in our industry.

Why?

Let's look at pitfalls for each characteristic.

## Isoloated: What prevents tests running in isolation?

## Reliable: What causes unreliable tests?

## Self-validating: When causes false positives?

## Timely: When is it the wrong time to write a test?

## Next

If we have tests which are not FIRST, we don't have to stay stuck with them.

In the next chapter, we'll look into design patterns to get us back on the FIRST track.
