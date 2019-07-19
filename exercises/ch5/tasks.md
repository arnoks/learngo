# Chapter 5 Exercises

## Open

1. Exercise 5.1 findlinksrecursive
2. Exercise 5.2 counttags & elementstats
3. Exercise 5.3 textnodes
4. Exercise 5.4 findothers
5. Exercise 5.5 countwordsandimages
6. Exercise 5.6 see countwordsandimages and surfaces
7. Exercise 5.7 htmlpretty
   Develop startElement and endElement into a general HTML
   pretty-printer. Print comment nodes, text nodes and attributes
   of each element (\<a href '...'>). USer short forms like \<img/\> 
   instead of \<img\>\</img\> when an element has no children. Write
   a test to ensure that the output can be pasrsed Successfully.
8. 
9.
10.
11.
12.
13.
14.
15.
16.
17.

## Closed

1. Exercise 5.1 - findlinksrecursive
   Change the findlinks program to traverse the
   n.FirstChild linked list using recursive calls
   to visit instead of a loop (findlinksrec)

2. Exercise 5.2 - counttags
   Write a function to populate a mapping from element
   names p, div, span and so on - to the number of elements
   with that name in HTML document tree.

3. Exercise 5.3 - textnodes
   Write a function to print the contents of all text nodes
   in an HTML document tree. Do not descend into \<script\> or
   \<style\> elemts, since their contents are not visisble in
   web browser.

4. Exercise 5.4 findothers
   Extend the visit function so that it extracts other kinds
   of links from the document, such as images, scripts and style
   sheets.

5. Exercise 5.5 countwordsandimages
   Implement countWordsAndImages. (See Exercise 4.9 for word-splitting).

6. Exercise 5.6
   Modify the corner function in ch3/surface to use named results
   and a bare return statement.
