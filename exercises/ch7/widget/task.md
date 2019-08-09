# Widget sort Vs Stable Sort

 Many Guis provide a table widget with a stateful multi-tier sort: the
 primary sort key is most recently clicked column head, the secondary sort
 key is the second-most recently clicked column head, and so on. Define an
 implementation of sort.Interface for use by such a table. Compare that 
 approach with repeated sorting using sort.Stable

## Soluiton Approach

The sort implementation maintains has list of Less function for all possible
headers rows. And executes them in order of the selection, applying all
pressed header in order using the customSort approach of chapter 7 sorting
example.
