
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
