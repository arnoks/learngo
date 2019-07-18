# Chapter 6 Exercises Oveview

## Open

1. Exercise 6.1 Done
2. Exercise 6.2 Done
3. Exercise 6.3 Done
4. Exercise 6.4 Done
5. Exercise 6.5 Done

## Closed

1. Excercise 6.1 Addititional Methods for Intset
   Add Len, Remove, Clear, Copy methods
2. Excercise 6.2 Variadic IntSet
   Add the AddAll method to the IntSet
3. Excercise 6.3 Add Intersections Difference etc to IntSet
4. Exercise 6.4
   Add a method Elems that returns a slice
   containing the elemets of the set, suitable for iterating
   over with a range loop.
5. Exercise 6.5
   The type of each word used by IntSet is uint64, but 64 bit
   arithmetic may be inefficient on a 32-bit platform. Modify
   the program to use the uint type, which is the most efficient
   unsigned integer type for the platform. Instead of deviding
   by 64, define a constant holding the effective size of uint
   in bits 32 or 64. You can use the perhaps too-clever expression
   32<<(^uint(0))>>63 for this purpose.
