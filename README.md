# AoC-2024

Advent of code 2024

For this year I'm learning golang and PHP, so I'm going to attempt my first try at solving with either of those two, and then python after to see if I can think of a better solution.

If I get completely stuck I will then try it in python "first"

Goddamn, this year ramped up fast. I gave up at day 3, I just don't have time to do the challenges this year.

## Insights

What I learned about each language as I went

### Golang

- Resources for go are garbage, looking up guides for how to do basic things will give you terrible advice
  - Trying to figure out how to sort a slice of integers was a nightmare. The builtin snippet on vscode returns code that is incomprehensible without google:
```go 
myslice := make([]int, 0)
sort.Slice(myslice, func(i, j int) bool {
    // Do something here?
})
```

- Why are there so many bait functions. Why is `sort.IntSlice()` a type instead of returning a sorted slice of ints?????? Of course I should understand naturally to try `sort.Ints()` a naming convention no other language uses!
  - Oh my god, same thing with `strings.Trim()` does not trim whitespace from a string, instead you have to use `strings.TrimSpace()`
- There is 0 consistency with how to convert types. For some values it's `value.(type)`, others it's `type(value)` others need a factory function with a different convention. Likewise they've made the necessity to change data types so often. To take the absolute value of 2 ints as ints I had to write `int(math.Floor(math.Abs(float64((1-0)))))`

### Python

- I've found myself more reliably completeing days in golang than python. Not sure If I'm just out of practice, or the static typing is helping me avoid parsing problems I'm too dumb to see in python

### PHP

- God why does this also have bait names. Why does `$data = readfile(filename);` not return the data. Instead you need to use `$data = file_get_contents("input.txt");`
- Alias `count()` to `len()` or something similar to be in line with the rest of the programming community (or add a `array->length` property)
