---
layout: post
title:  "Reject unrecognised parameters"
date:   2016-02-22 17:56:47 +0000
categories: code
---
**TL;DR:** your API or library should error if it is passed parameters it does
not recognise.

I recently spent a week debugging a friend's JavaScript code.
A datepicker library seemingly wasn't picking up changes to the date format,
supplied through options like so:

{% highlight javascript %}
$("#datepicker").datepicker({
  "dateFormat": "yy-mm-dd",
});
{% endhighlight %}

We arrived at the solution after 43 emails –
the code was using `bootstrap-datepicker` while the options were from `jQueryUI-datapicker`.
We should have been using `format` rather than `dateFormat`.
Both libraries use the same method, leading to a confusing half–working scenario.

Had the library failed noisily on unexpected parameters,
the error would have been picked up almost immediately.

Your API or library should return errors or fail noisily if it is passed unexpected arguments.
This helps your users understand that they are doing something wrong and lets them know how to fix it.
