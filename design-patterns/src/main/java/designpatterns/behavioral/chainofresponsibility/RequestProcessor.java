package designpatterns.behavioral.chainofresponsibility;

@FunctionalInterface
public interface RequestProcessor {
	ProcessorResult processRequest( Request request );
}
