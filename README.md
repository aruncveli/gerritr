# gerritr

Wrapping some [Git](https://git-scm.com/) for [Gerrit](https://www.gerritcodereview.com/)

A command line app with opinionated flows for a subset of Gerrit functionalities. Intended to ease people familiar with GitHub et al., to work with Gerrit. And hopefully help them remember to add reviewers to a change.

Inspired from:
* Lack of non-tedious way in [git-review](https://docs.opendev.org/opendev/git-review/latest/) to add reviewers
* Lack of Python tooling for cross compilation
* The ease of Golang tooling for cross compilation

## Installation
[Download the binary](https://github.com/aruncveli/gerritr/releases) and add to `$PATH`.

## Usage
`gerritr` is supposed to be used in the same context as `git` - a command line application from the repository root directory. `gerritr` has two subcommands: `push` and `patch`.

### Common flags
* `--branch` or `-b`: Target branch name. Required if the target branch is not main or master.
* `--message` or `-m`: Commit message.
* `--state` or `-s`: [Change state](https://gerrit-documentation.storage.googleapis.com/Documentation/3.6.2/user-upload.html#private). Ignored if the value is not one of:
	* `private`
	* `remove-private`
	* `wip`
	* `ready`

### `push`

#### Flags
* `--reviewers` or `-r`: Space separated list of reviewer email IDs or aliases

Push the latest commit to the target branch, thereby [creating a new change](https://gerrit-documentation.storage.googleapis.com/Documentation/3.6.2/intro-gerrit-walkthrough.html#_creating_the_review) in Gerrit.

If `--message` flag is present, a new commit is created initally - with the staged changes and the provided message - and then pushed.

Adding reviewers without verbosely typing in everyone's email ID, is possible in two ways:
#### `config.yml` or global config
* Go to `$XDG_CONFIG_HOME`, which is generally:
	* `~/.config` for Linux
	* `~/Library/Application Support` for macOS
	* `%LOCALAPPDATA%` for Windows
* Create a `gerritr` directory there.
* Add a `config.yml` in `gerritr` directory with content similar to the sample below:
	```YAML
	alias:
	  backend:
		- b1@org.com
		- b2@org.com
	  frontend:
		- f1@org.com
		- f2@org.com
	  someone:
	    - someone.with.long.email.id@somewhere.com
	  ...
	```
The configured aliases can then be used as values to `--reviewers`. Unresolvable values for `--reviewers` are ignored.

#### `REVIEWERS` or local config
Add a plaintext `REVIEWERS` file to the repository root directory, with a list of email IDs of people who should **always be added as reviewers to every change pushed from this repository**:
```Text
r1@org.com
r2@org.com
...
```

Both the ways to add reviewers are independent:
* With just the above global configuration in place, `gerritr push -r backend x1@org.com` will add `x1@org.com`, `b1@org.com` and `b2@org.com` as reviewers
* With just the above local configuration in place, `gerritr push` will add `r1@org.com` and `r2@org.com` as reviewers
* With both of the above sample configurations in place, `gerritr push -r x1@org.com frontend` will add `x1@org.com`, `b1@org.com`, `b2@org.com`, `r1@org.com` and `r2@org.com` as reviewers

### `patch`

Amend the latest commit with the staged changes and push it to the target branch, thereby [adding a patchset](https://gerrit-documentation.storage.googleapis.com/Documentation/3.6.2/intro-gerrit-walkthrough.html#_reworking_the_change) to an already existing change in Gerrit.

If `--message` flag is not specified, the commit message [will be preserved](https://git-scm.com/docs/git-commit#Documentation/git-commit.txt---no-edit). No change to the added reviewers.
