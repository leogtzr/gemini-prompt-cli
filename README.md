# gemini-prompt-cli

Calling Google Gemini AI from the terminal.

## Prerequisites

1. Get an API key: https://ai.google.dev/gemini-api/docs/api-key
2. Store the API key in anvironment variable. The code uses `GEMINI_API_KEY`
   3. You can put it in your .bashrc or .zshrc

## Build

`go build -o geminiprompt`

## Running

Using the standard input:
```
$ go build -o geminiprompt
$ ./geminiprompt                                                                            ✔  5s  base Py  18.18.1 Node  2 task  14:02:58 
> What is the golden number?
2024/05/03 14:04:17 debug: took 1.71 secs to get a response from AI

The golden number, also known as the golden ratio or divine proportion, is an irrational number approximately equal to 1.618. It is often found in nature and art, and is considered aesthetically pleasing. The golden number is represented by the Greek letter phi (φ).
> What is your name?
2024/05/03 14:04:26 debug: took 0.91 secs to get a response from AI

I am Gemini, a multimodal AI model, developed by Google.
> Escribe un poema de 5 líneas acerca de elefantes.
2024/05/03 14:04:44 debug: took 1.86 secs to get a response from AI

Gigantes grises, sabios y poderosos,
Trones de colmillos, amables y virtuosos.
Su trompa, un brazo, fuerte y delicado,
Guardias de la manada, siempre a su lado.
Sabios ancianos, pilares de la tierra,
Elefantes, un tesoro que siempre honramos.
> 
```

Using command line arguments:
```
$ ./geminiprompt "Write a simple child song about elephants"                            ✔  3m 46s  base Py  18.18.1 Node  2 task  14:06:47 
2024/05/03 14:07:29 debug: took 3.18 secs to get a response from AI

**The Elephant Song**

I saw an elephant today,
With his big gray trunk and ears.
He was walking down the street,
And I waved and he waved back.

(Chorus)
Oh, elephants are big and strong,
They're the gentlest giants we know.
They love to eat peanuts,
And take a bath in the river.

He had a big smile on his face,
And his eyes were twinkling bright.
He looked so happy and carefree,
It made me laugh with all my might.

(Chorus)
Oh, elephants are big and strong,
They're the gentlest giants we know.
They love to eat peanuts,
And take a bath in the river.
Done...

$ ./geminiprompt "Write a Golang function to calculate the Fibonacci serie"                     ✔  base Py  18.18.1 Node  2 task  14:12:13 
2024/05/03 14:12:53 debug: took 1.47 secs to get a response from AI

```go
func fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
Done...
```

