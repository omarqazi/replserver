# REPL Server Development Journal 001
__June 24, 2015 01:56 AM / Los Angeles, California__

## Goals
I am going to try and write a tool that I help might be useful for myself and other developers in a variety of situations. I don't have an official name for it yet, but I'm refering to it in my head as "REPL Anywhere".

## What is it?
The idea is pretty simple. It's a combination of a client library and server. I'll probably host a server and maybe have an option to compile and run your own server as well. The server will be implemented in Go, and I'll implement a client library in Ruby.

Basically the goal is that you should be able to add a simple method call anywhere Ruby code is running and then be able to open up a REPL at that point in the code remotely by maintaining a connection to the server and having an authenticated terminal user push lines of code for the listening client library to run. The client library will send back the results of each line so that the client library can display those results to the user.

## Why would anyone want that?
I don't know. Debugging? Tech support? Testing? I'm sure people will figure out uses for it that I can't think of.

## Isn't this dangerous?
Yeah, it could be really dangerous. I'll try and design it to be as secure as possible. Hopefully this doesn't get used irresponsibly too much.

## What are these journals?
These documents help serve two purposes:
1) To help me figure out what i'm doing
2) To document what I was thinking when I wrote the code

If you have any questions, send email to omar.qazi on iCloud.