Core C Standard Library Headers (found in C89/C90 and later):

<stdio.h> (Standard Input/Output):

Provides functions for file input/output (fopen, fclose, fread, fwrite, fseek, fprintf, fscanf, fgets, fputs, etc.) and console I/O (printf, scanf, puts, gets*, putchar, getchar).

Defines types like FILE, size_t, and constants like EOF, NULL, stdin, stdout, stderr.

*gets is extremely dangerous and should never be used.

<stdlib.h> (Standard Library Utilities):

Provides functions for memory allocation (malloc, calloc, realloc, free), number conversions (atoi, atol, atof, strtol, strtod), random number generation (rand, srand), process control (exit, abort, system, getenv), searching and sorting (bsearch, qsort).

Defines types like size_t, div_t, ldiv_t and constants like EXIT_SUCCESS, EXIT_FAILURE, RAND_MAX, NULL.

<string.h> (String Handling):

Provides functions for manipulating null-terminated character arrays (C strings): copying (strcpy*, strncpy, memcpy, memmove), concatenation (strcat*, strncat), comparison (strcmp, strncmp, memcmp), searching (strchr, strstr, strtok), getting length (strlen), and memory filling (memset).

Defines size_t and NULL.

*strcpy and strcat are dangerous due to potential buffer overflows; prefer strncpy, strncat, or memory functions with explicit sizes.

<math.h> (Mathematics):

Provides common mathematical functions: trigonometric (sin, cos, tan), hyperbolic (sinh, cosh), exponential/logarithmic (exp, log, log10, pow, sqrt), rounding (ceil, floor), absolute value (fabs).

Defines type double_t, float_t (since C99) and constants like HUGE_VAL, INFINITY, NAN (since C99).

<ctype.h> (Character Handling):

Provides functions for classifying characters (isalpha, isdigit, isspace, isupper, islower, isalnum, etc.) and converting character case (toupper, tolower).

<time.h> (Time and Date):

Provides functions for getting and manipulating time and date (time, clock, difftime, mktime, strftime, gmtime, localtime).

Defines types time_t, clock_t, struct tm and constant NULL.

<limits.h> (Integer Type Limits):

Defines constants specifying the limits of standard integer types (e.g., INT_MAX, INT_MIN, UINT_MAX, CHAR_BIT, LONG_MAX).

<float.h> (Floating-Point Type Limits):

Defines constants specifying the limits and properties of floating-point types (e.g., FLT_MAX, DBL_MIN, FLT_EPSILON, DBL_DIG).

<stddef.h> (Standard Definitions):

Defines fundamental types and macros like size_t, ptrdiff_t, wchar_t (though <wchar.h> is preferred for wide char stuff), and NULL. Also defines offsetof.

<errno.h> (Error Codes):

Defines the integer variable errno (which holds system error codes) and symbolic error constants like EDOM, ERANGE, etc.

<assert.h> (Assertions):

Provides the assert macro for verifying assumptions in code during debugging.

<locale.h> (Localization):

Provides functions (setlocale, localeconv) to handle culture-specific settings (like number formatting, currency symbols, date/time formats). Defines struct lconv, NULL.

<setjmp.h> (Non-Local Jumps):

Provides mechanisms for non-local jumps (setjmp, longjmp), allowing jumps between function calls (often used for error handling in C). Defines jmp_buf.

<signal.h> (Signal Handling):

Provides functions for handling signals (asynchronous notifications) like interrupts (signal, raise). Defines sig_atomic_t and signal constants like SIGINT, SIGTERM.

<stdarg.h> (Variable Arguments):

Provides macros (va_start, va_arg, va_copy, va_end) for handling functions that accept a variable number of arguments (like printf). Defines va_list.

Headers added in C99:

<stdbool.h> (Boolean Type):

Defines the bool type, true, and false constants, making boolean logic cleaner.

<stdint.h> (Standard Integer Types):

Defines fixed-width integer types (int8_t, uint32_t, intptr_t, etc.) for portability. Defines limits for these types (INT8_MAX, etc.).

<inttypes.h> (Format Conversion of Integer Types):

Includes <stdint.h> and provides macros for formatted I/O of fixed-width integers (e.g., PRId64, PRIu32) used with printf/scanf.

<complex.h> (Complex Arithmetic):

Defines types (double complex, float complex) and functions for complex number mathematics.

<fenv.h> (Floating-Point Environment):

Provides functions to control the floating-point environment (rounding modes, exception flags).

<tgmath.h> (Type-Generic Mathematics):

Provides type-generic macros for math functions (e.g., sqrt(x) works correctly whether x is float, double, or long double). Includes <math.h> and <complex.h>.

<wchar.h> (Wide Character Handling): (Enhanced in C99)

Functions and types for handling wide characters (beyond basic ASCII/single-byte).

<wctype.h> (Wide Character Classification): (Enhanced in C99)

Wide character versions of <ctype.h> functions.

Headers added in C11:

<stdalign.h> (Alignment):

Defines macros for specifying/querying alignment (alignas, alignof, _Alignas, _Alignof).

<stdatomic.h> (Atomics):

Provides types and functions for atomic operations, crucial for multi-threaded programming without locks.

<stdnoreturn.h> (Noreturn Functions):

Defines the _Noreturn and noreturn specifiers for functions that do not return to the caller (like exit).

<threads.h> (Threads):

Defines a standardized interface for creating and managing threads (thrd_create, mtx_lock, cnd_wait, etc.).

<uchar.h> (Unicode Characters):

Types (char16_t, char32_t) and functions for UTF-16/UTF-32 character handling.

This list covers the headers defined by the C standard itself. Actual libc implementations often provide additional, non-standard (but widely used) headers, especially on POSIX systems (e.g., <unistd.h>, <pthread.h>, <sys/types.h>, <sys/socket.h>).