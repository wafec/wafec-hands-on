package designpatterns.behavioral.chainofresponsibility;

import java.util.Optional;

public class HandlerImpl implements Handler {
	private final Handler nextHandler;
	private final RequestProcessor requestProcessor;

	public HandlerImpl( final Handler nextHandler, final RequestProcessor requestProcessor ) {
		this.nextHandler = nextHandler;
		this.requestProcessor = requestProcessor;
	}

	@Override
	public void execute( final Request request ) {
		final ProcessorResult processorResult = requestProcessor.processRequest( request );
		if ( processorResult.shouldCallNextHandler() ) {
			Optional.ofNullable( nextHandler ).ifPresent( nextHandler1 -> nextHandler1.execute( request ) );
		}
	}
}
