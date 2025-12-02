# day 1

## part 1

Dial is a circle from 0 - 99. Starts on 50, rotating left subtracts, rotating right adds.

Therefore in the input I can replace all 'L' characters with the minus sign, and replace all the 'R' characters with nothing.

I'm left with a list of numbers that I can convert to integers. Starting with 50, I can sum up each number one by one, and after each I can check if the result is divisible by 100 (ie num % 100 == 0). If it is, I'll add one more to the counter.