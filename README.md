# Shamir's Secret Sharing Scheme

Often we want to distribute a secret to multiple people and give each people a ```share``` of that secret in a way that it doesn't give away any information about the original secret.   


Highly confidential secrets shouldn't be known only by 2 people, because what if one of the party is unavailable or dies?  

This is where [**Shamir's Secret Sharing Scheme**](https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing) ```(SSSS)```comes in. It allows to ```divide``` a secret into ```n``` so called ```shares``` and specify a ```threshold``` which is a **minimum** number of shares required to reconstruct the original secret.  

## In-Depth

Let's propose the secret ```s = 832752``` and a prime```p, p > s```. We would like to divide this into ```n = 6``` parts and set the ```threshold``` value to ```k = 4```. note that ```n >= k```!

Our first important step is the realization that our secret is basically equivalent to the y co-ordinate in a 2D plane ```S``` where ```S```<sub>```x```</sub> = 0. In short, ```f(0)``` is the secret. (intersection with y-axis)

<img src="/assets/secret.jpg" alt="drawing" width="200"/>


Our secret then can be thought of, as a **[polynomial](https://en.wikipedia.org/wiki/Polynomial)** of ```k - 1```<sup>th</sup> degree, where ```f(0)``` is our secret. So our problem transposed into solving this polynomial. In fact, we don't even need to solve the entire polynomial we ought to find only ```f(0)```.  



This means that the knowledge of any ```k``` point will reconstruct the original secret. 

First step is to generate ```k - 1``` random coefficients, lets say


    1. a1 = 154
    1. a2 = 76424
    1. a3 = 1133421


Hence, our polynomial becomes: f(x) = 832752 + 154x + 76424x<sup>2</sup> + 1133421x<sup>3</sup>

Now, we can start to produce ```n``` shares:
```
1.  (463, f(463) mod p) 
2.  (88847, f(88847) mod p)
3.  (299, f(299) mod p) 
4.  (345672, f(345672) mod p)
5.  (33342, f(33342) mod p)
6.  (1734, f(1734) mod p)
```

Respresenting graphically a **3rd** degree polynomial in a 2D plane yields the result:

<img src="/assets/poly.png" alt="drawing" width="200"/>
 
Any 2 points in a 2D plane makes it possible to draw an **infinite** number of curves but only 3 points make up a specific 2nd degree polynomial curve. You can draw an **infinite** number of curves from 3 points but only 4 points make up a specific **3rd** degree polynomial, and so on.

**Note: there are infinite many points on a curve or line.**
###### The above description is using natural number integer arithmetic although please note that SSSS uses [finite field arithmetic](https://en.wikipedia.org/wiki/Finite_field).

These means that everything is taken ```mod p``` where p is a prime and p > S, P > n. The points are calculated as ```(x, f(x) mod p)``` and ```f(0)``` is reconstructed using [modular multiplicative inverse](https://en.wikipedia.org/wiki/Modular_multiplicative_inverse). Please note that there are only integers and no division operation in a finite field. Also note that a multiplicative inserve ```x``` modulo ```p``` is guaranteed to exist if ```p``` is prime. For example ```110``` and ```1832``` doesn't have a multiplicative inverse. Lastly, ```-x``` modulo ```p``` is equivalent to ```x = ((x % p) + p) % p```

Ultimately, any ```k``` points from the abovely generated set of ```n``` points will reconstruct the secret ```832752```, using [interpolation](https://en.wikipedia.org/wiki/Curve_fitting) 
Observation: if ```n == k``` then every piece of the secret is required to reconstruct the original secret.

## Original Paper
The original paper written by Adi Shamir in 1979 can be found [here.](http://web.mit.edu/6.857/OldStuff/Fall03/ref/Shamir-HowToShareASecret.pdf)
## Documentation

The program accepts 3  simple arguments, ```-s, --secret``` ```-n, --shares``` and ```-k, --threshold```

Example: ```go run main.go --secret 1234 -n 6 -k 3``` or build it with ```go build main.go``` and run it after ```./main.exe --secret 1234 -n 12 -k 3```


<img src="/assets/example.png" alt="drawing" width="1800" zoom="100"/>
  
## Feedback

If you have any feedback, please reach out to me at fezfamiliar@yahoo.com

  
## License

[MIT](https://choosealicense.com/licenses/mit/)

  
