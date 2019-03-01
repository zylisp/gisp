package repl

const ReplBanner string = `Okay, 3, 2, 1 - Let's jam!

Welcome to

/^^^^^^^^/^^ /^^      /^^ /^^       /^^ /^^^^^^^^ /^^^^^^^^^
       /^^    /^^    /^^  /^^       /^^ /^^       /^^    /^^
      /^^      /^^ /^^    /^^       /^^ /^^       /^^    /^^
    /^^          /^^      /^^       /^^ /^^^^^^^^ /^^^^^^^^^
   /^^           /^^      /^^       /^^       /^^ /^^
 /^^             /^^      /^^       /^^ /^^   /^^ /^^
/^^^^^^^^^^^     /^^      /^^^^^^^^ /^^ /^^^^^^^^ /^^
`

const CommonReplHelp string = `
        Docs: https://zylisp.github.io/zylisp/
     Project: https://github.com/zylisp/zylisp`

const AstReplHelp string = `
Instructions: Simply type any form to view the generated Go AST.
        Exit: <CONTROL><C>
`

const LispReplHelp string = `
      Exit: <CONTROL><D> or (exit) or (quit)
`

const ReplExitMsg string =`^D

See you space cowboy ...
`

const AstPrompt string = "AST> "

const GoPrompt string = "GO> "

const LispPrompt string = "𝛌> "
