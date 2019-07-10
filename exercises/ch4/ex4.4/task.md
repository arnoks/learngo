# Excercise 4.4

## Description

Write a version of the rotation that operates in a single pass.

## Implmentation

The rotation of n position of the array elements has been implemented
using three subsequent reverse operations. While this is a nice
demo for inplace operations it is probably not the most efficient way
to rotate and array.

In order to use a single pass one has to swap values with appropriate offsets:
assuming the rotation is 2 and the length is 10:


0 1 2 3 4 5 6 7 8 9 
3 1 2 0 4 5 6 7 8 9 
3 4 2 0 1 5 6 7 8 9 
3 4 5 0 1 2 6 7 8 9 
3 4 5 6 1 2 0 7 8 9
3 4 5 6 7 2 0 1 8 9
3 4 5 6 7 8 0 1 2 9
3 4 5 6 7 8 9 1 2 0
3 4 5 6 7 8 9 0 2 1
3 4 5 6 7 8 9 0 1 2

Two implmentations show different performance characteristics.
The discrete implementation with swap is fast for small arrays
and large offsets, while the copy and store approach is faster far small
offset in large arrays.
