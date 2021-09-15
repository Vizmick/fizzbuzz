
# Fizzbuzz app

A simple app exposing an endpoint solving the fizzbuzz exercise.  The endpoint /fizzbuzz accepts 5 parameters. int1,int2 and limit are integers; str1 et str2 are strings. It returns a list of strings  representing integers from 1 to limit, with multiple of int1 replaced by str1, multiples of int2 replace by str2 and multiples of both replaced by str1str2  
The endpoint /frequent returns the most frequent request addressed to /fizzbuzz, along with the count of calls. Only valid requests are counted.


### How to run the app

You need to have mysql installed. Logged in to mysql, create a database "fizzbuzz" with the command

>create database fizzbuzz;

Then, to initialize the database, cd to this folder and run:

>source data.sql

In the fizzbuzz_db.go file, replace username and password.

To run tests, run

> go test -v ./fb

To run the app, run

>go run .

You can now test the endpoints with

>curl "http://localhost:8080/fizzbuzz?int1=2&int2=3&limit=6&str1=two&str2=three"  
>curl "http://localhost:8080/frequent"

