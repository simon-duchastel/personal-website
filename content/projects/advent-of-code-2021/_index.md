---
title: "Advent of Code 2021"
date: 2022-01-18
---

[View the code](https://github.com/simon-duchastel/advent-of-code-2021)

This year for the Advent of Code challenge, I wanted to use the opportunity to try out a new technology. At the company I work for (98point6) some teams were starting to use Rust, and since it's a language that seems really interesting I was considering doing the problems in that. However, I decided instead to look into [Kotlin Multiplatform (KMP)](https://kotlinlang.org/docs/multiplatform.html) and better understand what it means to write Kotlin code that can run on Android, iOS, the web, etc. (Maybe I'll do Rust next year!).

## Kotlin Multiplatform

Kotlin Multiplatform utilizes the three different Kotlin compilers - Kotlin JVM for compiling to JVM bytecode, Kotlin Native for compiling to native binaries, and Kotlin Javascript for compiling to Javascript.

 You can define `common` code that's platform-agnostic (meaning you can't reference Java-specific APIs like JVM file system operations, for example). You can also specify platform-specific code that works with the common code but uses platform-specific functionality, and is only present when compiled to that platform. Common code can depend on platform-specific APIs by using the `expect` keyword, which tells Kotlin that _this isn't defined here, but each platform will have its own implementation that you can depend on_. Each platform you compile to must have corresponding `actual` declarations.

## The Project

For this project, I decided I'd start with creating a command-line program for executing each of the AoC problems. I could have written it as a simple Java command line tool like I'm used to, but I figured that if I'm going to be using Kotlin Multiplatform I might as well make use of the native functionality! Since I was developing on a windows machine, I targeted [MinGW](https://en.wikipedia.org/wiki/MinGW).

After a lot of trial and error with Gradle, I got to a point where building the command line tool became as easy as running:
```
gradlew linkReleaseExecutableMingw
```

Doing this would drop an executable _.exe_ file into my _\bin_ folder, and away I went!

I'm pretty pleased with how the command line tool turned out. There's definitely some rough edges, but the abstractions I built early on ended up enabling me to easily add additional problems to the tool. I set up a system where each problem declares a top-level command it responds to, as well as sub-commands the user can execute to configure that problem. The most annoying thing initially was probably getting the file system access right, since I had to use the native system functionality rather than the nicer stuff the JVM provides.

My goal in designing the tool's interface was to make it conversational. When you run the tool, it asks you which problem you want to execute.

![Command Line Tool In Action](command-line.gif#center "Command Line Tool")

Once you select a problem, there are a series of sub-commands — such as whether part 1 or part 2 should be executed. It's pretty basic, but works fairly well. I'm also pretty pleased with my solutions, although most of them are very naive implementations. They were fun to write though. I especially liked those problems that were about simulating things since I was able to model the problem with types and functions and use immutable state to keep track of the result.

You can find my code on Github here: https://github.com/simon-duchastel/advent-of-code-2021

## Things I've Learned

#### I still really like fold()

[As I noted last year](/projects/older-projects#advent-of-code-2020), I really like certain aspects of functional program and using the fold() function. I think I used fold() in almost every single AoC problem so far this year, and tried to use imperative patterns sparingly. There were many times where I think that spurning mutability helped me a lot in these problems! There were certainly a few times though where it led me to make mistakes and have a less clean solution than I would have otherwise, so I gained a better intuition about its limitations as well.

#### Kotlin Multiplatform is really meant to be used as a library

After working with Kotlin Multiplatform for a bit, it seems to me that it's intended use case is as building libraries for other native projects. This makes sense as a way to hav shared business logic that can be slotted into more "lower-level" platform-specific code, which seems to be Jetbrain's goal with KMP.

Anecdotally, I had difficulty getting Kotlin Multiplatform to compile to a standalone .jar rather than a library. I could get a _.jar_ built, but the normal functionality Kotlin has in gradle for defining an entry-point to the program didn't seem to quite work with Multiplatform. (I admit though that this was compounded by gradle errors on my part). Overall though, it seems that Kotlin Multiplatform tacitly encourages you to think of your code as a shared library, which seems reasonable to me but hampered some of my development during this project.

#### The build process is still a little rough around the edges

As I mentioned before, I fully admit that I am far from being a gradle wizard. That said though, I can generally make my way around gradle build systems, having worked with them a fair amount at my day-job as an Android developer. However, Kotlin Multiplatform has a very complex build system configuration, and the multiplatform gradle plugin that's included seems to be doing a lot of _* magic *_ under the covers. Wading through and understanding those configurations got pretty complex.

Hopefully, as KMP matures its build process will likewise get more understandable. When you have to build across three different platforms (and more when you consider all the sub-configurations for each platform), I imagine there's always going to be a certain level of complexity. However, for those of us not wanting to do anything too fancy that complexity should be more-or-less hidden if Jetbrains wants KMP to become more widely adopted.

## Next Steps

I plan to continue working on this project and have a few ideas for things to do to wrap it up:

- Finish the remaining eleven problems.
- Generate a basic app Android app to solve the problems
- Generate a basic website to solve the problems

This isn't meant to be actually useful and is just something I'm doing for fun, so I'll try to keep it up casually in my free time. TBD when it gets done, but I'm hoping to post more about my Advent of Code 2021 work soon!