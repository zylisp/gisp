package repl

// Constants for the repl package
const (
	REPLBannerGreeting string = `Okay, 3, 2, 1 - Let's jam!

Welcome to

/^^^^^^^^/^^ /^^      /^^ /^^       /^^ /^^^^^^^^ /^^^^^^^^^
       /^^    /^^    /^^  /^^       /^^ /^^       /^^    /^^
      /^^      /^^ /^^    /^^       /^^ /^^       /^^    /^^
    /^^          /^^      /^^       /^^ /^^^^^^^^ /^^^^^^^^^
   /^^           /^^      /^^       /^^       /^^ /^^
 /^^             /^^      /^^       /^^ /^^   /^^ /^^
/^^^^^^^^^^^     /^^      /^^^^^^^^ /^^ /^^^^^^^^ /^^
`

	CommonREPLHelp string = `
        Docs: https://zylisp.github.io/zylisp/
     Project: https://github.com/zylisp/zylisp`

	ASTREPLHelp string = `
Instructions: Simply type any form to view the generated Go AST.
        Exit: ^D or ^C
`

	GoGenREPLHelp string = ASTREPLHelp

	LispREPLHelp string = `
      Exit: ^D, ^C, (exit), or (quit)
`

	REPLCommonExitMsg string = `
See you space cowboy ...
`

	REPLCtlDExitMsg string = `^D
`

	REPLCtlCExitMsg string = `^C
`
	ASTPrompt string = "AST> "

	GoGenPrompt string = "GOGEN> "

	LispPrompt string = "Zy𝛌ISP> "

	LispDefaultPackage string = "user"
)
