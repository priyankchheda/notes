## Generic Programming
- In 1989, _Alexander Stepanov_ and _David Musser_ published a patper in which they coin the term generic programming and present a detailed definition.
- A programming paradigm where data types are not specified in the implementation of code. Instead, data types are specified in the lines of code that instantiate, or otherwise use those pieces of reusable code.

### Compile-Time Polymorphism
- Generic Programming is sometimes referred to as Compile-Time Polymorphism.
- Specifying data types at variable definition of object instantiation is, in essence, polymorphism resolved at compile time.
- Potentially more efficient than runtime polymorphism.
- Compile-Time polymorphism may be computationally less expensive than the traditional runtime polymorphism because most of the calculations and adaptations that need to be performed at achieve polymorphism will happend once during compile time.

## Class Template
Templates allows us to do Generic Programming in C++. Generic Programming is an approach where generic types are used as parameters in algorithms so that they work for a variety of suitable data types and data structures.

A template can be considered as a kind of macro. When an object of a specific type is defined for actual use, the template definition for that class is substituted with the required data type. Since a template is defined with a parameter that would be replaced by a specified data type at the time of actual use of the class or function, the templates are sometimes called parameterized classes or functions.

```cpp
#include <iostream>

template<typename T>
class Vector {
    T* v;
    int size;
public:
    //create a null Vector
    Vector(int m) {
        size = m;
        v = new T[size];
        for (int i = 0; i < size; i++)
            v[i] = 0;
    }

    // create a vector from an array
    Vector(T* arr, int m) {
        size = m;
        v = new T[size];
        for (int i = 0; i < size; i++)
            v[i] = arr[i];
    }

    // scalar product
    T operator * (Vector &other) {
        T sum = 0;
        for (int i = 0; i < size; i++)
            sum += this->v[i] * other.v[i];
        return sum;
    }
};

int main()
{
    int x[3] = {1, 2, 3};
    int y[3] = {4, 5, 6};
    Vector<int> v1_int(x, 3);
    Vector<int> v2_int(y, 3);

    int result_int = v1_int * v2_int;
    std::cout << "Result Int: " << result_int << "\n";

    float a[3] = {1.4, 2.2, 3.4};
    float b[3] = {4.1, 5.2, 6.0};
    Vector<float> v1_float(a, 3);
    Vector<float> v2_float(b, 3);

    float result_float = v1_float * v2_float;
    std::cout << "Result Float: " << result_float << "\n";

    return 0;
}
```

This process of creating a specific class from a class template is called instantiation. The compiler will perform the error analysis only when an instantiation takes place. It is, therefore, advisable to create and debug an ordinary class before converting it into a a template.


## Class Templates with Multiple Parameters

We can use more than one generic data type in a class template. They are declared as a comma-separated list within the template specification as shown below

```cpp
#include <iostream>

template<typename T1, typename T2>
class Test {
    T1 a;
    T2 b;
public:
    Test(T1 x, T2 y) { a = x; b = y; }
    void show() { std::cout << a << " and " << b << "\n"; }
};

int main() {
    Test<float, int> test1(1.23, 123);
    Test<int, char> test2(45, 'W');
    test1.show(); // 1.23 and 123
    test2.show(); // 45 and W
    return 0;
}
```


## Function Templates

Like class templates, we can also define function templates that could be used to create a family of functions with different argument types. The general format of a function template is:

```cpp
template<typename T>
void swap(T&a, T&b) {
    T temp = a;
    a = b;
    b = temp;
}
```

A function generated from a function template is called a template function. Below program demonstrates the use of two template functions in nested form for implementing the bubble sort algorithm.

```cpp
#include <iostream>

template<typename T>
void swap(T &a, T &b);

template<typename T>
void bubble(T a[], int n) {
    for (int i = 0; i < n - 1; i++)
        for (int j = n - 1; i < j; j--)
            if (a[j] < a[j - 1])
                swap(a[j], a[j-1]);
}

template<typename T>
void swap(T &a, T &b) {
    T temp = a; a = b; b = temp;
}

int main() {
    int x[5] = { 10, 50, 30, 20, 40 };
    float y[5] = { 1.1, 4.4, 5.5, 3.3, 2.2 };

    bubble(x, 5); // 10 20 30 40 50
    bubble(y, 5); // 1.1 2.2 3.3 4.4 5.5
}
```

## Function Templates with Multiple Parameters

Like template classes, we can use more than one generic data type in the template statement, using a comma-separated list as shown below:

```cpp
template<typename T1, typename T2>
void display(T1 a, T2 b) {
    std::cout << a << " " << b << "\n";
}
```





## Overloading of Template Functions

A template function may be overloaded either by template functions or ordinary functions of its name. In such cases, the overloading resolution is accomplished as follows:
- Call an ordinary function that has an exact match.
- Call a template function that could be created with an exact match.
- Try normal overloading resolution to ordinary functions and call the one that matches.

An error is generated if no match is found. Note that no automatic conversions are applied to arguments on the template functions. Below program shows how a template function is overloaded with an explicit function.

```cpp
#include <iostream>
template<typename T>
void display(T a) {
    std::cout << "Template display: " << a << "\n";
}

void display(int a) {
    std::cout << "Explicit display: " << a << "\n";
}

int main() {
    display(100); // Explicit display: 100
    display(10.45); // Template display: 10.45
    display('c'); // Template display: c
}
```

## Member Function Templates

When we created a class template for vector, all the member functions were defined as inline which was not necessary. We could have defined them outside the class as well. But remember that the member functions of the template classes themselves are parameterized by the type argument (to their template classes) and therefore these functions must be defined by the function templates.

The vector class template and its member functions are redefined as follows:

```cpp
// Class templates
template<typename T>
class Vector {
    T* v;
    int size;
public:
    Vector(int m);
    Vector(T* arr, int m);
    T operator * (Vector &other);
};

// Member function templates
template<typename T>
Vector<T>::Vector(int m) {
    size = m;
    v = new T[size];
    for (int i = 0; i < size; i++)
        v[i] = 0;
}

template<typename T>
Vector<T>::Vector(T* arr, int m) {
    size = m;
    v = new T[size];
    for (int i = 0; i < size; i++)
        v[i] = arr[i];
}

template<typename T>
T Vector<T>::operator * (Vector &other) {
    T sum = 0;
    for (int i = 0; i < size; i++)
        sum += this->v[i] * other.v[i];
    return sum;
}
```

**Recommendation:** Write C<T> as just C everywhere inside the scope of a class template C<T>.

## Non-Type Template Arguments

We have seen that a template can have multiple arguments. It is also possible to use non-type arguments. That is, in addition to the type argument T, we can also use other arguments such as strings, function names, constant expressions and built-in types. Consider the following example:

```cpp
template<typename T, int size>
class Array {
    T a[size];
    // ...
}
```

This template supplies the size of the array as an argument. This implies that the size of the array is known to the compiler at the compile time itself. The arguments must be specified whenever a template class is created. For example,

```cpp
Array<int, 10> a1; // Array of 10 integers
Array<float, 5> a1; // Array of 10 floats
Array<char, 20> a1; // String of size 20
```

The size is given as an argument to the template class.


## typename vs. class

You can declare a template type parameter using the keyword class instead of typename, as in

```cpp
template<class T> // means the same as template<typename T>
void swap(T &a, T &b);
```

This swap will still swap objects of class or non-class types.
The keywords class and typename aren't interchangeable elsewhere in C++.

_Why is this?_<br>
When C++ introduced templates way back in early 90s, the keyword typename didn't exist and Bjarne Stroustrup had a justifiable aversion to inventing new keywords so he tried to repurpose old keywords and class seems to be the one that make most sense.
But later on, standard committee saw that typename is better word in this syntax. 

## Template Arguments Revisited

A template argument can be the name of another tempate instantiation.
Here, ratios is a list with elements of type rational<int>:

```cpp
e.g. list <rational<int> > rations; // space between `>` is needed
```
In C++ 03, you had to leave a space between consecutive angle brackets:

```cpp
e.g. list<rational<int>> ratios; // error in C++03
                                 // OK since C++11
```
A C++03 compiler would have interpreted >> as shift operator.
C++11 fixed this annoyance.
C++ now recogizes >> as two separate brackets in this context.


## Headers and Source Files

With non-template functions and classes, you typically:
- place declarations in headers and
- definitions in source files.

You treat templates differently. You typically place all template declarations, including definitions, in headers.
- you rarely place template declarations in source files.

## Template Argument Deduction

- C++ often lets you omit the angle-bracketed template argument list from a call to a function template.
- That is, the call can use just the template name as the function name. Example `i = abs(j); // abs, not abs<?>`
- In this case, the compiler performs template argument deduction:
    - It deduces the template argument(s) from the function call argument(s).
- Template argument deduction makes a function template look more like an unbounded set of overloaded functions.
- Type deduction fails when the arguments don't match well enough
- Prior to C++17, type argument deduction applied only to function templates, not to class templates.

```cpp
rational<int> r1, r2;
swap(r1, r2); // OK
rational<int> r3(r2); // OK
rational r4(r1); // error until
```

## Two-Phase Translation
- The compiler is limited in what it can do when it first encounters a template definition such as:

```cpp
template<typename T>
T foo (T x) {...}
```

- It can't generate code for an instantiation. It doesn't know what T is, yet.
- For the most part, the compiler just scoops up the template and stores it in the symbol table.
- On this first reading, the compiler can't detect all possible errors.
- Nevertheless, it still tries to do as much checking as it can, so it can report errors as early as possible.
- Later, when it encounters a call such as abs(i), the compiler deduces the type argument from the function argument.
- It then substitutes the deduced type for T and does complete syntactic and semantic analysis as it compiles the instantiation.
- Thus, a compiler processes each template definition in two phases:
    - The 1st phase occurs when the compiler parses the template declaration. This happens just once for template.
    - Each 2nd phace occurs when the compiler instantiates the template for a particular combination of template arguments. This happens at each instantiation.