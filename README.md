# CP Tool

[![Build Status](https://travis-ci.org/jauhararifin/cptool.svg?branch=master)](https://travis-ci.org/jauhararifin/cptool)

Simple bash script to compile and run your solution when solving competitive programming problems.

## Motivation

When competing in programming competition, we always need to compile our program and run it. With c++ we can achieve this by running bash command like this `g++ solution.cpp -o solution && ./solution`. This command actually works pretty well, but wasting time to write. When you need to rerun your solution with different testcases, you normally use up arrow in your bash shell to run previous command. However running previous command is time comsuming because we need to recompile the solution. It would be great if there is command that can compile and run your solution with single simple command, and can rerun the solution without compiling it again. Cptool wants to achieve this by wrapping compilation and execution command for running your competitive programming solution.

## How To Install

Run this command in your bash shell for installing cptool:

```
curl -L -s https://raw.githubusercontent.com/jauhararifin/cptool/master/install.sh | bash
```

It will download this repository and copy to your computer. You may need to close your shell and open it again to start using this tool. Just run `cptool --version` command to check your installation.

## Quick Start

There are some term you need to understand in order to using this tool.

1. Language
The programming language that can be used for writing your solution. Cptool currently support c, c++ and pascal only. The language has definition that gives cptool information about its name, extension, how to compile and run solution in this language.

2. Solution
This is your solution source code. Your solution is a single file with a name and extension. The basename (filename without extension) of your solution file is considered as your solution name and the extension can be considered as your solution language.

3. Testcase
This is a pair of file which has extension .in and .out that represent testcase for your solution. You can test your solution using this testcase. A testcase should contain two file that has .in and .out extension and has exactly same basename (filename without extension).

For understanding how to use this tool, lets consider the most famous problem in competitive programming world. The "A+B" problem:
```
A + B

[Description]
Given integer A and B, compute the sum of A, B.

[Input Format]
The input is a single line that consist of integer A and B

[Output Format]
A single integer that is the sum of A and B

[Sample Input]
1 2

[Sample Output]
3

[Constraint]
- 1 <= A <= 1000
- 1 <= B <= 1000.
```

Lets write the solution using c++ language, create file `aplusb.cpp` and paste this code:

```
#include <iostream>

using namespace std;

int a,b;

int main() {
  cin>>a>>b;
  cout<<(a+b)<<endl;
  return 0;
}
```

At this point, you already have a solution named `aplusb` (because the filename is `aplusb.cpp` and the basename is `aplusb`). You can compile your solution using this command:

```
cptool compile aplusb
```

Your program will be compiled to executable file. You can find this file in .cptool directory that created in your working directory. Actually you can run the executable file from that directory but the easiest way to run that program is using `cptool run` command.

You can run your compiled program to test your solution before submitting it. To run your solution, just type this command in your bash shell:

```
cptool run aplusb
```

This command will check whether the compiled solution already up to date or not. It will recompile the solution if you made changes in your solution. When the compilation successfull it will run your program and you can start testing it. Actually you can omit the previous compilation process and start running your solution using this `cptool run` command without compile it first because this command will do the hard work for you.

Lets say you have found some testcases to test your solution. You can save this testcase in file `aplusb.sample_1.in` and `aplusb.sample_1.out`. Consider the following testcase files:
- `aplusb.sample_1.in`
```
1 2
```
- `aplusb.sample_1.out`
```
3
```
This two files represent testcase for your solution, and you can run and test your solution using this testcase by running this command:

```
cptool test aplusb aplusb
```

It will run (and compile if necessary) your solution using testcase as input and it will compare your solution output with the output file. It will inform you whether you passed the testcase or not.

## Compile Solution

Your solution should contain only one file with this format: `<solution-name>.<language-extension>`. The example of valid solution file is `helloworld.cpp` file, this file is considered as a c++ solution named `helloworld`.

To compile your solution, just run
```
cptool compile <solution-name>
```
This will compile your solution with default language, which is c++. The compiled program will placed inside `.cptool` directory inside your working directory. To compile with other language just run 
```
cptool compile <language-name> <solution-name>
```
This will compile your solution using specific language. Some language can be compiled in debug mode. For compiling your solution in debug mode, use `-d` flag in compilation command like this:

```
cptool run -d <solution-name>
```
or
```
cptool run -d <language-name> <solution-name>
```

The currently available language are `cpp`, `c`, and `pas` (for free pascal).

## Running Solution

To run your solution, you can use `cptool run` command. This command will compile your solution first if its not compiled yet. When you run this command twice, the compilation process is skipped. The solution will compiled again when your solution change or the `.cptool` directory is removed. Use this command to run your solution

```
cptool run <solution-name>
```
or you can specify the solution language using this command:
```
cptool run <language-name> <solution-name>
```

The running solution will have timeout, default is 10 seconds. Your solution will be killed when your solution is running over this amount of time. You can specify this time using `-t` or `--timeout` flag like this:

```
cptool run -t 5s <solution-name>
```
or
```
cptool run -t 5s <language-name> <solution-name>
```

use suffix 's' for second, 'm' for minutes and 'h' for hours in the timeout parameter.

## Testing Solution

For testing your solution, you need to have testcase file. A single testcase is consist of two files, the first file is input file and the second is output file. The input file is a plain text file that has `.in` extension and output has `.out` extension. In a single testcase, the input file and the output file must has the same basename (filename without extension). `tc1.in` and `tc`.out` is the valid example of a single testcase files.

For testing your program using predefined testcase, you can use `cptool test` command like this

```
cptool test <solution-name> <testcase-prefix>
```
or you can specify your solution language
```
cptool test <language-name> <solution-name> <testcase-prefix>
```

The testcase prefix is filename prefix of your testcase file. For example, the testcase `tc1` will has all of this prefix: "t", "tc", "tc1". You can run your solution like this `cptool run solution tc`. It will test your solution will all testcases that has "tc" as its prefix.

You can add timeout parameter too like this `cptool test <solution-name> --timeout 5s`, just like the `cptool run` command.

## List Languages

You can run `cptool lang` to list all available languages.

## Adding New Language

Cptool currently support three languages: C, C++, Pascal. The language is defined in `langs` directory in your installation folder (default is `~/.cptool`). You can add new language by adding new folder in that directory (`<installation-directory>/langs`). The folder name will be the language name. Inside that folder you need to add four files: `compile`, `debugcompile`, `run`, and `lang.conf`

- `compile` script
Cptool need to know how to compile the source code solution, The `compile` file contains bash script to compile the source code. You have to specify compilation command in this file. you will receive two parameters in order to locate the source code location and compiled target. First parameter contains the path of your solution file. The second parameter is location where you should put the compiled file. If the compilation is successfull this script must return 0 to the operating system.

- `debugcompile` script
Sometimes compiler has ability to compile source code in debug mode, you have to specify compilation command to compile the source code in debug mode. This script will also receive two parameters just like `compile` script.

- `run` script
You have to specify how to run your solution in this file. This script will receive a parameter. The parameter contain location of your compiled program path. When the program is exited normally, this script must return 0 to the operating system.

- `lang.conf` file
This file contain information about the language. Actually it just need two information: language verbose name and language extension. Language verbose name is just like your language displayed name, but your language name is the folder name. The language extension is the extension of your solution file, for example: the c language has `c` as language extension, pascal language has `pas` as language extension.
The format of this file is:

```
verbose_name=<language-verbose-name>
extension=<language-extension>
```
