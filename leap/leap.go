package leap

func IsLeapYear(yr int) bool {

    if yr % 4 == 0 && yr % 100 != 0 {
        return true;
    }

    if yr % 400 == 0 {
        return true;
    }

    return false;
}
