# 06: Moving to the next RGR cycle

TDD is all about forward motion, in tiny steps.

After one RGR iteration, our reward for such good work is ... more work!

The question is, of course, _what work_? What should we tackle next for our next test?

![Good work gets more work](images/more-work.png)

> It was ever thus

This chapter offers some useful ideas on thinking about how we drive our design forward, with intent.

## What will drive the design best?

With TDD, we use tests to drive our software design. We use the _order_ of tests we implement to control which details we add when. We are composing the overall solution from small steps, one test at a time.

The order we take these steps often leaves an imprint in our implementation. Some orders give better results than others.

### Start with a happy path

Where should we start? A simple _happy path_.

A happy path means the simplest example of the software being used successfully. Writing a test to specify this simplest case will drive out basic, useful functionality.

We can look at a shopping basket example.

We want calculate the total price of items in a shopping basket. The simplest happy path to check is that a basket with one item. The total must be the price of that item.

Our test is then:

```
TestTotalSingleItem
```

The code for that test might look like this:

```golang
func TestTotalSingleItem(t *testing.T) {

    basket := basket.New()
    basket.add( "Pizza", 495 )

    got := basket.total()

    want := 495
    if got != want {
        t.ErrorF("got %d, want %d", got, want)
    }
}
```

We check that the single item priced at 495 results in that figure as the total. We've specified that we will use an OOP approach, with a `total()` method on a basket object.

Our initial implementation can simply return the value 495 from `total()`. That gets our programming interface locked in, and our first test passing.

### Avoid weird stuff to start with

By _weird stuff_, we mean error cases, boundary conditions and any other complex rules.

All these things need developing, of course. Just not yet. If we _start_ there, we tend to bend our design out of shape. The code no longer succinctly describes the expected normal behaviour. Instead, it reads out as being mainly a list of stuff that can go wrong.

Starting with the simplest happy path makes for a sane foundation.

But after our first happy path, what next?

## Triangulation - iterating towards complete behaviour

In theory, we can write tests in any order, implement the code in any order, and clean up during the refactor steps.

In practice, adopting a structured appraoch can help us.

### Coding ZOMBIES

A useful approach is named Zombies, by James Grenning. The original article is [here](https://blog.wingman-sw.com/tdd-guided-by-zombies).

![zombies graphic](images/zombies.png)

ZOMBIES is a mnemonic:

- Z – Zero
- O – One
- M – Many (or More complex)
- B – Boundary Behaviors
- I – Interface definition
- E – Exercise Exceptional behavior
- S – Simple Scenarios, Simple Solutions

What do they mean?

#### Zero, One, Many/More (Z, O, M)

Following our happy path implementation - the One(O) case - we can work outwards in terms of quantity.

If we have one item in our basket, we will most likely also have many items in our basket.

Our next test can then be:

```
TestTotalMultipleItems
```

Writing a test for many (M) items makes for a sensible development. Our test will add three items to the basket and calculate the total.

This will drive out a more detailed data structure to hold more-than-one item (a slice in Go language). It also will drive the more advanced algorithm, that can total multiple items.

In a similar vein, the basket can be empty and the total should be zero. Coding for zero (Z) items rounds out our treatment of all the possibilities of quantity. We have worked in small steps, adding details iteratively.

(M) can also stand for More Complex. With our basket, we might add more complex logic to handle 2-for-1 discounts. This feature would give rise to a number of detailed tests to drive out that logic.

#### Boundary behaviours (B)

Boundary behaviours, also termed _edge cases_, refers to what the systems does when it hits a limit.

For our shopping basket, we might have a requirement to limit to 100 items maximum. In this case, we have a new boundary behaviour. When we attempt to add the 101st item, the `add()` method ought to fail.

This needs a new test. That means we need a piece of design thinking: How should we report this error to the calling code?

Once we decide, we can code up the test, then use that to drive an implementation.

> Q: How would you choose to report the error?

#### Interface definition

Paying attention to our interface design is something we do at each and every test.

The question to answer is: _is this easy to use?_

Imagine a shopping basket design where the test looked like this:

```golang
func TestTotalForTwoItems(t *t.Testing) {
    // Arrange
    additionEngine := AdditionEngine{algorithm: "inorder", rounding:2, maxItems:100}
    itemProvider := ItemProvider{}
    basket := Basket{additionEngine: additionEngine, itemProvider: itemProvider}

    // Act
    areMore := basket.CheckMoreItems()
    basket.FetchAndAddNextItem(areMore)

    areMore =  basket.CheckMoreItems()
    basket.FetchAndAddNextItem(areMore)

    total := basket.RoundAsPerRoundingStrategy(basket.Total)

    // Assert
    // code not listed ...
}
```

Now, that _is_ a programming interface to a shopping basket. And in much the same way, a photo of a fetid pile of dingo's excrement could be considered 'wildlife photography'. Anyway, I digress.

A core part of keeping software simple is to _amplify the essential, hide the irrelevant_.

What could we hide above?

- We could make basket create its own 'addition engine' with usable default setings
- If we continued with the "Item Provider" approach, we could pass only that in to a 'contructor' method
- We could siplify the next of `FetchAndAddNextItem` to `NextItem'
- 'NextItem' could do its own check for more items
- We can add a method to get the final total, including rounding

The improved code (and it is, really) would look like this:

func TestTotalForTwoItems(t \*t.Testing) {
// Arrange
itemProvider := ItemProvider{}
basket := Basket.New(itemProvider)

    // Act
    basket.NextItem()
    basket.NextItem()

    total := basket.TotalPrice()

    // Assert
    // code not listed ...

}

```

That has much less cognitive loading to figure out what's going on.

We could take it further:


func TestTotalForTwoItems(t *t.Testing) {
    // Arrange
    itemProvider := ItemProvider{}
    basket := Basket.New(itemProvider)

    // Act
    total := basket.TotalAllItems()

    // Assert
    // code not listed ...
}
```

The single method `TotalAllItems()` will fetch all items, total them, apply rounding and return the final answer.

You will notice that we've only shown the tests here. Interface design is a _design_ activity, and is fully captured in the test code. How it gets implemented is of no concern at the design stage.

> Always simplify the interface to your code. Make it easy to use in the wider program.

#### Exceptional behaviours

Exceptional behaviours are responses to error conditions, often unplanned. They include things like network connection loss, storage full and not available, errors in user data input and more.

For our shopping basket, an example would be attempting to add an unknown item to the basket:

```golang
basket.Add("Never Heard Of This Product", 999)
```

As ever, this needs a bit of designthinking. What should we do? It's up to us to decide.

One reasonable decision is to make the basket collaborate with something that knows our inventory. If the product is not found, our add() method should return a Go error.

We can write our test:

```golang
func TestRejectsUnknownItem(t *testing.T) {
    stockCheck := stockcheck.New()
    basket := basket.New(stockCheck)

    err := basket.Add("Never Heard Of This Product", 999)

    if err == nil {
        t.Error("Expected error, but none recieved")
    }
}
```

Obviously, we could make other decisions. Maybe `Add()` should not check anything, and that is done elsewhere. The basket could simply silently fail, and not add the unknown item. That sounds suitably only for a student exercise, but you never know.

Writing tests for exceptional behaviours - and the code that handles them - is a critical part of writing robust software.

### Observation: More specific tests, more general code

Working iteratively and traingulating detailed behaviour leads to a counter-intuitive result.

At the start, our tests are very general, and our implementation is very specific.

As we progress, we add more detailed tests. This drives out _more generalised_ code in our implementation.

---

BIG SUB HEAD: WHEN to change or delete tests

HOW to migrate an interface step by step

Eg PlaceShip( row col ) -> PlaceShip( Location ) work through

BIG SUB HEAD LAYERED TESTS and test pyramid

- Avoiding dupe tests
- Ideally one test per behaviour
- Balance close-in to boundaries
- Talk about test layer being a fixed point
