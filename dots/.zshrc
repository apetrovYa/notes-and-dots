# If you come from bash you might have to change your $PATH.
# export PATH=$HOME/bin:/usr/local/bin:$PATH

# Path to your oh-my-zsh installation.
export ZSH=/home/andov/.oh-my-zsh

# Set name of the theme to load. Optionally, if you set this to "random"
# it'll load a random theme each time that oh-my-zsh is loaded.
# See https://github.com/robbyrussell/oh-my-zsh/wiki/Themes
ZSH_THEME="theunraveler"

# Set list of themes to load
# Setting this variable when ZSH_THEME=random
# cause zsh load theme from this variable instead of
# looking in ~/.oh-my-zsh/themes/
# An empty array have no effect
# ZSH_THEME_RANDOM_CANDIDATES=( "robbyrussell" "agnoster" )

# Uncomment the following line to use case-sensitive completion.
# CASE_SENSITIVE="true"

# Uncomment the following line to use hyphen-insensitive completion. Case
# sensitive completion must be off. _ and - will be interchangeable.
# HYPHEN_INSENSITIVE="true"

# Uncomment the following line to disable bi-weekly auto-update checks.
# DISABLE_AUTO_UPDATE="true"

# Uncomment the following line to change how often to auto-update (in days).
# export UPDATE_ZSH_DAYS=13

# Uncomment the following line to disable colors in ls.
# DISABLE_LS_COLORS="true"

# Uncomment the following line to disable auto-setting terminal title.
# DISABLE_AUTO_TITLE="true"

# Uncomment the following line to enable command auto-correction.
# ENABLE_CORRECTION="true"

# Uncomment the following line to display red dots whilst waiting for completion.
# COMPLETION_WAITING_DOTS="true"

# Uncomment the following line if you want to disable marking untracked files
# under VCS as dirty. This makes repository status check for large repositories
# much, much faster.
# DISABLE_UNTRACKED_FILES_DIRTY="true"

# Uncomment the following line if you want to change the command execution time
# stamp shown in the history command output.
# The optional three formats: "mm/dd/yyyy"|"dd.mm.yyyy"|"yyyy-mm-dd"
# HIST_STAMPS="mm/dd/yyyy"

# Would you like to use another custom folder than $ZSH/custom?
# ZSH_CUSTOM=/path/to/new-custom-folder

# Which plugins would you like to load? (plugins can be found in ~/.oh-my-zsh/plugins/*)
# Custom plugins may be added to ~/.oh-my-zsh/custom/plugins/
# Example format: plugins=(rails git textmate ruby lighthouse)
# Add wisely, as too many plugins slow down shell startup.
plugins=(
  git aws docker pip 
)

source $ZSH/oh-my-zsh.sh

# User configuration

# export MANPATH="/usr/local/man:$MANPATH"

# You may need to manually set your language environment
# export LANG=en_US.UTF-8

# Preferred editor for local and remote sessions
if [[ -n $SSH_CONNECTION ]]; then
   export EDITOR='vim'
# else
#   export EDITOR='mvim'
 fi

# Export Maven configuration
# To remember: {{maven}} name is a soft link
export PATH="/opt/maven/bin:$PATH"
export PATH="/opt/apps/xmind/XMind_amd64:$PATH"

# Export Golang configuration
#
export PATH="/opt/go/bin:$PATH"
export GOPATH="$HOME/Documents/go"

# Compilation flags
# export ARCHFLAGS="-arch x86_64"

# ssh
# export SSH_KEY_PATH="~/.ssh/rsa_id"

# Set personal aliases, overriding those provided by oh-my-zsh libs,
# plugins, and themes. Aliases can be placed here, though oh-my-zsh
# users are encouraged to define aliases within the ZSH_CUSTOM folder.
# For a full list of active aliases, run `alias`.
#
# Example aliases
# alias zshconfig="mte ~/.zshrc"
# alias ohmyzsh="mate ~/.oh-my-zsh"
#

# Include and execute z.sh tool
. /opt/apps/z.sh


#
#
# Export the used keys
#
#
# The IKS -- working key
ssh-add ~/.ssh/iks_id_rsa
# The GSS -- working within customer's key
ssh-add ~/.ssh/generali_id_rsa

#
# Mount the Network File System of IKS on this computer
#
# This function requires the username and password
# of the user to mount the disk
#
function mount_iks_nfs () {
	local username="${NFS_DOMAIN_USERNAME}"
	local password="${NFS_DOMAIN_PASSWORD}"
	mkdir -p /home/apetrov/shares/iks_nfs
	sudo mount -t cifs \
		   -o username="${username}",password="${password}" \
		   --source //192.168.100.155/ \
		   --target /home/apetrov/shares/iks_nfs
}


#
# Docker shortcuts #
#
#
# A shortcut command to start docker daemon
#
function start_docker () {
	sudo systemctl start docker
}

#
# A shortcut command to show the docker daemon status
#
function status_docker () {
	sudo systemctl status docker
}

#
# A shortcut command to stop the docker daemon
#
function stop_docker () {
	sudo systemctl stop docker
}

#		
# GIT SHORTCATS 
#		

# Delete all current changes and revert the tracked branch to the latest commit
function get_rollback_to_last_commit () {
	git reset --hard HEAD@{1}
}

# Add all changes to Git Staging
function g2a () {
	git add *
}

# Get the current Git context status
function gs() {
       git status
}

#
# KUBERNETES KUBECTL CLI FOR INTERACTION WITH THE CLUSTER
#
source <(kubectl completion zsh)

#
# Autoenv file
#
source ~/.autoenv/activate.sh

#					     
# THE PERFORMANCE ANALYSIS SCRIPTS AND TOOLS 
# 					     
#
# The command below prints a line on the terminal
# with the following format. Reading from right to
# left, is following: z y x load average data.
# if z > x we are too late and the dister happened
# already.
# The time interval of the times
# 1. x is computed any 1 minute
# 2. y is computed any 5 minutes
# 3. z is computed anu 15 minutes
function perf_get_uptime_0 () {
	uptime
}

#
# Get the las 10 messages printed by the system control
# process.
# It is worth to check
function perf_get_system_messages_1 () {
	dmesg | tail
}

#
# Get the 1 second system virtual memory statistcs.
# Remember these are at server level.
# Check the following lines:
# 1. r: Number of processes running on the CPU and waiting for a turn.
#       if r > vcpus then system.isSaturated="true" fi
# 2. free: Free memory in kilobytes.
# 3. Swap-ins and Swap-outs.
#    if si and so != 0 then system.isOOM="true" fi
# 4. us,sy,id,wa,st Are the breakdowns of CPU time.
function perf_get_vmstat_2 () {
	vmstat 1
}


#
# This command prints CPU time breakdowns per CPU.
# A single hot CPU can be evidence of a single threaded application.
#
function perf_get_mpstat_3 () {
	mpstat -P ALL 1
}

#
# "pidstat" is like the top linux utility. It prints a rolling summary
# instead of clearing the screen. This may be useful to find time patterns.
# For example: in case of Java applications to find the processes that are
# responsible for consuming CPU.
function perf_get_pidstat_4 () {
	pidstat 1
}

#
# This is a tool to understand the block devices (workload applied and
# performance).
# The parameters to look are:
#   - r/s, w/s, rkB/s, wkB/s.
#   - await
#   - avgqu-sz
#   - %util
function perf_get_iostat_5 () {
	iostat -xz 1
}

#
# The right to parameters are the buffers and cache.
# These two should not be zero. If this is the case then
# we are having performance problems.
#
function perf_get_mem_usage_6 () {
	free -m
}

#
# This is a networking interface throughput:
# rxkB/s and txkB/s.
#
function perf_get_sar_7 () {
	sar -n DEV 1
}

#
# A summarized view of some key TCP metrics.
# - active/s
# - passive/s
# - retrans/s
#
function perf_get_sar_TCP_8 () {
	sar -n TCP,ETCP 1
}


#			#
# The Keybase shortcuts #
#			#
# How to decrypt a message. In output, STDOUT, will
# be printed the decrypted message.
function kb_decrypt_file () {
	if [ "$1" != "" ]; then
		keybase decrypt -i $1 -o decrypted_$2
	else
		echo "Provide the file to work on for decrypting..."
	fi
}




