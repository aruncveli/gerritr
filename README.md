# gerritr

Wrapping some [Git](https://git-scm.com/) for [Gerrit](https://www.gerritcodereview.com/)

A command line app with opinionated flows for a subset of Gerrit functionalities. Intended to ease people familiar with GitHub et al., to work with Gerrit. And hopefully help them remember to add reviewers to a change.

## Installation
Download the binary and add to `$PATH`.

## Usage
`gerritr` is supposed to be used in the same context as `git` - a command line application from the repository root directory. `gerritr` has two subcommands: `push` and `patch`.

### Common flags
* `--branch` or `-b`: Target branch name. Required if the target branch is not main or master.
* `--message` or `-m`: Commit message.
* `--state` or `-s`: Change state. Ignored if the value is not one of:
	* `private`
	* `remove-private`
	* `wip`
	* `ready`

### `push`

#### Flags
* `--reviewers` or `-r`: Space separated list of reviewer email IDs or teams

Push the latest commit to the target branch, thereby [creating a new change](https://gerrit-documentation.storage.googleapis.com/Documentation/3.6.2/intro-gerrit-walkthrough.html#_creating_the_review) in Gerrit.

Adding reviewers automatically is possible in two ways:
#### `config.yml` or global config
* Go to:
	* `~/.config` for Linux
	* `~/Library/Application Support` for MacOS
	* `%LOCALAPPDATA%` for Windows
* Create a `gerritr` directory there. Go inside `gerritr`.
* Add a `config.yml` there with content similar to the sample below:
	```YAML
	teams:
	  backend:
		- b1@org.com
		- b2@org.com
	  frontend:
		- f1@org.com
		- f2@org.com
	```

#### `REVIEWERS` or repository specific config
Add a plaintext `REVIEWERS` file to the repository root directory, with a list of email IDs of people who should **always be added to every change pushed from this repository**:
```Text
r1@org.com
r2@org.com
...
```

Both the ways to add reviewers are independent. They can work together and alone:
* After just the above global configuration in place, `gerritr push -r backend x1@org.com` will add `x1@org.com`, `b1@org.com` and `b2@org.com` as reviewers
* After just the above local configuration in place, `gerritr push` will add `r1@org.com` and `r2@org.com` as reviewers
* With both of the above sample configurations in place, `gerritr push -r x1@org.com frontend` will add `x1@org.com`, `b1@org.com`, `b2@org.com`, `r1@org.com` and `r2@org.com` as reviewers

### `patch`

Amend the latest commit and push it to the target branch, thereby [adding a patchset](https://gerrit-documentation.storage.googleapis.com/Documentation/3.6.2/intro-gerrit-walkthrough.html#_reworking_the_change) to an already existing change in Gerrit.

If `--message` flag is not specified, the commit message [will be preserved](https://git-scm.com/docs/git-commit#Documentation/git-commit.txt---no-edit). No change to the added reviewers.
