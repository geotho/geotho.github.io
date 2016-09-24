---
layout: page
title: Projects
permalink: /projects/
---

Personal projects I have done, presented most recent first:

### September

#### super-strava-boy

[Demo](https://geotho.github.io/super-strava-boy) ·
[GitHub](https://github.com/geotho/super-strava-boy)

A Super-Meat-Boy x Strava commute visualiser using the Strava API, Google Maps and some linear interpolation.

#### protobuf-to-typescript

[Demo](http://geotho.github.io/protobuf-to-typescript) ·
[GitHub](https://github.com/geotho/protobuf-to-typescript)

Converting Protocol buffer .proto files to TypeScript .d.ts files in browser. Essentially a more manual and error-prone `protoc`.

### August

#### strava-stealer

[Writeup](http://geotho.github.io/2016/09/04/recovering-privacy-zones-strava.html) ·
[Demo](http://geotho.github.io/strava-stealer) ·
[GitHub](https://github.com/geotho/strava-stealer)

How bicycle thieves trained in the art of gradient descent could calculate the origin of your Strava privacy zones.

#### ascii-graph

[GitHub](https://github.com/geotho/ascii-graph)

A Go library that takes graphs written as strings like:
```
1   2
 \ /
  5
 / \
3   4
```

and converts them to Go code like:
```go
g.AddEdge("1", "5")
g.AddEdge("2", "5")
g.AddEdge("3", "5")
g.AddEdge("4", "5")
```

Originally motivated by confusing unit test code for graph algorithms.

### July

#### kakuro-solver

[Writeup](http://geotho.github.io/code/2016/07/07/kakuro-solving.html) ·
[Demo](http://geotho.github.io/kakuro-solver) ·
[GitHub](https://github.com/geotho/kakuro-solver)

Kakuros are like Sudokus but are more entertaining. This JavaScript app solves them using different tricks.
