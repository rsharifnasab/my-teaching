# session 7

## review 
+ interface 
+ abstract class 
+ implements, multiple inheritance 
+ cast 
+ poly morphism 

+ static field, static method with poly morphism

## exception handling 
+ what if program doesn't go well?
+ int f(int a,int b){ return a/b; }
+ index out of bound 
+ what should we return? what should we do?
+ **let the caller manage it**
+ C: return null, set global error variables, terminate =)
-----
+ high level langauges: exception handling framework
+ throw, new Exception
+ pop from stack method by method!
+ we don't catch yet
-----
+ catch -> manage thrown exceptions
+ try, catch
+ print stack trace: please don't :)
+ ignore: please don't
----
+ finally
+ try finally
+ try catch finally 
+ return in finally
-------
+ checked exception vs unchecked exception 
+ throws keyword 
+ exception hierarchy
+ Throwable
![](https://i.stack.imgur.com/GsVNp.jpg)
-------
+ design exception class 
+ constructors


## generic 
+ method or class that can act on different types
+ same as poly-morphism! 
+ can't use specific behavior of type, because it's unknown 
+ type parameter `<V>`
+ generic method
+ why not use Object? (java 1-4)
```java
    public static <V> V f(V v){
        return v;
    }
```

+ now we use wrapper classes!

+ generic class
+ erasure, Object, cast

## collections 
+ we want to store objects 
+ arrays are not enough: add, remove, dynamic size
+ lists are better, same as linked list in ITP
+ get, set method
+ arrayList vs linkedList 
+ List is interface, implement it with arrayList and LinkedList
+ set, map: interface 
+ use set with iterator or `contains`
+ use map with `get(key)`
+ different implementations: hashSet, hashMap, treeSet, ..
+ iterator, iterable (interface), for-each

![](https://dzone.com/storage/temp/1821372-class-and-interface-hierarchy.png)

