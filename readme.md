# Cron Expression Parser

## Assignment
Write a command line application or script that parses a cron string and expands each field
to show the times at which it will run. You may use whichever language you feel most
comfortable with. The assignment must have an automated test suite.

## What is Cron and Cron Expression  

Cron is a general term used by most operating systems to describe job scheduling (known as Cron Jobs). The Cron command was originally made mainstream as part of the UNIX operating system back in the 1960s and 70s. Using Cron, a user can instruct a machine to execute a task in a specified interval. One way to describe this interval is using Cron Expressions, sometimes also called Cron Directives.

A cron expression is a string conforming to a pre-defined format that tell a computer program (or a person) how often a task needs to run. There are a couple of variants, depending on the program and/or operating system. Multiple operating systems like Linux and UNIX uses cron expressions. Likewise, various tools and libraries uses them, like Azure Functions and Quartz.NET. A cron expression in its basic form consist of 5 strings separated by spaces:

## * * * * *
At first, looking at 5 stars doesn't make much sense. Until you understand how to look at the characters in the magic string. Each block of text represents a unit in time for when to execute the job:

minute hour day month day_of_week
The placement of each * specity minute, hour, etc. A star meaning 'every' in this context. Rather using a star, specific times can be specified like this:

## Example
5 */4 12 7 1
This cron expression will execute the job every monday (1 day of the week), only on July 12, every fourth hour, at exactly 5 minutes over the hour.


## About the assignment
I tried to cover most of the cases mentioned in the assignment, error handling and edge cases are something which can be improved but due to time shortage I am not able to work much onto this part. 

A few things to double-check:

Error Handling – Right now program assumes all input values are valid. 
I have to add validations to handle incorrect cron formats or non-numeric values. For example: 
It would be good to handle and test following cases:
*/0 (invalid step)
60 in minutes (out of range)
13 in months (out of range)
Missing fields in the input.
