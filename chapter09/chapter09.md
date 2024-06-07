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

Simulations, games and statistical applications often use sources of random numbers. Code that relies on these sources is difficult to test. Given a random input, we can't predict the output. That means we can't write the assert section of our test. Oh dear.

## Managing dependencies by design

The problem here is direct coupling to a system we cannot easily control or inspect.

A key insight is that _we do not even care_ about that system. We are not testing that system. We only want to test drive code that _works with the results_ of that system. We're not bothered about reading data from a database. We are bothered about the logic our code does to that data. We're not bothered about reading the system time. We are bothered about what our code does in response to a specific time.

This insight hints at a solution:

- Test only the logic of our code
- Use a less-difficult dependency to replace our difficult one

What if we swapped our database with something under our full control? Something that _pretended_ to be a database, but had none of the issues of the real thing?

We can imagine replacing the database with something that always returned some pre-canned data for our User Profile code. Or a replacement fake random number generator, that always returned the number 4.

By doing this, our tests are simple. We avoid dealing with all the difficulties of our original dependency. No connections. No logins. No randomness. We can easily write a FIRST test. We can easily write an assert, now we can be certain of how our new dependency behaves.

It leaves only one question.

_How can we design our code to swap out a difficult dependency?_

### Dependency Inversion - Decoupling Dependencies

Fortunately for us, the answer comes from standard software design: Dependency Inversion.

The reason our `loadUserProfile` function was hard to test is because of the line

```golang
result := queryDatabase(id)
```

This line causes a direct coupling to the query database function. Assume here that inside this function, we connect to a real SQL database and query it.

_We need to break this direct connection._

Instead of directly calling `queryDatabase()`, we need that code to indirectly call it. It needs to call _something else_ that will allow us to swap out the actual call made.

We can do this in one of two ways, either using Object Oriented techniques or Functional Programming techniques.

#### Object Oriented Dependency Inversion

TODO TODO TODO

#### Functional Dependency Inversion

This approach uses function currying and closures. It's simpler to do than to describe.

- Close over a function reference
- Our code calls whatever function reference was supplied
- At run time, create this closure with the dependency we want to use

```golang
package main

import (
	"fmt"
)

type UserProfile struct {
	id            int
	name          string
	favouriteFood string
}

func createFetchProfileFunction(queryDatabase func(string, int) []string) func(int) UserProfile {
	return func(id int) UserProfile {
		query := "SELECT name, age, favouriteFood FROM Profiles WHERE id = ?"

		results := queryDatabase(query, id)

		return UserProfile{
			id:            id,
			name:          results[0],
			favouriteFood: results[1],
		}
	}
}

func fakeQueryDatabase(query string, param int) []string {
	return []string{"Alan", "curry"}
}

func main() {
	fetchUserProfile := createFetchProfileFunction(fakeQueryDatabase)

	profile := fetchUserProfile(3)
	fmt.Println(profile)
}
```

We first call `createFetchProfileFunction` to create our fetchUserProfile function. We have passed in our `fakeQueryDatabase` function in this case. That gets bound to the `queryDatabase` parameter. We return a function from this function - basic function currying. The returned function _closes over_ the queryDatabase parameter. Whenever we call this newly minted returned function, it will use whatever value was closed over in queryDatabase.

We can change which actual function gets called by the line reading `results := queryDatabase(query, id)`.

In the test example above, this means we will call the `fakeQueryDatabase` function. This isn't a database, of course. It returns pre-canned data.

You can see this code run on [this playground](https://goplay.tools/snippet/NLH6e1qCyvM)

### Hexagonal architecture

TODO TODO

Wikipedia has a good summary of [Hexagonal Architecture](<https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)>)

For more details about how TDD becomes more effective with heagonal architecture, see [this book](https://www.oreilly.com/library/view/test-driven-development-with/9781803236230).

## Working with Test Doubles

### Stubs - Testing sources

### Mocks - Testing sinks

## Next: TDD and agility

[How TDD assists true agility >>](/chapter10/chapter10.md)
