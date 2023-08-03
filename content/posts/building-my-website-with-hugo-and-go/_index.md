---
title: "Creating my website with Hugo + Go"
date: 2023-08-07
---

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

There are many static site generators out there, but I chose to use [Hugo](https://gohugo.io) because it's simple to use, highly-extensible, and well-supported.


## Go tool for uploading the website


## Next steps

I have a few ideas in mind to continue this process:

1. I'd like to keep doing cool software projects and reflecting on them via blog posts on this website. I'd like to both write more in my free time (both code and non-code) and I'm hoping that this website is a tool to encourage me to do that. I recently read [Atomic Habits by James Clear](https://jamesclear.com/atomic-habits) (which I really enjoyed), and one takeaway I have from that book is that an effective strategy for building habits is to have a small goal (something that takes you less than two minutes to do) which you hit consistently. My goal is to write one line of code and one line of non-code every day, and I'm hoping that publishing the results to this website helps in motivating me to that end.
2. I'd like to add SSL certificate rotation to my Go tool. The smart thing to do is migrate my website to a web hosting service which supports SSL certificates by default, like [Netlify](https://www.netlify.com). I'll probably end up doing that, but in the meantime I think it'd be fun to update my tool to manually create and upload the SSL certificate to GoDaddy using SSH.
3. I'd like to dip my toes into writing Science Fiction short stories. I read a lot of scifi already as I'm in love with the genre, and a life-goal of mine is to become a contributor to the scifi community. I'll probably wait on this one though until I'm more used to consistently writing every day.