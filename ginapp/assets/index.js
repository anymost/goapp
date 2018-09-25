function outer() {
    var count = 0;
    return function inner() {
        return count++;
    }
}
