# jump

We have lots of instances at work and due to the nature of the architecture, you'll often have to inspect the same logs on separate machines to find out issues in the message queue. I've gotten tired of having to manually move between machines, having to ssh to each separate machine, often having to go through a bastion. I could have just written some scripts, and I assume I could just save TMUX sessions as well, but that would be boring. I've wanted to use Bubble Tea to make a Terminal UI application for a while, so I thought this would be a good opportunity to do so.

The architecture for this isn't anything special. Bubble Tea is just a launcher where you can choose which instance to jump to and configure your launch options, like do you want to tail logs, or just get shell access to those machines. I thought this would also be a good place to change your configs, whether it's target machines or the layout of tmux panels and which target goes to which panel.
