---
title: "Building my website with Hugo + Go"
date: 2023-08-06
categories: ["Software"]
---

_[View the code](https://github.com/simon-duchastel/personal-website)_

Twenty years ago, my family bought the domain duchastel.com (our family's last name). Initially my family used it to host some bios but it hasn't been updated in years. I'm thankful to them for buying and holding onto that domain though, as it now gives me an opportunity to build and manage my own website.

## Goals

I had a few high-level goals which motivated me to build this website:

1. Make my resume publicly accessible on the internet so it's easier to share out
2. Encourage myself to write more by having a public place to display my work
3. Learn more about the web by building a website from scratch

After hosting a "Coming soon" page at simon.duchastel.com for more than a year, I finally got around to publishing a real website. I decided to write this post to reflect on my experience building the site.

## Building the website with Hugo

### Why a static site generator

I'm a big fan of learning from first principles. I learn better when I first understand what's happening under the hood before building up to higher-level concepts. Partially for that reason, I wanted to build my website as a static site so that I could study the resulting html/css/javascript and learn what the generator did. I also wanted the ability to tweak the output manually and run it in my local browser to further understand how changes would impact the final result.

There are a few other reasons I chose to build a static site. For one, it's very easy to host and requires only a basic web server to run. Additionally, if the framework ever becomes deprecated I still have static html I can host and modify indefinitely. Lastly, I don't feel like I need the additional features of more complex systems.

### Using Hugo

I chose to use [Hugo](https://gohugo.io) because it's simple to use, highly extensible, and well-supported. Also, while it's not a priority for me I do think it's cool how blazingly fast it runs (my site gets generated in ~70ms on average).

Overall I'm pretty happy with Hugo. I really like [Aditya Telange's PaperMod template](https://github.com/adityatelange/hugo-PaperMod). I also enjoy writing my posts as basic markdown files and having Hugo handle translating them into beautiful html. Once everything's set up, the process is effortless yet powerful. There are lots of options built into the theme to modify how the pages are rendered and if I ever want to further customize how things look I can overwrite the templates with my own html, css, and javascript.

You can find the code for [my Hugo website here](https://github.com/simon-duchastel/personal-website).

## Go tool for deploying the website

I wanted a way to easily deploy my website. Right now my family uses GoDaddy to host duchastel.com and GoDaddy doesn't have the best tooling. I probably could have deployed my website more easily by either using a third-party tool or writing bash scripts, but I decided to build my own command line tool since that seemed more fun and gave me an excuse to try out the [Go programming language](https://go.dev).

### Building the tool

My main goal for the tool was simple: deploy my website to GoDaddy. I broke down the problem into two sub-tasks: building the website with a `build` command, and uploading the output with an `upload` command. Here's a more specific breakdown:

* **Build the website (`build`)**
    1. Clear the output directory (to ensure old files aren't accidentally uploaded)
    2. Run the Hugo build command
* **Upload the website (`upload`)**
    1. Create a backup of the existing website by downloading a copy of the website from the web server (in case there's an issue and I want to rollback)
    2. Clear the directory on the web server (to ensure there are no old files left behind)
    3. Upload the website files to the web server

I also created a convenience command `deploy`, which calls `build` followed by `upload` (if there are no errors building the website).

First, I had to be able to execute commands on the shell so I could do things like run the Hugo build command. It turns out that [Go has a convenient API](https://pkg.go.dev/os/exec#Command) for doing this in their standard library. I also used [Go APIs to manipulate the filesystem](pkg.go.dev/os). With these, I wrote code to clear out the output directory and then execute Hugo on the command line to build the website. Since I want to be able to manage everything from my tool, I also created a `preview` command which executes the Hugo command that hosts the site at _localhost:1313_. With that, I was done with `build` and moved on to `upload`.

Next, I used [Go's SSH APIs](https://pkg.go.dev/golang.org/x/crypto/ssh) (maintained by the Go team but not part of the standard library) to connect to GoDaddy's webserver. My tool reads authentication credentials from a local file (which I don't source control) and uses them to create an SSH session. The way the APIs work, you first create an SSH connection and use it to start multiple sessions. You can only run one command per session, so I created helper functions like `runRemoteCommand("ls")` so I didn't have to worry about the underlying details.

With these helper functions I could now run `rm -rf` on the remote host to clear out the old website. But before I cleared out the old website, I wanted to make sure I had a local copy as a back up in case of a problem. (As an aside, a nice part of having my website source controlled is that even without a local copy I can still rollback by checking out an old commit and re-running the upload tool).

In order to transfer the files of my old website to my machine box, I turned to [Bram Vandenbogaerde's SCP Go module](https://github.com/bramvdbogaerde/go-scp). He wrote an easy-to-use wrapper on top of Go's SSH APIs to handle both downloading and uploading files. My method for creating the local backup uses shell commands for simplicity:

1. Run `find [dir] -type f` to get the paths of all the files I need to download
2.  Download all of the files to a local directory

My initial implementation was much slower. My first solution was to run `ls -la` on the root directory, run `test -d` and `test -f` on each item returned to see whether it was a file or a directory, and recursively repeat that for every directory. My implementation with `find` is much faster than the `ls` implementation as it runs in O(1) time rather than O(3n).

Finally, once I built the functionality to list out and download files on the remote web server I reversed the process to upload my website. In fact, building `upload` was easier because I could use built-in file system APIs to list out the files I needed to upload instead of relying on `find`.

With that, I can now deploy my website by running `./website deploy` on the command line and voila — my website gets updated with all the latest changes. If I make a mistake the old version is available locally for re-upload.

### Thoughts on Go

Overall I enjoy using Go. I like its simplicity and how easy it is to setup, and in particular I love the language feature where you can return multiple values from a function.

I also like Go's patterns. For example, it's cool that imports work by referencing the module's source code, as Go is big on building from source. I also appreciate that the compiler enforces strong style-guidelines (such as no unused imports and no unused variables). Strict adherence to the style-guide cans sometime feel unnecessary but makes the code feel more clean and consistent. Plus, being [strict on style is typical of Google](https://google.github.io/styleguide/cppguide.html) (Google develops Go).

One thing I don't like is how the language uses implicit naming conventions rather than explicit keywords for certain semantics. For example, whether a function is exported from a module depends on whether its name starts with an uppercase or lowercase letter. To me that convention is too subtle and will lead to mistakes exporting the wrong functions. I imagine this is something that tooling can make more obvious and that you get used to it with time, but it still feels too error-prone.

## Next steps

I had fun building this website. I have a next-steps in mind:

1. I want to write more (that includes code as well). I'll keep writing posts reflecting on my hobby of writing software and also plan to write book reviews.
2. I want to keep expanding my Go tool. I have a few ideas:
    * Add SSL certificate rotation. The smart thing to do would be to migrate my website to a hosting provider which supports SSL certificates by default like [Netlify](https://www.netlify.com). I'll probably eventually do that, but in the meantime it'd be fun to update my tool to manually upload an SSL certificate to GoDaddy.
    * Add automatic rollback functionality. This one should be easy — I should be able to re-use my `upload` command and point it to the local backup.
    * Integrate my tool with Github actions or another form of CI/CD. That way, I don't need to run the tool myself — I can just push to my `main` branch and the website gets deployed automatically.
3. I'm thinking of switching hosting providers. Right now I use GoDaddy but am not impressed by their tooling and the fact that they up-charge on a lot of features. For example, GoDaddy charges [$3-10/month for SSL encryption and https](https://www.godaddy.com/offers/ssl-certificate) (I found a way to get https on my website for free, but that's a future post). I could probably save money and get more features by switching to another hosting provider.