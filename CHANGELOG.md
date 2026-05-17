# Change Log
All notable changes to this project will be documented in this file.

### [1.0.4] - 2026-05-17
#### Added
> Added `-safe` to generate passwords using shell-safe special characters.
>
> Documented safe shell usage, command substitution quoting, and shell-safe character sets.

---

### [1.0.3] - 2026-05-17
#### Added
> Added a pixel art logo.
>
> Added a user manual covering all flags, valid combinations, invalid examples, exit codes, character sets, and build commands.
#### Changed
> Reduced the README to a concise project overview, quick start, documentation link, and build entry point.

---

### [1.0.2] - 2026-05-17
#### Changed
> Documented supported flag combinations, validation rules, invalid examples, and exit codes.
#### Fixed
> `-version` now rejects extra arguments and generation flags instead of ignoring them.

---

### [1.0.1] - 2026-05-17
#### Changed
> Updated GitHub Actions to Node 24-compatible action versions.
>
> Removed unnecessary Go dependency caching from workflows because this project has no external modules.

---

### [1.0.0] - 2026-05-17
#### Added
> Added release archives with SHA-256 checksums through `make release`.
>
> Added a tag-driven release workflow that publishes GitHub Release assets.
#### Changed
> Promoted the project from alpha to a stable release.
>
> Reworked build output to use `build/` for local binaries and `dist/` for release archives.
>
> Removed experimental wording and alpha install instructions from the README.
>
> Limited the default special-character pool to printable ASCII.
#### Removed
> Removed the `install` Make target.
>
> Removed the no-op `--silent` CLI flag.

---

### [0.5.0-alpha] - 2026-05-17
#### Added
> Added a Makefile with check, build, dist, install, and clean targets.
>
> Added GitHub Actions CI for formatting, tests, vet, and build.
>
> Added CLI tests for default output, help, version, and invalid length handling.
#### Changed
> Moved password policy validation into the generator package.
>
> Changed default CLI output to print only the generated password on stdout.
>
> Changed version reporting to support build-time metadata and Go module build info.
#### Fixed
> Fixed silent truncation of password lengths above the maximum.

---

### [0.4.0-alpha] - 2026-05-17
#### Added
> Added a clearer README with badges, installation instructions, usage examples, and experimental status.
>
> Added a `-version` flag.
#### Changed
> Aligned the Go module with `github.com/olelbis/pswg` and Go 1.26.3.
>
> Reworked CLI parsing using the standard `flag` package.
>
> Updated password generation to propagate crypto errors and use crypto-backed shuffling.
#### Fixed
> Fixed the failing test contract for short passwords.
>
> Fixed a possible infinite shuffle loop when generated passwords contained only numeric characters.

---

### [0.1.3.2] - 2022-09-12
#### Added
> Another long hiatus doing few test with 1.19
#### Changed
> _n/a_
#### Fixed
> _n/a_

### [0.1.3.1] - 2022-04-08
#### Added
> After a long hiatus i'm sarting write fist tests
#### Changed
> Review Melee function due to testing approach
#### Fixed
> _n/a_

---
### [0.1.3] - 2021-06-12
#### Added
> Start Play with ansi color 
#### Changed
> Move some const to var for a properties file implementation
#### Fixed
> _n/a_

---

### [0.1.2] - 2021-05-01
#### Added
> _n/a_ 
#### Changed
> Change default message again
#### Fixed
> Bug in maximum password lenght

---

### [0.1.1] - 2021-04-22
#### Added
> _n/a_ 
#### Changed
> Change default message 
#### Fixed
> _n/a_ 

---

### [0.1.0] - 2021-04-09
#### Added
> _n/a_ 
#### Changed
> Changed constant,variabels and function name (not all but now are more esplicative,and the code seem more readable i hope) 
#### Fixed
> _n/a_ 

---

### [0.0.9] - 2021-04-08
#### Added
> Find a more efficent way to re-call shuffle until fist character is a number
#### Changed
> _n/a_ 
#### Fixed
> _n/a_ 

---

### [0.0.8] - 2021-04-03
#### Added
> 1.16.3 Supported! Added a check to shuffle until fist character is a number (some vendor don't want this kind of password... But i don't know if i want to maintain this kind of request) Last but not least deprecate Pick Func
#### Changed
> _n/a_ 
#### Fixed
> _n/a_ 

---

### [0.0.7] - 2021-03-31
#### Added
> Added more err handling, but i need to do more test
#### Changed
> _n/a_ 
#### Fixed
> _n/a_ 

---

### [0.0.6] - 2021-03-18
#### Added
> _n/a_ 
#### Changed
> I'm a little busy in the latest weeks. As suggeted by Sfrisio i've switch from "math/rand" to "crypto/rand" but need to test it a little bit more. (crypto is implemented only in Pick* function)
#### Fixed
> _n/a_ 

---

### [0.0.5] - 2021-03-09
#### Added
> Create function DefPick to generate random password based on predefined rule
> 
> Parameter managment (need further work)
#### Changed
> _n/a_
#### Fixed
> _n/a_ 

---

### [0.0.4] - 2021-03-04
#### Added
> Define the default execution and added a lil message (can be better i know)
#### Changed
> _n/a_
#### Fixed
> _n/a_

---

### [0.0.3] - 2021-02-23
#### Added
> _n/a_
#### Changed
> Rename prj and a little cleanup
#### Fixed
> _n/a_

---

### [0.0.2] - 2021-02-22
#### Added
> _n/a_
#### Changed
> _n/a_
#### Fixed
> Some minor fix to the code :P

---

### [0.0.1] - 2021-02-18
#### Added
> _n/a_
#### Changed
> Move Melee function into a separate package to get familiarity with module :P
> 
> Change some comment from Italian to Esperanto(lol)...
> 
> Move constant...
> 
> Some lessons learned
#### Fixed
> _n/a_

---

### [0.0.0] - 2021-02-15
This is a personal excercise to create a simple stupid random password generator while learn go language
