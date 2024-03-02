## Brainfuck Interpreter / JIT Compiler
This project was inspired by the amazing [TsodingDaily](https://www.youtube.com/@TsodingDaily).

it was made in order to experiment with building basic compilers/ interpreters and
messing around with assembly.

NOTE:
- I did not implement the `,` instruction of the language for simplicity.
- Memory assigned to the brainfuck program is 8KB.
- I did not enforce validations over memory's out of bound access.

### HOW TO BUILD:
#### Prerequisites:

- Golang 1.18+

---- 

- ``cd`` to project's directory

Supported architectures:

- AMD64

Supported operating systems:

- Darwin ``make build-darwin``
- Linux ``make build-linux``


### HOW TO RUN:
``./bf interpret ./hello_world.bf``

``./bf compile ./hello_world.bf``