---
title: "Creating my website with Hugo + Go"
date: 2023-08-01
---

_[View the code](https://github.com/simon-duchastel/personal-website)_

Twenty years ago, my father and cousin bought the domain duchastel.com (my last name is Duchastel). Initially they hosted bios for some of our family members on it, but in the last few years we haven't done much with it. I'm thankful to them for buying and holding onto that domain all this time though, as it now gives me an opportunity to build and manage my own website with a sleek domain name.

I like the idea of having a personal website [firstname].[lastname].com. In fact, since I published this website a few of my family members even reached out to me asking if they could get their own subdomain too.

## Goals

I had a few high-level goals which motivated me to build this website:

1. Make my resume publicly accessible on the internet so that it's easier to share out
2. Encourage myself to write more by having a place to post publicly accessible blog posts (and in the future, maybe scifi short stories). I've always been interested in writing scifi stories, and as Stephen King said "If you want to be a writer, you must do two things above all others: read a lot and write a lot".
3. Learn more about the web by building a website more-or-less from scratch and managing my own web infrastructure

After more than a year of creating the simon.duchastel.com website and hosting a simple "Coming soon" page on there, I finally got around to publishing a real website. I decided to write this blog post both to commemorate the occasion and reflect on my experience building it.

## Building the website with Hugo

### Why a static site generator

I'm a big fan of learning from first principles, where I understand what's happening underneath the covers before building up to the higher-level concepts. Partially for that reason, I wanted to build my website as a static site because after I'm done building my website I could read through the html/javascript/css code which was generated to deeply understand what was happening to my website. On top of that, I could tweak the low-level code fairly easily and run it locally in my browser to further understand how changes impacted the final website.

There are other advantages to a static site. For one, it's very easy to host and requires only a basic web server to run (another consideration is that my web host, GoDaddy, charges extra). Additionally, if the tool becomes deprecated, unavailable, or otherwise stops working then you still have static html that you can host and tweak indefinitely.

### Using Hugo

There are many static site generators out there, but I chose to use [Hugo](https://gohugo.io) because it's simple to use, highly-extensible, and well-supported. Also, even though it's not a priority for me I like that it's blazingly fast (my site gets generated in ~70ms). You can find the code for my Hugo website [here](https://github.com/simon-duchastel/personal-website).

Overall I'm pretty happy with Hugo. I really like the PaperMod template that [Aditya Telange](https://github.com/adityatelange) created. I also like how I can create all of my blog posts as simple markdown files and Hugo handles creating the html from there. There are lots of options built into the theme to modify how the pages are rendered and what features are included, and if I ever want to further customize how things look it's easy for me to overwrite the html templates/css/javascript.

## Go tool for uploading the website

Separately from building my website, I wanted to have a way to easily upload it to my web host. Right now my family uses GoDaddy to host duchastel.com, and they lock down a lot of things. For example, GoDaddy charges [$3-10/month for SSL encryption](https://www.godaddy.com/offers/ssl-certificate) (I found a way to do it for free, but that's a future blog post). I didn't find a way to easily upload my static site to their webserver, but then again I thought it'd be more fun to build a tool to do that for me. It would have been more efficient to at least build it via bash scripts or something, but this seemed like more fun and gave me an excuse to try out the [Go programming language](https://go.dev).

### Building the Tool



### Thoughts on Go

Overall I really like Go as a language. I really like its simplicity and how easy it is to setup, and in particular I love the language feature where you can return multiple things from a function.

I also like the patterns and idioms it uses. The fact that all of their imports download the source and build from source is cool. I think it's a little weird but overall like the fact that the compiler enforces strong style-guidelines (like no unused imports or no unused variables). Strictness on style-guide feels a bit unnecessary and sometimes overly constraining, but at the same time makes the code feel more clean and consistent and having a [strict style guide is a very google thing to do](https://google.github.io/styleguide/cppguide.html) (Google being the developers of Go).

One thing I didn't like is how much implicit name conventions are used for semantics rather than explicit keywords. For example, whether a function or type is exported from a module is entirely dependent on whether its name starts with an uppercase or lowercase letter. I guess like anything once you get used to it it probably comes naturally, but it still feels like a way to make things error-prone.

## Next steps

I have a few ideas in mind to continue this process:

1. I'd like to keep doing cool software projects and reflecting on them via blog posts on this website. I'd like to both write more in my free time (both code and non-code) and I'm hoping that this website is a tool to encourage me to do that. I recently read [Atomic Habits by James Clear](https://jamesclear.com/atomic-habits) (which I really enjoyed), and one takeaway I have from that book is that an effective strategy for building habits is to have a small goal (something that takes you less than two minutes to do) which you hit consistently. My goal is to write one line of code and one line of non-code every day, and I'm hoping that publishing the results to this website helps in motivating me to that end.
2. In addition to writing software blog posts, I think it'd be cool to write reviews of the books I read.
3. I'd like to add SSL certificate rotation to my Go tool. The smart thing to do is migrate my website to a web hosting service which supports SSL certificates by default, like [Netlify](https://www.netlify.com). I'll probably end up doing that, but in the meantime I think it'd be fun to update my tool to manually create and upload the SSL certificate to GoDaddy using SSH.
4. My goal is to eventually dip my toes into writing Science Fiction short stories. I read a lot of scifi already as I'm in love with the genre, and a life-goal of mine is to become a contributor to the scifi community. I'll probably wait on this one though until I'm used to consistently writing every day.