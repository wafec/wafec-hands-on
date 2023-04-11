package designpatterns.behavioral.command;

class Invoker {
	private final Command command;

	public Invoker( final Command command ) {
		this.command = command;
	}

	public void executeCommand() {
		command.execute();
	}
}
