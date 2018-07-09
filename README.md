# CP Tool

Simple bash script to compile and run your solution when solving competitive programming problems.

## Motivation

When competing in programming competition, we always need to compile our program and run it. With c++ we can achieve this by running bash command like this `g++ solution.cpp -o solution && ./solution`. This command actually works pretty well, but wasting time to write. When you need to rerun your solution with different testcases, you normally use up arrow in your bash shell to run previous command. However running previous command is time comsuming because we need to recompile the solution. It would be great if there is command that can compile and run your solution with single simple command, and can rerun the solution without compiling it again. Cptool wants to achieve this by wrapping compilation and execution command for running your competitive programming solution.

## How To Install

Run this command in your bash shell for installing cptool:

```
curl -L -s https://raw.githubusercontent.com/cp-itb/cptool/master/install.sh | bash
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

Your solution should contain only one file with this format: `<solution-name>.<language-extension>`. To compile your solution, just run `cptool compile <solution-name>`, this will compile your solution with default language, which is c++. To compile with other language just run `cptool compile <language-name> <solution-name>`, this will compile your solution using specific language.

## Running Solution

After compiling solution, the executable file should exists in folder `.cptool`. You can find your compiled solution there and run it. However, the easiest way to run your solution is using this command `cptool run <solution-name>` or `cptool run <language-name> <solution-name>` to specify the language. The solution will be compiled first if not yet compiled. After that, this command will run your solution in command line interface.

## Testing Solution

TBD

## Adding New Language

TBD