package repl

const ReplBannerGreeting string = `Okay, 3, 2, 1 - Let's jam!

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

const ReplCommonExitMsg string =`
See you space cowboy ...
`

const RepCtlDlExitMsg string =`^D
`

const RepCtlClExitMsg string =`^C
`
const AstPrompt string = "AST> "

const GoPrompt string = "GO> "

const LispPrompt string = "𝛌> "

const LispDefaultPackage string = "user"
