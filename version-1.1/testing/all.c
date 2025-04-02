#include <stdio.h>  // For standard input/output functions like printf
#include <stdlib.h> // For memory allocation (malloc, free), exit
#include <string.h> // For string manipulation functions (strcpy, strlen)
#include <stdbool.h> // For boolean type (C99 and later)

#define PI 3.14159       // Defining a constant macro
#define GREETING "Hello" // Defining a string macro

// Conditional compilation
#define DEBUG_LEVEL 1

// Function prototype (declaration)
int add(int a, int b);
void print_message(const char *message);

// --- Section 13: Structs Definition ---
// Define a structure blueprint
struct Point {
    int x;
    double y;
    char label[10];
};

// --- Section 14: Enums Definition ---
// Define an enumeration
enum Status {
    PENDING,
    PROCESSING,
    COMPLETED,
    FAILED
};

// --- Section 15: Typedef ---
typedef unsigned long long ULL; // Typedef for a built-in type
typedef struct Point Coordinate; // Typedef for a user-defined structure
typedef enum Status TaskStatus;  // Typedef for an enum

// --- Main Function: Entry point of the program ---
int main() {
    printf("--- C Syntax Showcase Start ---\n\n");

    // --- Section 1: Variables, Data Types & Basic Assignment ---
    printf("--- Section 1: Variables & Data Types ---\n");
    int integerVar = 10;
    float floatVar = 2.5f; // 'f' suffix for float literals
    double doubleVar = 5.123456789;
    char charVar = 'A';
    bool boolVar = true; // Requires <stdbool.h>
    const int constantVar = 100; // Value cannot be changed

    // constantVar = 101; // This would cause a compile error

    printf("int: %d, float: %.2f, double: %lf, char: %c, bool: %d\n",
           integerVar, floatVar, doubleVar, charVar, boolVar);
    printf("Constant int: %d\n", constantVar);
    printf("Macro PI: %f\n", PI);
    printf("\n");


    // --- Section 2: Operators ---
    printf("--- Section 2: Operators ---\n");
    int a = 15, b = 4;
    int sum_res = a + b;      // Addition
    int diff_res = a - b;     // Subtraction
    int prod_res = a * b;     // Multiplication
    int quot_res = a / b;     // Integer Division
    int mod_res = a % b;      // Modulo (remainder)
    printf("Arithmetic: %d + %d = %d, %d / %d = %d, %d %% %d = %d\n",
           a, b, sum_res, a, b, quot_res, a, b, mod_res);

    a++; // Increment
    b--; // Decrement
    printf("Increment/Decrement: a is now %d, b is now %d\n", a, b);

    // Relational and Logical Operators (often used with if)
    bool isGreater = (a > b);       // Greater than
    bool isEqual = (sum_res == 19); // Equal to
    bool logicalAnd = (isGreater && isEqual); // Logical AND
    bool logicalOr = (isGreater || (b < 0)); // Logical OR
    bool logicalNot = !isEqual;     // Logical NOT
    printf("Relational/Logical: isGreater: %d, isEqual: %d, AND: %d, OR: %d, NOT: %d\n",
           isGreater, isEqual, logicalAnd, logicalOr, logicalNot);

    // Bitwise Operators
    int x = 5; // 0101
    int y = 3; // 0011
    int bitwiseAnd = x & y; // 0001 -> 1
    int bitwiseOr = x | y;  // 0111 -> 7
    int bitwiseXor = x ^ y; // 0110 -> 6
    int bitwiseNot = ~x;    // 1...11111010 (depends on int size)
    int leftShift = x << 1; // 1010 -> 10
    int rightShift = x >> 1;// 0010 -> 2
    printf("Bitwise: AND: %d, OR: %d, XOR: %d, NOT: %d, <<1: %d, >>1: %d\n",
           bitwiseAnd, bitwiseOr, bitwiseXor, bitwiseNot, leftShift, rightShift);

    // Assignment Operators
    a += 5; // a = a + 5;
    printf("Assignment Operator: a += 5 -> a = %d\n", a);

    // Ternary Operator
    int maxVal = (a > b) ? a : b;
    printf("Ternary Operator: max(%d, %d) = %d\n", a, b, maxVal);

    // sizeof Operator
    size_t intSize = sizeof(int);
    size_t pointSize = sizeof(struct Point);
    printf("sizeof: int is %zu bytes, struct Point is %zu bytes\n", intSize, pointSize);
    printf("\n");


    // --- Section 3: Control Flow - Nested If-Else If-Else ---
    printf("--- Section 3: Nested If-Else If-Else ---\n");
    int score = 85;
    if (score >= 90) {
        printf("Grade: A\n");
    } else if (score >= 80) {
        printf("Grade: B\n");
        if (score == 85) {
            printf("  (Exactly 85!)\n"); // Nested if
        }
    } else if (score >= 70) {
        printf("Grade: C\n");
    } else {
        printf("Grade: D or F\n");
    }
    printf("\n");


    // --- Section 4: Control Flow - Switch Case ---
    printf("--- Section 4: Switch Case ---\n");
    char grade = 'B';
    switch (grade) {
        case 'A':
            printf("Excellent!\n");
            break; // Exit switch
        case 'B':
            printf("Good Job!\n");
            // Fall-through is possible if 'break' is omitted
        case 'C':
            printf("Okay. (Reached from B or C)\n");
            break;
        case 'D':
        case 'F': // Multiple cases can lead to the same block
            printf("Needs Improvement.\n");
            break;
        default: // Optional default case
            printf("Invalid Grade.\n");
            break; // Good practice to have break in default too
    }
    printf("\n");


    // --- Section 5: Control Flow - While Loop ---
    printf("--- Section 5: While Loop ---\n");
    int whileCounter = 0;
    while (whileCounter < 3) {
        printf("  whileCounter = %d\n", whileCounter);
        whileCounter++;
    }
    printf("\n");


    // --- Section 6: Control Flow - Do-While Loop ---
    printf("--- Section 6: Do-While Loop ---\n");
    int doWhileCounter = 0;
    do {
        printf("  doWhileCounter = %d\n", doWhileCounter);
        doWhileCounter++;
    } while (doWhileCounter < 3); // Condition checked *after* the first iteration
    printf("\n");


    // --- Section 7: Control Flow - For Loop ---
    printf("--- Section 7: For Loop ---\n");
    // Initialization; Condition; Increment
    for (int i = 0; i < 4; i++) {
        printf("  For loop i = %d\n", i);
    }
    printf("\n");


    // --- Section 8: Control Flow - Break and Continue ---
    printf("--- Section 8: Break and Continue ---\n");
    for (int j = 0; j < 10; j++) {
        if (j == 2) {
            printf("  (Skipping j=2 with continue)\n");
            continue; // Skip the rest of this iteration, go to next j
        }
        if (j == 5) {
            printf("  (Exiting loop at j=5 with break)\n");
            break; // Exit the loop entirely
        }
        printf("  Looping with j = %d\n", j);
    }
    printf("\n");


    // --- Section 9: Functions (Calling) ---
    printf("--- Section 9: Functions ---\n");
    int num1 = 25, num2 = 17;
    int resultSum = add(num1, num2); // Call function defined below main
    printf("  Result of add(%d, %d) = %d\n", num1, num2, resultSum);
    print_message("This message comes from a function call."); // Call void function
    printf("\n");


    // --- Section 10: Arrays ---
    printf("--- Section 10: Arrays ---\n");
    int numbers[5] = {10, 20, 30, 40, 50}; // Declaration and initialization
    numbers[1] = 25; // Modify an element
    printf("  Array elements: ");
    for (int k = 0; k < 5; k++) {
        printf("%d ", numbers[k]); // Access elements
    }
    printf("\n");

    // Multi-dimensional array
    int matrix[2][3] = { {1, 2, 3}, {4, 5, 6} };
    printf("  Matrix element [1][1]: %d\n", matrix[1][1]);
    printf("\n");


    // --- Section 11: Pointers ---
    printf("--- Section 11: Pointers ---\n");
    int value = 99;
    int *pointerToValue = &value; // Pointer holds the address of 'value'

    printf("  Value: %d, Address of value: %p\n", value, (void*)&value);
    printf("  Pointer points to address: %p\n", (void*)pointerToValue);
    printf("  Value via pointer (dereference *): %d\n", *pointerToValue);

    *pointerToValue = 101; // Change the value using the pointer
    printf("  Value changed via pointer to: %d\n", value);

    // Pointer to array (points to the first element)
    int *arrayPtr = numbers; // Same as int *arrayPtr = &numbers[0];
    printf("  First array element via pointer: %d\n", *arrayPtr);
    printf("  Second array element via pointer arithmetic: %d\n", *(arrayPtr + 1));

    // Null pointer
    int *nullPtr = NULL;
    printf("  A null pointer: %p\n", (void*)nullPtr);
    // Dereferencing NULL causes undefined behavior (often a crash)
    // if (nullPtr != NULL) { *nullPtr = 10; } // Always check before dereferencing
    printf("\n");


    // --- Section 12: Strings (as character arrays) ---
    printf("--- Section 12: Strings ---\n");
    char str1[] = "Hello"; // String literal initialization (null-terminated)
    char str2[20];         // Declare a char array to hold a string

    strcpy(str2, ", World!"); // Copy string into str2 (requires string.h)
    // Note: strcpy is unsafe; strncpy or other functions are preferred in modern code.

    printf("  String 1: %s\n", str1);
    printf("  String 2: %s\n", str2);

    char str3[50];
    sprintf(str3, "%s%s (Length: %zu)", str1, str2, strlen(str1) + strlen(str2)); // Formatted string output
    printf("  Combined: %s\n", str3);
    printf("\n");


    // --- Section 13: Structs (Usage) ---
    printf("--- Section 13: Structs ---\n");
    struct Point p1; // Declare a variable of struct type
    p1.x = 10;
    p1.y = 20.5;
    strcpy(p1.label, "Start");

    printf("  Point p1: label='%s', x=%d, y=%.1f\n", p1.label, p1.x, p1.y);

    // Pointer to struct
    struct Point *pPtr = &p1;
    pPtr->x = 15; // Access member via pointer using -> operator
    printf("  Point p1 via pointer: x=%d\n", pPtr->x);
    printf("\n");


    // --- Section 14: Enums (Usage) ---
    printf("--- Section 14: Enums ---\n");
    enum Status currentStatus = PROCESSING;
    if (currentStatus == PROCESSING) {
        printf("  Current status is PROCESSING (value %d)\n", currentStatus);
    }
    // Assign another value
    currentStatus = COMPLETED;
    printf("  New status value: %d\n", currentStatus);
    printf("\n");


    // --- Section 15: Typedef (Usage) ---
    printf("--- Section 15: Typedef ---\n");
    ULL largeNumber = 123456789012345ULL; // Using the ULL typedef
    Coordinate c1 = {5, -2.5, "Center"};    // Using the Coordinate typedef for struct Point
    TaskStatus jobState = PENDING;          // Using the TaskStatus typedef for enum Status

    printf("  Typedef ULL: %llu\n", largeNumber);
    printf("  Typedef Coordinate: label='%s', x=%d, y=%.1f\n", c1.label, c1.x, c1.y);
    printf("  Typedef TaskStatus: %d\n", jobState);
    printf("\n");


    // --- Section 16: Preprocessor Directives (Usage) ---
    printf("--- Section 16: Preprocessor Directives ---\n");
    printf("  Using #defined PI: %f\n", PI);
    printf("  Using #defined GREETING: %s\n", GREETING);

    #if DEBUG_LEVEL == 1
        printf("  Debug Level 1 is enabled.\n");
    #elif DEBUG_LEVEL == 2
        printf("  Debug Level 2 is enabled.\n");
    #else
        printf("  Debug Level is not 1 or 2.\n");
    #endif

    #ifdef GREETING
        printf("  GREETING macro is defined.\n");
    #endif

    #ifndef NON_EXISTENT_MACRO
        printf("  NON_EXISTENT_MACRO is not defined.\n");
    #endif
    printf("\n");


    // --- Section 17: Dynamic Memory Allocation ---
    printf("--- Section 17: Dynamic Memory Allocation ---\n");
    int *dynamicArray = NULL;
    int size = 5;

    // Allocate memory for 'size' integers
    dynamicArray = (int*) malloc(size * sizeof(int));

    if (dynamicArray == NULL) {
        fprintf(stderr, "  Error: Failed to allocate memory!\n");
        // exit(EXIT_FAILURE); // Exit if allocation fails in real app
    } else {
        printf("  Memory allocated successfully at %p\n", (void*)dynamicArray);
        // Initialize and use the allocated memory
        for (int i = 0; i < size; i++) {
            dynamicArray[i] = i * 100;
        }
        printf("  Dynamic Array[2] = %d\n", dynamicArray[2]);

        // Free the allocated memory when done
        free(dynamicArray);
        printf("  Memory freed.\n");
        dynamicArray = NULL; // Good practice to NULL pointer after free
    }
    printf("\n");

    // --- Section 18: Basic File I/O ---
    printf("--- Section 18: Basic File I/O ---\n");
    FILE *filePtr = NULL;
    const char *filename = "c_syntax_output.txt";

    // Write to a file
    filePtr = fopen(filename, "w"); // Open for writing ("w")
    if (filePtr == NULL) {
        perror("  Error opening file for writing");
    } else {
        fprintf(filePtr, "This is line 1 written from the C program.\n");
        fprintf(filePtr, "The value of integerVar was %d.\n", integerVar);
        fclose(filePtr); // Close the file
        printf("  Successfully wrote to %s\n", filename);
    }

    // Read from a file (simple example reading one line)
    char buffer[100];
    filePtr = fopen(filename, "r"); // Open for reading ("r")
     if (filePtr == NULL) {
        perror("  Error opening file for reading");
    } else {
        if (fgets(buffer, sizeof(buffer), filePtr) != NULL) {
            printf("  Read from file: %s", buffer); // fgets includes newline if it fits
        } else {
            printf("  Could not read line from file or file empty.\n");
        }
        fclose(filePtr); // Close the file
    }
    printf("\n");


    // --- Section 19: Goto Statement (Use Sparingly!) ---
    printf("--- Section 19: Goto Statement ---\n");
    int gotoCounter = 0;
    start_loop_label: // Label definition
        if (gotoCounter < 2) {
            printf("  Goto loop iteration %d\n", gotoCounter);
            gotoCounter++;
            goto start_loop_label; // Jump back to the label
        }
    printf("  Goto loop finished.\n");
    // Note: goto is generally discouraged for creating loops or complex jumps.
    //       Use loops (for, while, do-while) and functions instead for clarity.
    printf("\n");


    // --- End of Showcase ---
    printf("--- C Syntax Showcase End ---\n");
    return 0; // Indicate successful execution
}

// --- Section 9: Functions (Definition) ---
// Function definition: Adds two integers
int add(int a, int b) {
    return a + b; // Return statement
}

// Function definition: Prints a message (void return type)
void print_message(const char *message) {
    // 'const' indicates the function won't modify the string via the pointer
    printf("  Function print_message: %s\n", message);
}