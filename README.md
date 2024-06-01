# CRM Backend
A simple CRM software developed in Go for the final project of the Udacity Golang course.

## Getting Started

This CRM backend project provides a basic HTTP server implementation for the following API endpoints:

HTTP Method | URL Path        | Description
----------- | --------------- | ----------------------------
GET         | /customers      | Gets a list of all customers
GET         | /customers/{id} | Gets a single customer
POST        | /customers      | Adds a new customer
PUT         | /customers/{id} | Updates an existing customer
DELETE      | /customers/{id} | Deletes a customer

## Local Setup

When setting up the project on your computer, first of all make sure you have **Go** installed. You can download the latest version 
[here](https://go.dev/dl/) and verify the installation by executing the `go version` command in a terminal window.

To run the CRM backend on your local machine, navigate to the project's root directory and execute the command `go run main.go`. The CRM backend HTTP server can then be accessed at http://localhost:3000.
