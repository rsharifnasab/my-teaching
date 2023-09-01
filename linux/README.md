---
theme: gaia
paginate: true
---

# Linux!

> Linux is a family of open-source Unix-like operating systems based on the Linux kernel, an *operating system kernel* first released on September 17, 1991, by **Linus Torvalds**.


**By Roozbeh Sharifnasab**
rsharifnasab@gmail.com

------------

# Linus Torvalds

+ Published Linux in 1991

+ Just for fun [book](http://linuxstory.ir/)

![bg right vertical 105%](./images/Linus-Torvalds.jpg)
![bg right vertical 75%](./images/linus.jpg)


------------

# Linux Birthday

> … I’m doing a (free) operating system (just a hobby, won’t be big and professional like gnu) for 386(486) AT clones. …

Linus (torvalds@kruuna.helsinki.fi) - August 25, 1991


https://linuxiac.com/linux-birthday/


------------------


# Why Linux? (Philosophy)
+ Free software vs Proprietary software?
+ Why GNU/Linux?

https://en.wikipedia.org/wiki/GNU/Linux_naming_controversy

-------------------

# Free software's 4 Freedoms

+ The freedom to run the program as you wish, for any purpose.
+ The freedom to study how the program works, and change it, so it does your computing as you wish. Access to the source code is a precondition for this.
+ The freedom to redistribute copies, so you can help others.
+ The freedom to distribute copies of your modified versions to others.

https://www.gnu.org/philosophy/free-sw.en.html

-----------------


# Why Linux as a Desktop OS
+ ~47% of Professional Devs
+ Free
+ Customizable
+ Fast & Lightweight
+ Security? Privacy?
+ Gaming?

https://www.mygreatlearning.com/blog/linux-vs-windows/

--------------

# Why Linux as a Server OS
+ ~96.3% of servers
+ All top 500 super computers
+ Low-spec (cheap) servers
+ Security
+ Pricing

https://truelist.co/blog/linux-statistics/

-------------

# Other Unix-like OSs
+ Android: Based on Linux kernel (by now)
+ BSD: allows commercialization
+ macOS: Darwin Kernel
+ PS4 kernel based on FreeBSD
+ Minix ([intel ME](https://itsfoss.com/fact-intel-minix-case/))

--------------

# Distro

+ [About disto](https://en.wikipedia.org/wiki/Linux_distribution)
+ [Distro timeline](https://upload.wikimedia.org/wikipedia/commons/1/1b/Linux_Distribution_Timeline.svg)

+ How to choose? [Distrowatch](https://distrowatch.com/index.php)

+ [Rolling vs fixed release?](https://averagelinuxuser.com/rolling-vs-fixed-linux-release/)

--------------

# Distro parts

+ Bootloader (GRUB)
+ Kernel
+ Init system (System D, Runit, OpenRC)
+ Daemons (background services)
+ Graphical server (X Server)
+ Display manager
+ Desktop environment (GNOME, Cinnamon, Mate, XFCE, KDE)
+ Applications

------------

# Terminal
+ Text-based
+ not interesting
+ Terminal vs Shell?
+ Terminal emulator?
https://superuser.com/questions/96628/why-are-things-like-gnome-terminal-called-terminal-emulators-instead-of-just

-----------


# basic commands
+ pwd, cd, ls
+ cp, mv
+ cat, head, tail, tailf
+ `|`, grep
+ `>`, `>>`, `<`

https://dev.to/awwsmm/101-bash-commands-and-tips-for-beginners-to-experts-30je

https://www.educative.io/blog/bash-shell-command-cheat-sheet


----------

# SSH
+ why?
+ why not graphical?

----------

# Edit files on terminal
+ Vim
+ Nano

---------

# File-system heirarchy

https://linuxhandbook.com/linux-directory-structure/

-----------

# Root?
+ Just a user. Why it is bad?
+ Root of file system


-----------

# Sudo
+ a prefix for some commands? No
+ Sudoer accounts?
https://www.linuxfordevices.com/tutorials/linux/sudo-command-in-linux-unix

---------------

# Package managing

+ Install/Upgrade/Remove
+ Apt/Pacman/Yum
+ Software Repository

### Good cheatsheet [here](https://blog.packagecloud.io/apt-cheat-sheet/)

-------------


# Get help
+ man
+ online man
https://www.man7.org/linux/man-pages/man1/uptime.1.html

+ [TLDR](https://tldr.sh/)

----------

# Bash scripting

+ hashbang

+ syntax: https://learnxinyminutes.com/docs/bash/

---------

# What if it isn't in software repo? 
+ Compiling from source [for example](https://github.com/rofl0r/ncdu)
+ Download from ppa [for example](https://persepolisdm.github.io/)
+ Install `.deb`

-----------

# What about Docker
+ libc: glibc, musl
+ utils: gnucoreutils, busybox
+ distro: ubuntu, alpine, scratch?


------- 

# APPENDIX-1: unix (philosophy)
+ [develop with unix philo in mind](https://monkey.org/~marius/unix-tools-hints.html)
+ [How C was born?](https://www.redhat.com/en/command-line-heroes/season-3/the-c-change?sc_id=701f20000012rt4AAA)
+ [Ritchie talk about Unix & C](https://www.youtube.com/watch?v=yY6YY81P3lE)
+ [Why Unix philosophy](https://www.youtube.com/watch?v=XvDZLjaCJuw)

--------

# APPENDIX-2: Torvalds-Tanenbaum debate
+ https://stackoverflow.com/questions/46410886/modular-kernel-vs-micro-kernel-monolitic-kernel
+ https://unix.stackexchange.com/questions/6409/how-does-linux-kernel-compare-to-microkernel-architectures
+ https://matt-rickard.com/16-lessons-from-the-tanenbaum-torvalds-debates
+ https://www.oreilly.com/openbook/opensources/book/appa.html

------------------


Made with [Marp](https://marp.app/)
