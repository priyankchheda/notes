# C++ Functions

## Function Prototyping
The prototype describes the function interface to the compiler by giving details such as the number and type of arguments and the type of return values. With function prototyping, a template is always used when declaring and defining a function. When a function is called, the compiler uses the template to ensure that proper arfuments are passed, and the return balue is treated correctly. Any violation in matching the arguments or the return types will be caught by the compiler at the time of compilation itself.

Function prototype is a declaration statement in the calling program and is of the following form:
```cpp
type function-name (argument-lists);
```
The argument-list contains the types and names of arguments that must be passed to the function.
```cpp
float volumn(int x, float y, float z);
```

## Passing Arguments
### Pass by value
- The call/pass by value method of passing arguments to a function copies the actual value of an argument into the formal parameter of the fucntion.
- In this case, changes made to the parameter inside the function have no effect on the argument.
- By default, C++ uses call by value to pass arguments. In general, this means that code within a function cannot alter the arguments used to call the function.

### Pass by Reference
- The call/pass by reference method of passing arguments to a function copies the reference of an argument into the formal parameter.
- Inside the function, the reference is used to access the actual argument used in the call. Thus means that changes made to the parameter affect the passed argument.
- To pass the value by reference, argument reference is passed to the functions just like any other value.

### Pass by Address
- The call by pointer method of passing arguments to a function copies the address of an argument into the formal parameter.
- Inside the function, the address is used to access the actual argument used in the call. This means that changes made to the parameter affect the passed argument.
- To pass the value by pointer, argument pointers are passed to the functions just like any other value.

## Return by Reference
A function can also return a reference. Consider the following functions:
```cpp
int& max(int &x, int &y) {
    if (x > y) return x;
    return y;
}
```
Since the return type of max() is `int&`, the function returns reference to x and y (and not the values). Then a function call such as max(a, b) will yield a reference to either a or b depending on their values. This means that this function call can appear on the left-hand side of an assigment statement. That is, the statement
```cpp
max(a, b) = -1;
```
is legal and assigns -1 to a if it is larger, otherwise -1 to b.

## Inline Function
- **Defination:** If a function is inline, the compiler places a copy of the code of that function at each point where the function is called at compile time.
- Any changes to an inline function could require the function to be recompiled because compiler would need to replace all the code once again otherwise it will continue with old functionality.
```cpp
inline return-type function-name(parameters) {
    // function code
}
```
We should exercise care before making a function `inline`. The speed benefits of inline functions diminish as the function grows in size. At some point the overhead of the function call becomes small compared to the execution of the function, and the benefits of inline functions may be lost. In such cases, the use of normal functions will be more meaningful. Usually, the functions are made inline when they are small enough to be defined in one or two lines. Example:
```cpp
inline double cube(double a) { return a * a * a; }
```
Remember that the inline keyword merely sends a request, not a command, to the compiler. The compiler may ignore this request if the funciton definition is too long or too complicated and compile the function as a normal function.

Some of the situations where inline expression may not work are:
- For functions returning values, if a loop, a switch, or a goto exists.
- For functions not returning values, if a return statment exists.
- If function contain static variables.
- If inline functions are recursive.

## Default Arguments
- A default argument is a value provided in function declaration that is automatically assigned by the compiler if caller of the function doesn't provide a value for the argument with default value.
- Allows a function to be called without providing one or more trailing arguments.
- Default Arguments should always start from right to left
```cpp
int sum (int x, int y, int z = 0, int w = 0) {
    return (x + y + z + w);
}
```

## const Arguments
In C++, an argument to a function can be declared as const as shown below:
```cpp
int strlen(const char *p);
int strlen(const string &s);
```
The qualifier **const** tells the compiler that the function should not modify the argument. The compiler will generate an error when this condition is violated. This type of declaration is significant only when we pass arguments by reference or pointers.

## Function Overloading
- Function overloading is a feature in C++ where two or more functions can have the same name but different parameters.
- Function overloading can be considered as an example of polymorphism feature in C++.
```cpp
// Declarations
int add(int a, int b) // prototype 1
int add(int a, int b, int c) // prototype 2
double add(double a, double b) // prototype 3
double add(int a, double b) // prototype 4
double add(double a, int b) // prototype 5

// Function calls
cout << add(5, 10); // uses prototype 1
cout << add(5, 10.0); // uses prototype 4
cout << add(5.0, 10.2); // uses prototype 3
cout << add(5, 10, 1); // uses prototype 2
cout << add(0.5, 10); // uses prototype 5
```
A function call first matches the prototype having the same number and type of arguments and then calls the appropriate function for execution. A best match must be unique. The function selection involves the following steps:
1. The compiler first tries to find an exact match in which the types of actual arguments are the same, and use that function.
2. If an exact match is not found, the compiler uses the integral promotions to the actual arguments, such as **char** to **int** or **float** to **double** to find a match.
3. When either of them fails, the compiler tries to use the built-in conversions (the implicit assignment conversions) to the actual arguments and then uses the function whose match is unique. If the conversion is possible to have multiple matches, then the compiler will generate an error message. Suppose we use the following two functions:
```cpp
long square(long n);
double square(double x);
```
A function call such as `square(10)` will cause an error because int argument can be converted to either long or double, thereby creating an ambiguous situation as to which version of square() should be used.
4. If all of the steps fail, then the compiler will try the user-defined conversions in combination with integral promotions and built-in conversions to find a unique match. User-defined conversions are often used in handling class objects.

Overloading of the functions should be done with caution. We should not overload unrelated functions and should reserve function overloading for functions that perform closely related tasks. Sometimes, the default arguments may be used instead of overloading. They may reduce the number of functions to be defined.
