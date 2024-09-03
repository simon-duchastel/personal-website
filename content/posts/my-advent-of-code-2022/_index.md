---
title: "My Advent of Code 2022"
date: 2022-12-28
layout: single
categories: ["Software"]
---

_[View the code](https://github.com/simon-duchastel/advent-of-code-2022)_

_[View the website for computing solutions](https://simon.duchastel.com/adventofcode2022)_

This year I participated in the [Advent of Code](https://adventofcode.com) for the third year in a row. I wrote my solutions in Kotlin Multiplatform (KMP) like I did in 2021, but this time using with the Javascript compiler and Jetpack Compose to render a website written in Kotlin.

[I shared some thoughts](../advent-of-code-2021/) on Kotlin Multiplatform last year. My thoughts mostly haven't changed. The ecosystem appears to be maturing, which is good. There's still a ways to go before it's at a point where I think it will be ready for broad adoption but I'm excited for KMP to continue to evolve.

## Jetpack Compose

[Jetpack Compose](https://developer.android.com/jetpack/compose) is a declarative UI framework developed by Google for use on Android. [Jake Wharton argues](https://jakewharton.com/a-jetpack-compose-by-any-other-name/) it's not _just_ a UI framework but a "general-purpose framework for managing trees of nodes". While I agree with him, most people are familiar with it as a UI framework.

Jetpack Compose takes immutable state and computes a declaratively-defined UI tree from that state. Everything is defined as functions, and when state changes the runtime re-calls your functions and diffs the result in order to determine how the UI should change. It uses some neat tricks to do this efficiently and limit the number of function calls. This blog post does a good job of going into the details if you're interested: http://intelligiblebabble.com/compose-from-first-principles/.

Compose code looks something like this (not quite compilable but close enough):
```
var myState = true
if (myState) { // if it's true, show the Big Red Button
    BigRedButton() // renders the Big Red Button

    delay(30) // wait 30 seconds
    myState = false
} else { // if it's false, show the Small Blue Button
    SmallBlueButton() // renders the Big Blue Button
}
```

The above example would show a Big Red Button for 30 seconds, then show a Small Blue Button. It's intuitive to write, easy to read, and simple to test - just pass in your state and assert that the right elements do or do not exist on the screen as you expect.

### Jetpack Compose for the web

Jetbrains (the company behind Kotlin) also develops a [version of Jetpack Compose for the web](https://jb.gg/compose-web) which is built off of Google's version. This lets you write Jetpack Compose which gets translated into a web UI using the Kotlin Javascript compiler.

I really like Jetpack Compose for its easy-to-reason-about UIs and was excited to use it for web development. My hope was that if factored my code well, 90% of my UI could be re-used between Android and the web simply by re-using the same Composable functions.

That turned out to not be the case. Jetpack Compose for the web has a different set of primitive Composable functions and still requires you to use html tags like `div`, `h1`, and `a`.

To be fair, using Compose has its benefits aside from portability. In my project it was easy and fun to implement a loading UI followed by the final result once loading was complete. That said, I was hoping for more re-usability.

I'm told that somewhere there exists a version of Jetpack Compose for Javascript which actually implements all of the standard Jetpack Compose primitives (like `Column()`, `Row()`, `Text`, etc.). If so, I haven't found it yet.

## The Project

My project is a very basic web UI that lets you select which problem you want to solve, which part of the problem to solve (part 1 or 2), and which input to use (the sample input or the real input).

The cool part is that when you click a button to see the solution, it doesn't just show you a pre-computed answer. The Javascript code actually runs in the browser, compiled from Kotlin code.

### Extensibility

I'm pretty proud of [my factoring](https://github.com/simon-duchastel/advent-of-code-2022/blob/492f2467a0c0816ac917e093172efa2edbfe4d04/src/jsmain/kotlin/com/duchastel/simon/adventofcode2022/problems/Problem.kt#L17). My Compose UI takes a list of `Problem` objects. Each of these objects tells the UI its name and a few other parameters and exposes lambdas to solve the problem.

As a result, all I need to do to add a new problem to the UI is add a new `Problem` object to the global list and point it to a lambda that does the computation. The UI automatically handles calling the computation when the problem is selected and rendering the solution to the screen. Even better, while my UI is Javascript-specific my computations are all written in common Kotlin code and can be re-used in any Kotlin project regardless of platform.

### The UI

The UI and code for this year's solution runs live in your browser and can be viewed at https://simon.duchastel.com/adventofcode2022.

## Thoughts

### I'm disappointed by Jetpack Compose for Javascript

I'm disappointed that Jetpack Compose for Javascript is tightly coupled to the web and not portable to Android. While it was cool to use declarative patterns for my web UI I still ended up building a website using html tags.

Since I started this project a coworker showed me [Redwood](https://jakewharton.com/native-ui-with-multiplatform-compose/), a tool built by Jake Wharton and his team at Cash App. This seems a lot closer in spirit to what I was hoping for out of Compose for Javascript, but requires you to create all the cross-platform components yourself. It seems better suited to mature UI systems at large companies than quick iteration on hobbyist multiplatform projects. Hopefully Jetbrains and/or Google eventually gets around to making true multiplatform Compose (although I'd prefer they fully mature the rest of the Kotlin Multiplatform ecosystem first).

### Kotlin Multiplatform continues to mature

The build system for KMP has really improved since the last time I used it. I think part of it is the fact that I've been getting more familiar with the intricacies of Kotlin's gradle framework, but it does seem like Kotlin is getting much simpler to use. It was much quicker to get my builds up-and-running, installing the IDE and development environment was a less fiddly, and I ran into fewer seemingly random errors.

The error messages from the build system are still unhelpful though. The documentation is getting better, but if you do something wrong in your build it's often hard to understand how to fix it.

## Takeaways

Overall, Advent of Code 2022 was a lot of fun. I didn't get as far as I intended to get with the problems due to family, vacation, and work (I only got to day 4) but it was cool to use KMP in a new way and try my hand at more web development.