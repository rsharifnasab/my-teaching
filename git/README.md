---
theme: gaia
paginate: true
---

# Git!

> Git (/git/)[7] is a distributed version-control system for tracking changes in source code during software development.

------------------

# Manual version control 

![bg right 100% ](images/trad.png)

------------------

# Problems

-   save-as whole project 
-   collaboration
-   last stable version
-   space inefficiency
-   tracking changes

------------------

# Benefits of version control

-   easily manage collaboration on a project
-   ability to have an unlimited number of developers
-   revert changes if something goes wrong

------------------

# SVN (by Apache)

![bg auto](https://www.paulocollares.com.br/wp-content/uploads/2019/04/213-2134650_svn-logo-png-apache-subversion-logo.jpg)

------------------

# Visual Studio Team Services code (by Microsoft)

![bg 30% down](https://logodix.com/logo/719944.png)

------------------

# Git (by Linus Torvalds)

![bg contain](https://files.virgool.io/upload/users/4458/posts/q85kpw57vte6/quvbhkgvelb5.jpeg)

------------------

# As of 2020, the 5.6 release of the Linux kernel had around 33 million lines of code.

![bg contain](images/tux.png)

------------------

# Git features

-   free and open source
-   distributed
-   non-linear (branches)
-   handle large projects efficiently

![bg contain right 125%](./images/Linus-Torvalds.jpg)

------------------


![bg contain](https://homes.cs.washington.edu/~mernst/advice/version-control-fig2.png)

------------------


![bg contain](https://homes.cs.washington.edu/~mernst/advice/version-control-fig3.png)

------------------

# how to use git

1. Search!
2. I Search, too
3. everybody else does Search too.

------------------

![bg auto](https://slideplayer.com/slide/13332433/80/images/17/Git+file+lifecycle.jpg)

------------------

# Hands-on

+ initialize git
+ write some code
+ add, commit
+ status
+ checkout to another commit

------------------

# In-depth

+ Is Git **running** in the background?

-   What if we change a file twice?
-   What does **tracking** mean?
-   What if I uninstall git?
-   What about deleting `.git`?

------------------

# Best practices

+ I just wrote a line of code; should I commit? ([atomic commits](https://dev.to/paulinevos/atomic-commits-will-help-you-git-legit-35i7))
+ Should I commit on a timely basis?
+ What to write in the commit message?  ([fun](https://whatthecommit.com/)) ([conventional](https://gist.github.com/qoomon/5dfcdf8eec66a051ecd85625518cfd13))
+ aliases

------------------

# Branch

+ working on a feature
+ separate different programmers' work
+ do not mess up the main/master branch

![bg h:309 right](./images/git-branches-merge.png)

------------------

![bg fit](https://res.cloudinary.com/practicaldev/image/fetch/s--MEKaM3dY--/c_imagga_scale,f_auto,fl_progressive,h_900,q_auto,w_1600/https://cl.ly/430Q2w473e2R/Image%25202018-04-30%2520at%25201.07.58%2520PM.png)

------------------

# Merge conflict

-   What if we can't merge?
-   some developers changed the same file
-   somebody should merge 2 versions 
-   with the help of a merge-tool

![bg  right 100%](https://lh6.googleusercontent.com/proxy/EXZtnMuZcVrMmQ1YJ1vdyoadiEy-FQtUocRc5mWiOqUgcxp5SlJ-T-Bs8dFERfxym7E7U6SebY1PJRx9OYPJ5gtFrDPMMFF-)

------------------

# Git HEAD

+ a pointer to a commit 
+ the current state of working dir
+ where are the other commits?

![bg fit right](https://static.javatpoint.com/tutorial/git/images/git-head.png)

------------------

# Checkout

- switch to another branch (prefer `switch`)

- reset a file to a specific commit version

![bg fit right](https://static.javatpoint.com/tutorial/git/images/git-checkout.png)

------------------

# Branching best pratcies
- git flow
  - stable master branch
  - development branch
  - feature branch
  - release
  - hot-fix branch

------------------

# Branching best pratcies
- Trunk-based
  - DevOps friendly
  - merge small, frequent updates
  - core trunk = main branch
    checkout [this link](https://www.atlassian.com/continuous-delivery/continuous-integration/trunk-based-development)


![bg fit right](./images/trunk-vs-flow.jpg)


------------------


# Hands-on

+ create another branch 
+ checkout to another branch
+ add commits to different branches
+ merge changes
+ solve a merge conflict

------------------

# GitHub

- Instagram for git

- a place to keep git projects, review them, fork them, and star them.

- alternatives: Gitlab, Bitbucket, and more

![bg fit right](https://devmountain.com/wp-content/uploads/2022/01/Gitvs_Github-1a-1.jpg)

------------------

# Gist
+  some parts of code to share with others
+  not the whole program, just code snippets
+  it is not common
+  nor big projects, nor small ones do not have gists.
+  for example [a cheat sheet](https://gist.github.com/mayazdi/9c3fc4f6e9828a803be757d177cea8e1)

------------------

# Push? Remote? Clone?

-   remote: where should I upload my git projects
-   push: the act of uploading the git project
-   clone: download the whole git project
-   pull: check for updates in the remote git

![bg contain horizontal right](images/push-meme.png)


------------------

# fork

![fork](images/fork.png)

------------------

# PR

![bg contain ](images/PR.png)


------------------

# Hands-on


+ Want a review? [Check this](https://medium.com/hackernoon/a-gentle-introduction-to-git-and-github-the-eli5-way-43f0aa64f2e4)

![bg contain horizontal right](https://res.cloudinary.com/practicaldev/image/fetch/s--NUSRQ-3J--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://i.redd.it/5iphhycu0io11.png)


------------------


![bg fit contain ](https://github.blog/wp-content/uploads/2018/05/new-issue-page-with-multiple-templates.png?fit=1604%2C694)

------------------

# .gitignore, .git

-   .git: local and hidden folder that contains git internal files, don't open it!

-   .gitignore: ignore these sorts of files

```
*.class
.idea/
__pycache__/
```

Good site: [gitignore.io](https://www.gitignore.io/)

####  Do not commit large and binary files! Why?

------------------


# tag

-   commits don't have a name
-   to specify particular commits
-   which commit is the last working version?
-   how to handle versions?

- read more [here](https://www.atlassian.com/git/tutorials/inspecting-a-repository/git-tag)

![bg fit right](https://miro.medium.com/max/638/1*CW-G-2zVFpEfEZNOTBvT0g.jpeg)

------------------

# amend

-   re-writing history
-   change commit message
-   change committed files

### don't re-write history (especially for published commits)

------------------

# stash

-   save uncommitted changes
-   different options for uncommitted, ignored, staged

![fit right bg](https://static.javatpoint.com/tutorial/git/images/git-stash.png)

------------------

# revert

-   undo last commit changes

-   with a brand-new commit

![fit right bg](https://www.blog.nakulrajput.com/wp-content/uploads/2018/10/Git-Reverting-Resetting.jpg)

------------------

# bisect

![fit bg 80%](https://files.virgool.io/upload/users/195916/posts/kiybjsojv6cf/amufcdey5v9x.png)

------------------

# rebase

-   instead of merge
-   combine history with adding old commits
-   interactive mode: select some commits
-   be careful: might re-write history
-   don't do this in public branches or junior development teams.

![fit right bg](https://miro.medium.com/max/603/1*75GIyh005_Wl9vRQF3tgTQ.png)

### learn more [here](https://www.atlassian.com/git/tutorials/merging-vs-rebasing)

------------------

# squash

-   not a git command
-   a term in using rebase
-   group-specific changes to one commit

------------------

# cherry-pick

-   add other branch's commit to the current branch

![](https://www.syntevo.com/doc/download/attachments/6979705/workflows-cherrypick.png?version=1&modificationDate=1443735398000&api=v2)


------------------


# common commands

```bash
# first time initialize
git config --global user.name "Bugs Bunny"
git config --global user.email bugs@gmail.com
git init


# regulary code and commit
git status
git add -A # or git add filename
git commit -m 'commit message'
```

----------------

```bash
# work with remote
git remote add origin https://github.com/your-account/your-repository.git
git push origin master # push from master branch to origin remote
git pull # get new updates
git clone https://github.com/sb-acc/some-repo.git

# see old commits (and other beautiful versions of this command)
git log
git log --abbrev-commit --pretty=oneline
git log --graph --oneline --decorate --abbrev-commit

```

----------------

```bash
# work with history
git checkout branchname # change HEAD to another commit/branch
git reset --hard HEAD # revert to last commit
git diff HEAD # what is changed from head
git checkout HEAD filename # revert filename to last commit

git blame filename # who changed this file

```

----------------

```bash
# stash
git stash # save uncommited changes and revert repository to commited sate
git stash list
git stash pop # pop last saved stash


# tag
git tag tagname -am 'tag message' # create new tag
git tag # list all tags
git push --tags


```

----------------

```bash
# merge
git config --global merge.tool meld # meld or another app
git checkout destination
git merge source
git mergetool # open the merge tool

# get help
curl http:/cheat.sh/git
man git-add
info git-add


rm -rf .git # get rid of git!
```

------------------

# conclusion

-   search & use help & read documentation
-   don't fear
-   everything messed up: re-clone!
-   test new commands in a toy repository
-   don't commit large files
-   don't re-write public history
-   pull before push, even better, pull before starting coding
-   alias for your useful command could be good

------------------

# further read

-   [Interactive learning](https://learngitbranching.js.org/)
-   [Amir's awesome slides](https://github.com/amirhallaji/My-Learning/tree/master/git)
-   [perfect cheat sheet](https://github.com/arslanbilal/git-cheat-sheet)
-   [this GitHub io page](https://rachelcarmena.github.io/2018/12/12/how-to-teach-git.html)
-   [command by command, explain](https://recompilermag.com/issues/issue-1/how-to-teach-git/)
-   [Jadi's videos](https://faradars.org/courses/fvgit9609-git-github-gitlab)
-   [step by step](https://dont-be-afraid-to-commit.readthedocs.io/en/latest/git/commandlinegit.html)
-   [this good slide](https://courses.cs.washington.edu/courses/cse403/13au/lectures/git.ppt.pdf)
-   [cheatsheet](https://www.pauline-vos.nl/git-legit-cheatsheet/)

------------------


made with [Marp](https://marp.app/)
