let reduce = func(arr, init, reducer) {
    let curr = init;
    for (a in arr) {
        curr = reducer(curr, a);
    }
    return curr;
};

let arr = [1, 2, 3, 4, 5];
let sum = reduce(arr, 0, func(acc, curr) {
    return acc + curr;
});

// expect 15
println("Reduced sum:", sum);

let factorial = reduce(arr, 1, func(acc, curr) {
    return acc * curr;
});

// expect 120
println("Reduced factorial:", factorial);

let calcFactorial = func(num) {
    if (num <= 1) {
        return num;
    }
    return num * calcFactorial(num - 1);
};

println(calcFactorial(5));
