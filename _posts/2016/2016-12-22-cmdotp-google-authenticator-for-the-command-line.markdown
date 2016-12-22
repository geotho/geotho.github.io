---
layout: "post"
title: "cmdotp - Google Authenticator for the command line"
date: "2016-12-22 16:03"
---

When my phone broke, I no longer had access to the 2FA codes stored in
Google Authenticator. This can be a pretty perilous situation, especially if
you haven't written down your backup codes.

I've written [`cmdotp`](https://github.com/geotho/cmdotp), which you can use to
store your generate 2FA codes on your computer. They are stored encrypted in
your home folder. It requires a password any time you want to access the codes.

You could use something like [Authy](https://www.authy.com/) but there's
something about cloud-syncing my 2FA codes that doesn't sit right with me.

Is this a dangerous thing to do? Isn't the whole point that your phone is a
second factor, separate from your laptop?

Good second-factor authentication relies on something you know (your password)
and something you have (your phone). The time-based one-time passwords are an
implementation detail: a way to prove you have access to a particular item, be
it a phone, or now, a laptop.

One threat model is that someone might steal my phone, somehow crack my pin, and
then have access to my 2FA codes. By using `cmdotp`, you introduce another
threat model which is that someone could do the same to your laptop. I believe
the risk from these two is negligible compared to the risk of an attacker: (1)
socially engineering my phone provider to get SMS backup codes forwarded to some
other phone number or (2) [threatening to break my kneecaps](https://xkcd.com/538/).
