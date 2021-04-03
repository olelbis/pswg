# pswg 

It's an personal excercise to create a simple stupid random password generator while learn go language

-   18/02/2021  Move Melee function into a separate package to get familiarity with module :P
  
                -   Change some comment from Italian to Esperanto(lol)...
                -   Move constant...
                -   Some lessons learned
  
-   22/02/2021  Some minor fix to the code :P
-   23/02/2021  Rename prj and a little cleanup
-   04/03/2021  Define the default execution and added a lil message (csn be better i know)
-   09/03/2021  Some mods:
  
                -   Create function DefPick to generate random password based on predefined rule
                -   Review constants
                -   Parameter managment (need further work)

-   18/03/2020  I'm a little busy in the latest weeks. As suggeted by Sfrisio i've switch from "math/rand" to "crypto/rand" but need to test it a little bit more. (crypto is implemented only in Pick* function)
-   23/03/2021  Tested with 1.16.2
-   31/03/2021  Added more err handling, but i need to do more test
-   03/04/2021  1.16.3 Check! Added a check to shuffle until fist character is a number (some vendor don't want this kind of password... But i don't know if i want to maintain this kind of request) Last but not least deprecate Pick Func