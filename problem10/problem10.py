#! /usr/bin/env python

'''
$ time python problem10.py
142913828922

real    96m42.166s
user    96m46.879s
sys 0m1.682s
'''

def isPrime(x):
    if x == 0 or x == 1:
        return False

    for i in xrange(2, x):
       if x % i == 0:
           return False

    return True

sum = 0
for i in xrange(2000000):
    if isPrime(i):
        sum += i

print sum
