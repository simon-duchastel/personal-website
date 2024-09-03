---
title: "Advent of Code 2023"
date: 2024-07-19
layout: single
categories: ["Software"]
---

_[View the code](https://github.com/simon-duchastel/advent-of-code-2023)_

I finally got around to writing about how I used [Rust](https://www.rust-lang.org) 6 months ago for the 2023 [Advent of Code](https://adventofcode.com). I've been meaning to try out Rust in a small personal project for a few years now but have never gotten around to it.

## What is Rust

Rust is a new(ish) systems language often touted as a replacement for C++. It's become very popular in the developer community over the last few years. Notably, Amazon announced that they were moving to Rust for a lot of their systems.

The language's focus is on memory safety. One of the most unique features of the language is its borrow checker, which is a set of strict compile-time checks that enforce that a pointer in your code can only ever be referenced in one place (with some exceptions if it's a read-only pointer). I'm omitting some details here but in essence, the effect is to prevent the dreaded "Null-Pointer Exception" by ensuring that there's only ever one piece of code writing to a piece of memory.

On the one hand, this is cool - Rust gives you safety by ensuring that your code won't read and write inconsistent data, and it does all of this at compile time so you still have super fast code! The downside though is that because this is a rigorous compile-time check that prioritizes safety and "zero-cost abstractions" (meaning heavy compile-time checking), Rust's type system can feel quite onerous. The way Rust's borrow checker can ensure there are no memory issues is by requiring you the programmer to specify in great detail how their data is being stored, how it's being used, and writing their types in such a way that it becomes physically impossible for unsafe memory calls to happen.

All that to say Rust's strong memory safety, fast speed, and strong focus on developer experience (such as support of macros in the language and built-in package manager) have made it a very popular language amongst programmers.

## The Project

My project this year was a command line program, simpler than prior year's programs in terms of user experience and options for interacting with it. The result of the project was quite boring. Unlike in previous years, I didn't try to do anything cool like [generate a website that computes the problems in the browser](../advent-of-code-2022/). My program was a very simple command line interface for calling functions that execute the Advent of Code problems. The program itself wasn't that extensible, as the whole project was more of my attempt at teaching myself basic Rust.

The project helped me achieve my goal though. I can now say I've written Rust code. I didn't use many of the more advanced features of the language like macros, but I was exposed to them and now feel like I have a greater appreciation of why Rust is as popular as it is - and why some people are skeptical.

## My Thoughts on Rust

### The developer experience is wonderful.

In particular, [Rust's handbook](https://doc.rust-lang.org/book/) and its built-in package manager [Crate](https://doc.rust-lang.org/book/ch07-01-packages-and-crates.html) are stand-out examples of making it easy for developers to get up and running with a language.

### The borrow-checker is frustrating to use.

This is hardly an original thought - it's one of the most commonly-remarked aspects of Rust! While it certainly helps you write pure code free of memory bugs, the borrow checker was very unintuitive. I felt pretty confident in the fundamentals of how it worked, but even so would bang my head against a wall trying to get my code to compile. It turns out that memory is unintuitive and hard to grasp, and writing pure code that never runs into memory issues is hard!

You could argue that the borrow checker is good as it prevents you from writing nasty memory-leak bugs into your code. At the same time though, there were several times where I knew my code would execute correctly given my inputs, but the code technically could have a memory bug given the right circumstance. It took a lot of fighting with the borrow-checker in these cases to constrain my solution down to just the inputs I knew wouldn't fail or to abstract my variables such that there was no chance of memory issues, even though in practice I knew I'd never get segfaults in my simple Advent of Code program.

### Rust slows you down, but your code's memory management is likely to be correct at the end.

It's hard to write simple code quickly in Rust. Rust puts a lot of roadblocks in front of you before allowing you to compile your program successfully. However, once the program does compile, you can be pretty confident that a lot of nasty bugs won't be present (null-pointer exceptions being the one people point to most often). This is what I love about Kotlin, and Rust appears to be doing the same thing for low-level languages that manipulate pointers more directly.

However, whereas I find Kotlin to be fairly forgiving and quick to write with, Rust is not forgiving. It forces you to specify all the ways your program is memory-safe and gives you very few escape hatches when your goal is to write moderately unsafe code. Compare that to Kotlin, where it's easy to use operators like `!!` or `?:` to escape out of nullability checks in a way that might yield bad behavior, but which you as the programmer expect never to happen (or if it does, you don't care so much about it). Escape hatches such as these help with quick development, both in cases where you truly don't care about writing 100% correct code (like in a hobby Advent of Code program) or in the case of quick prototypes where you clean up the code once you're ready for production.

## Takeaways

Overall I enjoyed doing the Advent of Code! I'm happy that I tried a different language than Kotlin. I'd already used Kotlin in my [2022](../advent-of-code-2022/) and [2021](../advent-of-code-2021/) Advent of Code projects. I plan to do 2024 Advent of Code as well, and not just because it's become sort of a tradition for me. I enjoy the excuse Advent of Code gives me to try something new, so I'll see what new language or framework I try next year.

The other thing I found doing Advent of Code this year is that I don't prefer Rust as a language. I've enjoyed using Go and Kotlin more, and feel these are more suited to the kind of code I tend to write in my free time. That said, I understand why so many people like Rust and I'm glad I tried it. I also think that judging the language off on just a few programming challenges isn't a fair evaluation of the language, so I'll look for opportunities to use Rust in other personal projects to see if my opinion changes.