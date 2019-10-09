<div class="info">
  <h1 class="name" align="center">🔥<br>filetinder</h1>
</div>


## Summary
Tinderify you files. Reveiew your files, keep what you liked, and delete what you hated.

## Installation

Download latest release from [release page](https://github.com/shrotavre/filetinder/releases). These releases should work nicely in linux based distros.

Extract it anywhere and it'll be ready for use.

## User Guide

Once you downloaded the latest release, you will be presented to a `filetinder` executable. This cli app provide these commands you can use:

### Start from the CLI

```bash
# Basic CLI
$ filetinder start # this will start filetinder session in 
                   # background and give you link to open FileTinder's UI
                   # default will be available at http://localhost:17763/ui/

$ filetinder stop  # this will stop all running filetinder processes

# Adding/Including Files
$ filetinder add ./somedirectory/somefile.txt # Add single file
$ filetinder add ./somedirectory/ # Add all files in directory
$ filetinder add /home/somefile.jpg # Add using absolute path

# List Included Files
$ filetinder list # this will lists all included files in terminal

# Removing Included Files
$ filetinder remove 4 # this will remove 4th included file from list
```

### Tagging and removing files via GUI

Once you started filetinder, open the link provided by the cli, defaultly will be at [localhost:17763/ui](http://localhost:17763/ui/).

You will be presented with interface below.

![Contribution guidelines for this project](docs/images/ui-preview.gif)

## Additional Notes

FileTinder support any type of files to include in, **but FileTinder UI's Preview only support previewing image-typed files (currently)**

Contributions are open!

## Contributing

FileTinder is an OPEN Open Source Project. This means that:

Individuals making significant and valuable contributions are given
commit-access to the project to contribute as I see fit.