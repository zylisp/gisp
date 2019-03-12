package repl

const REPLBannerGreeting string = `Okay, 3, 2, 1 - Let's jam!

Welcome to

/^^^^^^^^/^^ /^^      /^^ /^^       /^^ /^^^^^^^^ /^^^^^^^^^
       /^^    /^^    /^^  /^^       /^^ /^^       /^^    /^^
      /^^      /^^ /^^    /^^       /^^ /^^       /^^    /^^
    /^^          /^^      /^^       /^^ /^^^^^^^^ /^^^^^^^^^
   /^^           /^^      /^^       /^^       /^^ /^^
 /^^             /^^      /^^       /^^ /^^   /^^ /^^
/^^^^^^^^^^^     /^^      /^^^^^^^^ /^^ /^^^^^^^^ /^^
`

const CommonREPLHelp string = `
        Docs: https://zylisp.github.io/zylisp/
     Project: https://github.com/zylisp/zylisp`

const ASTREPLHelp string = `
Instructions: Simply type any form to view the generated Go AST.
        Exit: <CONTROL><C>
`

const GoGenREPLHelp string = ASTREPLHelp

const LispREPLHelp string = `
      Exit: <CONTROL><D> or (exit) or (quit)
`

const REPLCommonExitMsg string = `
See you space cowboy ...
`

const REPLCtlDExitMsg string = `^D
`

const REPLCtlCExitMsg string = `^C
`
const ASTPrompt string = "AST> "

const GoGenPrompt string = "GOGEN> "

const LispPrompt string = "Zy𝛌ISP> "

const LispDefaultPackage string = "user"
