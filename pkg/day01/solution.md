# day 1

## part 1

Dial is a circle from 0 - 99. Starts on 50, rotating left subtracts, rotating right adds.

Therefore in the input I can replace all 'L' characters with the minus sign, and replace all the 'R' characters with nothing.

I'm left with a list of numbers that I can convert to integers. Starting with 50, I can sum up each number one by one, and after each I can check if the result is divisible by 100 (ie num % 100 == 0). If it is, I'll add one more to the counter.

## part 2

Ugh... okay, so, this is essentially the modulus and keep in range thing. Basically this: https://dev.to/timothee/using-modulo-to-shift-a-value-and-keep-it-inside-a-range-8fm

Anyways, the super easy bits:
* if rotation is 0, nothing to do, yay
* if rotation is more than 100, we can immediately have however many hundreds there are and only deal with the remainder. So for 251, I know I'll touch 2 zeros, and I can calculate with 51
* with the help of The Formula™, I can always deal with 50 + 80 = 30 situations. And 50 - 80 = 70 situations. No biggie
* I can also deal with 0 + 87 = 87 situations
* I can also calculate whether 50 + 80 or 50 - 80 touched a zero, again no issue!
* The hardest part was starting on 0, and then doing a negative turn... I was missing some logic checks, because 0 - 50 = 50, but I did not touch an extra zero

Anyways, let me clean up the code...