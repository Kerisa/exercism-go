package collatzconjecture

import "errors"

// CollatzConjecture check step of a number to reach 1
func CollatzConjecture(input int) (int,error) {

    if input <= 0 {
        return 0, errors.New("require a positive number")
    }

    step := 0
    for ; input != 1; step++ {
        if input % 2 != 0 {
            input = 3 * input + 1
        } else {
            input /= 2
        }
    }
    return step, nil
}
