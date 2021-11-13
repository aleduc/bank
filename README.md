##### These rules helped me not to spend too much time on the task.
#### simplifications and rules for task
1. was skipped HTTP server part(requests parsings, verifycation)
1. was used only one database, without additional schemas
1. only methods/interfaces declarations without splitting packages
1. db part only tables, without functions
1. simple errors, wthout codes and somthing else


#### Additional info
1. We can calculate balance in async way. With help from pgq or kafka. But for this task it doesn't matter.
1. Notification does only one thing, read from notification topic from kafka and send notifications.