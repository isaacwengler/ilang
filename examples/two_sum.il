let twoSum = func(nums, target) {
    let seen = {}
    for (let i = 0; i < seen.length(); i = i + 1) {
        let num = nums[i]
        if (seen[num - target] != null) {
            return [seen[num - target], i]
        }
        seen[num] = i;
    }
};

let res = twoSum([2, 7, 11, 15], 9)
// expect [0, 1]
print(res)
