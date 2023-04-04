package designpatterns.behavioral.chainofresponsibility;

/*
	Main idea: pass along a request to multiple handlers in a sequence
	Pros:
		1. Open/Closed Principle: new handlers do not modify existent handlers
		2. Single Responsibility: a handler does what it is meant to do only
		3. Decoupling: invocation is decoupled from processing
		4. Reuse: the chain can be built at runtime
	Cons:
		1. Requests might end up unhandled
 */
public interface Handler {
	void execute( Request request );
}
