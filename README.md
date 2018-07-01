# CPTOOL

Simple bash script to compile and run your solution when solving competitive programming problems.

## How To Install

Run this command in your bash shell

```
curl -L -s https://raw.githubusercontent.com/jauhararifin/cptool/master/install.sh | bash
```

## Compile Solution

Your solution should contain only one file with this format: `<solution-name>.<language-extension>`. To compile your solution, just run `cptool compile <solution-name>`, this will compile your solution with default language, which is c++. To compile with other language just run `cptool compile <language-name> <solution-name>`, this will compile your solution using specific language.

## Running Solution

After compiling solution, the executable file should exists in folder `.cptool`. You can find your compiled solution there and run it. However, the easiest way to run your solution is using this command `cptool run <solution-name>` or `cptool run <language-name> <solution-name>` to specify the language. The solution will be compiled first if not yet compiled. After that, this command will run your solution in command line interface.