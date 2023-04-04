package designpatterns.behavioral.chainofresponsibility;

public interface Handler {
	void execute( Request request );
}
