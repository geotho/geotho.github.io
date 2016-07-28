## Missing the forest for the trees in code reviews

Reading through a pull request earlier this week, I came across some code that read:

```javascript

function addPanel(p) {
  panels.push(p);
    
  if (panels.length > 0) {
    p.open();
  } 
}
```

Feeling smug, I commented that `if (panels.length > 0)` should be shortened to `if (panels.length)`. 

In hindsight, a more helpful comment would have been that `panels.length > 0` will always be true in this scenario and therefore the code is likely to be erroneous.

I'm reflecting on this in the hope I can keep sight of what the code is trying to do rather than getting caught up in the syntax.