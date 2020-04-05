## Introduction
- In 1989, _Alexander Stepanov_ and _David Musser_ published a patper in which they coin the term generic programming and present a detailed definition.
- Later in 1994, _Alexander Stepanov_ and _Meng Lee_ of Hewlett-Packard developer a set of general-purpose templatized classes (data structures) and functions (algorithms) that could be used as a standard approach for storing and processing of data. The collection of these generic classes and functions is called Standard Template Library (STL).
- The STL has now become a part of ANSI startard C++ class library. STL components which are now part of the Standard C++ Library are defined in the **namespace std**.

## Components of STL
The STL contains several components. But at its core are three key components. They are:
- **Containers**: A container is an object that actually stores data. It is a way data is organized in memory. The STL containers are implemented by template classes and therefore can be easily customized to hold different types of data.
- **Algorithms**: A algorithm is a procedure that is used to process the data contained in the containers. The STL includes many different kinds of algorithms to provide support to tasks such as initializing, searching, copying, sorting, and merging. Algorithms are implemented by template functions.
- **Iterators**: An iterator is an object (like a pointer) that points to an element in a container. We can use iterators to move through the contents of containers. Iterators are handled just like pointers. We can increment or decrement them. Iterators connect algorithms with containers and play a key role in the manipulation of data stored in the containers.

These three components work in conjunction with one another to provide support to a variety of programming solutions.

## Containers
Containers are objects that hold data (of same type). There are three categories of containers:

### Sequence Containers:
Sequence containers store elements in a linear sequence. Each element is related to other elements by its position along the line. They all expand themselves to allow insertion of elements and all of them support a number of operations on them.

- Vector: A dynamic array. Allows insertions and deletions at back. Permits direct access to any element.
- List: A bidirectional, linear list (Doubly Linked List). Allows insertions and deletions anywhere.
- Forward List: A singly linked List. Allows insertions and deletions anywhere.
- Deque: A double-ended queue. Allows insertions and deletions at both the ends. Permits direct access to any element.

Elements in all these containers can be accessed using an iterator. The difference between the three of them is related to only their performance.

### Associative Containers
Associative containers are designed to support direct access to elements using keys. They are not sequential.
- Set: An associative container for storing unique sets. Allows rapid lookup. (No duplicates allowed).
- MultiSet: An associative container for storing non-unique sets. (Duplicates allowed).
- Map: An associative container for storing unique key/value pairs. Each key is associated with only one value (one-to-one mapping). Allows key-based lookup.
- MultiMap: An associative container for storing key/value pairs in which one key may be associated with more than one value (one-to-many mapping). Allows key-based lookup.

There are two sub categories of Associative Containers:
- Ordered Associative Containers
- Unordered Associative Containers

The main difference is that ordered/regular associative containers are implemented with _self-balancing binary tree_ under the hood, meaning that they can be traversed, iterated through in a particular order and are relatively fast to read or write. Whereas, the unordered variants are implemented with _hash tables_. So they are even faster on average but have a very bad worse case performance.

Containers **set** and **multiset** can store a number of items and provide operations for manipulating them using the values as the keys. For example, a set might store objects of the student class which are ordered alphabetically using names as keys. We can search for a desired student using his name as the key. The main difference between a set and a multiset is that a multiset allows duplicate items while a set does not.

Container **map** and **multimap** are used to store pairs of items, one called a key and the other called the value. We can manipulate the values using the keys associated with them. The values are sometimes called _mapped values_. The difference between a map and a multimap is that a map allows only one value for a given key to be stored while multimap permits multiple values.

### Derived Containers
This are also known as container adapters. Container adapters are interfaces that implement some special functionality on yop of a sequence container. The C++ STL suppoers the following container.
- Stack: A standard stack. Last-in-first-out (LIFO). The ideal choice for the underlying container here is the deque. Although a vector or list could be used here too.
- Queue: A standard queue. First-in-first-out (FIFO). This is also typically implemented on top of a deque but a list or a vector could be used as well.
- Priority Queue: A priority queue. The first element out is always the highest priority element. This is naturally performed by an array based data structure called a heap. Thus, its underlying container is usually a vector. But it could also be a deque.

All of these underlying container options are possible because container adapters are templates too.

Stacks, Queues and Priority Queues can be created from different sequence containers. The derived containers do not support iterators and therefore we cannot use them for data manipulation. However, they support two member functions pop() and push() for implementing deleting and inserting operations.

## Algorithms
Algorithms are functions that can be used generally across a variety of containers for processing their contents. Although each container provides functions for its basic operations, STL provides more than sixty standard algorithms to support more extended or complex opertions. Standard algorithms also permit us to work with two different types of containers at the same time. Remember, STL algorithms are not member functions or friends of containers. They are standalone template functions.

STL algorithms reinforce the philosophy of reusability. To have access to the STL algorithms, we must include \<algorithm\> in out programs.

STL algorithms, based on the nature of operations they perform, maybe categorized as under:
- Retrieve or non-mutating algorithms
- Mutating algorithms
- Sorting algorithms
- Set algorithms
- Relational algorithms

## Iterators
Iterators behave like pointers and are used to access container elements. They are often used to traverse from one element to another, a process known as _iterating_ through the container.

There are five types of iterators as describe below:

| Iterator | Access Method | Direction of movement | I/O capability | Remark |
| -------- | ------------- | --------------------- | -------------- | ------ |
| Input | Linear | Forward Only | Read Only | Cannot be saved |
| Output | Linear | Forward Only | Write Only | Cannot be saved |
| Forward | Linear | Forward Only | Read/Write | Can be saved |
| Bidirectional | Linear | Forward and backward | Read/Write | Can be saved |
| Random | Random | Forward and backward | Read/Write | Can be saved |

Iterators in the C++ STL are defined as templates, and they must comply with a very specific set of rules in order to qualify as one of many types of iterators. Different types of iterator must be used with the different types of containers. Note that only sequence and associative containers are traversable with iterators.

### Types of iterators:
Based upon the functionality of the iterators, they can be classified into five major categories:

- **Input Iterators:** They are the weakest of all the iterators and have very limited functionality. They can only be used in a single-pass algorithms, i.e., those algorithms which process the container sequentially such that no element is accessed more than once.
- **Output Iterators:** Just like input iterators, they are also very limited in their functionality and can only be used in single-pass algorithm, but not for accessing elements, but for being assigned elements.
- **Forward Iterator:** They are higher in hierarachy than input and output iterators, and contain all the features present in these two iterators. But, as the name suggests, they also can only move in forward direction and that too one step at a time.
- **Bidirectional Iterators:** They have all the features of forward iterators along with the fact that they overcome the drawback of forward iterators, as they can move in both the directions, that is why their name is bidirectional.
- **Random-Access Iterators:** They are the most powerful iterators. They are not limited to moving sequentially, as their name suggests, they can randomly access any element inside the container. They are the ones whose functionality is same as pointers.

Each type of iterator is used for performing certain functions. Below figures gives the functionality Venn diagram of the iterators. It illustrates the level of functionality provided by different categories of iterators.

![iterator_venn_diagram](images/iteratorVennDiagram.png)

The input and output iterators support the least functions. They can be used only to traverse in a container. The forward iterator supports all operations of input and output iterators and also retains its position in the container. A bidirectional iterator, while supporting all forward iterator operations, provides the ability to move in the backward direction in the container. A random access iterator combines the functionality of a bidirectional iterator with an ability to jump to an arbitrary location.

### Containers and their corresponding Iterators:
|Iterator|Containers|
|---|---|
|Forward|forward_list<br>unorderer_set, unorderer_multiset<br>unorderer_map, unorderer_multimap|
|Bidirectional|list<br>set, multiset<br>map, multimap|
|Random Access|vector<br>deque<br>array|

### Four important Iterator functions
Here are four important functions in the C++ STL supported by most containers. They all return an iterator, which is supposed to serve as the starting point of the loop that will somehow traverse the container.
- **begin()**: The begin function returns an iterator to the beginning of the container.
- **end()**: The end function is similar to begin, but it returns an iterator to the slot past the last element so that if begin and end return equal iterators, then the container is empty.
- **rbegin()**: Rbegin returns an iterator to the last element in the container, which is the first element in the reversed container.
- **rend()**: Rend returns an iterator to the slot before the first element in the container. As you can imagine, these functions make sense when the container can be traversed backwards.

_Note_: begin() and end() are available for all containers in the STL. Rbegin() and rend() are implemented for all containers except forward list.

### Iterator Invalidation
- When there are changes, such as insertions or deletions in containers that are currently using iterators, these iterators may become invalid.
- Invalid doesn't mean that it will cause a runtime error, but rather, that the iterator is no longer guaranteed to have access to the same element it was referring to before the change.
- There are many well-documented rules for iterator invalidation in the description of each container and its member functions.
- Depending on the implementation of containers, invalidation may or may not happen. For instance, the insert function does not invalidate iterators in linked lists, because the iterators are pointers to the actual objects, and the insertion does not relocate any elements in linked lists. So keep in mind, that some functions may invalidate iterators, so you must always read the documentation.

## Function Objects or Functors
A function objects or functors are classes that overload the parenthesis operator so that they can be used as functions that maintain state and can be parameterized. There's some tension on the naming of these elements. Some programmer like to call them function objects while other prefer functors

Function objects are often used as arguments to certain containers and algorithms. For example,the statement

```cpp
sort(array, array+5, greater<int>())
```

uses the function object `greater<int>()` to sort the elements contained in array in descending order.

Besides comparisons, STL provides many other predefined function objects for performing arithmetical and logical operations as shown in table below. Note that there are function objects corresponding to all the major C++ opertors. For using function objects, we must include \<functional\> header file.

|Function Object|Type|Description|
|---|---|:---:|
|divides<T>|arithmetic|x / y|
|equal_to<T>|relational|x == y|
|greater<T>|relational|x > y|
|greater_equal<T>|relational|x >= y|
|less<T>|relational|x < y|
|less_equal<T>|relational|x <= y|
|logical_and<T>|logical|x && y|
|logical_not<T>|logical|!x|
|logical_or<T>|logical|x || y|
|minus<T>|arithmetic|x - y|
|modulus<T>|arithmetic|x % y|
|negate<T>|relational|-x|
|not_equal_to<T>|arithmetic|x != y|
|plus<T>|arithmetic|x + y|
|multiplies<T>|arithmetic|x * y|

**Note:** The variables x and y represent objects of class T passed to the function object as arguments.