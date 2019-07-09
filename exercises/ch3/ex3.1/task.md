# Excersice

## Desciption

If the function f returns a non-finite float64 value, the SVG file
will contain invalid \<polygon\> elements (Although many SVG renderers handle
this gracefully) Modify the program to skip invalid polygons.

## Implemetation

a. Write some test code to reproduce `math.NaN()` and `math.Inf()`
b. extend f to return a bool indating wether the calculation was ok.
c. check corenr to return ok else continue with next polygon

## Test 

