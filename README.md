# What's gorest

`gorest` is a resource-oriented HTTP handler that lets you create HTTP server focusing on your resources rather than handling.

[![Go Report Card](https://goreportcard.com/badge/github.com/fredmaggiowski/gorest)](https://goreportcard.com/report/github.com/fredmaggiowski/gorest)&nbsp;
[![Build Status](https://travis-ci.org/fredmaggiowski/gorest.svg?branch=master)](https://travis-ci.org/fredmaggiowski/gorest)&nbsp;
[![GoDoc](https://godoc.org/github.com/fredmaggiowski/gorest?status.svg)](https://godoc.org/github.com/fredmaggiowski/gorest)&nbsp;
[![codecov](https://codecov.io/gh/fredmaggiowski/gorest/branch/master/graph/badge.svg)](https://codecov.io/gh/fredmaggiowski/gorest)&nbsp;
[![GitHub issues](https://img.shields.io/github/issues/fredmaggiowski/gorest.svg "GitHub issues")](https://github.com/fredmaggiowski/gorest)

All you need to do when using `gorest` is to create your resources and, once created the handler, assign them to the desired routes.

## Example

```go
package main

import (
    "net/http"
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
    // Remeber to pass the pointer of the resource and not the resource itself
    // otherwise the server will not work!
    handler.SetRoutes([]*gorest.Route{
        gorest.NewRoute(&resource1, "/resource/1"),
        gorest.NewRoute(&resource2, "/resource/2"),
        gorest.NewRoute(&resource3, "/resource/3"),
    })

    // Get the handler for your HTTP(S) server.
    router := handler.GetMuxRouter(nil)
    http.ListenAndServe("localhost:80", router)
}
```

## Concept

The concept beneath the gorest core is to use Go structures and associated-functions in order to provide the means to create a simple and well-structured API server.

In order to achieve this gorest uses Go interfaces to understand whether a user-defined resource supports the request HTTP method and act accordingly.

You just have to define your structure and implement only the methods you need (like: `Get`, `Post`, `Put`, etc.)

## Anatomy of a `Resource`

If you look through the code you might notice that a `Resource`, for GoRest, is pretty much everything (since it is a `interface{}` type).

What makes gorest pick the right method for the right resource is another set of interfaces (`GetSupported`, `PostSupported`, etc.) that are used to detect whether your resource implements the method used in the HTTP request.

### A `Resource` example

Suppose you have to provide an API in order to interact with your blog posts, You might want the following features to be available:
 - `GET`: retrieves a post based on provided identifier;
 - `POST`: creates a new post using data provided in the request body;
 - `DELETE`: deletes a post based on provided identifier;

In order to implement these features you can do as follows:

```go
package myresources

import(
    "net/http"
    "github.com/fredmaggiowski/gorest"
)

type PostResource struct {}

func (p *PostResource) Get(r *http.Request) (int, gorest.Response) {
    // Parse the request body or the URI in order to 
    // retrieve the post ID.

    // Get the post from your database.
    if postHasNotBeenFound {
        return http.StatusNotFound, nil
    }

    // Insert the post in a Response object
    response := gorest.NewStandardResponse()
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

As of now `gorest` does not implement the `http.Handler` interface making it impossible to be used directly as handler in the `ListenAndServe` function.

In order to grant the power of multiplexing requests we decided to keep a dependcy on the [gorilla/mux](https://github.com/gorilla/mux) package.

This will be the first thing that will be removed in the next release since we desire GoRest to be a standalone framework for API definition and management.
