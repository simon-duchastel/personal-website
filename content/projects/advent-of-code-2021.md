---
title: "Advent of Code 2021"
date: 2022-01-15
project-tags: ["Advent of Code 2021"]
---

[Source](https://github.com/simon-duchastel/advent-of-code-2021)

This year for the Advent of Code challenge, I wanted to use it as an opportunity to try out something new. At the company I work for (98point6) some teams were starting to use Rust, and since it's a language that seems really interesting I was considering doing the problems in that. However, I decided instead to look into [Kotlin Multiplatform](https://kotlinlang.org/docs/multiplatform.html) (KMP) and better understand what it means to write Kotlin code that can run on an Android device, an iOS device, a web browser, and more. (Maybe I'll do Rust next year!).

## Kotlin Multiplatform

Kotlin Multiplatform utilizes the three different Kotlin compilers Jetbrains has for Kotlin - Kotlin JVM for compiling to JVM bytecode, Kotlin Native for compiling to native binaries, and Kotlin Javascript for compiling into Javascript. You can define `common` code that's platform-agnostic (meaning you can't reference Java-specific APIs like JVM file system operations, for example).

You can also specify platform-specific code that works with the common code but uses platform-specific functionality, and is only present when compiled to that platform. Common code can depend on platform-specific APIs by using the `expect` keyword, which tells Kotlin "This function/variable/etc. isn't defined here, but each platform will have its own implementation that you can depend on". Each platform you compile to must have corresponding `actual` declarations.

## The Project

For this project, I decided I'd start with creating a command-line program for executing each of the AoC problems. I could do it as a simple Java command line tool like I'm used to, but I figured that if I'm going to be using Kotlin Native I might as well make use of the native functionality! Since I was developing on a windows machine, I targeted MinGW.

After a lot of trial and error and coercing Gradle to do what I wanted it to, building the command line tool became as easy as running:
```
gradlew linkReleaseExecutableMingw
```

I'm pretty pleased with how the command line tool turned out. There's definitely some rough edges, but the abstractions I built early on ended up enabling me to effortlessly add additional problems to the tool. I setup a system where each problem could declare a top-level command to respond to, as well as sub-commands the user can execute. The most annoying thing initially was probably getting the file system access right, since I had to use the native system functionality rather than the nicer stuff the JVM provides.

My goal in designing the tool was to make it conversational. When you run the tool, it asks you which problem you want to execute.

![Command Line Tool In Action](command-line.gif#center "Command Line Tool")

Once you select a problem, there are a series of sub-questions — such as whether part 1 or part 2 should be executed. It's pretty basic, but works fairly well. I'm also pretty pleased with my solutions, although most of them are fairly straight-forward. I especially liked those problems that were about simulating things, since I was able to model the problem with types and functions and use immutable state to keep track of the result.

You can find my code on Github here: https://github.com/simon-duchastel/advent-of-code-2021

## Things I've Learned

#### I still really like fold()

[As I noted last year](/projects/older-projects#advent-of-code-2020), I really like certain aspects of functional program and using the fold() function. I think I used fold() in almost every single AoC problem so far this year, and tried to use imperative patterns sparingly. There were many times where I think that spurning mutability helped me a lot in these problems! There were certainly a few times though where it led me to make mistakes and have a less clean solution than I would have otherwise, so I gained a better intuition about its limitations as well.

#### Kotlin Multiplatform is really meant to be used as a library

After working with Kotlin Multiplatform for a bit, it seems to me that it's intended to be used as a library by other native projects on their respective platforms. This makes sense given its use case as reusable code that slots into wider native projects, which seems to be Jetbrain's goal with Kotlin.

Anecdotally, I had difficulty getting KMP to compile to a standalone .jar rather than a library. I could get a .jar, but the normal functionality for Kotlin has in gradle for defining an entry-point to the program didn't seem to quite work with Multiplatform. (I admit though that this was likely due to some build process and gradle errors on my part too). Overall though, it seems that Kotlin Multiplatform tacitly encourages you to think of your code as a shared library, which seems reasonable to me.

#### The build process is still a little rough around the edges

As I mentioned before, I fully admit that I am far from being a gradle wizard. That said though, I can generally make my way around gradle build systems, having worked with them quite a bit at my day-job as an Android developer. However, Kotlin Multiplatform has a very complex build system configuration, and the multiplatform gradle plugin that's included to get things to build seems to be doing a lot of * magic * under the covers. Wading through and understanding those configurations gets really complex.

Hopefully, as KMP matures its build process likewise gets more understandable. When you have to build across three different platforms (and more when you consider all the sub-configurations for each platform), I imagine there's always going to be a certain level of complexity. For those not wanting to do anything fancy though and just compile a basic command line executable, app, or website, that complexity should be more-or-less hidden if Jetbrains wants KMP to become more widely adopted.

## Next Steps

I plan to continue working on this project and have a few ideas for things to do to wrap it up:

- Finish the remaining 11 problems.
- Generate a basic app Android app to solve the problems
- Generate a basic website to solve the problems

This isn't meant to be actually useful and is just something I'm doing for fun, so I'll do it casually in my free time. TBD when it gets done, but I'm hoping to post more about my Advent of Code 2021 work soon!