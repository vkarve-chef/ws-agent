# ws-agent
Prototype code to test gRPC calls to LaunchAgent based process.

## Development Setup
1. Run `go build`
2. Copy `ws-agent` to `/tmp/'
3. Copy `io.chef.workstation.plist` to `~/Library/LaunchAgents/`
4. Run `launchctl load ~/Library/LaunchAgents/io.chef.workstation.plist`
5. Run `launchctl start ~/Library/LaunchAgents/io.chef.workstation.plist`

Observe the log file `/tmp/ws-agent.out.log`. The dummy output from process will be appended here.
Stop and unload the service so that the log file size doesn't keep growing, as this only for testing.

1. Run `launchctl stop ~/Library/LaunchAgents/io.chef.workstation.plist`
2. Run `launchctl unload ~/Library/LaunchAgents/io.chef.workstation.plist`
