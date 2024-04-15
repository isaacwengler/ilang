let twoSum = func(nums, target) {
    let seen = {};
    for (let i = 0; i < nums.length(); i = i + 1) {
        let num = nums[i];
        if (seen[target - num] != null) {
            return [seen[target - num], i];
        }
        seen[num] = i;
    }
    return [-1, -1];
};

let a = [];
for (let i = 0; i < 100000; i = i + 1) {
    a.push(i);
}
let res = twoSum(a, 5892);
println("result", res);
