---
title: "Advent of Code 2022"
date: 2022-12-28
---

[View the code](https://github.com/simon-duchastel/advent-of-code-2022)

[View the website for calculating solutions](https://simon.duchastel.com/adventofcode2022)

This year I again participated in the Advent of Code, for the third year in a row. I wrote my solutions in Kotlin leveraging Kotlin Multiplatform (KMP), this time messing with the javascript compiler and the javascript Jetpack Compose Can you tell I really like Kotlin?

I used Kotlin Multiplatform last year and gave [some of my thoughts here](../advent-of-code-2021/#kotlin-multiplatform). My thoughts on KMP overall mostly haven't changed. I will say that the ecosystem appears to be maturing, which is good. There are still a lot of rough edges and a ways to go, but I'm excited for KMP to continue to evolve!

## Jetpack Compose

[Jetpack Compose](https://developer.android.com/jetpack/compose) is a declarative UI framework developed by Google for use on Android. If you're [Jake Wharton you'll also argue](https://jakewharton.com/a-jetpack-compose-by-any-other-name/) it's not _just_ a UI framework but a "general-purpose framework for managing trees of nodes". While I agree with him, most people are familiar with it as a UI framework so I'll stick to that here.

Jetpack Compose lets you take an immutable state and declaratively define your UI from that state. Everything is defined as functions, and when the state changes at any point the framework handles re-calling your functions to re-render your final UI. It does some "fancy magic"© to do this efficiently and only re-call the necessary functions.

In code, that might look something like this (not 100% compilable but close enough):
```
var myState = true
if (myState) { // if it's true, show the Big Red Button
    BigRedButton()

    delay(30) // wait 30 seconds
    myState = false
} else { // if it's false, show the Small Blue Button
    SmallBlueButton()
}
```

The above example would show a Big Red Button on the screen for 30 seconds, then show the Small Blue Button. This is pretty intuitive to write, easy to read, plus super easy to test - just pass in your state and assert that the rendered UI matches what you expect. No need to try increasingly complex combinations of actions to find invalid UI states.

### Jetpack Compose for the web

 Jetbrains (the company behind Kotlin) also develops a [version of Jetpack Compose for the web](https://jb.gg/compose-web) which is built off of Google's version. This lets you write Web UI in Jetpack Compose code, using the Kotlin Javascript compiler!

At first I was super excited to use this! I really like Jetpack Compose for its ability to write clean, easy to reason about UIs and was excited to use this same UI framework for the web. I was imagining this cool world where you could write generic Jetpack Compose code and have it more-or-less work on both Android and the Web. I knew you'd always make some tweaks to get things to show up beautifully in the web, but my hope was that if you wrote your code generically and factored it well, 90% could be re-used between Android and the web.

That turned out to not be the case at all. Jetpack Compose for the web allowed you to re-us the declarative pattern inherit to Compose, but that was it. You still wrote in `div`, `body`, `h1`, etc. Just this time, you had declarative Compose functions to do it in so that you could generate UIs using immutable state.

Don't get me wrong, I love the declarative UIs written in Jetpack Compose! In my project the rendering of loading state followed by the final result once loading was completed was very easy using this pattern. That said, I was hoping for something more when I first saw "Jetpack Compose for the web".

I'm told that somewhere there exists a version of Jetpack Compose for Javascript which actually implements all of the standard Jetpack Compose standard elements (like `Column()`, `Row()`, etc.) but if so, I haven't found it yet.

## The Project

The project itself is a very basic web-ui which lets you select which problem you wish to solve, then select which part of the problem to solve (part 1 or 2) and which input to use (the sample input or the real input).

Now the cool part is that when you click a button to see the solution, it doesn't just show you a pre-computed answer. The Javascript code actually runs in your browser, compiled from Kotlin code! The computation happens pretty quickly, but then again the problems I'm solving aren't particularly intensive so it shouldn't come as a big surprise.

### Extensibility

I got a bit fancy and was pretty proud of [my factoring](https://github.com/simon-duchastel/advent-of-code-2022/blob/492f2467a0c0816ac917e093172efa2edbfe4d04/src/jsmain/kotlin/com/duchastel/simon/adventofcode2022/problems/Problem.kt#L17). I had my Compose UI take a list of `Problem` objects. Each of these objects told the UI its name and a few other parameters, and then gave it some lamdbas for executing the problem. As a result, once I set my UI I could

Moreover, since my UI used re-usable Jetpack Compose elements it was easy and intuitive for me to add new problems as needed. All I needed to do was add a new `Problem` object to a global list, point it to some Kotlin lambda that did my computation, and my UI would automatically render it on the screen and handle calling the computation when the user clicked the right button and rendering the result on screen. Even better - while my UI was Javascript-specific, my computations were all Kotlin common and could be re-used in any Kotlin project (for Android, iOS, etc).

Now whether or not this would actually hold-up in a large-scale environment, I'm not sure. But it was fun to do and made it easy to add my solutions to subsequent days in the Advent of Code challenge, so I see it as a win!

### The UI

In previous posts I've attached a gif or screenshot of the UI, but why show you a still image of the UI when you can see the UI for real! [Click here](https://simon.duchastel.com/adventofcode2022) to see the project live in your browser.

## Thoughts

#### I'm disappointed by Javascript Jetpack Compose

Like I shared above, I was disappointed by the limitations of Jetpack Compose for the web and Javascript. It was cool to use declarative patterns for my web UI, but fundamentally I was still creating UI with low-level web building blocks instead of the higher-level components I wanted. I don't want to build with `div` and `h1` - I want to build with `Column()` and `Row()`.

Since I started this project a coworker showed me [Redwood](https://jakewharton.com/native-ui-with-multiplatform-compose/), a tool built by the incredible Jake Wharton and his team at Cash App. This seems a lot closer in spirit to what I was hoping, but requires you basically do 90% of the heavy lifting of creating generic components yourself and is more suited to well-developed UI systems at mature companies than it is for quick iteration on multiplatform projects. Hopefully Jetbrains and/or Google eventually gets around to making true multiplatform Compose (after they fully mature the rest of the Kotlin Multiplatform ecosystem of course!).

#### Kotlin Multiplatform continues to mature

The build system for KMP seems to have really improved since the last time I used it. I think part of it is also the fact that I've been getting much more familiar with the intricacies of Kotlin's gradle framework both in and out of work, but it really does seem like Kotlin is getting much easier to use! I had an easier time getting my builds up-and-running, installing the IDE and development environment was a lot less fiddly, and the number of seemingly random errors I ran into was a lot less.

The error messages in the build system are still very unhelpful though. The documentation is getting better, but if you do something wrong it can be hard to understand how to fix it.

## Takeaways

Overall, this project was a lot of fun. I didn't get nearly as far as I intended to get with the problems (I only got to day 4) because the month of December became very busy planning a trip to Spain, busyness at work and outside of it, and my family visiting from the East coast. That said, it was cool to poke some more at a technology I'm interested in and I had fun solving this years' problems. Let's see if I can break my streak of doing Advent of Code in Kotlin next year!