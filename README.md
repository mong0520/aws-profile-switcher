# aws-profile-switcher

Switch your AWS profile easily

<img src=./res/demo.gif>

## Prerequisite

Configure your AWS profile via aws cli

## Setup

```sh
git clone https://github.com/mong0520/aws-profile-switcher.git
cd aws-profile-switcher
make install

# reload or restart your terminal
```

Adding the following to your `.bashrc` or `.zshrc` config in order to take effect immediately, and then reload your shell

```sh
# make sure new session takes the new aws profile
source ~/.aws_exports
# make sure the current session takes the new aws profile immediately
alias aws-profile-switcher="source run_aws-profile-switcher.sh"
```

## Usage

```
aws-profile-switcher
```
