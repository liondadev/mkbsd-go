# MKBSD-go
This is a faster, concurrent implementation of [MKBSD](https://github.com/nadimkobeissi/mkbsd).

## Usage
1. Download the executable for your system from the releases page.
2. Copy the executable into an empty folder. 
   - The program will make the necessary `downloads` folder in here
3. Run the program
4. Check the newly created `downloads` folder

## Benchmarks

### Windows

| Stat | Value |
|------|-------|
| Seconds Taken| ~20 |

Raw Execution:
```
PS C:\Users\liond\Downloads\mkbsd> Measure-Command { .\mkbsd-win-x64.exe | Out-Default }
2024/09/27 16:08:03 Got API Response. (Version: 1)
2024/09/27 16:08:04 Downloaded wallpaper 1646651327.jpeg (112273 bytes)
...
2024/09/27 16:08:22 Finished downloading!


Days              : 0
Hours             : 0
Minutes           : 0
Seconds           : 20
Milliseconds      : 368
Ticks             : 203684262
TotalDays         : 0.000235745673611111
TotalHours        : 0.00565789616666667
TotalMinutes      : 0.33947377
TotalSeconds      : 20.3684262
TotalMilliseconds : 20368.4262
```
