## C Structures Revisited
C structures provide a method for packing together data of different types. It is a user-defined data type with a template that serves to define its data properties.

```cpp
struct student {
    char name[20];
    int roll_number;
    float total_marks;
};
```

The keyword struct declares student as a new data type that can hold three fields of different data types. These fields are known as _structure members_ or _elements_. The identifier student, which referred to as _strucure name_ or _structure tag_, can be used to create variables of type student.

```cpp
struct student A;

strcpy(A.name, "John");
A.roll_number = 99;
A.total_marks = 595.5;
FinalTotal = A.total_marks + 5;
```

Structures can have arrays, pointers or structures as members.

### Limitations of C Structures
- Can not perform arithmetic operations on two similar structures `c = c1 + c2` where c, c1 and c2 are complex numbers.
- They do not permit data hiding. Structure members can be directly accessed by the structure variables by any function anywhere in their scope. In other words, the structure members are public members.

## Extensions to Sturctures
- C++ supports all the features of structures as defined in C. But C++ has expanded its capabilities further to suit its OOP philosophy.
- In C++, a structure can have both variables and functions as members. It can also declare some of its members as `private`  so that they can not be accessed directly by the external functions.
- In C++, the structure names are stand-alone and can be used like any other type names. In other words, the keyword struct can be omitted in the declaration of structure variables. For example, we can declare the student variable A as
```cpp
student A; // C++ declaration. This is an error in C.
```
- C++ incoporates all these extension in another user-defined type known as **class**. There is very little syntactical difference between structures and classes in C++ and, therefore, they can be used interchangeably with minor modifications.
- Since class is a specially introduced data type in C++, most of the C++ programmers tend to use the structure for holding only data, and classes to hold both the data and functions.
- The only difference between a structure and a class in C++ is that, by default, the members of a class are _private_, while, by default, the member of a structure are _public_.

## Specifying a Class
A class is a way to bind the data and its associated functions together. It allows the data (and functions) to be hidden, if necessary, for external use.

The general form of a class declaration is:
```cpp
class class_name {
    private:
        variable declarations;
        function declaration;
    public:
        variable declarations;
        function declaration;
};
```
- The functions and variables are collectively called as _class members_.
- The keywords private and public are known as _visibility labels_.
- The variables declared inside the class are known as _data members_ and the functions are known as _member functions_. Only member functions can have access to the private data members and member functions. However, the public members (both functions and data) can be accessed from outside the class.
- The binding of data and functions together into a single class-type variable is referred to as _encapsulation_.

### Creating Objects
- Once a class is declared, we can create variables of that type by using class name (like any other built-in type variable). For example,

```cpp
class_name x; // memory for x is created
```
- In C++, the class variables are known as objects. Therefore, x is called an object of type *class_name*. We may also declare more than one object in one statement. Example `class_name x, y, z;`
- The declaration of an object is similar to that of a variable of any basic type. The necessary memory space is allocated to an object at this stage. Note that class specification, like a structure, provides only a template and does not create any memory space for the objects.

## Accessing Class Members
- You can access public members from outside class directly using following syntax.
```cpp
object-name.function_name(actual-arguments)
```

## Defining Member Functions
Member functions can be defined in two places:
- Outside the class definition
- Inside the class definition

It is obvious the, irrespective of the place of definition, the function should perform the same task. Therefore, the code for the function body would be identical in both the cases.

### Outside the Class Definition
- Member functions that are declared inside a class have to be defined separately outside the class. Their definitions are very much like the normal functions.
- The important difference between a member function and a normal function is that a member function incorporates a membership 'identity label' in the header. This 'label' tells the compiler which class the functions belongs to. The general form of a member function definition is:
```cpp
return-type class-name :: function-name(argument-declaration) {
    function-body;
}
```
- The membership label _class-name ::_ tells the compiler that the function _function-name_ belongs to the class _class-name_. That is, the scope of the function is restricted to the _class-name_ specified in the header line. The symbol :: is called the _scope resolution_ operator.
- The member functions have some special characteristics that are often used in the programs development. These characteristics are:
    - Several different classes can use the same function name. The 'membership label' will resolve their scope.
    - Member functions can access the private data of the class. A non-member function cannot do so. (However, an expection to this rule is a friend function)
    - A member function can call another member function directly, without using the dot operator.

### Inside the Class Definition
- Another method of defining a member function is to replace the function declaration by the actual function definition inside the class. For example, we could define the item class as follows:
```cpp
class item {
    int number;
    float cost;
public:
    void getdata(int a, int b); // declaration
    void putdata() { // definition inside the class aka inline function
        cout << number << "\n";
        cout << cost << "\n";
    }
};
```
- When a function is defined inside a class, it is treated as an inline function. Therefore, all the restrictions and limitations that apply to an incline function are also applicable here. Normally, only small functions are defined inside the class definition.

### Making an Outside Function Inline
- One of the objective of OOP is to separate the details of implementation from the class definition. It is therefore good practice to define the member functions outside the class.
- We can define a member function outside the class definition and still make it inline by just using the qualifier inline in the header line of function definition.
```cpp
inline void item::getData(int a, float b) { // definition
    number = a;
    cost = b;
}
```

## Nesting of Member Functions
We have mentioned that a member function of a class can be called only by an object of that class using a dot operator. However, there is an exception to this. A member function can be called by using its name inside another member function of the same class. This is known as _nesting of member functions_.

## Private Member Functions
- Although it is normal practice to place all the data items in a private section and all the functions in public, some situations may require certain functions to be hidden (like private data) from outside calls. Tasks such as deleting an account in a customer file, or providing increment to an employee are events of serious consequences and therefore the functions handling such tasks should have restrictied access. We can place these functions in the private sections.
- A private member function can only be called by another function that is a member of its class. Even an object cannot invoke a private function using the dot operator.

Consider a class as defined below:
```cpp
class sample {
    int m;
    void read(); // private member function
public:
    void update();
    void write();
}
```
If s1 is an object of sample, then `s1.read()` won't work because objects cannot access private members. However, the function **read()** can be called by the function **update()** to update the value of m.

## Memory Allocation for Objects
We have stated that the memory space for objects is allocated when they are declared and not when the class is specified. This statement is only partly true. Actually, the member functions are created and placed in the memory space only once when they are defined as a part of class specification. Since all the objects belonging to that class use the same member functions, no separate space is allocated for member functions when the objects are created. Only space for member variables is allocated separately for each object. Separate memory locations for the objects are essential, because the member variables will hold different data values for different objects.

## Static Data Members
- A data member of a class can be qualified as static. The properties of a static member variable are similar to that of a C static variable.
- A static member variable has certain special characteristics. These are:
    - It is initialized to zero when the first object of its class is created. No other initialization is permitted.
    - Only one copy of that member is created for the entire class and is shared by all the objects of that class, no matter how many objects are created.
    - It is visible only within the class, but its lifetime is the entire program.
- Static variables are normally used to maintain values common to the entire class. For example, a static data member can be used as a counter that records the occurrences of all the objects.
```cpp
class item {
    static int count;
    int number;
public:
    void getdata(int a) { number = a; count++; }
    void getcount() { cout << "count: " << count << "\n"; }
};

int item::count; // definition of static data member

int main() {
    item a, b, c;
    a.getcount(); // count: 0
    b.getcount(); // count: 0
    c.getcount(); // count: 0

    a.getdata(100);
    b.getdata(200);
    c.getdata(300);

    a.getcount(); // count: 3
    b.getcount(); // count: 3
    c.getcount(); // count: 3
}
```
- Note that the type and scope of each static member variable must be defined outside the class definition. This is necessary because the static data members are stored separately rather than as a part of an object. Since they are associated with the class itself rather than with any class object, they are also known as **class variables**.
- Static variables are like non-inline member functions as they are declared in a class declaration and defined in the source file. While defining a static variable, some initial value can also be assigned to the variable.
For instance, the following definition gives count the initial value 10.
```cpp
int item::count = 10;
```

## Static Member Functions
Like static member variable, we can also have static member functions. A member function that is declared static has the following properties:
- A static function can have access to only other static members (functions or variables) declared in the same class.
- A static member function can be called using the class name (instead of its objects) as `class-name::function-name`.
Non-static member fucntions can access all data members of the class: static and non-static. Whereas, static member functions can only operate on the static data members.

## Objects as Function Arguments
Like any other data type, an object may be used as a function argument. This can be done in two ways:
- A copy of the entire object is passed to the function (_pass-by-value_).
- Only the address of the object is transffered to the function (_pass-by-reference_).

## const Member Functions
If a member function does not alter any data in the class, then we may declare it as a const member function as follows:
```cpp
void mul(int, int) const;
double get_balance() const;
```
The qualifier const is appended to the function prototypes (in both declaration and definition). The compiler will generate an error message if such functions try to alter the data values.

## Pointers to Members
It is possible to take the address of a member of a class and assign it to a pointer. The address of a member can be obtained by applying the operator & to a "fully qualified" class member name. A class member pointer can be declared using the operator `::*` with the class name. For example, given the class
```cpp
class A {
private:
    int m;
public:
    void show();
};
```
We can define a pointer to the member m as follows:
```cpp
int A ::* ip = &A :: m;
```
The ip pointer created thus acts like a class member in that it must be invoked with a class object. In the statement above, the phrase `A::*` means "pointer-to-member of A class". The phrase `&A::m` means the "address of the m member of A class".

Remember, the following statement is not valid:
```cpp
int *ip = &m; // won't work
```

This is because m is not simply an int type data. It has meaning only when it is associated with the class to which it belongs. The scope operator must be applied to both the pointer and the member.

The pointer ip can now be used to access the member m inside member functions(or friend functions). Let us assume that a is an object of A declared in a member function. We can access m using the pointer ip as follows:
```cpp
cout << a.*ip; // display
cout << a.m;   // same as above
```

## Local Classes
Classes can be defined and used inside a function or a block. Such classes are called local classes. Examples:
```cpp
void test(int a) { // function
    ...
    ...
    ...
    class student { // local class
        ...
        ...     // class definition
    };
    ...
    student s1(a); // create student object
    ...            // use student object
}
```

Local classes can use global variables (declared above the function) and static variables declared inside the function but cannot use automatic local variables. The global variables should be used with the scope operator (::).

There are some restrictions in constructing local classes. They cannot have static data members and member functions must be defined inside the local classes. Enclosing function cannot access the pricate members of a local class. However, we can achieve this by declaring the enclosing function as a friend.