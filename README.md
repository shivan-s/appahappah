# Appahappah

[![Deploy](https://github.com/shivan-s/appahappah/actions/workflows/main.yml/badge.svg)](https://github.com/shivan-s/appahappah/actions/workflows/main.yml)

As a Sri Lankan Tamil, _Appahappah_ is what I would call my paternal grandfather. On the 4th of November 2023, he passed away in hospital.

This project is dedicated to him.

## Motivation

Honouring my grandfather is one aspect, but education is also another. This could have been written in simple HTML and CSS. However, I wanted to include Golang in the mix, as that is the language I want to learn after extensive experience (relative to myself) with Typescript and Python.

## How it works

Put simply, this is a miniature static site generator.

It's easy to begin with the `/static` directory, where we have `css/styles.css`, which contains our CSS rules, along with an `/img` directory which contains our images. These files are copied by the Go code and placed in a `/public` directory on the build step before the site is distributed via a CDN.

Also, in the `/static` directory is a template HTML file, `layout.html`. This is used by the Go code, where in the build step, the markdown files in `/content` are translated into HTML and then inserted into `layout.html`. This creates `index.html` in the `/public` directory.

## Deployment

As mentioned, Golang is used in the build step to create files in the `/public` directory. Notice that this directory is not committed to the source repository. Instead, GitHub Actions will initialise this build process with every commit. With a `/public` directory set as an artifact, this can be used by GitHub Pages for hosting. In addition to GitHub Pages acting as a CDN host, Cloudflare is also being used to - a Go worker is used to complete the build step before deploying the output onto their CDN.

## Local Development

Only [Go (version 1.20 or above)](https://go.dev/) is required, but other tools can be useful.

These include:

- `[air](https://github.com/cosmtrek/air)` - This runs the Go build step command every time there is a change in the code written. This saves having to run `go run main.go` after every single code change. All you need to do is run `air` in command line as the configuration handles the rest.
- `[live-server](https://github.com/tapio/live-server)` - `cd ./public` and once in this directory run `npx live-server .`. This will server the `index.html` file via HTTP server and can be accessed usually via `http://localhost:8080`.

## Feedback

Thanks for reading this and I hope you find it useful and feel free to get in touch if you have any feedback, questions or suggestions.
