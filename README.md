# jump

We have lots of instances at work and due to the nature of the architecture, you'll often have to inspect the same logs on separate machines to find out issues in the message queue. I've gotten tired of having to manually move between machines, having to ssh to each separate machine, often having to go through a bastion. I could have just written some scripts, and I assume I could just save TMUX sessions as well, but that would be boring. I've wanted to use Bubble Tea to make a Terminal UI application for a while, so I thought this would be a good opportunity to do so.

## Goals
- initial configuration options
- launch tmux sessions
- build launcher for session management
- build configuration manager for the launcher
- build template options for the launcher
