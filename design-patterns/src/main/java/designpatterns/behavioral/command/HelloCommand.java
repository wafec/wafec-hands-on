package designpatterns.behavioral.command;

public class HelloCommand implements Command {
	private final HelloNotifier helloNotifier;

	public HelloCommand( final HelloNotifier helloNotifier ) {
		this.helloNotifier = helloNotifier;
	}

	@Override
	public void execute() {
		helloNotifier.sendHello( "Hello!" );
	}
}
