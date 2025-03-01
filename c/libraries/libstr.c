char *reverse(char *string, int length)
{
    int l = 0, r = length - 1;
    while (l < r)
    {
        char temp = string[l];
        string[l] = string[r];
        string[r] = temp;
        l++;
        r--;
    }

    return string;
}