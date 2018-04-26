### Hashtable solution
* The key here is creating a hash.
* Then doing an initial iteration through the passed array and:
  * Recording the element as the __key__.
  * Recording the index of the element in the array as the __value__.
* Doing the initial pass through the array to store the array elements and their
  indices prevents us from having to nest a loop to get access to two elements
  at one time.
  * This makes the run time `O(n)` rather than `O(n^2)`

__Key Takeaway__: Always try to look for ways to __NOT__ nest a loop. `O(n^2)`
  should be *worst-case scenario*
