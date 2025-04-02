@.str.0 = constant [4 x i8] c"%d\0A\00"
@.str.1 = constant [4 x i8] c"%f\0A\00"
@.str.2 = constant [4 x i8] c"%s\0A\00"
@.str.3 = constant [5 x i8] c"true\00"
@.str.4 = constant [6 x i8] c"false\00"
@.str.5 = constant [13 x i8] c"Hello, %s!\5Cn\00"
@.str.6 = constant [5 x i8] c"Okkk\00"
@.str.7 = constant [15 x i8] c"sounds good !!\00"
@.str.8 = constant [5 x i8] c"%s\5Cn\00"
@.str.9 = constant [7 x i8] c"okkk\5Cn\00"

declare i32 @printf(i8* %format, ...)

define void @print_int(i32 %value) {
0:
	%1 = getelementptr [4 x i8], [4 x i8]* @.str.0, i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i32 %value)
	ret void
}

define void @print_float(float %value) {
0:
	%1 = getelementptr [4 x i8], [4 x i8]* @.str.1, i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, float %value)
	ret void
}

define void @print_bool(i1 %value) {
0:
	%1 = getelementptr [4 x i8], [4 x i8]* @.str.2, i32 0, i32 0
	%2 = getelementptr [5 x i8], [5 x i8]* @.str.3, i32 0, i32 0
	%3 = getelementptr [6 x i8], [6 x i8]* @.str.4, i32 0, i32 0
	%4 = select i1 %value, i8* %2, i8* %3
	%5 = call i32 (i8*, ...) @printf(i8* %1, i8* %4)
	ret void
}

define void @print_string(i8* %value) {
0:
	%1 = getelementptr [4 x i8], [4 x i8]* @.str.2, i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %value)
	ret void
}

define i32 @add(i32 %x, i32 %y) {
0:
	%x.addr = alloca i32
	store i32 %x, i32* %x.addr
	%y.addr = alloca i32
	store i32 %y, i32* %y.addr
	%1 = load i32, i32* %x.addr
	%2 = load i32, i32* %y.addr
	%3 = add i32 %1, %2
	ret i32 %3
}

define void @greet(i8* %name) {
0:
	%name.addr = alloca i8*
	store i8* %name, i8** %name.addr
	%1 = getelementptr [13 x i8], [13 x i8]* @.str.5, i32 0, i32 0
	%2 = load i8*, i8** %name.addr
	%3 = call i32 (i8*, ...) @printf(i8* %1, i8* %2)
	ret void
}

define i32 @main() {
0:
	%value1 = alloca i32
	store i32 42, i32* %value1
	%value2 = alloca float
	store float 0x40091EB840000000, float* %value2
	%value4 = alloca i1
	store i1 true, i1* %value4
	%result = alloca i32
	%1 = call i32 @add(i32 5, i32 3)
	store i32 %1, i32* %result
	%2 = getelementptr [5 x i8], [5 x i8]* @.str.6, i32 0, i32 0
	call void @greet(i8* %2)
	%value3 = alloca i8*
	%3 = getelementptr [15 x i8], [15 x i8]* @.str.7, i32 0, i32 0
	store i8* %3, i8** %value3
	%4 = getelementptr [5 x i8], [5 x i8]* @.str.8, i32 0, i32 0
	%5 = load i8*, i8** %value3
	%6 = call i32 (i8*, ...) @printf(i8* %4, i8* %5)
	br label %while.cond

while.cond:
	br label %while.end

while.body:
	%7 = getelementptr [7 x i8], [7 x i8]* @.str.9, i32 0, i32 0
	%8 = call i32 (i8*, ...) @printf(i8* %7)
	br label %while.cond

while.end:
	ret i32 0
}
