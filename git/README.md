---
theme: gaia
paginate: true
---

# Git!

> Git (/git/)[7] is a distributed version-control system for tracking changes in source code during software development.


**By Roozbeh Sharifnasab**
rsharifnasab@gmail.com

------------------

# Manual version control

![bg right 100%](images/trad.png)

------------------

# Problems

-   Save-as whole project
-   Collaboration
-   Last stable version
-   Space inefficiency
-   Tracking changes

------------------

# Benefits of version control

-   Manage collaboration on a project
-   Unlimited number of developers
-   Revert changes if something went wrong

------------------

# SVN (by Apache)

![bg auto](./images/svn.jpg)

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

-   Free and open source
-   Distributed
-   Non-linear history (branches)
-   Handle large projects efficiently

![bg contain right 125%](./images/Linus-Torvalds.jpg)

------------------


![bg contain](./images/version-control-fig2.png)

------------------


![bg contain](./images/version-control-fig3.png)



---------

![bg auto](./images/local.jpg)

------------------

# How to use Git

![bg center 30%](./images/search-icon.png)

------------------

![bg auto](https://slideplayer.com/slide/13332433/80/images/17/Git+file+lifecycle.jpg)


------

![bg contain](./images/add.jpg)

------------------

# Hands-on

+ Initialize Git
+ Write code
+ Add to staging, Commit
+ Status
+ Git log/checkout

![bg contain right](./images/git-init-meme.jpg)

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
+ What to write in the commit message?  ([conventional commit message](https://gist.github.com/qoomon/5dfcdf8eec66a051ecd85625518cfd13))

-----------------

![bg contain horizontal](./images/commit-meme.jpg)
![bg contain horizontal](./images/commit-meme-2.jpg)

------------------

# Branch

+ Working on a feature
+ Split different programmers' work
+ Do not mess up the main/master branch

![bg h:309 right](./images/git-branches-merge.png)

----------

![bg contain](./images/branching.jpg)

------------------

![bg fit](./images/merge-commit.jpg)

------------------

# Merge conflict

-   What if we can't merge?
-   Two developers changed the same file
-   Somebody should merge 2 versions
-   Merge-tool

![bg  right 100%](./images/conflict-2.png)

------------------

# Git HEAD

+ Pointer to a commit
+ Current state of the working dir
+ Where are the other commits?

![bg fit right](https://static.javatpoint.com/tutorial/git/images/git-head.png)

------------------

# Checkout

- Switch to another branch (prefer `switch`)

- Reset a file to a specific commit version

![bg fit right](https://static.javatpoint.com/tutorial/git/images/git-checkout.png)

------------------

# Hands-on

+ Create another branch
+ Checkout to another branch
+ Add commits to different branches
+ Merge changes
+ Solve a merge conflict

------------------

# Branching best pratcies

- Topic branches
  - Avoid merging to master directly (PR & code review)
  - Keep master branch up-to-date and high quality

![vertical bg right 105%](./images/topic.png)

------------------

# Branching best pratcies

- Release branches
  - Stabilize a release of code
  - Long-lived
  - Lock after end of support
  - Do not merge to master
  - Port updates from master

- Read more [here](https://learn.microsoft.com/en-us/azure/devops/repos/git/git-branching-guidance?view=azure-devops)

![bg contain right](./images/release-flow.png)

------------------

# Branching best pratcies

- **Git flow**
  - Stable master branch
  - Feature branch
  - Hot-fix branch
  - Release branch
  - Development branch
![bg fit right](./images/develop.png)

------------------

# Branching best pratcies
- **Trunk-based**
  - DevOps friendly
  - merge small, frequent updates
  - core trunk = main branch

Read more [here](https://www.atlassian.com/continuous-delivery/continuous-integration/trunk-based-development)


![bg fit right](./images/trunk-vs-flow.jpg)


------------------

# GitHub

- Instagram for git

- A place to keep git projects, review them, fork them, and star them.

- Alternatives: Gitlab, Bitbucket, and more

![bg fit right](https://devmountain.com/wp-content/uploads/2022/01/Gitvs_Github-1a-1.jpg)

------------------

# Gist
+  A code snippet
+  It is not common
+  Nor big projects, nor small ones
+  like [a cheat sheet](https://gist.github.com/mayazdi/9c3fc4f6e9828a803be757d177cea8e1)

![bg contain right](./images/gist.webp)

------------------

# Push? Remote? Clone?

-   remote: where should I upload my git projects
-   push: the act of uploading the git project
-   pull: check for updates in the remote git
-   clone: download the whole git project

![bg contain horizontal right](images/push-meme.png)


------------------

# Fork

![fork](images/fork.png)

------------------

# Pull Request (PR)

![bg contain ](images/PR.png)


------------------

# Hands-on

+ Want a review? [Check this](https://medium.com/hackernoon/a-gentle-introduction-to-git-and-github-the-eli5-way-43f0aa64f2e4)

+ A cheat sheet? [Check this](https://github.com/arslanbilal/git-cheat-sheet)

![bg contain horizontal right](https://res.cloudinary.com/practicaldev/image/fetch/s--NUSRQ-3J--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://i.redd.it/5iphhycu0io11.png)


------------------

<!-- Issue -->
![bg fit contain](https://github.blog/wp-content/uploads/2018/05/new-issue-page-with-multiple-templates.png?fit=1604%2C694)

------------------

# .gitignore

- .git: git internal files
    - don't modify it!

- .gitignore: which files should be ignored by git?
```
*.class
node_modules/
__pycache__/
```

[gitignore.io](https://www.gitignore.io/)

![bg contain right](./images/ignore.jpg)


-------------------

## File to ignore
+ Blob files
+ Auto-generated files
+ IDE-specific files
+ Assets?
+ Passwords?

![bg fit contain right vertical](./images/onde_modules.png)
![bg fit contain right vertical](./images/ignore2.webp)

------------------


# Tag

- Commits don't have a name
- How to specify commits
- Which commit was the last stable version?
- How to handle versions?

- Read more [here](https://www.atlassian.com/git/tutorials/inspecting-a-repository/git-tag)

![bg fit right](./images/tag.webp)

------------------

# Fetch VS Pull

+ Pull = Fetch + merge
+ Just merge?
+ What if we don't merge

![bg fit right vertical](https://rachelcarmena.github.io/img/cards/posts/how-to-teach-Git/fetch.png)
![bg fit right vertical](https://rachelcarmena.github.io/img/cards/posts/how-to-teach-Git/pull.png)

-----------------

# History
+ Once again, why are we controlling version?
+ What was my last version?
+ Who/When did this change?

![bg fit right contain](https://res.cloudinary.com/practicaldev/image/fetch/s---vCBqatm--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://thepracticaldev.s3.amazonaws.com/i/48schr7iqawvz6skkkbb.png)

------------------

# Revert

- Undo last commit changes

- With a brand-new commit

![fit right bg](https://www.blog.nakulrajput.com/wp-content/uploads/2018/10/Git-Reverting-Resetting.jpg)

----------------

# Amend

Change latest commit
    - commit message
    - or even committed files

![bg fit contain right](https://www.w3docs.com/uploads/media/default/0001/03/998b9219a5a05a184c9e3905b0714880537ee1a7.png)

#rewrite_history

------------------

# Rebase

- Another way of merge
- Combine history by adding old commits
- Not suitable for juniors
- learn more [here](https://www.atlassian.com/git/tutorials/merging-vs-rebasing)
- Linear history

![fit right contain bg](./images/rebase.png)

#rewrite_history

-----------------

![bg contain](./images/rebase.jpg)

------------------

# Interactive Rebase

- Select recent some commits to edit
- Squash
    - not a git command
    - a term in using rebase
    - group specific changes to one commit

![bg contain right](./images/rebase-i.jpg)

------------------

# Force push

+ Don't re-write history
    + Published commits
    + Master branch

![bg contain fit right](https://www.mememaker.net/static/images/memes/3914636.jpg)

-----------------

# Force with lease push
+ Force push but safer
+ Check if is OK or not
+ In new git version

![bg contain right](./images/force-meme.jpeg)

---

# Stash

- save uncommitted changes
- What to do for
    + uncommitted
    + ignored
    + staged

![fit right bg](https://static.javatpoint.com/tutorial/git/images/git-stash.png)

------------------

# Bisect

![fit bg 80%](https://files.virgool.io/upload/users/195916/posts/kiybjsojv6cf/amufcdey5v9x.png)

------------------

# Cherry-pick

- add other branch's commit to the current branch

![bg fit contain right](https://res.cloudinary.com/practicaldev/image/fetch/s--jmZDSObz--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/i/dmb185yxo5umin72abyr.jpg)

------------------

# Conclusion

- Search & use help & read documentation
- Don't panic
- Test new commands in a toy repository
- Don't commit large files
- Don't re-write public history
- Pull before push, even better, pull before starting coding
- Aliasing for your useful command

--------------

![bg contain](./images/torvalds-ignore.webp)

------------------

![bg contain](./images/re-clone.jpg)

-----------------

![bg contain](./images/collaboration.jpg)

------------------

# further read

- [Interactive learning](https://learngitbranching.js.org/)
- [Amir's awesome slides](https://github.com/amirhallaji/My-Learning/tree/master/git)
- [perfect cheat sheet](https://github.com/arslanbilal/git-cheat-sheet)
- [this GitHub io page](https://rachelcarmena.github.io/2018/12/12/how-to-teach-git.html)
- [command by command, explain](https://recompilermag.com/issues/issue-1/how-to-teach-git/)
- [Jadi's videos](https://faradars.org/courses/fvgit9609-git-github-gitlab)
- [Roadmap](https://devopscube.com/git-for-devops/)
- [this good slide](https://courses.cs.washington.edu/courses/cse403/13au/lectures/git.ppt.pdf)
- [cheatsheet](https://www.pauline-vos.nl/git-legit-cheatsheet/)

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


made with [Marp](https://marp.app/)

+ run with:
```bash
marp README.md --allow-local-files -w
```
