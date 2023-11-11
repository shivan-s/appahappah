# Appahappah

[![Deploy](https://github.com/shivan-s/appahappah/actions/workflows/main.yml/badge.svg)](https://github.com/shivan-s/appahappah/actions/workflows/main.yml)

Appahappah is what I would call my paternal grandfather. He passed away on the 4th of November 2023.

This project is dedicated to him.

## The Project

This is a miniature static site generator written in Golang. The inspiration was not only to honour my Appahappah but also to learn Golang (I usually used Typescript for work and have used Python in the past quite extensively). The project takes markdown files and converts this into HTML with the use of the standard library as well as some third party packages.

Github actions is responsible for running the Go code, which turns the markdown into HTML as well as moving some images and CSS files. This is then hosted statically using Github pages.

## How it works ++

It starts with the `/static` directory which contains a `layout.html` file; this file is boiler plate for the markdown files in the `/content` directory for when they are compiled into HTML. There are three different markdown files, `header.md`, `index.md`, and `footer.md` so that styling can be applied to them individually inside their `<header>`, `<main>` and `<footer>` HTML tags respectively.

The `static/css/style.css` file is copied to the `/public` directory along with a freshly created `index.html` file from the Go code. Images with extensions `jpg`, `jpeg`, and `png` are also copied to the `/public/img` directory. The `/public` directory is not committed to the repository but this is what GitHub pages uses to deliver content.
