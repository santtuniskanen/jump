# jump

We have lots of instances at work and due to the nature of the architecture, you'll often have to inspect the same logs on separate machines to find out issues in the message queue. I've gotten tired of having to manually move between machines, having to ssh to each separate machine, often having to go through a bastion. I could have just written some scripts, and I assume I could just save TMUX sessions as well (and I know there are apps like tmuxinator that could do this already?), but that would be boring. I've wanted to use Bubble Tea to make a Terminal UI application for a while, so I thought this would be a good opportunity to do so.

## Goals
- ~~basic `toml` config~~
- ~~"jump" into an instance~~
- jump into multiple instances at once
- sqlite for some sort of persistence, like templates?
- add configuration options to Terminal UI launcher, like the ability to configure windows, commands, change templates?
- separate sessions?
