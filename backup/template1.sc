func add(x: int, y: int) -> int {
    return x + y;
}

func greet(name: string) -> void {
    printf("Hello, %s!\n", name);
}

func main() -> int {

    // supported types
    int value1 = 42;
    float value2 = 3.14;
    string value3 = "Awesome";
    bool value4 = true;

    int result = add(5, 3);


    printf("%s\n", greet("Okkk"));


    return 0;
}