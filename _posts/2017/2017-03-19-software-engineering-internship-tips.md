---
layout: "post"
title: "Tips for software engineering internships"
date: "2017-03-19 21:20"
---

I've done four internships, two at startups and two at Google. Since joining Ravelin full time, I've had to interview, hire and mentor interns. Here are some tips I wish someone had told me before starting my internships, and indeed that I continually need reminders of even as a full time software engineer:

## Remember Hofstadter's Law: it always takes longer than you expect, even when you take into account Hofstadter's Law
Writing production-ready software takes a great deal longer than writing it for personal or school use.

At home, writing a program that does X is usually a one-step process:
1. Write program that does X.

Whereas in production, you might have some additional considerations:
1. Ascertain whether X really meets business requirements, or whether Y might be better.
2. Find a reasonable way to do Y on the large amount of data before the heat death of the universe.
3. Write program that does Y.
4. Write unit tests and integration tests.
5. Realise you broke Z while writing a thing that does Y, so fix that.
6. Add metrics and logging.
7. Write deployment scripts.
8. Work out how to update / backfill all the old data that hasn't had Y done to it.

And then there are all the unknown unknowns that will throw a spanner in the works that are, by definition, impossible to plan for.

## Limit project scope as much as possible

Because things will take longer than you think, it is important to limit your project's scope as much as possible. It is easy for mentors to forget how much more productive they can be, because of familiarity with the codebase, experience programming and proper scoping. You should be working on something that at first seems trivial to accomplish: the unforeseen technical challenges or unspecified business cases will eat up the majority of your time.

## Review your own code first

Before asking for code reviews from your peers or mentors, have a read through your own pull request first. Have you left debug statements in there? Is the code style consistent and readable? Did you copy and paste code but forget to update the comments? Try to spot all the issues you are able to spot before sending it on for review. That way, comments on your code review will focus on stuff you didn't already know.

## Obey the 10 minute rule
I stole the 10 minute rule from [someone who works at DeepMind (who, in turn, stole it from elsewhere).](https://www.reddit.com/r/MachineLearning/comments/4w6tsv/ama_we_are_the_google_brain_team_wed_love_to/d6diast/) If you get stuck, try to solve the problem for 10 minutes. If, after 10 minutes you are stuck, you have to ask someone else. If you ask sooner, you are wasting other people's time. If you ask later, you are wasting your own. And yours is important, too!

## Read and review other people's code 
One of the best ways to learn is to do code reviews for other members of the team. How do they organise their code? How do they name variables? How do they write testable code and tests? It's also a great opportunity to ask questions about stuff that is unclear to you, and find out about advanced language or library features you didn't know existed.

Interns and junior engineers often feel underqualified to review the code of more experienced engineers. But I would urge you to anyway because:
- you can still ask questions and learn something;
- experienced engineers still make simple mistakes (e.g. copy and pasting errors) that you can spot;
- code that is unclear to you might be able to be improved for readability;
- even comments along the lines of "I have read this code and have spotted no obvious errors" are useful to receive, and make engineers feel loved.
