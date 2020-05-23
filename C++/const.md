# Const Keyword

A compile time constraint that an object can not be modified.

```cpp
const int i = 9;
i = 6; // compilation error: can't modify read-only variable
```

In the case of const fundamental data types, initialization can be done through copy, direct, or uniform initialization:
```cpp
const int value1 = 5; // copy initialization
const int value2(7); // direct initialization
const int value3 { 9 }; // uniform initialization (C++11)
```

## Constant Parameters
```cpp
const int* ptr = &i; // data is const, pointer is not
*ptr = 5 // compilation error: can't modify read-only variable
ptr++ // successfully compiled

int* const p2 = nullptr; // pointer is const, data is not

const int* const p3; // data and pointer both are const
```
`int const * p4 = &i;` is name as `const int* p4 = &i;`

If const is on the left of `*`, data is const<br>
If const is on the right of `*`, pointer is const

## Why use const?
- Guards against inadvertent write to the variable
- self documenting
- enables compiler to do more optimization, making code tigher
- const means the variable can be put in ROM


## Const and Functions

### const parameters
const reference parameter prevents function from modifing that parameter. It is widely used in C++ functions.
```cpp
// const reference is passed as a parameter to a function.
void setAge(const int& a);
```

Using const parameter (without reference) is not that much useful because the parameter itself is a copy of actual argument. Hence, modifing parameter `a` in function will not affect main argument in any way.
```cpp
void setAge(const int a); // not much useful as 'a' is just a copy.
```

### Const Return Value
returning const reference from function.
```cpp
const string& getName() { return name; }
```

const is completely useless as return value is just a copy.
```cpp
const string getName() { return name; }
```

### Const Function
```cpp
void printDogName() const { cout << name << endl; }
```
- This function will not change any of the member variable of this class.
- Const function can only call another const functions in order to maintain the const correctness.

printDogName can be overloaded to non const version like
```cpp
void printDogName() { cout << name << endl; }
```

Const function will be invoked when the dog object is const. Other function will be invoked when dog object is not const.
```cpp
Dog d; d.printDogName(); // calls normal non const version
const Dog d2; d2.printDogName(); // calls const version
```

## Const objects
Instantiated class objects can also be made const by using the const keyword. Initialization is done via class constructors:
```cpp
const Date date1; // initialize using default constructor
const Date date2(2020, 10, 16); // initialize using parameterized constructor
const Date date3 { 2020, 10, 16 }; // initialize using parameterized constructor (C++11)
```
- Once a const class object has been initialized via constructor, any attempt to modify the member variables of the object is disallowed, as it would violate the const-ness of the object. This includes both changing member variables directly (if they are public), or calling member functions that set the value of member variables
- Const class object can only call const member functions. Non-const class object can call both const and non-const member functions.
- Note that constructors cannot be marked as const. This is because constructors need to be able to initialize their member variables, and a const constructor would not be able to do so. Consequently, the language disallows const constructors.
> Rule: Make any member function that does not modify the state of the class object const, so that it can be called by const objects

## Overloading const and non-const function

Finally, although it is not done very often, it is possible to overload a function in such a way to have a const and non-const version of the same function:
```cpp
#include <string>
class Something {
private:
    std::string m_value;
public:
    Something(const std::string &value=""): m_value{ value } {}
    const std::string& getValue() const { return m_value; } // getValue() for const objects
    std::string& getValue() { return m_value; } // getValue() for non-const objects
};
```
The const version of the function will be called on any const objects, and the non-const version will be called on any non-const objects:
```cpp
int main() {
	Something something{};
	something.getValue() = "Hi"; // calls non-const getValue();
	const Something something2{};
	something2.getValue(); // calls const getValue();
	return 0;
}
```
Overloading a function with a const and non-const version is typically done when the return value needs to differ in constness. In the example above, the non-const version of getValue() will only work with non-const objects, but is more flexible in that we can use it to both read and write m_value (which we do by assigning the string “Hi”).

The const version of getValue() will work with either const or non-const objects, but returns a const reference, to ensure we can’t modify the const object’s data.

This works because the const-ness of the function is considered part of the function’s signature, so a const and non-const function which differ only in const-ness are considered distinct.

**Summary:** Because passing objects by const reference is common, your classes should be const-friendly. That means making any member function that does not modify the state of the class object const!

# Mutable Keyword

Mutable keyword can only be applied to non-static and non-const data members of a class. If a data member is declared mutable, then it is legal to assign a value to this data member from a const member function.

For example, the following code will compile without error because m_accessCount has been declared to be mutable, and therefore can be modified by GetFlag even though GetFlag is a const member function.

```cpp
class X {
private:
   bool m_flag;
   mutable int m_accessCount;
public:
    bool GetFlag() const {
        m_accessCount++;
        return m_flag;
    }
};
```
> TODO: mutable keyword in lambda

**Reference:**
- [Advanced C++: const](https://www.youtube.com/watch?v=7arYbAhu0aw)
- [Advanced C++: const and Functions](https://www.youtube.com/watch?v=RC7uE_wl1Uc)
- [CONST in C++](https://www.youtube.com/watch?v=4fJBrditnJU)
- [The Mutable Keyword in C++](https://www.youtube.com/watch?v=bP9z3H3cVMY)
- [8.10 — Const class objects and member functions](https://www.learncpp.com/cpp-tutorial/810-const-class-objects-and-member-functions/)