# Working with Test Doubles

Using Dependency Inversion allows us to _replace_ the real dependencies with _fake_ dependencies for testing.

![Using test doubles instead of real dependencies](images/test-doubles-used.png)

This makes no difference to our test.

When we execute our application logic, we are only concerned with what it does. We write code that will take data and transform it, or make decisions and act on them. This logic is independent of any source of data or target for actions.

That's the secret of _Test Doubles_.

## What are Test Doubles?

The name _Test Double_ comes from the world of movies.

When an actor is required to perform stunt work, usually a _stunt double_ is substituted for the real actor. They look similar. The specialist stunt performer can safely act out the hazardous action in the scene. The real actor will be unharmed - and most likely will not have the required skills to perform the stunt.

In software, a Test Double is a dummy component that simulates the real component for testing.

## Kinds of Test Doubles

### Stubs - Testing sources

### Mocks - Testing sinks

### Other kinds of doubles

## Example: Stubbing the System Clock

## Next: TDD and agility

[How TDD assists true agility >>](/chapter10/chapter10.md)
