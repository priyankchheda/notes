# Constructors
- A constructor is a special type of member function that initializes an object automatically when it is created.
- Constructor has the same name as that of the class and it does not have any return type.
- The constructor is always public (almost always).
- A constructor that accepts no parameters is called the _default constructor_. The default constructor for class A is A::A(). If no such constructor is defined, then the compiler supplies a default constructor.
- A class can have multiple constructors with different signatures.

The constructor dunctions have some special characteristics:
- They should be declared in the public section.
- They are involed automatically when the objects are created.
- They do not have return types, not even void and therefore, they cannot return values.
- They cannot be inherited, though a derived class can call the base class constructor.
- Like other C++ functions, they can have default arguments.
- Constructors cannot be virtual.
- We cannot refer to their addresses.
- An object with a constructor (or destructor) can not be used as a member of a union.

## Parameterized Constructors
The constructor that can take arguments are called _parameterized constructors_.
```cpp
class Integer {
    int m, n;
public:
    Integer(int x, int y) // parameterized constructor
    { m = x; n = y; }
};
```
There are two ways to call the parameterized constructor
- By calling the constructor explicitly
```cpp
    Integer int1 = Integer(0, 100); // explicit call
```
- By calling the constructor implicitly
```cpp
    Integer int1(0, 100); // implicit call
```

## Constructors with Default Arguments
It is possible to define constructors with default arguments. For example, the constructor `Complex()` can be declared as follows:
```cpp
Complex(float real, float imag=0);
```
It is important to distinguish between the default constructor `A::A()` and the default argument constructor `A::A(int = 0)`. The default argument constructor can be called with wither one argument or no arguments. When called with no arguments, it becomes a default constructor. When both these forms are used in a class, it causes ambiguity for a statement such as `A a;`. The ambigiuity is whether to call `A::A()` or `A::A(int = 0)`.

## Copy Constructor
The parameters of a construstor can be of any type except that of the class to which it belongs. For example,
```cpp
class class-name {
public:
    class-name(class-name) // illegal
};
```
However, a construstor can accept a reference to its own class as a parameter. Thus the statement
```cpp
class class-name {
public:
    class-name(class-name&) // legal
};
```
Such constructor is called _copy constructor_.

A copy constructor is used to declare and initialize an object from another object. For example, the statement `class-name obj2(obj1);` would define the object obj2 and at the same time initialize it to the values of obj1. Another form of this statement is `classname obj2 = obj1`.

The process of initializing through a copy constructor is known as _copy initialization_. Remember, the statement `obj2 = obj1` will not invoke the copy constructor. However, if obj1 and obj2 are objects, this statement is legal and simply assigns the values og obj1 to obj2, member-by-member.

When no copy constructor is defined, the compiler supplies its own copy constructor.

## const Objects
We may create and use constant objects using **const** keyword before object declaration. For example, we may create X as a constant object of the class matrix as follows:
```cpp
const matric X(m, n) // object X is constant
```
Any attempt to modify the values of m and n will generate compile-time error. Further, a constant object can call only const member functions. As we know, a const member is a function prototype or function definition where the keyword const appears after the function's signature.

Whenever const objects try to invoke non-const member functions, the compiler generates errors.


# Destructors
A destructor, as the name implies, is used to destroy the objects that have been reated by a constructor. Like a constructor, the destructor is a member function whose name is the same as the class name but is preceded by a tilde. For example, the destructor for the class `A` can be defined as `~A(){}`.

A destructor never takes any argument nor does it return any value. It will be invoked implicitly by the compiler upon exit from the program (or block or function as the case may be) to clean up storage that is no longer accessible. It is a good practice to declare destructors in a program since it releases memory space for future use.

Whenever new is allocated memory in the construtors, we should use delete to free that memory.