# GoRest

GoRest is a powerful resource-oriented HTTP handler.

[![Go Report Card](https://goreportcard.com/badge/github.com/fredmaggiowski/gorest)](https://goreportcard.com/report/github.com/fredmaggiowski/gorest)&nbsp;
[![Build Status](https://travis-ci.org/fredmaggiowski/gorest.svg?branch=master)](https://travis-ci.org/fredmaggiowski/gorest)&nbsp;
[![GoDoc](https://godoc.org/github.com/fredmaggiowski/gorest?status.svg)](https://godoc.org/github.com/fredmaggiowski/gorest)&nbsp;
[![codecov](https://codecov.io/gh/fredmaggiowski/gorest/branch/master/graph/badge.svg)](https://codecov.io/gh/fredmaggiowski/gorest)&nbsp;
[![GitHub issues](https://img.shields.io/github/issues/fredmaggiowski/gorest.svg "GitHub issues")](https://github.com/fredmaggiowski/gorest)


Using GoRest extremely simple, all you need to do is create the handler, create your resources and decide on which path you want them to be available and register your desired _Routes_.

## Example

```go
package main

import (
    "github.com/fredmaggiowski/gorest"
)

func main() {
    // Create a new handler
    handler := gorest.NewHandler()

    // Define and setup your custom structures.
    var resource1 Resource1
    var resource2 Resource2
    var resource3 Resource2

    // Register the routes.
    handler.SetRoutes([]*gorest.Route{
        gorest.NewRoute(resource1, "/resource/1"),
        gorest.NewRoute(resource2, "/resource/2"),
        gorest.NewRoute(resource3, "/resource/3"),
    })

    // Get the handler for your HTTP(S) server.
    router := gorest.GetMuxRouter(nil)
    http.ListenAndServe("localhost:80", router)
}
```

## Concept

The concept beneath the gorest core is to use Go structures and methods in order to provide the means to create a simple and well-structured API server.

In order to achieve this gorest uses Go interfaces to understand whether a user-defined resource supports the request HTTP method and act accordingly.

You just have to define your structure and implement only the methods you need (like: `Get`, `Post`, `Put`, etc.)

## Anatomy of a `Resource`

If you look through the code you might notice that a `Resource`, for GoRest, is pretty much everything (since it is a `interface{}` type).

What makes gorest pick the right method for the right resource is another set of interfaces (`GetSupported`, `PostSupported`, etc.) that are used to detect whether your resource implements the method used in the HTTP request.

### A `Resource` example

Let's suppose you have to provide an API in order to interact with your blog posts.

You might the following features:
 - GET: retrieves a post based on provided identifier;
 - POST: creates a new post using data provided in the request body;
 - DELETE: deletes a post based on provided identifier;

In order to implement these features we could create the following:

```go
package myresources

type PostResource struct {}

func (p *PostResource) Get(r *http.Request) (int, gorest.Response) {
    // Parse the request body or the URI in order to 
    // retrieve the post ID.

    // Get the post from your database.
    if postHasNotBeenFound {
        return http.StatusNotFound, nil
    }

    // Insert the post in a Response object
    response := NewResponse()
    response.SetBody(post)

    return http.StatusOK, response
}

func (p *PostResource) Post(r *http.Request) (int, gorest.Response) {
    // Parse the request body.

    // Update the post

    // Return status and response
}


func (p *PostResource) Delete(r *http.Request) (int, gorest.Response) {
    // Parse the request body.

    // Delete the post.

    // Return status and response
}
```

## Anatomy of a `Response`

TBD

## Roadmap

As of now GoRest does not implement the `http.Handler` interface making it impossible to be used directly as handler in the `ListenAndServe` function.

In order to grant the power of multiplexing requests we decided to keep a dependcy on the [gorilla/mux](https://github.com/gorilla/mux) package.

This will be the first thing that will be removed in the next release since we desire GoRest to be a standalone framework for API definition and management.
