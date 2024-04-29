# 08: TDD and Software Design

(intro)

what does TDD push?

- decoupling
- encapsulation
- abstractions

relevant to all paradigms oo fp proc

more important than normal
aim cont redesign

## key techniques

Let's summarise the key technqiues for achieveing thos design goals.

### Encapsulation

outside-in thinking
What do I need to be told?
What do I need to tell others?

nb: this is why TDD feels backwards. Its not TDD doing that, it is thinking in an abstract, encapsulated way.

implementation thinking
What am I allowed to remember?
What should I know how to do myself?
What do I need to translate?

### Abstraction

### Decoupling

Two things A and B are coupled together when a change in A requires a matching change in B, otherwise B will stop working.

Sometimes, these couplings are obvious and the code explicitly states what they are. Othertimes, they are implicit - sneaky little things that must change together, despite that not being obvious.

An example
examples to be aware of -

magic constants

implicit
executeQuery( queryString ) - implicitly, the syntax of the query string will be coupled to what the datasource needs. If we change a SQL source for a NoSQL source, then the queryString would also need to change. Avoid this by using the [Query](https://martinfowler.com/eaaCatalog/queryObject.html) pattern.

API parameters

> How many more examples of coupling can you think of from your past work?

### Accidental v essential complexity

> KISS - Keep it simple!

### Beck's Four Rules of Simple Design

https://martinfowler.com/bliki/BeckDesignRules.html

## How TDD helps surface design flaws

messy AAA gives feedback

Biology - systems grow from simplicity

## Three general principles

These old ideas are evergeen: keep things simple, avoid duplication and defer decisions until a need becomes apparent.

### KISS - Keep it Simple

Originally termed -Keep it Simple, Stupid- the KISS idea should be self-explanatory. Prefer simple solutions!

Because of human ego, it isn't quite as obvious as it should be, so we'll spell it out!

> Make things as simple as possible

Generally, we can build things in multiple ways. Some of those ways will be more difficult to understand. Some ways will be easier to understand. But all approaches are interchangeable, as they all do the exact same thing.

Choose the one that is simplest to understand.

### DRY - Don't Repeat Yourself

It's very difficult to safely change code when multiple copies of one piece of information exist.

We've all seen this:

- Code blocks copy-pasted everywhere
- Magic numbers throughout the code
- Multiple different functions, all doing the same thing
- Multiple sources of data (ugh, that one is the nastiest of all)

The solution is to have a single source of truth for everything. One constant. One function. One library. One data source.

This is given the name -Don't Repeat Yourself- and means just that. For each piece of truth, define it once and refer to it throughout the code.

This is also a key to finding abstractions: discovering where some fact or logic keeps turning up. Why? What about the problem makes it do that?

### YAGNI - You ain't gonna need it

This is a way of thinking central to working in short iterations.

It's tempting to plan ahead, looking far into the future, and guessing possible future requirements and possible future changes. For an engineer, it is then tempting to design flexibility into the code to handle these guesses.

DO NOT DO THIS!

The guesses will be wrong, generally. You then are stuck with bloated code that has some unpleasant properties:

- The code is way more complex than it ever needed to be
- The design fails to support the actual change
- The design you added actively gets in the way of the change!
- Future maintainers cannot tell that this over-design serves no purpose. It looks like real code.

The best plan is to restrict future planning. Accommodate short-term known growth, yes. But resist long term "flexible, configurable" rabbit holes.

## Generalising code

When doing TDD and using triangulation, we start with very specific implementations. As we add more specific tests, the implementation gets more general.

> Tests get more specific - implementation gets more general

There are several techniques that allow us to generalise code, in steps.

Once the variations are known, you can revise the concrete code to be more flexible, summarised next.

### Add variables

Replaced a fixed constant with a calculated variable:

```golang
func (b Basket) totalPrice() int {
  return 100 // get the first test to pass
}
```

Generalise this by adding the real calculation in a variable:

```golang
func (b Basket) calculateTotalPrice() int {
  total := 0

  for _, item := range b.items {
    total += item.Price
  }

  return total
}
```

Variables, and code to compute them, seem an almost trivial example.

### Add configuration

Replace hard coded constants and limits with data pulled from a configuration source:

```golang
func connect() {
  connection := connectTo( "https://somewhere.com/server")
}
```

hard-codes a URL. We can replace this with some configuration data:

```golang
func connect(config Configuration) {
  connection := connectTo(config.ServerURL)
}
```

Now the URL is passed via a Configuration struct, which can be initialised elsewhere.

### Add plugin points

Where behaviour needs to change, rather than data, we can use a plugin.

Plugins can take many forms. They are all ways to inject a block of code, rather than a piece of data. To do that, we can pass in a function containing the code, or pass in an object implementing the Strategy Design Pattern.

#### Varying the onclick behaviour of a button

Perhaps the easiest example to think of is a simple html button:

```html
<button onclick="alert('Hello world')">Click Me!</button>
```

This html fragment will display a button. When that button is clicked, an alert box will pop up saying 'Hello world'.

We might take this for granted, but it is an example of a powerful technique for varying behaviour. The button code here has two pieces of variability:

- The text to display - Click Me!
- The function to call when clicked - alert() in this case

With html, we use the attribute _onclick_ to define a function that will be called when the button is clicked. This function is passed as an argument to the button code.

This is a general mechanism we can use across languages: pass in a function to call in response to an event.

Here's a somewhat contrived example. We pass in a user `id` and and a function into `fetchUsername()`. This attempts to look up some user data and return the name. If there is an error in the lookup, it will call a function that we pass in as parameter `onError`:

```go
func fetchUsername( id int, onError func() ) string {
  profile, err := fetchUserProfile( id )

  if err != nil {
    onError() // call error handler function
    return ""
  }

  return profile.Name
}
```

This is an example of varying behaviour by passing in a function. It is analogous to the Object Oriented _Strategy_ Design Pattern. This is where we pass an object instead of a function, and some method on the passed-in object gets called. For more details on Strategy, see [here](https://refactoring.guru/design-patterns/strategy)

Further notes on this topic of generalising code can be found [here](https://www.quora.com/How-do-I-start-developing-a-software-framework/answer/Alan-Mellor)

## Further Reading

pre-reqs academy

- Refactoring
- DIP, DI and IoC
- Clean Code
- Ousterhout book

https://martinfowler.com/bliki/BeckDesignRules.html
