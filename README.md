# Git Repos Management Tool

This project is a command-line tool designed to manage files from external Git repositories. It allows users to add, validate, and manage different Git contexts with ease. Version management and update synchronization is provided by this tool.

## Features

- **Add Repositories**: Add new repositories to your list with the `add` command. You can specify the repository URL, subpath, version, local path, and whether it should be ignored.

- **Validate Repositories**: Each repository entry is validated to ensure it has all the necessary details and that the URL is correctly formatted.

- **Manage Repositories**: The tool allows you to manage your repositories by writing the details to a file. This makes it easy to keep track of all your repositories and their details.

- **Declarative State Management**: The tool uses a declarative approach to manage the state of your repositories. This makes it easy to add, update, and remove repositories as needed. Add your modifications to the root `.gitrepos` file and run the `apply` command to synchronize the repositories. That should be all you need to do. Since it is best to use specific git hashes tied to a specific version, the tool will not automatically update the repositories. This is to prevent any unexpected changes to your codebase.

## Usage

Initialize the current directory as root for any supplied .gitrepos entries.

```bash
git repos init
```

To add a new repository, use the `add` command with the necessary flags. For example:

```bash
git repos add --repo-url https://github.com/user/repo --version 1.0.0 --local-path ./my-repo
```

This will add a new repository with the specified details to your list.

If you wish to see the current status of the subdirectories, use the `status` command.

```bash
git repos status
```

This should supply the following output if all files are synchronized:

```bash
Up to date
```

Otherwise, something like this:

```bash
Synchronization needed

  my-repo current=1.0.1 local=1.0.0

Use `git repos apply` to synchronize
```

Clone the files from remote and track synchronization.

```bash
git repos apply
```

This will add a `.gitrepos` file to the root of each entry in the root `.gitrepos` file. This file will contain the version and timestamp of the repository that was last synchronized.

`.gitrepos` files that exist at the root of the directory have each entry begin with `-` and contain the following fields:

- `ignore`: Whether the repository should be ignored.
  - `?`: Do nothing
  - `!`: Present in .gitignore file
- `repo-url`: The URL of the repository.
  - Must begin with `https://` or `git@`
- `subpath`: The subpath of the repository.
  - Leading `/` will be stripped
- `version`: The version of the repository.
  - Can be a git branch, tag, or commit hash
- `local-path`: The local path of the repository.
  - Leading `/` will be stripped
- `timestamp`: The timestamp of the last synchronization.
