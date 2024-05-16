# 06: Moving to the next RGR cycle

TDD is all about forward motion, in tiny steps.

After one RGR iteration, our reward for such good work is ... more work!

The question is, of course, _what work_? What should we tackle next for our next test?

![Good work gets more work](images/more-work.png)

> It was ever thus

This chapter offers some useful ideas on thinking about how we drive our design forward, with intent.

## What will drive the design best?

We use tests to drive our software design as well as the order we add details to our implementation. We are composing the overall solution from small steps, one test at a time.

### Start with a happy path

The best test to write first is a _happy path_ test. This means the simplest example of the software being used successfully.

Take code to calculate the total price of items in a shopping basket. The simplest happy path to check is that a basket with one item in it calculates the total as being the price of that one item:

```
TestTotalForSingleItem
```

This might look like this:

```
func TestTotalForSingleItem(t *testing.T) {

    basket := basket.New()
    basket.add( "Pizza", 495 )

    got := basket.total()

    want := 495
    if got != want {
        t.ErrorF("got %d, want %d", got, want)
    }
}
```

We check that the single item priced at 495 results in that figure as the total.

### Avoid weird stuff to start with

By _weird stuff_, we mean error cases, boundary conditions and any other complex rule.

All these things need developing, of course. But if we _start_ there, we tend to bend our design out of shape. The code no longer succinctly describes the expected normal behaviour. Instead, it reads out as being mainly a list of stuff that can go wrong.

Starting withthe simplest happy path ensures that our implementation has a sane foundation.

But after this, what next?

## Triangulation - iterating towards complete behaviour

describe

more specific tests - more generalised code

## ZOMBIES - inspiring triangulation

notes for rest--------->
Small steps
100% legit coverage
Start happy then zombies

BIG SUB HEAD: WHEN to change or delete tests

HOW to migrate an interface step by step

Eg PlaceShip( row col ) -> PlaceShip( Location ) work through

BIG SUB HEAD LAYERED TESTS and test pyramid

- Avoiding dupe tests
- Ideally one test per behaviour
- Balance close-in to boundaries
- Talk about test layer being a fixed point
