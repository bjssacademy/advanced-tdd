# advanced-tdd

## Introduction

Test-Driven development is a technique for writing _executable specifications_ for what a piece of code will do, and how we want that code to be called:

![TDD overview](/images/tdd-overview.png)

This enables us to work in very small, well-understood steps. TDD provides the opportunity to think about how our code will present itself to the rest of the program. We can decide what to encapsulate, and what to expose. We are doing _software design_ with TDD.

As a bonus, we get a suite of regression tests that give us confidence to change code and be sure we have not broken anything.

TDD is a key technique to _being agile_ as it leads to well-understood, highly modular code. This kind of code is the easiest to change.

## TDD Benefits

TDD enables software to be written so that it:

- is easier to change
- is safe to modify later
- has fewer defects
- can be broken into small pieces
- is easier to continuously integrate

This guide will start from the basic rhythms of TDD. We will learn the idea behind executable specifications, and how to write them. We will learn the regular workflows behind developing test-first. We will see that this prioritises continuous improvement of code and enables continuous delivery.

As TDD is all about encouraging and experimenting with software design, we will be covering aspects of software design as they relate to TDD. These include areas such as:

- Designing software components
- Dependency Inversion
- Test Doubles - simulating hard to control dependencies
- Decoupling from external systems
- Code as storytelling
- Separation of concerns
- Decoupling interface from implementation

Test-Driven Development is a _technique_ (not a framework or library) that applies to all programming efforts, across all languages. It works for front end and back end development.

By the end, we should be comfortable in understanding what TDD is, how it can be applied to our work, and the benefits it can bring.

## Pre requisites

- [Go in a Day](https://github.com/bjssacademy/goinaday) for an intro to Go as a second language
- [Refactoring and Code Smells]() covering iterative re-design
- [DIP, DI and IoC]() describes the design technique supporting Test Doubles

The Go programming language is used for examples in this version. The technique applies to all languages.

An introduction to the fundamentals of programming using Go can be found [here](https://github.com/bjssacademy/fundamentals1)

This material supports the content in the BJSS Academy Intro to TDD in the Engineering Foundation Course, and also the Level 6 TDD Apprenticeship sessions.

## Start here

Right, let's crack on - what has TDD ever done for us, anyway?

![What has TDD ever done for us?](images/tdd-ever-done-for-us.jpg)

[Chapter 1 >>](/chapter01/chapter01.md)
