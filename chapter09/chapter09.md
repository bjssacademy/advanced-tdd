# Difficult Dependencies - Designing with Test Doubles

As we grow our system using TDD, we will run into a major problem.

How do we test the code that interacts with external systems?

External systems like payment processors form difficult dependencies to test. How would we test code that sends a pyment of ten pounds to a payment processor, without actually sending ten pounds every time we run the test?

## What are difficult dependencies?

Technically, a 'difficult dependency' is a connection to a system that we don't have full control over.

Our code depends on this system to do its work. Our test will have to setup this dependency and quite likely assert against some observable change in its behaviour.

Both the setup, done in the arrange step, and the verification in the assert steo can be problematic.

Let's look at some common cases to see why.

#### Databases

The most common problem is testing code that accesses a database.

Here's a function that will read a database table called `Profiles` and return a struct of user profile data:

```golang
func loadUserProfile(id int) UserProfile {
    query := "SELECT name, age, favouriteFood FROM Profiles WHERE id = ?"

    results := queryDatabase(query, id)

    return UserProfile{
        id: id,
        name: results[0],
        age: results[1],
        favouriteFood: results[2]
        }
}
```

With a bit of artistic licence assuming the function `queryDatabase` exists and returns a slice of column values. You can see how that gets mapped into a struct.

This is hard to test. But why?

- We need a database server running
- It must have a valid connection
- The network must work
- We must have a database user account and password
- The database must have a table called Profiles
- The table must have the correct columns
- We must either have some test data in that table, or be able to add it

All this causes several difficulties:

- The test will be very slow to run
- A test fail might be because the database is down, and not because our code failed the test
- What if we add data to the production database during the test?
- If we have pre-set data in a test database, the test is hard to read: why are we looking for a user called "Daffy Duck"?

This is how we get slow and flaky tests. This is a violation of FIRST test principles.

#### User interfaces

Testing through user interfaces can be slow. We need to simulate key presses, clicks and navigation. We need to scrape screen content to check the results.

#### System time

Any code that directly access the system clock to perform time-sensitive actions is hard to test. We would have to change the system time in the test.
This is often impractical. It may cause other running applications to fail, or data to be stored incorrectly.

#### Random number generators

Simulations, games and statistical applications often use sources of random numbers. Code that relies on these sources is difficult to test. Given a random input, we can't predict the output. That means we can't write the assert section of our test.

## Managing dependencies by design

The problem here is direct coupling to a system we cannot easily control or inspect.

A key insight is that _we do not even care_ about that system. We are not testing that system. We only want to test drive code that _works with the results_ of that system. We're not bothered about reading data from a database. We are bothered about the logic our code does to that data. We're not bothered about reading the system time. We are bothered about what our code does in response to a specific time.

This insight hints at a solution:

- Test only the logic of our code
- Use a less-difficult dependency

What if we swapped our database with something under our full control? Something that _pretended_ to be a database, but had none of the issues of the real thing?

How could we design our code to swap out a difficult dependency?

### Dependency Inversion - Decoupling Dependencies

### Hexagonal architecture

Wikipedia has a good summary of [Hexagonal Architecture](<https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)>)

For more details about how TDD becomes more effective with heagonal architecture, see [this book](https://www.oreilly.com/library/view/test-driven-development-with/9781803236230).

## Working with Test Doubles

### Stubs - Testing sources

### Mocks - Testing sinks

## Next: TDD and agility

[How TDD assists true agility >>](/chapter10/chapter10.md)
